package docker

import (
	"fmt"
	"github.com/EternallyAscend/GoToolkits/pkg/container/docker/dockerCompose"
)

func GenerateCaServices(imageVersion string, name string, tls bool, serverPort uint, listenPort uint, volume string, password string, networks ...string) *dockerCompose.Service {
	return &dockerCompose.Service{
		Image:         fmt.Sprintf("hyperledger/fabric-ca:%s", imageVersion),
		Environment:   []string{
			"FABRIC_CA_HOME=/etc/hyperledger/fabric-ca-server",
			fmt.Sprintf("FABRIC_CA_SERVER_CA_NAME=%s", name),
			fmt.Sprintf("FABRIC_CA_SERVER_TLS_ENABLED=%v", tls),
			fmt.Sprintf("FABRIC_CA_SERVER_PORT=%d", serverPort),
			fmt.Sprintf("FABRIC_CA_SERVER_OPERATIONS_LISTENADDRESS=0.0.0.0:%d", listenPort),
		},
		Ports:         []string{
			fmt.Sprintf("%d:%d", serverPort, serverPort),
			fmt.Sprintf("%d:%d", listenPort, listenPort),
		},
		Command:       fmt.Sprintf("sh -c 'fabric-ca-server start -b admin:%s -d'", password),
		Volumes:       []string{
			fmt.Sprintf("%s:/etc/hyperledger/fabric-ca-server", volume),
		},
		ContainerName: name,
		Networks:      networks,
	}
}
