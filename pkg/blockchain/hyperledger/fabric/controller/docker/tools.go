package docker

import (
	"fmt"
	"github.com/EternallyAscend/GoToolkits/pkg/container/docker/dockerCompose"
)

func GenerateToolsService(imageVersion string, name string, dependOn []string, networks ...string) *dockerCompose.Service {
	return &dockerCompose.Service{
		Image:         fmt.Sprintf("hyperledger/fabric-tools:%s", imageVersion),
		Environment:   []string{
			"GOPATH=/opt/gopath",
			"CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock",
			"FABRIC_LOGGING_SPEC=INFO",
		},
		Command:       "/bin/bash",
		Volumes:       []string{
			"/var/run/:/host/var/run/",

		},
		ContainerName: name,
		Networks:      networks,
		WorkingDir:    "/opt/gopath/src/github.com/hyperledger/fabric/peer",
		DependsOn:     dependOn,
		Tty:           true,
		StdinOpen:     true,
	}
}
