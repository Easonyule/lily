package util

import (
	"bytes"
	"context"
	"strings"

	"github.com/filecoin-project/go-address"
	builtin "github.com/filecoin-project/lotus/chain/actors/builtin"
	"github.com/filecoin-project/lotus/chain/state"
	"github.com/filecoin-project/lotus/chain/store"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/vm"
	"github.com/filecoin-project/specs-actors/actors/util/adt"
	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	"github.com/ipld/go-ipld-prime"
	"golang.org/x/xerrors"

	builtininit "github.com/filecoin-project/lily/chain/actors/builtin/init"
	"github.com/filecoin-project/lily/lens"
	"github.com/filecoin-project/lily/tasks/messages"
	"github.com/filecoin-project/lily/tasks/messages/fcjson"
)

var log = logging.Logger("lens/util")

// GetMessagesForTipset returns a list of messages sent as part of pts (parent) with receipts found in ts (child).
// No attempt at deduplication of messages is made. A list of blocks with their corresponding messages is also returned - it contains all messages
// in the block regardless if they were applied during the state change.
func GetExecutedAndBlockMessagesForTipset(ctx context.Context, cs *store.ChainStore, next, current *types.TipSet) (*lens.TipSetMessages, error) {
	if !types.CidArrsEqual(next.Parents().Cids(), current.Cids()) {
		return nil, xerrors.Errorf("child tipset (%s) is not on the same chain as parent (%s)", next.Key(), current.Key())
	}

	getActorCode, err := MakeGetActorCodeFunc(ctx, cs.ActorStore(ctx), next, current)
	if err != nil {
		return nil, err
	}

	// Build a lookup of which blocks each message appears in
	messageBlocks := map[cid.Cid][]cid.Cid{}
	for blockIdx, bh := range current.Blocks() {
		blscids, secpkcids, err := cs.ReadMsgMetaCids(bh.Messages)
		if err != nil {
			return nil, xerrors.Errorf("read messages for block: %w", err)
		}

		for _, c := range blscids {
			messageBlocks[c] = append(messageBlocks[c], current.Cids()[blockIdx])
		}

		for _, c := range secpkcids {
			messageBlocks[c] = append(messageBlocks[c], current.Cids()[blockIdx])
		}

	}

	bmsgs, err := cs.BlockMsgsForTipset(current)
	if err != nil {
		return nil, xerrors.Errorf("block messages for tipset: %w", err)
	}

	pblocks := current.Blocks()
	if len(bmsgs) != len(pblocks) {
		// logic error somewhere
		return nil, xerrors.Errorf("mismatching number of blocks returned from block messages, got %d wanted %d", len(bmsgs), len(pblocks))
	}

	count := 0
	for _, bm := range bmsgs {
		count += len(bm.BlsMessages) + len(bm.SecpkMessages)
	}

	// Start building a list of completed message with receipt
	emsgs := make([]*lens.ExecutedMessage, 0, count)

	// bmsgs is ordered by block
	var index uint64
	for blockIdx, bm := range bmsgs {
		for _, blsm := range bm.BlsMessages {
			msg := blsm.VMMessage()
			// if a message ran out of gas while executing this is expected.
			toCode, found := getActorCode(msg.To)
			if !found {
				log.Warnw("failed to find TO actor", "height", next.Height().String(), "message", msg.Cid().String(), "actor", msg.To.String())
			}
			// we must always be able to find the sender, else there is a logic error somewhere.
			fromCode, found := getActorCode(msg.From)
			if !found {
				return nil, xerrors.Errorf("failed to find from actor %s height %d message %s", msg.From, next.Height(), msg.Cid())
			}
			emsgs = append(emsgs, &lens.ExecutedMessage{
				Cid:           blsm.Cid(),
				Height:        current.Height(),
				Message:       msg,
				BlockHeader:   pblocks[blockIdx],
				Blocks:        messageBlocks[blsm.Cid()],
				Index:         index,
				FromActorCode: fromCode,
				ToActorCode:   toCode,
			})
			index++
		}

		for _, secm := range bm.SecpkMessages {
			msg := secm.VMMessage()
			toCode, found := getActorCode(msg.To)
			// if a message ran out of gas while executing this is expected.
			if !found {
				log.Warnw("failed to find TO actor", "height", next.Height().String(), "message", msg.Cid().String(), "actor", msg.To.String())
			}
			// we must always be able to find the sender, else there is a logic error somewhere.
			fromCode, found := getActorCode(msg.From)
			if !found {
				return nil, xerrors.Errorf("failed to find from actor %s height %d message %s", msg.From, next.Height(), msg.Cid())
			}
			emsgs = append(emsgs, &lens.ExecutedMessage{
				Cid:           secm.Cid(),
				Height:        current.Height(),
				Message:       secm.VMMessage(),
				BlockHeader:   pblocks[blockIdx],
				Blocks:        messageBlocks[secm.Cid()],
				Index:         index,
				FromActorCode: fromCode,
				ToActorCode:   toCode,
			})
			index++
		}

	}

	// Retrieve receipts using a block from the child tipset
	rs, err := adt.AsArray(cs.ActorStore(ctx), next.Blocks()[0].ParentMessageReceipts)
	if err != nil {
		return nil, xerrors.Errorf("amt load: %w", err)
	}

	if rs.Length() != uint64(len(emsgs)) {
		// logic error somewhere
		return nil, xerrors.Errorf("mismatching number of receipts: got %d wanted %d", rs.Length(), len(emsgs))
	}

	// Create a skeleton vm just for calling ShouldBurn
	vmi, err := vm.NewVM(ctx, &vm.VMOpts{
		StateBase:   current.ParentState(),
		Epoch:       current.Height(),
		Bstore:      cs.StateBlockstore(),
		NtwkVersion: DefaultNetwork.Version,
	})
	if err != nil {
		return nil, xerrors.Errorf("creating temporary vm: %w", err)
	}

	parentStateTree, err := state.LoadStateTree(cs.ActorStore(ctx), current.ParentState())
	if err != nil {
		return nil, xerrors.Errorf("load state tree: %w", err)
	}

	// Receipts are in same order as BlockMsgsForTipset
	for _, em := range emsgs {
		var r types.MessageReceipt
		if found, err := rs.Get(em.Index, &r); err != nil {
			return nil, err
		} else if !found {
			return nil, xerrors.Errorf("failed to find receipt %d", em.Index)
		}
		em.Receipt = &r

		burn, err := vmi.ShouldBurn(ctx, parentStateTree, em.Message, em.Receipt.ExitCode)
		if err != nil {
			return nil, xerrors.Errorf("deciding whether should burn failed: %w", err)
		}

		em.GasOutputs = vm.ComputeGasOutputs(em.Receipt.GasUsed, em.Message.GasLimit, em.BlockHeader.ParentBaseFee, em.Message.GasFeeCap, em.Message.GasPremium, burn)

	}
	blkMsgs := make([]*lens.BlockMessages, len(next.Blocks()))
	for idx, blk := range next.Blocks() {
		msgs, smsgs, err := cs.MessagesForBlock(blk)
		if err != nil {
			return nil, err
		}
		blkMsgs[idx] = &lens.BlockMessages{
			Block:        blk,
			BlsMessages:  msgs,
			SecpMessages: smsgs,
		}
	}

	return &lens.TipSetMessages{
		Executed: emsgs,
		Block:    blkMsgs,
	}, nil
}

