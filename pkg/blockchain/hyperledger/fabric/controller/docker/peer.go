package docker

import (
	"fmt"

	"github.com/EternallyAscend/GoToolkits/pkg/container/docker/dockerCompose"
)

func GeneratePeerService(imageVersion string, peerName string, orgName string, domainRoot string, network string, tls bool, tlsPath string, profile bool, msp string, mspPath string, peerPort uint, chaincodePort uint, operationsPort uint) *dockerCompose.Service {
	service := &dockerCompose.Service{
		Image: fmt.Sprintf("hyperledger/fabric-peer:%s", imageVersion),
		Environment: []string{
			"CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock",
			fmt.Sprintf("CORE_VM_DOCKER_HOSTCONFIG_NETWORKMODE=%s", network),
			"FABRIC_LOGGING_SPEC=INFO",
			fmt.Sprintf("CORE_PEER_TLS_ENABLED=%v", tls),
			fmt.Sprintf("CORE_PEER_PROFILE_ENABLED=%v", profile),
			fmt.Sprintf("CORE_PEER_ID=%s.%s.%s", peerName, orgName, domainRoot),
			fmt.Sprintf("CORE_PEER_ADDRESS=%s.%s.%s:%d", peerName, orgName, domainRoot, peerPort),
			fmt.Sprintf("CORE_PEER_LISTENADDRESS=0.0.0.0:%d", peerPort),
			fmt.Sprintf("CORE_PEER_CHAINCODEADDRESS=%s.%s.%s:%d", peerName, orgName, domainRoot, chaincodePort),
			fmt.Sprintf("CORE_PEER_CHAINCODELISTENADDRESS=0.0.0.0:%d", chaincodePort),
			fmt.Sprintf("CORE_PEER_GOSSIP_BOOTSTRAP=%s.%s.%s:%d", peerName, orgName, domainRoot, peerPort),
			fmt.Sprintf("CORE_PEER_GOSSIP_EXTERNALENDPOINT=%s:%d", peerName, orgName, domainRoot, peerPort),
			fmt.Sprintf("CORE_PEER_LOCALMSPID=%s", msp),
			fmt.Sprintf("CORE_OPERATIONS_LISTENADDRESS=%s.%s.%s:%d", peerName, orgName, domainRoot, operationsPort),
		},
		Ports: []string{
			fmt.Sprintf("\"%d:%d\"", peerPort, peerPort),
			fmt.Sprintf("\"%d:%d\"", chaincodePort, chaincodePort),
			fmt.Sprintf("\"%d:%d\"", operationsPort, operationsPort),
		},
		Command: "peer node start",
		Volumes: []string{
			"/var/run/docker.sock:/host/var/run/docker.sock",
			fmt.Sprintf("%s:/etc/hyperledger/fabric/msp", mspPath),
			fmt.Sprintf("%s:/etc/hyperledger/fabric/tls", tlsPath),
			fmt.Sprintf("%s.%s.%s:/var/hyperledger/production", peerName, orgName, domainRoot),
		},
		ContainerName: fmt.Sprintf("%s.%s.%s", peerName, orgName, domainRoot),
		Networks: []string{
			network,
		},
		WorkingDir: "/opt/gopath/src/github.com/hyperledger/fabric/peer",
		Tty:        true,
	}
	if tls {
		service.Environment = append(service.Environment, fmt.Sprintf("CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/fabric/tls/server.crt"))
		service.Environment = append(service.Environment, fmt.Sprintf("CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/fabric/tls/server.key"))
		service.Environment = append(service.Environment, fmt.Sprintf("CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/fabric/tls/ca.crt"))
	}
	return service
}
