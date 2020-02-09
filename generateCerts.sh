#rm -R crypto-config/*

./bin/cryptogen generate --config=crypto-config.yaml

rm config/*

./bin/configtxgen -profile dsctOrgOrdererGenesis -outputBlock ./config/genesis.block

./bin/configtxgen -profile dsctOrgChannel -outputCreateChannelTx ./config/dsctchannel.tx -channelID dsctchannel
