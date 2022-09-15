package fabric

import (
	"fmt"
	"github.com/EternallyAscend/GoToolkits/pkg/blockchain/hyperledger/fabric/controller/config"
)

// Create, Update, Read, Delete, CURD.

func getBaseFolderPath() string {
	return config.FabricDataPath
}

func getOrgSubPathByOrderer(orderer bool) string {
	if orderer {
		return "ordererOrganizations"
	} else {
		return "peerOrganizations"
	}
}

func CheckCryptogenCommand() []string {
	// TODO Check cryptogen.
	return []string{}
}

func CryptoGenerateCommand(configPath string, orderer bool, domain string, user string, pwd string, port uint, caName string, caOrg string) []string {
	// TODO Remove peer/ordererOrganization Folder if exists.
	orgGroup := getOrgSubPathByOrderer(orderer)
	var cmds []string
	// Cryptogen
	cmds = append(cmds, fmt.Sprintf("cryptogen generate --config=%s --output=\"%sorganizations\"", configPath, getBaseFolderPath()))

	// Create Org Folder.
	cmds = append(cmds, fmt.Sprintf("mkdir %sorganizations/%s/%s/", getBaseFolderPath(), orgGroup, domain))
	// Env
	cmds = append(cmds, fmt.Sprintf("export FABRIC_CA_CLIENT_HOME=%sorganizations/%s/%s/", getBaseFolderPath(), orgGroup, domain))
	// CA cert
	cmds = append(cmds, fmt.Sprintf("fabric-ca-client enroll -u https://%s:%s@%s:%d --caname %s --tls.certfiles \"%sorganizations/fabric-ca/%s/ca-cert.pem\"", user, pwd, domain, port, caName, getBaseFolderPath(), caOrg)) // 7054 9054

	return cmds
}

func GenerateMspConfigYamlCommand(orderer bool, domain string, port uint, caOrg string) []string {
	orgGroup := getOrgSubPathByOrderer(orderer)
	var cmds []string
	certFile := fmt.Sprintf("cacerts/%s-%d-ca-%s.pem", domain, port, caOrg)
	// Maybe need to make msp dir first.
	cmds = append(cmds, fmt.Sprintf("echo 'NodeOUs:\n  Enable: true\n  ClientOUIdentifier:\n    Certificate: %s\n    OrganizationalUnitIdentifier: client\n  PeerOUIdentifier:\n    Certificate: %s\n    OrganizationalUnitIdentifier: peer\n  AdminOUIdentifier:\n    Certificate: %s\n    OrganizationalUnitIdentifier: admin\n  OrdererOUIdentifier:\n    Certificate: %s\n    OrganizationalUnitIdentifier: orderer' > \"%sorganizations/%s/%s/msp/config.yaml\"", certFile, certFile, certFile, certFile, getBaseFolderPath(), orgGroup, domain))

	return cmds
}

func GenerateCopyCertFileCommand(orderer bool, domain string, caOrg string) []string {

	orgGroup := getOrgSubPathByOrderer(orderer)
	var cmds []string

	cmds = append(cmds, fmt.Sprintf("  mkdir -p \"%sorganizations/%s/%s/msp/tlscacerts\"\n && cp \"%sorganizations/fabric-ca/%s/ca-cert.pem\" \"%sorganizations/%s/%s/msp/tlscacerts/ca.crt\"", getBaseFolderPath(), orgGroup, domain, getBaseFolderPath(), caOrg, getBaseFolderPath(), orgGroup, domain))

	cmds = append(cmds, fmt.Sprintf("  mkdir -p \"%sorganizations/%s/%s/tlsca\"\n && cp \"%sorganizations/fabric-ca/%s/ca-cert.pem\" \"%sorganizations/%s/%s/tlsca/tlsca.%s-cert.pem\"", getBaseFolderPath(), orgGroup, domain, getBaseFolderPath(), caOrg, getBaseFolderPath(), orgGroup, domain, domain))

	cmds = append(cmds, fmt.Sprintf("  mkdir -p \"%sorganizations/%s/%s/ca\"\n && cp \"%sorganizations/fabric-ca/%s/ca-cert.pem\" \"%sorganizations/%s/%s/ca/ca.%s-cert.pem\"", getBaseFolderPath(), orgGroup, domain, getBaseFolderPath(), caOrg, getBaseFolderPath(), orgGroup, domain, domain))

	return cmds
}

func CreateOrganizationCommands(configPath string, domain string, orderer bool, adminPw string, port uint, caName string, caOrg string) []string {

	// Generate CCP files
	return nil
}

func RegisterUserViaCaCommand(caName, userName, userPwd, caOrg string) []string {
	var cmds []string

	// Register
	cmds = append(cmds, fmt.Sprintf("fabric-ca-client register --caname %s --id.name %s --id.secret %s --id.type %s --tls.certfiles \"%s/organizations.fabric-ca/%s/ca-cert.pem\"", caName, userName, userPwd, getBaseFolderPath(), caOrg))
	return cmds
}

