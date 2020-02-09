echo "Setting up the network.."

echo "Creating channel genesis block.."

# Create the channel
docker exec -e "CORE_PEER_LOCALMSPID=HospMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/hosp.dsct.com/users/Admin@hosp.dsct.com/msp" -e "CORE_PEER_ADDRESS=peer0.hosp.dsct.com:7051" cli peer channel create -o orderer.dsct.com:7050 -c dsctchannel -f /etc/hyperledger/configtx/dsctchannel.tx


sleep 5

echo "Channel genesis block created."

echo "peer0.hosp.dsct.com joining the channel..."
# Join peer0.hosp.dsct.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=HospMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/hosp.dsct.com/users/Admin@Hosp.dsct.com/msp" -e "CORE_PEER_ADDRESS=peer0.hosp.dsct.com:7051" cli peer channel join -b dsctchannel.block

echo "peer0.hosp.dsct.com joined the channel"

echo "peer0.manufacturer.dsct.com joining the channel..."

# Join peer0.manufacturer.dsct.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=ManufacturerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.dsct.com/users/Admin@manufacturer.dsct.com/msp" -e "CORE_PEER_ADDRESS=peer0.manufacturer.dsct.com:7051" cli peer channel join -b dsctchannel.block

echo "peer0.manufacturer.dsct.com joined the channel"

echo "peer0.delear.dsct.com joining the channel..."
# Join peer0.Delear.dsct.com to the channel.
docker exec -e "CORE_PEER_LOCALMSPID=DelearMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/delear.dsct.com/users/Admin@delear.dsct.com/msp" -e "CORE_PEER_ADDRESS=peer0.delear.dsct.com:7051" cli peer channel join -b dsctchannel.block
sleep 5

echo "peer0.delear.dsct.com joined the channel"

echo "Installing dsct chaincode to peer0.hosp.dsct.com..."

# install chaincode
# Install code on Hosp peer
docker exec -e "CORE_PEER_LOCALMSPID=HospMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/hosp.dsct.com/users/Admin@hosp.dsct.com/msp" -e "CORE_PEER_ADDRESS=peer0.hosp.dsct.com:7051" cli peer chaincode install -n dsctcc -v 1.0 -p github.com/dsct/go -l golang

echo "Installed dsct chaincode to peer0.hosp.dsct.com"

echo "Installing dsct chaincode to peer0.manufacturer.dsct.com...."

# Install code on Manufacturer peer
docker exec -e "CORE_PEER_LOCALMSPID=ManufacturerMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/manufacturer.dsct.com/users/Admin@manufacturer.dsct.com/msp" -e "CORE_PEER_ADDRESS=peer0.manufacturer.dsct.com:7051" cli peer chaincode install -n dsctcc -v 1.0 -p github.com/dsct/go -l golang

echo "Installed dsct chaincode to peer0.manufacturer.dsct.com"

echo "Installing dsct chaincode to peer0.delear.dsct.com..."
# Install code on Delear peer
docker exec -e "CORE_PEER_LOCALMSPID=DelearMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/delear.dsct.com/users/Admin@delear.dsct.com/msp" -e "CORE_PEER_ADDRESS=peer0.delear.dsct.com:7051" cli peer chaincode install -n dsctcc -v 1.0 -p github.com/dsct/go -l golang

sleep 5

echo "Installed dsct chaincode to peer0.Manufacturer.dsct.com"

echo "Instantiating dsct chaincode.."

docker exec -e "CORE_PEER_LOCALMSPID=HospMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/hosp.dsct.com/users/Admin@hosp.dsct.com/msp" -e "CORE_PEER_ADDRESS=peer0.hosp.dsct.com:7051" cli peer chaincode instantiate -o orderer.dsct.com:7050 -C dsctchannel -n dsctcc -l golang -v 1.0 -c '{"Args":[""]}' -P "OR ('HospMSP.member','ManufacturerMSP.member','DelearMSP.member')"

echo "Instantiated dsct chaincode."

echo "Following is the docker network....."

docker ps