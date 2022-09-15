package docker

import "fmt"

func RunDockerMonitorCommand(port uint, network string) []string {
	return []string{
		fmt.Sprintf("docker run -d --rm --name=\"logspout\" \\\n\t--volume=/var/run/docker.sock:/var/run/docker.sock \\\n\t--publish=127.0.0.1:%d:80 \\\n\t--network  %s \\\n\tgliderlabs/logspout", port, network),
	}
}
