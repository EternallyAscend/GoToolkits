package fabric

import "fmt"

func CreateChannelArtifactsFolderCommand() []string {
	return []string{
		fmt.Sprintf("mkdir %schannel-artifacts", getBaseFolderPath()),
	}
}

func GenerateSystemChannelBlock(consortium string) []string {
	var cmds []string
	cmds = append(cmds, fmt.Sprintf("mkdir %ssystem-gensis-block/gensis.b lock", getBaseFolderPath()))
	cmds = append(cmds, fmt.Sprintf("configtxgen -profile %s -channelID system-channel -outputBlock %ssystem-gensis-block/gensis.block", consortium, getBaseFolderPath()))
	return cmds
}

// After version 2.2

//func CreateChannelGenesisBlockCommand(fabricName string, channelName string) []string {
//	var cmds []string
//	cmds = append(cmds, fmt.Sprintf("FABRIC_CFG_PATH=%sconfigtx", getBaseFolderPath()))
//	cmds = append(cmds, fmt.Sprintf("configtxgen -profile %s -outputBlock %schannel-artifacts/%s.block -channelID %s", fabricName, getBaseFolderPath(), channelName, channelName))
//	return cmds
//}

func CreateChannelTx(consortium string, channelName string) []string {
	var cmds []string
	cmds = append(cmds, fmt.Sprintf("FABRIC_CFG_PATH=%sconfigtx", getBaseFolderPath()))

	cmds = append(cmds, fmt.Sprintf("configtxgen -profile %s -outputCreateChannelTx .%schannel-artifacts/%s.tx -channelID %s", consortium, getBaseFolderPath(), channelName, channelName))
	return cmds
}

func CreateChannelCommand(ordererPeerName, ordererOrgName, ordererDomainRoot string, ordererPort uint, channelName string) []string {
	var cmds []string
	// Here use configtx as well for lack config.
	cmds = append(cmds, fmt.Sprintf("FABRIC_CFG_PATH=%sconfigtx", getBaseFolderPath()))
	ordererDomain := fmt.Sprintf("%s.%s.%s", ordererPeerName, ordererOrgName, ordererDomainRoot)
	blockFile := fmt.Sprintf("BLOCKFILE=\"%schannel-artifacts/%s.block\"", getBaseFolderPath(), channelName)
	cmds = append(cmds, fmt.Sprintf("BLOCKFILE=%s", blockFile))
	cmds = append(cmds, fmt.Sprintf("peer channel create -o %s:%d -c %s --ordererTLSHostnameOverride %s -f .%schannel-artifacts/%s.tx --outputBlock %s --tls --cafile %sorganizations/ordererOrganizations/%s/orderers/%s/msp/tlscacerts/tlsca.%s-cert.pem", ordererDomain, ordererPort, channelName, ordererDomain, getBaseFolderPath(), channelName, blockFile, getBaseFolderPath(), ordererDomain, ordererDomain, ordererDomainRoot))

	// TODO Check tlscacert file: "tlsca.orderer.examle.com" name.
	return cmds
}

func SetGlobals(orderer bool, org string, corePeerAddress string, corePeerPort uint, peerName string, orgName string, domainRoot string, peer string, userAdmin string, pwd string) []string {
	orgGroup := getOrgSubPathByOrderer(orderer)
	domain := fmt.Sprintf("%s.%s.%s", peerName, orgName, domainRoot)
	var cmds []string
	cmds = append(cmds, fmt.Sprintf("export CORE_PEER_LOCALMSPID=\"%sMSP\"", org))
	cmds = append(cmds, fmt.Sprintf("export CORE_PEER_TLS_ROOTCERT_FILE=%sorganizations/%s/%s/peers/%s.%s/tls/ca.crt", getBaseFolderPath(), orgGroup, domain, peer, domain))

	// TODO Check this command.
	cmds = append(cmds, fmt.Sprintf("export CORE_PEER_MSPCONFIGPATH=%sorganizations/%s/%s/users/%s@%s/msp", getBaseFolderPath(), orgGroup, userAdmin, pwd, domain))

	cmds = append(cmds, fmt.Sprintf("export CORE_PEER_ADDRESS=%s:%d", corePeerAddress, corePeerPort))

	return cmds
}

func JoinChannelCommand(channelName string) []string {
	var cmds []string
	// TODO config folder
	cmds = append(cmds, fmt.Sprintf("peer channel join -b %schannel-artifacts/%s.block", getBaseFolderPath(), channelName))
	return cmds
}

func SetGlobalCliCommand(corePeerAddress string, corePeerPort uint) []string {
	var cmds []string
	cmds = append(cmds, fmt.Sprintf("export CORE_PEER_ADDRESS=%s:%d", corePeerAddress, corePeerPort))
	return cmds
}

// SetAnchorPeer Using in hyperledger-tools container
func SetAnchorPeer(ordererPeerName, ordererOrgName, ordererDomainRoot string, ordererPort uint, channelName, ordererCaPath string, corePeerMspId string) []string {
	var cmds []string
	ordererDomain := fmt.Sprintf("%s.%s.%s", ordererPeerName, ordererOrgName, ordererDomainRoot)

	blockFile := fmt.Sprintf("BLOCKFILE=\"%schannel-artifacts/%s.block\"", getBaseFolderPath(), channelName)
	cmds = append(cmds, fmt.Sprintf("BLOCKFILE=%s", blockFile))

	cmds = append(cmds, fmt.Sprintf("peer channel create -o %s:%d -c %s --ordererTLSHostnameOverride %s -f %schannel-artifacts/%s.tx --outputBlock %s --tls --cafile %s", ordererDomain, ordererPort, channelName, ordererDomain, getBaseFolderPath(), channelName, blockFile, ordererCaPath))

	// TODO Combine OrdererCA Path.

	cmds = append(cmds, fmt.Sprintf("jq '.channel_group.groups.Application.groups.'%sMSP'.values += {\"AnchorPeers\":{\"mod_policy\": \"Admins\",\"value\":{\"anchor_peers\": [{\"host\": \"%s\",\"port\": %d}]},\"version\": \"0\"}}' %sMSPconfig.json > %sMSPmodified_config.json", corePeerMspId, ordererDomain, ordererPort, corePeerMspId, corePeerMspId))

	// TODO Set this command......
	cmds = append(cmds, fmt.Sprintf("  configtxlator proto_encode --input \"${ORIGINAL}\" --type common.Config >original_config.pb\n  configtxlator proto_encode --input \"${MODIFIED}\" --type common.Config >modified_config.pb\n  configtxlator compute_update --channel_id \"${CHANNEL}\" --original original_config.pb --updated modified_config.pb >config_update.pb\n  configtxlator proto_decode --input config_update.pb --type common.ConfigUpdate >config_update.json\n  echo '{\"payload\":{\"header\":{\"channel_header\":{\"channel_id\":\"'$CHANNEL'\", \"type\":2}},\"data\":{\"config_update\":'$(cat config_update.json)'}}}' | jq . >config_update_in_envelope.json\n  configtxlator proto_encode --input config_update_in_envelope.json --type common.Envelope >\"${OUTPUT}\""))

	cmds = append(cmds, fmt.Sprintf("peer channel update -o %s:%d --ordererTLSHostnameOverride %s -c %s -f %sMSPanchors.tx --tls --cafile %s", ordererDomain, ordererPort, ordererDomain, channelName, corePeerMspId, ordererCaPath))

	return cmds
}
