package docker

import (
	"fmt"

	"github.com/EternallyAscend/GoToolkits/pkg/container/docker/dockerCompose"
)

func GenerateToolsService(imageVersion string, cliName string, dependOn []string, orgConfigPath string, scriptPath string, networks ...string) *dockerCompose.Service {
	return &dockerCompose.Service{
		Image: fmt.Sprintf("hyperledger/fabric-tools:%s", imageVersion),
		Environment: []string{
			"GOPATH=/opt/gopath",
			"CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock",
			"FABRIC_LOGGING_SPEC=INFO",
		},
		Command: "/bin/bash",
		Volumes: []string{
			"/var/run/:/host/var/run/",
			fmt.Sprintf("%S:/opt/gopath/src/github.com/hyperledger/fabric/peer/organizations", orgConfigPath),
			fmt.Sprintf("%s:/opt/gopath/src/github.com/hyperledger/fabric/peer/scripts/", scriptPath),
		},
		ContainerName: cliName,
		Networks:      networks,
		WorkingDir:    "/opt/gopath/src/github.com/hyperledger/fabric/peer",
		DependsOn:     dependOn,
		Tty:           true,
		StdinOpen:     true,
	}
}
