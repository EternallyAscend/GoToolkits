package cmds

import (
	"fmt"
)

func DebianInstallDocker() []string {
	return []string{
		"apt-get install -y docker.io docker-compose",
	}
}

func PullFabricDockerImages(version string, versionCA string) []string {
	template := "docker pull hyperledger/fabric"
	return []string{
		"echo '{\n\"registry-mirrors\": [\"https://registry.docker-cn.com\"]\n}\n' >> /etc/docker/daemon.json",
		"systemctl restart docker",
		fmt.Sprintf("%s-%s:%s", template, "ca", versionCA),
		fmt.Sprintf("%s-%s:%s", template, "baseos", version),
		fmt.Sprintf("%s-%s:%s", template, "ccenv", version),
		fmt.Sprintf("%s-%s:%s", template, "orderer", version),
		fmt.Sprintf("%s-%s:%s", template, "peer", version),
		fmt.Sprintf("%s-%s:%s", template, "tools", version),
	}
}

func RunDockerContainer() []string {
	return []string{}
}

func GenerateCaYaml(path string) error {
	return nil
}
