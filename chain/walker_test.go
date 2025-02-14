package chain

import (
	"context"
	"testing"
	"time"

	itestkit "github.com/filecoin-project/lotus/itests/kit"
	"github.com/go-pg/pg/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/lily/chain/actors/builtin"
	"github.com/filecoin-project/lily/model/blocks"
	"github.com/filecoin-project/lily/storage"
	"github.com/filecoin-project/lily/testutil"
)

func TestWalker(t *testing.T) {
	if testing.Short() {
		t.Skip("short testing requested")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	db, cleanup, err := testutil.WaitForExclusiveDatabase(ctx, t)
	require.NoError(t, err)
	defer func() { require.NoError(t, cleanup()) }()

	t.Logf("truncating database tables")
	err = truncateBlockTables(t, db)
	require.NoError(t, err, "truncating tables")

	t.Logf("preparing chain")
	full, miner, _ := itestkit.EnsembleMinimal(t)

	nodeAPI := testutil.NewAPIWrapper(full)

	bm := itestkit.NewBlockMiner(t, miner)
	bm.MineUntilBlock(ctx, full, nil)

	head, err := full.ChainHead(ctx)
	require.NoError(t, err, "chain head")

	t.Logf("collecting chain blocks from tipset before head")

	beforeHead, err := full.ChainGetTipSet(ctx, head.Parents())
	require.NoError(t, err, "get tipset before head")

	bhs, err := collectBlockHeaders(nodeAPI, beforeHead)
	require.NoError(t, err, "collect chain blocks")

	cids := bhs.Cids()
	rounds := bhs.Rounds()

	strg, err := storage.NewDatabaseFromDB(ctx, db, "public")
	require.NoError(t, err, "NewDatabaseFromDB")

	tsIndexer, err := NewTipSetIndexer(nodeAPI, strg, builtin.EpochDurationSeconds*time.Second, t.Name(), []string{BlocksTask})
	require.NoError(t, err, "NewTipSetIndexer")
	t.Logf("initializing indexer")
	idx := NewWalker(tsIndexer, nodeAPI, 0, int64(beforeHead.Height()))

	t.Logf("indexing chain")
	err = idx.WalkChain(ctx, nodeAPI, head)
	require.NoError(t, err, "WalkChain")

	// TODO NewTipSetIndexer runs its processors in their own go routines (started when TipSet() is called)
	// this causes this test to behave nondeterministicly so we sleep here to ensure all async jobs
	// have completed before asserting results
	time.Sleep(time.Second * 3)

	t.Run("block_headers", func(t *testing.T) {
		var count int
		_, err := db.QueryOne(pg.Scan(&count), `SELECT COUNT(*) FROM block_headers`)
		require.NoError(t, err)
		assert.Equal(t, len(cids), count)

		var m *blocks.BlockHeader
		for _, cid := range cids {
			exists, err := db.Model(m).Where("cid = ?", cid).Exists()
			require.NoError(t, err)
			assert.True(t, exists, "cid: %s", cid)
		}
	})

	t.Run("block_parents", func(t *testing.T) {
		var count int
		_, err := db.QueryOne(pg.Scan(&count), `SELECT COUNT(*) FROM block_parents`)
		require.NoError(t, err)
		assert.Equal(t, len(cids), count)

		var m *blocks.BlockParent
		for _, cid := range cids {
			exists, err := db.Model(m).Where("block = ?", cid).Exists()
			require.NoError(t, err)
			assert.True(t, exists, "block: %s", cid)
		}
	})

	t.Run("drand_block_entries", func(t *testing.T) {
		var count int
		_, err := db.QueryOne(pg.Scan(&count), `SELECT COUNT(*) FROM drand_block_entries`)
		require.NoError(t, err)
		assert.Equal(t, len(rounds), count)

		var m *blocks.DrandBlockEntrie
		for _, round := range rounds {
			exists, err := db.Model(m).Where("round = ?", round).Exists()
			require.NoError(t, err)
			assert.True(t, exists, "round: %d", round)
		}
	})
}
