package fabric

import "fmt"

func PackageChaincodeCommand(ccName, ccPath, ccLang, ccVersion string) []string {
	return []string{
		fmt.Sprintf("peer lifecycle chaincode package %s.tar.gz --path %s --lang %s --label %s_%s", ccName, ccPath, ccLang, ccName, ccVersion),
	}
}

func InstallChaincodeCommand(ccName string) []string {
	return []string{
		fmt.Sprintf("peer lifecycle chaincode install %s.tar.gz", ccName),
	}
}

func QueryInstalledChaincodeCommand() []string {
	return []string{
		// TODO Get Chaincode Package ID.
		fmt.Sprintf("peer lifecycle chaincode queryinstalled"),
	}
}

func ApproveChaincodeInOrgCommand(peerName, orgName, domainRoot string, port uint, ordererPeerName, ordererOrgName, ordererDomainRoot string, ordererPort uint, ordererCaPath string, channelId string, ccName, ccVersion, packageId, ccSequence, initRequired, ccEndPolicy, ccCollConfig string) []string {
	// 7050
	ordererDomain := fmt.Sprintf("%s.%s.%s", ordererPeerName, ordererOrgName, ordererDomainRoot)
	return []string{
		fmt.Sprintf("peer lifecycle chaincode approveformyorg -o %s:%d --ordererTLSHostnameOverride %s --tls --cafile %s --channelID %s --name %s --version %s --package-id %s --sequence %s %s %s %s", ordererDomain, ordererPort, ordererDomain, ordererCaPath, channelId, ccName, ccVersion, packageId, ccSequence, initRequired, ccEndPolicy, ccCollConfig),
	}
}

func CheckCommitReadinessCommand(channelId string, ccName, ccVersion, ccSequence, initRequired, ccEndPolicy, ccCollConfig string) []string {
	// TODO Run SetGlobal first.

	// TODO Run envVar.sh commitChaincodeDefinition first.
	return []string{
		fmt.Sprintf("peer lifecycle chaincode checkcommitreadiness --channelID %s --name %s --version %s --sequence %s %s %s %s --output json", channelId, ccName, ccVersion, ccSequence, initRequired, ccEndPolicy, ccCollConfig),
	}
}

func CommitChaincodeDefinitionCommand(ordererPort uint, ordererPeerName, ordererOrgName, ordererDomainRoot string, ordererCaPath string, channelId string, ccName, ccVersion, peerConnParms, ccSequence, initRequired, ccEndPolicy, ccCollConfig string) []string {
	// 7050
	ordererDomain := fmt.Sprintf("%s.%s.%s", ordererPeerName, ordererOrgName, ordererDomainRoot)
	return []string{
		fmt.Sprintf("peer lifecycle chaincode commit -o %s:%d --ordererTLSHostnameOverride %s --tls --cafile %s --channelID %s --name %s %s --version %s --sequence %s %s %s %s", ordererDomain, ordererPort, ordererDomain, ordererCaPath, channelId, ccName, peerConnParms, ccVersion, ccSequence, initRequired, ccEndPolicy, ccCollConfig),
	}
}

func QueryCommitCommand(channelName, ccName string) []string {
	return []string{
		fmt.Sprintf("peer lifecycle chaincode querycommitted --channelID %s --name %s", channelName, ccName),
	}
}

func ChaincodeInvokeInitCommand(ordererPort uint, ordererPeerName, ordererOrgName, ordererDomainRoot string, ordererCaPath string, channelId string, ccName, ccVersion, peerConnParms, ccSequence, initRequired, ccEndPolicy, ccCollConfig string, ccInitFunction, argsString string) []string {
	// 7050
	ordererDomain := fmt.Sprintf("%s.%s.%s", ordererPeerName, ordererOrgName, ordererDomainRoot)
	return []string{
		fmt.Sprintf(" peer chaincode invoke -o %s:%d --ordererTLSHostnameOverride %s --tls --cafile %s -C %s -n %s %s --isInit -c {\"function\":\"'%s'\",\"Args\":[%s]}", ordererDomain, ordererPort, ordererDomain, ordererCaPath, channelId, ccName, peerConnParms, ccInitFunction, argsString),
	}
}

func ChaincodeQueryCommand(channelName, ccName, argsString string) []string {
	return []string{
		fmt.Sprintf("peer chaincode query -C %s -n %s -c '{\"Args\":[\"%s\"]}'", channelName, ccName, argsString),
	}
}
