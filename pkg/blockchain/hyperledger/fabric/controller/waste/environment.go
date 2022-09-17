package waste

import (
	"github.com/EternallyAscend/GoToolkits/pkg/blockchain/hyperledger/fabric/controller"
	"github.com/EternallyAscend/GoToolkits/pkg/command"
	"github.com/EternallyAscend/GoToolkits/pkg/network/ssh"
	"log"
)

func InstallEnvironment(target *ssh.IPv4Client) bool {
	return controller.ExecuteControllerCommand(target, installGolang("1.19", "linux", "amd64"))
}

func CheckEnvironment(target *ssh.IPv4Client) bool {
	res := command.ExecuteMultiCommands(checkEnvironmentCommand())
	for i := range res {
		log.Println(res[i].GetOutputAsString(), res[i].GetErrorAsString(), res[i].GetErr())
	}
	return true
	//return ExecuteControllerCommand(target, checkEnvironmentCommand())
}

func PullDockerImages(target *ssh.IPv4Client) bool {
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

func PullFabricBinaryFiles(target *ssh.IPv4Client) bool {
	return controller.ExecuteControllerParallelCommand(target, pullFabricBinaryFilesCommand("2.2.0", "1.4.8"))
}
