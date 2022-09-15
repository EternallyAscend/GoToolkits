package fabric

import "fmt"

func CreateChannelArtifactsFolderCommand() []string {
	return []string{
		fmt.Sprintf("mkdir %schannel-artifacts", getBaseFolderPath()),
	}
}

func GenerateSystemChannelBlock(fabricName string) []string {
	var cmds []string
	cmds = append(cmds, fmt.Sprintf("mkdir %ssystem-gensis-block/gensis.b lock"))
	cmds = append(cmds, fmt.Sprintf("configtxgen -profile %s -channelID system-channel -outputBlock %ssystem-gensis-block/gensis.block", fabricName, getBaseFolderPath()))
	return cmds
}

// After version 2.2

//func CreateChannelGenesisBlockCommand(fabricName string, channelName string) []string {
//	var cmds []string
//	cmds = append(cmds, fmt.Sprintf("FABRIC_CFG_PATH=%sconfigtx", getBaseFolderPath()))
//	cmds = append(cmds, fmt.Sprintf("configtxgen -profile %s -outputBlock %schannel-artifacts/%s.block -channelID %s", fabricName, getBaseFolderPath(), channelName, channelName))
//	return cmds
//}

func CreateChannelTx(fabricName string, channelName string) []string {
	var cmds []string
	cmds = append(cmds, fmt.Sprintf("FABRIC_CFG_PATH=%sconfigtx", getBaseFolderPath()))

	cmds = append(cmds, fmt.Sprintf("configtxgen -profile %s -outputCreateChannelTx .%schannel-artifacts/%s.tx -channelID %s"), fabricName, getBaseFolderPath(), channelName, channelName)
	return cmds
}

func CreateChannelCommand(channelName string, ordererDomain string) []string {
	var cmds []string
	// Here use configtx as well for lack config.
	cmds = append(cmds, fmt.Sprintf("FABRIC_CFG_PATH=%sconfigtx", getBaseFolderPath()))

	cmds = append(cmds, fmt.Sprintf("BLOCKFILE=\"%schannel-artifacts/%s.block\"", getBaseFolderPath(), channelName))
	cmds = append(cmds, fmt.Sprintf("peer channel create -o localhost:7050 -c %s --ordererTLSHostnameOverride orderer.%s -f .%schannel-artifacts/%s.tx --outputBlock $BLOCKFILE --tls --cafile %sorganizations/ordererOrganizations/%s/orderers/orderer.%s/msp/tlscacerts/tlsca.%s-cert.pem", channelName, ordererDomain, getBaseFolderPath(), channelName, getBaseFolderPath(), ordererDomain, ordererDomain, ordererDomain))
	return cmds

}

func SetGlobals(orderer bool, org string, corePeerAddress string, corePeerPort uint, domain string, peer string) []string {
	orgGroup := getOrgSubPathByOrderer(orderer)
	var cmds []string
	cmds = append(cmds, fmt.Sprintf("export CORE_PEER_LOCALMSPID=\"%sMSP\"", org))
	cmds = append(cmds, fmt.Sprintf("export CORE_PEER_TLS_ROOTCERT_FILE=%sorganizations/%s/%s/peers/%s.%s/tls/ca.crt", getBaseFolderPath(), orgGroup, domain, peer, domain))

	cmds = append(cmds, fmt.Sprintf("export CORE_PEER_MSPCONFIGPATH=%sorganizations/%s/%s/users/Admin@%s/msp", getBaseFolderPath(), orgGroup, domain, domain))

	cmds = append(cmds, fmt.Sprintf("export CORE_PEER_ADDRESS=%s:%s", corePeerAddress, corePeerPort))

	return cmds
}

func JoinChannelCommand(channelName string) []string {
	var cmds []string
	cmds = append(cmds, fmt.Sprintf("peer channel join -b %schannel-artifacts/%s.block", getBaseFolderPath(), channelName))
	return cmds
}

func SetGlobalCliCommand(corePeerAddress string, corePeerPort uint) []string {
	var cmds []string
	cmds = append(cmds, fmt.Sprintf("export CORE_PEER_ADDRESS=%s:%d", corePeerAddress, corePeerPort))
	return cmds
}

func SetAnchorPeer() []string {
	var cmds []string

	return cmds
}
