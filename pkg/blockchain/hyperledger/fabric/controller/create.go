package controller

import (
	"fmt"
	"log"

	"github.com/EternallyAscend/GoToolkits/pkg/blockchain/hyperledger/fabric/controller/config"
	fconfig "github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func loadClient(configPath string) (*fabsdk.FabricSDK, error) {
	sdk, err := fabsdk.New(fconfig.FromFile(configPath))
	if err != nil {
		log.Panicf("failed to create fabric sdk: %s", err)
	}
	return sdk, err
}

func GenerateMspID(orgName string) string {
	return fmt.Sprintf("%sMSP", orgName)
}

func GenerateClientTLSCertPath(isPeer bool, peerName, orgName, domainRoot string) string {
	if isPeer {
		return fmt.Sprintf("%sorganizations/peerOrganizations/%s.%s/peers/%s.%s.%s/tls/server.crt", config.FabricDataPath, orgName, domainRoot, peerName, orgName, domainRoot)
	} else {
		return fmt.Sprintf("%sorganizations/ordererOrganizations/%s.%s/orderers/%s.%s.%s/tls/server.crt", config.FabricDataPath, orgName, domainRoot, peerName, orgName, domainRoot)
	}
}

func GenerateServerTLSCertPath(isPeer bool, peerName, orgName, domainRoot string) string {
	if isPeer {
		return fmt.Sprintf("%sorganizations/peerOrganizations/%s.%s/peers/%s.%s.%s/tls/server.crt", config.FabricDataPath, orgName, domainRoot, peerName, orgName, domainRoot)
	} else {
		return fmt.Sprintf("%sorganizations/ordererOrganizations/%s.%s/orderers/%s.%s.%s/tls/server.crt", config.FabricDataPath, orgName, domainRoot, peerName, orgName, domainRoot)
	}
}

func GenerateTlsCertPath(isPeer bool, peerName, orgName, domainRoot string) string {
	if isPeer {
		return fmt.Sprintf("%sorganizations/peerOrganizations/%s.%s/peers/%s.%s.%s/tls", config.FabricDataPath, orgName, domainRoot, peerName, orgName, domainRoot)
	} else {
		return fmt.Sprintf("%sorganizations/ordererOrganizations/%s.%s/orderers/%s.%s.%s/tls", config.FabricDataPath, orgName, domainRoot, peerName, orgName, domainRoot)
	}
}

func GenerateMspPath(isPeer bool, peerName, orgName, domainRoot string) string {
	if isPeer {
		return fmt.Sprintf("%sorganizations/peerOrganizations/%s.%s/peers/%s.%s.%s/msp", config.FabricDataPath, orgName, domainRoot, peerName, orgName, domainRoot)
	} else {
		return fmt.Sprintf("%sorganizations/ordererOrganizations/%s.%s/orderers/%s.%s.%s/msp", config.FabricDataPath, orgName, domainRoot, peerName, orgName, domainRoot)
	}
}

func GenerateBlockPath() string {
	return fmt.Sprintf("%ssystem-genesis-block/genesis.block", config.FabricDataPath)
}

func GenerateDomain(peerName, orgName, domainRoot string) string {
	return fmt.Sprintf("%s.%s.%s", peerName, orgName, domainRoot)
}
