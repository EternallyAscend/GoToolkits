package docker

import (
	"fmt"
	"github.com/EternallyAscend/GoToolkits/pkg/container/docker/dockerCompose"
)

func GenerateCouchdbCommand(dbPeerName, domainRoot, user, pwd string, port uint) *dockerCompose.Service {
	return &dockerCompose.Service{
		ContainerName: fmt.Sprintf("%s.%s", dbPeerName, domainRoot),
		Image:         "couchdb:3.1.1",
		Environment: []string{
			fmt.Sprintf("COUCHDB_USER=%s", user),
			fmt.Sprintf("COUCHDB_PASSWORD=%s", pwd),
		},
		Ports: []string{
			fmt.Sprintf("\"%d:5984\"", port),
		},
		Tty: true,
	}
}