func RegisterPeerViaCaCommand(caName, peerName, pwd, caOrg string) []string {
	var cmds []string

	// Register
	cmds = append(cmds, fmt.Sprintf("fabric-ca-client register --caname %s --id.name %s --id.secret %s --id.type %s --tls.certfiles \"%s/organizations.fabric-ca/%s/ca-cert.pem\"", caName, peerName, pwd, getBaseFolderPath(), caOrg))
	return cmds
}

func GeneratePeerMspViaCaCommand(orderer bool, peer string, pwd string, domain string, port uint, caName string, caOrg string) []string {
	orgGroup := getOrgSubPathByOrderer(orderer)
	var cmds []string
	cmds = append(cmds, fmt.Sprintf("fabric-ca-client enroll -u https://%s:%s@%s:%d caname ca-%s -M \"%sorganizations/%s/%s/peers/%s.%s/msp\" --csr.hosts %s.%s --tls.certfiles \"%sorganizations/fabric-ca/%s/ca-cert.pem\"", peer, pwd, domain, port, caName, getBaseFolderPath(), orgGroup, domain, peer, domain, peer, domain, getBaseFolderPath(), caOrg))
	// Copy
	cmds = append(cmds, fmt.Sprintf(" cp \"%sorganizations/%s/%s/msp/config.yaml\" \"%sorganizations/%s/%s/peers/%s.%s/msp/config.yaml\"\n", getBaseFolderPath(), orgGroup, domain, getBaseFolderPath(), orgGroup, domain, peer, domain))
	return cmds
} // 7054

func GeneratePeerTlsCertCommand(orderer bool, peer string, pwd string, domain string, port uint, caName string, caOrg string) []string {
	orgGroup := getOrgSubPathByOrderer(orderer)
	var cmds []string
	// Generate
	cmds = append(cmds, fmt.Sprintf("fabric-ca-client enroll -u https://%s:%s@%s:%d --caname ca-%s -M \"%sorganizations/%s/%s/peers/%s.%s/tls\" --enrollment.profile tls --csr.hosts %s.%s --csr.hosts localhost --tls.certfiles \"%sorganizations/fabric-ca/%s/ca-cert.pem\"", peer, pwd, domain, port, caName, getBaseFolderPath(), orgGroup, domain, peer, domain, peer, domain, getBaseFolderPath(), caOrg))

	// Copy
	cmds = append(cmds, fmt.Sprintf("cp \"%sorganizations/%s/%s/peers/%s.%s/tls/tlscacerts/\"* \"%sorganizations/%s/%s/peers/%s.%s/tls/ca.crt\"", getBaseFolderPath(), orgGroup, domain, peer, domain, getBaseFolderPath(), orgGroup, domain, peer, domain))

	cmds = append(cmds, fmt.Sprintf("cp \"%sorganizations/%s/%s/peers/%s.%s/tls/signcerts/\"* \"%sorganizations/%s/%s/peers/%s.%s/tls/server.crt\"", getBaseFolderPath(), orgGroup, domain, peer, domain, getBaseFolderPath(), orgGroup, domain, peer, domain))

	cmds = append(cmds, fmt.Sprintf("cp \"%sorganizations/%s/%s/peers/%s.%s/tls/keystore/\"* \"%sorganizations/%s/%s/peers/%s.%s/tls/server.key\"", getBaseFolderPath(), orgGroup, domain, peer, domain, getBaseFolderPath(), orgGroup, domain, peer, domain))
	return cmds
}

func GenerateUserMspViaCaCommand(orderer bool, user string, pwd string, domain string, port uint, caName string, caOrg string) []string {
	orgGroup := getOrgSubPathByOrderer(orderer)
	var cmds []string
	cmds = append(cmds, fmt.Sprintf("fabric-ca-client enroll -u https://%s:%s@%s:%d --caname ca-%s -M \"%sorganizations/%s/%S/users/%s@o%s/msp\" --tls.certfiles \"%sorganizations/fabric-ca/%s/ca-cert.pem\"", user, pwd, domain, port, caName, getBaseFolderPath(), orgGroup, domain, user, domain, getBaseFolderPath(), caOrg))

	cmds = append(cmds, fmt.Sprintf("cp \"%sorganizations/%s/%s/msp/config.yaml\" \"%sorganizations/%s/%s/users/%s@%s/msp/config.yaml\"", getBaseFolderPath(), orgGroup, domain, getBaseFolderPath(), orgGroup, domain, user, domain))
	return cmds
}

func GenerateOrganizationCcpCommand() []string {
	var cmds []string
	// organizations/ccp-generate.sh
	return cmds
}
