package controller

import (
	"fmt"
	"github.com/EternallyAscend/GoToolkits/pkg/command"
	"github.com/EternallyAscend/GoToolkits/pkg/network/ssh"
	"log"
)

func ExecuteControllerCommand(target *ssh.Client, commands []*command.Command) bool {
	if nil == target {
		return false
	}
	res, errs, err := target.ExecuteMultiCommands(commands)
	if nil != err {
		log.Println(err)
		return false
	}
	log.Println(len(res), res)
	log.Println(len(errs), errs)
	log.Println(err)
	return true
}

func ExecuteControllerParallelCommand(target *ssh.Client, commands []*command.Command) bool {
	if nil == target {
		return false
	}
	res, errs, err := target.ExecuteMultiParallelCommands(commands)
	if nil != err {
		log.Println(err)
		return false
	}
	log.Println(len(res), res)
	log.Println(len(errs), errs)
	log.Println(err)
	return true
}

func installEnvironmentCommand() []*command.Command {
	return command.GenerateCommands([]string{
		"apt-get install -y golang docker.io docker-compose wget",
		"go version",
		"docker -v",
		"docker-compose -v",
		//"git clone https://github.com/hyperledger/fabric-samples.git",
	})
}

func InstallEnvironment(target *ssh.Client) bool {
	return ExecuteControllerCommand(target, installEnvironmentCommand())
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

func PullDockerImages(target *ssh.Client) bool {
	if nil == target {
		return false
	}
	res, errs, err := target.ExecuteMultiCommands(command.GenerateCommands([]string{
		"echo '{\n\"registry-mirrors\": [\"https://registry.docker-cn.com\"]\n}\n' >> /etc/docker/daemon.json",
		"systemctl restart docker",
		"docker pull hyperledger/fabric-ca:1.4.8",
		"docker pull hyperledger/fabric-baseos:2.2.0",
		"docker pull hyperledger/fabric-ccenv:2.2.0",
		"docker pull hyperledger/fabric-orderer:2.2.0",
		"docker pull hyperledger/fabric-peer:2.2.0",
		"docker pull hyperledger/fabric-tools:2.2.0",
	}))
	if nil != err {
		log.Println(err)
		return false
	}
	log.Println(len(res), res)
	log.Println(len(errs), errs)
	log.Println(err)
	return true
}

func pullFabricBinaryFilesCommand(version string, versionCA string) []*command.Command {
	common := "wget https://github.com/hyperledger/fabric/releases/download/v%s/hyperledger-fabric-linux-amd64-%s.tar.gz"
	ca :="wget https://github.com/hyperledger/fabric-ca/releases/download/v%s/hyperledger-fabric-ca-linux-amd64-%s.tar.gz"
	return command.GenerateCommands([]string{
		fmt.Sprintf(common, version, version),
		fmt.Sprintf(ca, versionCA, versionCA),
	})
}

func PullFabricBinaryFiles(target *ssh.Client) bool {
	return ExecuteControllerParallelCommand(target, pullFabricBinaryFilesCommand("2.2.0", "1.4.8"))
}
