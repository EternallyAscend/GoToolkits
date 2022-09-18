package fabric

import (
	"fmt"

	"github.com/EternallyAscend/GoToolkits/pkg/blockchain/hyperledger/fabric/controller/config"
)

func GenerateGenesisBlockCommand(fabricName, channelID string) []string {
	var cmds []string
	cmds = append(cmds, fmt.Sprintf("configtxgen -profile %s -channelID %s -outputBlock %schannel-artifacts/genesis.block", fabricName, channelID, config.FabricDataPath))
	return cmds
}

func GenerateChannelTxCommand(fabricName, channelID string) []string {
	var cmds []string
	cmds = append(cmds, fmt.Sprintf("configtxgen -profile %s -outputCreateChannelTx %schannel-artifacts/channel.tx -channelID %s", fabricName, config.FabricDataPath, channelID))
	return cmds
}

func SetChannelAnchorPeerCommand(fabricName, anchorMSP, channelID string) []string {
	var cmds []string
	cmds = append(cmds, fmt.Sprintf("configtxgen -profile %s -outputAnchorPeersUpdate %schannel-artifacts/%sanchors.tx -channelID %s -asOrg %s", fabricName, config.FabricDataPath, anchorMSP, channelID, anchorMSP))
	return cmds
}

func CreateChannelInDockerCommand() []string {
	var cmds []string
	cmds = append(cmds, fmt.Sprintf("peer channel create -o orderer.example.com:7050 -c mychannel -f ./channel-artifacts/channel.tx --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem"))
	return cmds
}

func JoinChannelInDockerCommand(channelName string) []string {
	var cmds []string
	cmds = append(cmds, fmt.Sprintf("peer channel join -b %s.block", channelName))
	return cmds
}

func UpdateAnchorPeerInDockerCommand() []string {
	var cmds []string
	cmds = append(cmds, fmt.Sprintf("peer channel update -o orderer.example.com:7050 -c mychannel -f ./channel-artifacts/Org1MSPanchors.tx --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"))
	return cmds
}

func PackageChaincodeInDockerCommand() []string {
	var cmds []string
	cmds = append(cmds, fmt.Sprintf("peer lifecycle chaincode package sacc.tar.gz \\  --path github.com/hyperledger/fabric-cluster/chaincode/go/sacc/ \\  --label sacc_1"))
	return cmds
}

func InstallChaincodeInDockerCommand() []string {
	var cmds []string
	cmds = append(cmds, fmt.Sprintf("peer lifecycle chaincode install sacc.tar.gz"))
	cmds = append(cmds, fmt.Sprintf("peer lifecycle chaincode queryinstalled"))
	// Need result query last step.
	cmds = append(cmds, fmt.Sprintf("peer lifecycle chaincode approveformyorg --channelID mychannel --name sacc --version 1.0 --init-required --package-id sacc_1:1d9838e6893e068a94f055e807b18289559af748e5196a79a640b66305a74428 --sequence 1 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem"))
	cmds = append(cmds, "peer lifecycle chaincode commit -o orderer.example.com:7050 --channelID mychannel --name sacc --version 1.0 --sequence 1 --init-required --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses peer0.org2.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crtv")
	cmds = append(cmds, fmt.Sprintf("peer chaincode invoke -o orderer.example.com:7050 --isInit --ordererTLSHostnameOverride orderer.example.com --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n sacc --peerAddresses peer0.org1.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses peer0.org2.example.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{\"Args\":[\"a\",\"bb\"]}' "))
	return cmds
}
