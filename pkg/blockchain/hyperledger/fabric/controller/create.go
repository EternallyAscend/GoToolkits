package controller

import (
	"fmt"
	"github.com/EternallyAscend/GoToolkits/pkg/blockchain/hyperledger/fabric/controller/config"
	fconfig "github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"log"
)

func loadClient(configPath string) (*fabsdk.FabricSDK, error) {
	sdk, err := fabsdk.New(fconfig.FromFile(configPath))
	if err != nil {
		log.Panicf("failed to create fabric sdk: %s", err)
	}
	return sdk, err
}

func GenerateMSPID(orgName string) string {
	return fmt.Sprintf("%sMSP", orgName)
}
func GenerateClientTLSCertPath(isPeer bool, peerName, orgName, domainRoot string) string {
	if isPeer {
		return fmt.Sprintf("%s/organizations/peerOrganizations/%s.%s/peers/%s.%s.%s/tls/server.crt", config.FabricDataPath, orgName, domainRoot, peerName, orgName, domainRoot)
	} else {
		return fmt.Sprintf("%s/organizations/ordererOrganizations/%s.%s/orderers/%s.%s.%s/tls/server.crt", config.FabricDataPath, orgName, domainRoot, peerName, orgName, domainRoot)
	}

}
func GenerateServerTLSCertPath(isPeer bool, peerName, orgName, domainRoot string) string {
	if isPeer {
		return fmt.Sprintf("%s/organizations/peerOrganizations/%s.%s/peers/%s.%s.%s/tls/server.crt", config.FabricDataPath, orgName, domainRoot, peerName, orgName, domainRoot)
	} else {
		return fmt.Sprintf("%s/organizations/ordererOrganizations/%s.%s/orderers/%s.%s.%s/tls/server.crt", config.FabricDataPath, orgName, domainRoot, peerName, orgName, domainRoot)
	}
}

func GenerateDomain(peerName, orgName, domainRoot string) string {
	return fmt.Sprintf("%s.%s.%s", peerName, orgName, domainRoot)
}
