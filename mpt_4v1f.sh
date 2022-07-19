#!/usr/bin/env bash
DATAPATH="$HOME/Downloads"

db_backend="goleveldb"

while getopts "i" opt; do
  case $opt in
  i)
    NAME="exchaind"
    MYNAME="mpt_4v1f.sh"
    ps -ef|grep "$NAME"|grep -v grep |grep -v $MYNAME |awk '{print "kill -9 "$2", "$8}'
    ps -ef|grep "$NAME"|grep -v grep |awk '{print "kill -9 "$2}' | sh
    echo "All <$NAME> killed!"

    echo "INIT"
    exchaind unsafe-reset-all --home ${DATAPATH}/cache/node0/exchaind
    exchaind unsafe-reset-all --home ${DATAPATH}/cache/node1/exchaind
    exchaind unsafe-reset-all --home ${DATAPATH}/cache/node2/exchaind
    exchaind unsafe-reset-all --home ${DATAPATH}/cache/node3/exchaind
    exchaind unsafe-reset-all --home ${DATAPATH}/cache/node4/exchaind
    ;;
  esac
done

sleep 2

# val0 with seed
nohup exchaind start  --chain-id exchain-67 --log_level "state:info,main:info,root-multi:info" --db_backend $db_backend --pruning=everything \
                --rpc.unsafe --disable-abci-query-mutex=true --consensus.timeout_commit 1s \
                --prof_laddr="localhost:6060" \
                --enable-batch-tx=true \
                --mempool.max_tx_num_per_block=3000 \
                --mempool.size=100000 \
                --p2p.private_peer_ids="3813c7011932b18f27f172f0de2347871d27e852" \
                --p2p.sentry_partner="3813c7011932b18f27f172f0de2347871d27e852" \
                --sentry-node=true \
                --p2p.seed_mode=false --p2p.allow_duplicate_ip  --p2p.pex=false --p2p.addr_book_strict=false --p2p.laddr tcp://127.0.0.1:20056 \
                --rest.laddr tcp://0.0.0.0:8545 --rpc.laddr tcp://0.0.0.0:20057 --home ${DATAPATH}/cache/node0/exchaind > ${DATAPATH}/cache/0.log &

# val1
nohup exchaind start  --chain-id exchain-67 --log_level "state:info,main:info,root-multi:info" --db_backend $db_backend --pruning=everything \
                --rpc.unsafe --disable-abci-query-mutex=true --consensus.timeout_commit 1s \
                --prof_laddr="localhost:6070" \
                --enable-batch-tx=true \
                --mempool.max_tx_num_per_block=3000 \
                --mempool.size=100000 \
                --p2p.sentry_partner="0b066ca0790f27a6595560b23bf1a1193f100797" \
                --p2p.seed_mode=false --p2p.allow_duplicate_ip --p2p.pex=false --p2p.addr_book_strict=false --p2p.laddr tcp://127.0.0.1:20156 \
                --rest.laddr tcp://0.0.0.0:20159 --rpc.laddr tcp://0.0.0.0:20157 --home ${DATAPATH}/cache/node1/exchaind > ${DATAPATH}/cache/1.log &

# val2
nohup exchaind start  --chain-id exchain-67 --log_level "state:info,main:info,root-multi:info" --db_backend $db_backend --pruning=nothing \
                --rpc.unsafe --disable-abci-query-mutex=true --consensus.timeout_commit 1s \
                --prof_laddr="localhost:6080" \
                --enable-batch-tx=true \
                --mempool.max_tx_num_per_block=3000 \
                --mempool.size=100000 \
                --p2p.private_peer_ids="bab6c32fa95f3a54ecb7d32869e32e85a25d2e08" \
                --p2p.sentry_partner="bab6c32fa95f3a54ecb7d32869e32e85a25d2e08" \
                --sentry-node=true \
                --p2p.seed_mode=false --p2p.allow_duplicate_ip --p2p.pex=false --p2p.addr_book_strict=false --p2p.laddr tcp://127.0.0.1:20256 \
                --rest.laddr tcp://0.0.0.0:20259 --rpc.laddr tcp://0.0.0.0:20257 --home ${DATAPATH}/cache/node2/exchaind > ${DATAPATH}/cache/2.log &

# val3
nohup exchaind start  --chain-id exchain-67 --log_level "state:info,main:info,root-multi:info" --db_backend $db_backend --pruning=everything \
                --rpc.unsafe --disable-abci-query-mutex=true --consensus.timeout_commit 1s \
                --prof_laddr="localhost:6090" \
                --enable-batch-tx=true \
                --mempool.max_tx_num_per_block=3000 \
                --mempool.size=100000 \
                --p2p.sentry_partner="6ea83a21a43c30a280a3139f6f23d737104b6975" \
                --p2p.seed_mode=false --p2p.allow_duplicate_ip --p2p.pex=false --p2p.addr_book_strict=false --p2p.laddr tcp://127.0.0.1:20356 \
                --rest.laddr tcp://0.0.0.0:20359 --rpc.laddr tcp://0.0.0.0:20357 --home ${DATAPATH}/cache/node3/exchaind > ${DATAPATH}/cache/3.log &

# full0
#nohup exchaind start  --chain-id exchain-67 --log_level "state:info,main:info,root-multi:info,mempool:info" --db_backend $db_backend --pruning=nothing \
#                --rpc.unsafe --disable-abci-query-mutex=true --consensus.timeout_commit 1s \
#                --enable-batch-tx=true \
#                --mempool.max_tx_num_per_block=3000 \
#                --mempool.size=100000 \
#                --p2p.seed_mode=false --p2p.allow_duplicate_ip --p2p.pex=false --p2p.addr_book_strict=false --p2p.laddr tcp://127.0.0.1:20456 --p2p.seeds 0b066ca0790f27a6595560b23bf1a1193f100797@127.0.0.1:20056 \
#                --rest.laddr tcp://localhost:8545 --rpc.laddr tcp://0.0.0.0:26657 --home ${DATAPATH}/cache/node4/exchaind > ${DATAPATH}/cache/4.log &

# --halt-height 150