func MethodAndParamsForMessage(m *types.Message, destCode cid.Cid) (string, string, error) {
	var params ipld.Node
	var method string
	var err error

	// fall back to generic cbor->json conversion.
	params, method, err = messages.ParseParams(m.Params, int64(m.Method), destCode)
	if method == "Unknown" {
		return "", "", xerrors.Errorf("unknown method for actor type %s: %d", destCode.String(), int64(m.Method))
	}
	if err != nil {
		log.Warnf("failed to parse parameters of message %s: %v", m.Cid(), err)
		// this can occur when the message is not valid cbor
		return method, "", xerrors.Errorf("failed to parse parameters of message %s: %w", m.Cid(), err)
	}
	if params == nil {
		return method, "", nil
	}

	buf := bytes.NewBuffer(nil)
	if err := fcjson.Encoder(params, buf); err != nil {
		return "", "", xerrors.Errorf("json encode message params: %w", err)
	}

	encoded := string(bytes.ReplaceAll(bytes.ToValidUTF8(buf.Bytes(), []byte{}), []byte{0x00}, []byte{}))

	return method, encoded, nil
}

func ActorNameAndFamilyFromCode(c cid.Cid) (name string, family string, err error) {
	if !c.Defined() {
		return "", "", xerrors.Errorf("cannot derive actor name from undefined CID")
	}
	name = builtin.ActorNameByCode(c)
	if name == "<unknown>" {
		return "", "", xerrors.Errorf("cannot derive actor name from unknown CID: %s (maybe we need up update deps?)", c.String())
	}
	tokens := strings.Split(name, "/")
	if len(tokens) != 3 {
		return "", "", xerrors.Errorf("cannot parse actor name: %s from tokens: %s", name, tokens)
	}
	// network = tokens[0]
	// version = tokens[1]
	family = tokens[2]
	return
}

func MakeGetActorCodeFunc(ctx context.Context, store adt.Store, next, current *types.TipSet) (func(a address.Address) (cid.Cid, bool), error) {
	nextStateTree, err := state.LoadStateTree(store, next.ParentState())
	if err != nil {
		return nil, xerrors.Errorf("load state tree: %w", err)
	}

	// Build a lookup of actor codes that exist after all messages in the current epoch have been executed
	actorCodes := map[address.Address]cid.Cid{}
	if err := nextStateTree.ForEach(func(a address.Address, act *types.Actor) error {
		actorCodes[a] = act.Code
		return nil
	}); err != nil {
		return nil, xerrors.Errorf("iterate actors: %w", err)
	}

	nextInitActor, err := nextStateTree.GetActor(builtininit.Address)
	if err != nil {
		return nil, xerrors.Errorf("getting init actor: %w", err)
	}

	nextInitActorState, err := builtininit.Load(store, nextInitActor)
	if err != nil {
		return nil, xerrors.Errorf("loading init actor state: %w", err)
	}

	return func(a address.Address) (cid.Cid, bool) {
		// Shortcut lookup before resolving
		c, ok := actorCodes[a]
		if ok {
			return c, true
		}

		ra, found, err := nextInitActorState.ResolveAddress(a)
		if err != nil || !found {
			log.Warnw("failed to resolve actor address", "address", a.String())
			return cid.Undef, false
		}

		c, ok = actorCodes[ra]
		if ok {
			return c, true
		}

		// Fall back to looking in current state tree. This actor may have been deleted.
		currentStateTree, err := state.LoadStateTree(store, current.ParentState())
		if err != nil {
			log.Warnf("failed to load state tree: %v", err)
			return cid.Undef, false
		}

		act, err := currentStateTree.GetActor(a)
		if err != nil {
			log.Warnw("failed to find actor in state tree", "address", a.String(), "error", err.Error())
			return cid.Undef, false
		}

		return act.Code, true
	}, nil
}
