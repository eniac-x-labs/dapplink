#!/bin/bash

KEY="mykey"
CHAINID="dapplink_9000-1"
MONIKER="mymoniker"
DATA_DIR=$(mktemp -d -t dapplink-datadir.XXXXX)

echo "create and add new keys"
./dapplinkd keys add $KEY --home $DATA_DIR --no-backup --chain-id $CHAINID --algo "eth_secp256k1" --keyring-backend test
echo "init DappLink with moniker=$MONIKER and chain-id=$CHAINID"
./dapplinkd init $MONIKER --chain-id $CHAINID --home $DATA_DIR
echo "prepare genesis: Allocate genesis accounts"
./dapplinkd add-genesis-account \
"$(./dapplinkd keys show $KEY -a --home $DATA_DIR --keyring-backend test)" 1000000000000000000abit,1000000000000000000stake \
--home $DATA_DIR --keyring-backend test
echo "prepare genesis: Sign genesis transaction"
./dapplinkd gentx $KEY 1000000000000000000stake --keyring-backend test --home $DATA_DIR --keyring-backend test --chain-id $CHAINID
echo "prepare genesis: Collect genesis tx"
./dapplinkd collect-gentxs --home $DATA_DIR
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./dapplinkd validate-genesis --home $DATA_DIR

echo "starting dapplink node $i in background ..."
./dapplinkd start --pruning=nothing --rpc.unsafe \
--keyring-backend test --home $DATA_DIR \
>$DATA_DIR/node.log 2>&1 & disown

echo "started dapplink node"
tail -f /dev/null