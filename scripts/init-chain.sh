#!/bin/sh
rm -Rf ~/.medipoint
medipointd config keyring-backend test 
medipointd init $MONIKER --chain-id $CHAIN_ID --overwrite 
echo $VALIDATOR_MNEMONIC | medipointd keys add $VALIDATOR_ACCOUNT --hd-path "m/44'/118'/0'/0" --recover --interactive=false
medipointd add-genesis-account $VALIDATOR_ACCOUNT $AMOUNT
medipointd keys list
medipointd gentx $VALIDATOR_ACCOUNT 100000000stake --chain-id $CHAIN_ID
medipointd collect-gentxs
(sleep 10 && medipointd tx medipoint create-super-admin $SUPER_ADMIN_ADDRESS --from $VALIDATOR_ACCOUNT --chain-id $CHAIN_ID -y) | medipointd start --rpc.laddr tcp://127.0.0.1:26657 --rpc.pprof_laddr tcp://127.0.0.1:8002 --grpc.address 0.0.0.0:9090