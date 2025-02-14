# Changelog
All notable changes to this project will be documented in this file.

The format is a variant of [Keep a Changelog](https://keepachangelog.com/en/1.0.0/) combined with categories from [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)

This project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html). Breaking changes should trigger an increment to the major version. Features increment the minor version and fixes or other changes increment the patch number.

<a name="v0.7.7"></a>
## [v0.7.7] - 2021-08-24

### Feat
- GapFind and GapFill task implementation (#598)

<a name="v0.7.6"></a>
## [v0.7.6] - 2021-08-24

### Feat
- implement consensus task (#615)
- revamp help command and add topic based help text (#546)
- implement verifreg actor tracking (#539)

### Fix
- report separate metric for actor state extraction (#633)
- replace current epoch estimation with new function (#631)
- add missing models to schema verification (#643)
- internal message bugs (#639)
- avoid missing block when miner loads receipt (#648)
- record correct height for messages and block_messages (#650)
- consistent processing report heights (#653)
- avoid reporting processing error for malformed messages (#671)
- avoid reading state of deleted actors (#672)

### Chore
- remove unused finality field (#632)
- track start and end times for jobs (#638)
- implement chain reading methods (#657)
- include window and storage in job params output (#630)

<a name="v0.7.5"></a>
## [v0.7.5] - 2021-08-09

### Feat
- Add option to omit writing csv headers (#617)
- Add basic templating of csv filenames (#622)

### Fix
- Only process message executions for tipsets on same chain (#618)
- Correctly handle error for internal message processing (#627)
- Use schema config for migrations (#628)

### Chore
- Log timings for message extraction (#626)
- Add newlines to walk and watch output, print command errors to stdout (#625)
- Use lowercase for id option for consistency (#623)
- Make from and to options required for walk command (#624)

<a name="v0.7.4"></a>
## [v0.7.4] - 2021-08-02

### Feat
 - Implement internal message task (#533)

### Fix
 - Export watch_height metric (#606)
 - Remove space in lint comment (#607) 

<a name="v0.7.3"></a>
## [v0.7.3] - 2021-08-02

### Feat
 - Add v1 schema migration capability (#578)
 - Add metrics for tipset cache used by watches (#599)
 - Upgrade to lotus@v1.11.0 (#601)

<a name="v0.7.2"></a>
## [v0.7.2] - 2021-07-02

### Fix
 - Skip actor and message tasks if tipset is not direct child of parent
 - Change failure metrics to counts and export tipset_skip metric 

### Feat
 - Add metrics for job lifecycle and model persistence

### Build
 - Cache vectors in CI

<a name="v0.7.1"></a>
## [v0.7.1] - 2021-06-29

### Fix
 - Avoid indexer blocking head events
 - Add batch persist support for v0 schema message tables
 - Increase API cache size 


<a name="v0.7.0"></a>
## [v0.7.0] - 2021-06-24

This version introduces a revised database schema, support for filecoin network version 13 and actors 5.

### Fix
 - Add usage for --api flag
 - Update amt package with diffing fix
 - Bare migrate command uses supplied schema name
 - Use actorname function generated for v5 actors
 - Better errors, equals method, correct cid
 - Give head notifier a buffer
 - Correctly handle the genesis state
 - log-level option should apply to all loggers
 - Add hamt map opts for v5 actors
 - Use correct log level regex
 - Ensure id_addresses contains height

### Feat
 - Introduce major schema versioning and version 1 schema
 - Add stop command
 - Add net commands for inspecting daemon connectivity
 - Decode message bitfields to json


### Other
 - Update lotus dependency to include v1.10 changes
 - Add support for v5 actors
 - Automated Calibnet Docker Builds
 - Removed unused visor processing models
 - Support building visor for devnets
 - Sort list of log levels
 - Replace golangci-lint by staticcheck
 - Support building and executing on M1-based Macs


<a name="v0.6.7"></a>
## [v0.6.7] - 2021-06-10

### Fix
 - Restart watcher on failure in daemon mode
 - Update statediff to include v4 miners
 - Parse plain value transfer messages
 - Fall back to using config in repo directory if unspecified
 - Expand the config path when starting daemon 

### Feat
 - Add logging command and sane default log levels

### Chore
 - Remove s3 lens
 - Remove dependency on statediff
 - Update to use latest version of Lotus

<a name="v0.6.6"></a>
## [v0.6.6] - 2021-05-28

### Fix
 - Track miner sector size

### Feat
 - Introduce major schema versioning
 - Upgrade HAMT and AMT diffing to more efficient implementation when able to
 - Implement optimized state tree diffing

### Chore
 - Add actor-shim code generation
 - Add more tracing calls to critial path


<a name="v0.6.5"></a>
## [v0.6.5] - 2021-05-20

### Change
 - fail parsed message task when an unknown method is encountered

### Fix
 - update statediff to v0.0.24 to properly support actor v4 message methods and parameters

<a name="v0.6.4"></a>
## [v0.6.4] - 2021-05-19

### Chore
 - Update Lotus dependency to latest master to include v1.9.0 changes

<a name="v0.6.3"></a>
## [v0.6.3] - 2021-05-18

### Fix
 - setup logging, metrics and tracing in daemon
 - import correct v4 actor package for msapprovals 
 - jsonrpc only supports returning 2 params 
 - message task tracks unexecuted messages and their blocks
 - only exit once all scheduled jobs complete 
 - allow scheduler to exit if scheduled jobs are complete 

### Feat
 - allow database urls to use environment variables in config
 - add wait-api command to wait for the visor api to come online
 - add init command

### Chore
 - shorten indexhead-confidence option name on watch command
 - fix typos in init command usage


<a name="v0.6.2"></a>
## [v0.6.2] - 2021-04-28

### Fix
 - improve reliability of reconnecting to lotus after disconnect

### Feat
 - Upgrade lotus dependency to v1.8.0
 - Import more lotus commands

### Perf
 - Remove redundant tipset lookups

### Chore
 - Download vectors before running tests in CI


<a name="v0.6.1"></a>
## [v0.6.1] - 2021-04-21

### Feat
 - Upgrade actor support to version 4

### Chore
 - Fix dev docker image
 - Add circleci check for docker builds


<a name="v0.6.0"></a>
## [v0.6.0] - 2021-04-20

### BREAKING CHANGE

This changes the cli interface to make `walk` and `watch` subcommands of a new `run` command. Command line options
that are specific to `walk` and `watch` must now be specified after the subcommand.

Where before a walk would be started like this:

```sh
$ visor --db=foo --lens=lotus walk --from=1000 --to=1001
```

it must now be started like this:

```sh
$ visor run walk --db=foo --lens=lotus --from=1000 --to=1001
```

And for `watch`, before:

```sh
$ visor --db=foo --lens=lotus watch --headindexer-confidence=100
```

after:

```sh
$ visor run watch --db=foo --lens=lotus --headindexer-confidence=100
```

The `migrate` command now also expects options to be specified after the command:

```sh
$ visor migrate --db=foo --latest
```


In addition: 

 - the `--api` option has been renamed to `--lens-lotus-api` since it is used by the lotus lens to specify the api that visor will connect to. This reduces confusion from the `daemon` command's `--api` option which specifies the api address of visor's daemon.
 - the `--repo` option has been renamed to `--lens-repo` since it is used by various lenses to specify the location of the data file or directory that visor will read from. This distinguishes it from the `daemon` command's `--repo` option which specifies the path where the visor daemon should write its data.



### Feat
 - Add a new long running daemon mode for visor
 - Add new daemon command
 - Add new job command
 - Add new run command and move watch and walk to be subcommands of run
 - Reorganise CLI options to be associated with relevant command or subcommand

<a name="v0.5.7"></a>
## [v0.5.7] - 2021-04-09

### Fix
 - Ensure persistence semaphore channel is drained on close

<a name="v0.5.6"></a>
## [v0.5.6] - 2021-04-07

### Chore
 - update to Lotus 1.6.0


<a name="v0.5.5"></a>
## [v0.5.5] - 2021-03-25

### Fix
 - only close TIpSetIndexer in walk and watch
 - embed genesis block in executable
 - ignore false positive gosec failure in wait package
 - close TipSetObs in walker and watcher

### Feat
 - add benchmarking of vectors
 - record tipset cache metrics during watch
 - Support current sqlotus schema
 - sqlotus dag prefetch option

### Chore
 - update for Lotus 1.5.3
 - increase linter timout in ci
 - remove lint github action
 - fix the linter



<a name="v0.5.4"></a>
## [v0.5.4] - 2021-03-09

### Fix
- guard concurrent accesses to node api ([#412](https://github.com/filecoin-project/sentinel-visor/issues/412))
- avoid deadlock in indexer when processor errors ([#407](https://github.com/filecoin-project/sentinel-visor/issues/407))

<a name="v0.5.3"></a>
## [v0.5.3] - 2021-03-02

### Feat
- support Lotus 1.5.0 and Actors 3.0.3 ([#373](https://github.com/filecoin-project/sentinel-visor/issues/373))

### Fix
- wait for persist routines to complete in Close ([#374](https://github.com/filecoin-project/sentinel-visor/issues/374))

<a name="v0.5.2"></a>
## [v0.5.2] - 2021-02-22

### Feat
- record multisig approvals ([#389](https://github.com/filecoin-project/sentinel-visor/issues/389))
- implement test vector builder and executer ([#370](https://github.com/filecoin-project/sentinel-visor/issues/370))

### Fix
- msapprovals missing pending transaction ([#395](https://github.com/filecoin-project/sentinel-visor/issues/395))
- correct docker image name; simplify build pipeline ([#393](https://github.com/filecoin-project/sentinel-visor/issues/393))
- set chainstore on repo lens util

### Chore
- add release process details and workflow ([#353](https://github.com/filecoin-project/sentinel-visor/issues/353))

<a name="v0.5.1"></a>
## [v0.5.1] - 2021-02-09

### Feat
- record actor task metrics ([#376](https://github.com/filecoin-project/sentinel-visor/issues/376))

### Chore
- increase lens object cache size ([#377](https://github.com/filecoin-project/sentinel-visor/issues/377))

### Schema
- remove use of add_drop_chunks_policy timescale function ([#379](https://github.com/filecoin-project/sentinel-visor/issues/379))


<a name="v0.5.0"></a>
## [v0.5.0] - 2021-02-09

No changes from v0.5.0-rc2


<a name="v0.5.0-rc2"></a>
## [v0.5.0-rc2] - 2021-01-27

Required schema version: `27`

### Notable for this release:
- update specs-actors to support v3 upgrade
- CSV exporting for easier ingestion into the DB of your choice
- bug fix for incorrect gas outputs (which changed after FIP-0009 was applied)
- inline schema documentation

### Feat
- remove default value for --db parameter ([#348](https://github.com/filecoin-project/sentinel-visor/issues/348))
- abstract model storage and add csv output for walk command ([#316](https://github.com/filecoin-project/sentinel-visor/issues/316))
- allow finer-grained actor task processing ([#305](https://github.com/filecoin-project/sentinel-visor/issues/305))
- record metrics for watch and walk commands ([#312](https://github.com/filecoin-project/sentinel-visor/issues/312))
- **db:** allow model upsertion
- **gas outputs:** Add Height and ActorName ([#270](https://github.com/filecoin-project/sentinel-visor/issues/270))
- **lens:** Optimize StateGetActor calls. ([#214](https://github.com/filecoin-project/sentinel-visor/issues/214))

### Fix
- persist market deal states ([#367](https://github.com/filecoin-project/sentinel-visor/issues/367))
- improve inferred json encoding for csv output ([#364](https://github.com/filecoin-project/sentinel-visor/issues/364))
- csv output handles time and interface values ([#351](https://github.com/filecoin-project/sentinel-visor/issues/351))
- adjust calculation of gas outputs for FIP-0009 ([#356](https://github.com/filecoin-project/sentinel-visor/issues/356))
- reject names that exceed maximum postgres name length ([#323](https://github.com/filecoin-project/sentinel-visor/issues/323))
- don't restart a walk if it fails ([#320](https://github.com/filecoin-project/sentinel-visor/issues/320))
- close all processor connections to lotus on fatal error ([#309](https://github.com/filecoin-project/sentinel-visor/issues/309))
- use migration database connection when installing timescale extension ([#304](https://github.com/filecoin-project/sentinel-visor/issues/304))
- **ci:** Pin TimescaleDB to v1.7 on Postgres v12 ([#340](https://github.com/filecoin-project/sentinel-visor/issues/340))
- **migration:** don't recreate miner_sector_deals primary key if it is correct ([#300](https://github.com/filecoin-project/sentinel-visor/issues/300))
- **migrations:** CREATE EXTENSION deadlocks inside migrations global lock ([#210](https://github.com/filecoin-project/sentinel-visor/issues/210))
- **miner:** extract miner PoSt's from parent messages

### Chore
- update imports and ffi stub for lotus 1.5.0-pre1 ([#371](https://github.com/filecoin-project/sentinel-visor/issues/371))
- fix some linting issues ([#349](https://github.com/filecoin-project/sentinel-visor/issues/349))
- **api:** trim the lens API to required methods
- **lint:** fix linter errors
- **lint:** fix staticcheck linting issues ([#299](https://github.com/filecoin-project/sentinel-visor/issues/299))
- **sql:** user numeric type to represent numbers ([#327](https://github.com/filecoin-project/sentinel-visor/issues/327))

### Perf
- replace local state diffing with StateChangeActors API method ([#303](https://github.com/filecoin-project/sentinel-visor/issues/303))

### Test
- **actorstate:** unit test actorstate actor task
- **chain:** refactor and test chain economics extraction ([#298](https://github.com/filecoin-project/sentinel-visor/issues/298))

### Docs
- table and column comments ([#346](https://github.com/filecoin-project/sentinel-visor/issues/346))
- Update README and docker-compose to require use of TimescaleDB v1.7 ([#341](https://github.com/filecoin-project/sentinel-visor/issues/341))
- document mapping between tasks and tables ([#369](https://github.com/filecoin-project/sentinel-visor/issues/369))

### Polish
- **test:** allow `make test` to "just work"

## [v0.5.0-rc1] - Re-released as [v0.5.0-rc2](#v0.5.0-rc2)

<a name="v0.4.0"></a>
## [v0.4.0] - 2020-12-16
### Chore
- remove test branch and temp deploy config

### Fix
- Make visor the entrypoint for dev containers


<a name="v0.4.0-rc2"></a>
## [v0.4.0-rc2] - 2020-12-16
### Feat
- **ci:** Dockerfile.dev; Refactor docker push steps in circleci.yaml


<a name="v0.4.0-rc1"></a>
## [v0.4.0-rc1] - 2020-12-02
### DEPRECATION

The CLI interface has shifted again to deprecate the `run` subcommand in favor of dedicated subcommands for `indexer` and `processor` behaviors.

Previously the indexer and procerror would be started via:

```sh
  sentinel-visor run --indexhead
  sentinel-visor run --indexhistory
```

After this change:

```sh
  sentinel-visor watch
  sentinel-visor walk
```

The `run` subcommand will be removed in v0.5.0.

### Feat
- extract basic account actor states ([#278](https://github.com/filecoin-project/sentinel-visor/issues/278))
- add watch and walk commands to index chain during traversal ([#249](https://github.com/filecoin-project/sentinel-visor/issues/249))
- functions to convert unix epoch to fil epoch ([#252](https://github.com/filecoin-project/sentinel-visor/issues/252))
- add repo-read-only flag to enable read or write on lotus repo ([#250](https://github.com/filecoin-project/sentinel-visor/issues/250))
- allow application name to be passed in postgres connection url ([#243](https://github.com/filecoin-project/sentinel-visor/issues/243))
- limit history indexer by height ([#234](https://github.com/filecoin-project/sentinel-visor/issues/234))
- extract msig transaction hamt

### Fix
- optimisable height functions ([#268](https://github.com/filecoin-project/sentinel-visor/issues/268))
- don't update go modules when running make
- gracefully disconnect from postgres on exit
- truncated tables in tests ([#277](https://github.com/filecoin-project/sentinel-visor/issues/277))
- tests defer database cleanup without invoking ([#274](https://github.com/filecoin-project/sentinel-visor/issues/274))
- totalGasLimit and totalUniqueGasLimit are correct
- missed while closing [#201](https://github.com/filecoin-project/sentinel-visor/issues/201)
- include height with chain power results ([#255](https://github.com/filecoin-project/sentinel-visor/issues/255))
- avoid panic when miner has no peer id ([#254](https://github.com/filecoin-project/sentinel-visor/issues/254))
- Remove hack to RestartOnFailure
- Reorder migrations after merging latest master ([#248](https://github.com/filecoin-project/sentinel-visor/issues/248))
- multisig actor migration
- lotus chain store is a blockstore
- panic in multisig genesis task casting
- **actorstate:** adjust account extractor to conform to new interface ([#294](https://github.com/filecoin-project/sentinel-visor/issues/294))
- **init:** extract idAddress instead of actorID
- **schema:** fix primary key for miner_sector_deals table ([#291](https://github.com/filecoin-project/sentinel-visor/issues/291))

### Refactor
- **cmd:** Modify command line default parameters ([#271](https://github.com/filecoin-project/sentinel-visor/issues/271))

### Test
- add multisig actor extractor tests
- power actor claim extration test
- **init:** test coverage for init actor extractor

### Chore
- Avoid ingesting binary and unused data ([#241](https://github.com/filecoin-project/sentinel-visor/issues/241))
- remove unused tables and views

### CI
- **test:** add code coverage
- **test:** run full testing suite

### Build
- **ci:** add go mod tidy check ([#266](https://github.com/filecoin-project/sentinel-visor/issues/266))

### Docs
- expand getting started guide and add running tests section ([#275](https://github.com/filecoin-project/sentinel-visor/issues/275))

### Polish
- Avoid duplicate work when reading receipts
- use new init actor diffing logic
- **mockapi:** names reflect method action
- **mockapi:** remove returned errors and condense mockTipset
- **mockapi:** accepts testing.TB, no errors

<a name="v0.3.0"></a>
## [v0.3.0] - 2020-11-03
### Feat
- add visor processing stats table ([#96](https://github.com/filecoin-project/sentinel-visor/issues/96))
- allow actor state processor to run without leasing ([#178](https://github.com/filecoin-project/sentinel-visor/issues/178))
- rpc reconnection on failure ([#149](https://github.com/filecoin-project/sentinel-visor/issues/149))
- add dynamic panel creation based on tags ([#159](https://github.com/filecoin-project/sentinel-visor/issues/159))
- add dynamic panel creation based on tags
- make delay between tasks configurable ([#151](https://github.com/filecoin-project/sentinel-visor/issues/151))
- convert processing, block and message tables to hypertables ([#111](https://github.com/filecoin-project/sentinel-visor/issues/111))
- set default numbers of workers to zero in run subcommand ([#116](https://github.com/filecoin-project/sentinel-visor/issues/116))
- add dashboard for process completion
- add changelog generator
- log visor version on startup ([#117](https://github.com/filecoin-project/sentinel-visor/issues/117))
- Add heaviest chain materialized view ([#97](https://github.com/filecoin-project/sentinel-visor/issues/97))
- Add miner_sector_posts tracking of window posts ([#74](https://github.com/filecoin-project/sentinel-visor/issues/74))
- Add historical indexer metrics ([#92](https://github.com/filecoin-project/sentinel-visor/issues/92))
- add message gas economy processing
- set application name in postgres connection ([#104](https://github.com/filecoin-project/sentinel-visor/issues/104))
- **miner:** compute miner sector events
- **task:** add chain economics processing ([#94](https://github.com/filecoin-project/sentinel-visor/issues/94))

### Fix
- Make ChainVis into basic views
- failure to get lock when ExitOnFailure is true now exits
- use hash index type for visor_processing_actors_code_idx ([#106](https://github.com/filecoin-project/sentinel-visor/issues/106))
- fix actor completion query
- visor_processing_stats queries for Visor processing dash ([#156](https://github.com/filecoin-project/sentinel-visor/issues/156))
- remove errgrp from UnindexedBlockData persist
- migration table name
- correct typo in derived_consensus_chain_view name and add to view refresh ([#112](https://github.com/filecoin-project/sentinel-visor/issues/112))
- avoid panic when miner extractor does not find receipt ([#110](https://github.com/filecoin-project/sentinel-visor/issues/110))
- verify there are no missing migrations before migrating ([#89](https://github.com/filecoin-project/sentinel-visor/issues/89))
- **lens:** Include dependencies needed for Repo Lens ([#90](https://github.com/filecoin-project/sentinel-visor/issues/90))
- **metrics:** export the completion and batch selection views ([#197](https://github.com/filecoin-project/sentinel-visor/issues/197))
- **migration:** message gas economy uses bigint
- **migrations:** migrations require version 0
- **schema:** remove blocking processing indexes and improve processing stats table ([#130](https://github.com/filecoin-project/sentinel-visor/issues/130))

### Build
- add prometheus, grafana and dashboard images

### Chore
- Incl migration in CI test
- Include RC releases in push docker images ([#195](https://github.com/filecoin-project/sentinel-visor/issues/195))
- add metrics to leasing and work completion queries
- add changelog ([#150](https://github.com/filecoin-project/sentinel-visor/issues/150))
- update go.mod after recent merge ([#155](https://github.com/filecoin-project/sentinel-visor/issues/155))
- add issue templates
- add more error context reporting in messages task ([#133](https://github.com/filecoin-project/sentinel-visor/issues/133))

### Deps
- remove unused docker file for redis

### Perf
- ensure processing updates always include height in criteria ([#192](https://github.com/filecoin-project/sentinel-visor/issues/192))
- include height restrictions in update clauses of leasing queries ([#189](https://github.com/filecoin-project/sentinel-visor/issues/189))
- **db:** reduce batch size for chain history indexer ([#105](https://github.com/filecoin-project/sentinel-visor/issues/105))

### Polish
- update miner processing logic

### Test
- ensure docker-compose down runs on test fail


<a name="v0.2.0"></a>
## [v0.2.0] - 2020-10-11
### BREAKING CHANGE

this changes the cli interface to remove the run subcommand.

Previously the indexer and procerror would be started via:

```sh
  sentinel-visor run indexer
  sentinel-visor run processor
```

After this change:

```sh
  sentinel-visor index
  sentinel-visor process
```

### Feat
- add standard build targets ([#18](https://github.com/filecoin-project/sentinel-visor/issues/18))
- add licenses and skeleton readme ([#5](https://github.com/filecoin-project/sentinel-visor/issues/5))
- instrument with tracing ([#15](https://github.com/filecoin-project/sentinel-visor/issues/15))
- add a configurable delay between task restarts ([#71](https://github.com/filecoin-project/sentinel-visor/issues/71))
- compute gas outputs ([#67](https://github.com/filecoin-project/sentinel-visor/issues/67))
- add tests for indexer ([#12](https://github.com/filecoin-project/sentinel-visor/issues/12))
- add schema migration capability ([#40](https://github.com/filecoin-project/sentinel-visor/issues/40))
- add LILY_TEST_DB environment variable to specify test database ([#35](https://github.com/filecoin-project/sentinel-visor/issues/35))
- respect log level flag and allow per logger levels ([#34](https://github.com/filecoin-project/sentinel-visor/issues/34))
- remove run subcommand and make index and process top level
- embed version number from build
- support v2 actor codes ([#84](https://github.com/filecoin-project/sentinel-visor/issues/84))
- add test for create schema ([#3](https://github.com/filecoin-project/sentinel-visor/issues/3))
- **api:** wrap lotus api and store with wrapper
- **debug:** Process actor by head without persistance ([#86](https://github.com/filecoin-project/sentinel-visor/issues/86))
- **genesis:** add task for processing genesis state
- **scheduler:** Refactor task scheduler impl ([#41](https://github.com/filecoin-project/sentinel-visor/issues/41))
- **task:** add actor, actor-state, and init actor processing ([#14](https://github.com/filecoin-project/sentinel-visor/issues/14))
- **task:** implement message processing task
- **task:** add market actor task
- **task:** add reward actor processing ([#16](https://github.com/filecoin-project/sentinel-visor/issues/16))
- **task:** add power actor processing task ([#11](https://github.com/filecoin-project/sentinel-visor/issues/11))
- **task:** Create chainvis views and refresher ([#77](https://github.com/filecoin-project/sentinel-visor/issues/77))

### Fix
- use debugf logging method in message processor ([#82](https://github.com/filecoin-project/sentinel-visor/issues/82))
- chain history indexer includes genesis ([#72](https://github.com/filecoin-project/sentinel-visor/issues/72))
- use context deadlines only if task has been assigned work ([#70](https://github.com/filecoin-project/sentinel-visor/issues/70))
- fix failing chain head indexer tests ([#66](https://github.com/filecoin-project/sentinel-visor/issues/66))
- add migration to remove old chainwatch schema constraints ([#48](https://github.com/filecoin-project/sentinel-visor/issues/48))
- use noop tracer when tracing disabled ([#39](https://github.com/filecoin-project/sentinel-visor/issues/39))
- ensure processor stops scheduler when exiting ([#24](https://github.com/filecoin-project/sentinel-visor/issues/24))
- **build:** ensure deps are built befor visor
- **indexer:** don't error on empty blocks_synced table
- **model:** replace BeginContext with RunInTransaction ([#7](https://github.com/filecoin-project/sentinel-visor/issues/7))
- **task:** correct index when computing deal state
### Chore
- add tests for reward and power actor state extracters ([#83](https://github.com/filecoin-project/sentinel-visor/issues/83))
- fail database tests if LILY_TEST_DB not set ([#79](https://github.com/filecoin-project/sentinel-visor/issues/79))
- use clock package for time mocking ([#65](https://github.com/filecoin-project/sentinel-visor/issues/65))
- remove unused redis-based scheduler code ([#64](https://github.com/filecoin-project/sentinel-visor/issues/64))
- Push docker images on [a-z]*-master branch updates ([#49](https://github.com/filecoin-project/sentinel-visor/issues/49))
- Remove sentinel prefix for local dev use ([#36](https://github.com/filecoin-project/sentinel-visor/issues/36))
- push docker tags from ci ([#26](https://github.com/filecoin-project/sentinel-visor/issues/26))
- tighten up error propagation ([#23](https://github.com/filecoin-project/sentinel-visor/issues/23))
- fix docker hub submodule error ([#22](https://github.com/filecoin-project/sentinel-visor/issues/22))
- add circle ci ([#20](https://github.com/filecoin-project/sentinel-visor/issues/20))
- add docker build and make targets ([#19](https://github.com/filecoin-project/sentinel-visor/issues/19))

### Dep
- add fil-blst submodule

### Perf
- minor optimization of market actor diffing ([#78](https://github.com/filecoin-project/sentinel-visor/issues/78))
- use batched inserts for models ([#73](https://github.com/filecoin-project/sentinel-visor/issues/73))

### Pg
- configurable pool size

### Polish
- **processor:** parallelize actor change collection
- **publisher:** receive publish operations on channel
- **redis:** configure redis with env vars ([#21](https://github.com/filecoin-project/sentinel-visor/issues/21))

### Refactor
- prepare for specs-actors upgrade
- replace panic with error return in indexer.Start ([#4](https://github.com/filecoin-project/sentinel-visor/issues/4))

### Test
- **storage:** add test to check for duplicate schema migrations ([#80](https://github.com/filecoin-project/sentinel-visor/issues/80))

[v0.5.4]: https://github.com/filecoin-project/sentinel-visor/compare/v0.5.3...v0.5.4
[v0.5.3]: https://github.com/filecoin-project/sentinel-visor/compare/v0.5.2...v0.5.3
[v0.5.2]: https://github.com/filecoin-project/sentinel-visor/compare/v0.5.1...v0.5.2
[v0.5.1]: https://github.com/filecoin-project/sentinel-visor/compare/v0.5.0...v0.5.1
[v0.5.0]: https://github.com/filecoin-project/sentinel-visor/compare/v0.5.0-rc2...v0.5.0
[v0.5.0-rc2]: https://github.com/filecoin-project/sentinel-visor/compare/v0.4.0...v0.5.0-rc1
[v0.4.0]: https://github.com/filecoin-project/sentinel-visor/compare/v0.4.0-rc2...v0.4.0
[v0.4.0-rc2]: https://github.com/filecoin-project/sentinel-visor/compare/v0.4.0-rc1...v0.4.0-rc2
[v0.4.0-rc1]: https://github.com/filecoin-project/sentinel-visor/compare/v0.3.0...v0.4.0-rc1
[v0.3.0]: https://github.com/filecoin-project/sentinel-visor/compare/v0.2.0...v0.3.0
[v0.2.0]: https://github.com/filecoin-project/sentinel-visor/compare/b7044af...v0.2.0
