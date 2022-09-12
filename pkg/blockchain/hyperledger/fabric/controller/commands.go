package controller

import (
	"fmt"
	"github.com/EternallyAscend/GoToolkits/pkg/command"
)

func checkEnvironmentCommand() []*command.Command {
	return command.GenerateCommands([]string{
		"echo $PATH",
		"source /etc/profile" +
			"&& faberGo=$(go version)" +
			"&& faberGo=$(echo ${faberGo##*version })" +
			"&& faberGoVersion=$(echo ${faberGo##*go})" +
			"&& faberGoVersion=$(echo ${faberGoVersion% *})" +
			"&& faberGoArch=$(echo ${faberGo##* })" +
			"&& echo $faberGoVersion" +
			"&& echo $faberGoArch" +
			"&& echo $faberGoVersion $faberGoArch",
	})
}

func enterOrMakeBaseFolder() string {
	return "if [ ! -d '/root/faber' ]; then\nmkdir /root/faber\nfi\ncd /root/faber"
}

func installGolang(version string, os string, arch string) []*command.Command {
	tarFile := fmt.Sprintf("go%s.%s-%s.tar.gz", version, os, arch)
	return command.GenerateCommands([]string{
		enterOrMakeBaseFolder(),
		fmt.Sprintf("cd /root/faber\nif [ ! -f '%s' ]; then\n wget https://go.dev/dl/%s\nfi", tarFile, tarFile),
		fmt.Sprintf("cd /root/faber\nif [ ! -d 'go' ]; then\ntar -zxvf %s\nfi", tarFile),
		fmt.Sprintf("echo 'export PATH=/root/faber/go/bin:$PATH' > /etc/profile.d/faber.sh"),
		fmt.Sprintf("source /etc/profile && go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct"),
	})
}

func installEnvironmentCommand() []*command.Command {
	return command.GenerateCommands([]string{
		"apt-get install -y golang docker.io docker-compose wget",
		"go version",
		"docker -v",
		"docker-compose -v",
		//"https://go.dev/dl/go1.19.linux-amd64.tar.gz",
		//"git clone https://github.com/hyperledger/fabric-samples.git",
	})
}

func pullDockerImagesCommand(version string, versionCA string) []*command.Command {
	template := "docker pull hyperledger/fabric"
	return command.GenerateCommands([]string{
		"echo '{\n\"registry-mirrors\": [\"https://registry.docker-cn.com\"]\n}\n' >> /etc/docker/daemon.json",
		"systemctl restart docker",
		fmt.Sprintf("%s-%s:%s", template, "ca", versionCA),
		fmt.Sprintf("%s-%s:%s", template, "baseos", version),
		fmt.Sprintf("%s-%s:%s", template, "ccenv", version),
		fmt.Sprintf("%s-%s:%s", template, "orderer", version),
		fmt.Sprintf("%s-%s:%s", template, "peer", version),
		fmt.Sprintf("%s-%s:%s", template, "tools", version),
	})
}

func pullFabricBinaryFilesCommand(version string, versionCA string) []*command.Command {
	common := "wget https://github.com/hyperledger/fabric/releases/download/v%s/hyperledger-fabric-linux-amd64-%s.tar.gz"
	ca := "wget https://github.com/hyperledger/fabric-ca/releases/download/v%s/hyperledger-fabric-ca-linux-amd64-%s.tar.gz"
	return command.GenerateCommands([]string{
		fmt.Sprintf(common, version, version),
		fmt.Sprintf(ca, versionCA, versionCA),
	})
}

func makeFabricBinaryFilesCommand(versiion string, versionCA string) []*command.Command {
	return command.GenerateCommands([]string{
		"sudo apt-get -y install build-essential libtool",
		fmt.Sprintf(""),
	})
}
