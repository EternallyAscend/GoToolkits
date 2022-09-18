package docker

import (
	"fmt"

	"github.com/EternallyAscend/GoToolkits/pkg/container/docker/dockerCompose"
)

func GenerateOrdererService(imageVersion string, peerName string, orgName string, domainRoot string, generalPort uint, msp string, mspPath string, operationPort uint, network string, tls bool, tlsPath string, blockPath string, kafkaVerbose bool) *dockerCompose.Service {
	service := &dockerCompose.Service{
		Image: fmt.Sprintf("hyperledger/fabric-orderer:%s", imageVersion),
		Environment: []string{
			"FABRIC_LOGGING_SPEC=INFO",
			"ORDERER_GENERAL_LISTENADDRESS=0.0.0.0",
			fmt.Sprintf("ORDERER_GENERAL_LISTENPORT=%d", generalPort),
			"ORDERER_GENERAL_GENESISMETHOD=file",
			"ORDERER_GENERAL_GENESISFILE=/var/hyperledger/orderer/orderer.genesis.block",
			fmt.Sprintf("ORDERER_GENERAL_LOCALMSPID=%s", msp),
			"ORDERER_GENERAL_LOCALMSPDIR=/var/hyperledger/orderer/msp",
			fmt.Sprintf("ORDERER_OPERATIONS_LISTENADDRESS=%s.%s.%s:%d", peerName, orgName, domainRoot, operationPort),
			fmt.Sprintf("ORDERER_GENERAL_TLS_ENABLED=%v", tls),
			"ORDERER_KAFKA_TOPIC_REPLICATIONFACTOR=1",
			fmt.Sprintf("ORDERER_KAFKA_VERBOSE=%v", kafkaVerbose),
		},
		Ports: []string{
			fmt.Sprintf("\"%d:%d\"", generalPort, generalPort),
			fmt.Sprintf("\"%d:%d\"", operationPort, operationPort),
		},
		Command: "orderer",
		Volumes: []string{
			fmt.Sprintf("%s:/var/hyperledger/orderer/orderer.genesis.block", blockPath),
			fmt.Sprintf("%s:/var/hyperledger/orderer/tls", tlsPath),
			fmt.Sprintf("%s:/var/hyperledger/orderer/msp", mspPath),
			fmt.Sprintf("%s.%s.%s:/var/hyperledger/production/orderer", peerName, orgName, domainRoot),
		},
		ContainerName: fmt.Sprintf("%s.%s.%s", peerName, orgName, domainRoot),
		Networks: []string{
			network,
		},
		WorkingDir: "/opt/gopath/src/github.com/hyperledger/fabric",
		Tty:        true,
	}
	if tls {
		service.Environment = append(service.Environment, fmt.Sprintf("ORDERER_GENERAL_TLS_PRIVATEKEY=/var/hyperledger/orderer/tls/server.key"))
		service.Environment = append(service.Environment, fmt.Sprintf("ORDERER_GENERAL_TLS_CERTIFICATE=/var/hyperledger/orderer/tls/server.crt"))
		service.Environment = append(service.Environment, fmt.Sprintf("ORDERER_GENERAL_TLS_ROOTCAS=[/var/hyperledger/orderer/tls/ca.crt]"))
		service.Environment = append(service.Environment, fmt.Sprintf("ORDERER_GENERAL_CLUSTER_CLIENTCERTIFICATE=/var/hyperledger/orderer/tls/server.crt"))
		service.Environment = append(service.Environment, fmt.Sprintf("ORDERER_GENERAL_CLUSTER_CLIENTPRIVATEKEY=/var/hyperledger/orderer/tls/server.key"))
		service.Environment = append(service.Environment, fmt.Sprintf("ORDERER_GENERAL_CLUSTER_ROOTCAS=[/var/hyperledger/orderer/tls/ca.crt]"))
	}
	return service
}
