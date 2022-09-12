package environment

import "fmt"

func InstallDockerCommand() []string {
	return []string{
		"apt-get install -y docker.io docker-compose",
	}
}

func CheckDockerCommand() []string {
	return []string{}
}

func CheckDockerComposeCommand() []string {
	return []string{}
}

func SetDockerImageOriginCommand() []string {
	return []string{
		"echo '{\n\"registry-mirrors\": [\"https://registry.docker-cn.com\"]\n}\n' >> /etc/docker/daemon.json",
	}
}

func RestartDockerCommand() []string {
	return []string{
		"systemctl restart docker",
	}
}

func PullFabricDockerImagesCommand(version, versionCA string) []string {
	template := "docker pull hyperledger/fabric-%s:%s"
	return []string{
		fmt.Sprintf(template, "ca", versionCA),
		fmt.Sprintf(template, "baseos", version),
		fmt.Sprintf(template, "ccenv", version),
		fmt.Sprintf(template, "orderer", version),
		fmt.Sprintf(template, "peer", version),
		fmt.Sprintf(template, "tools", version),
	}
}
