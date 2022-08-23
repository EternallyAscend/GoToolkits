package controller

import (
	"github.com/EternallyAscend/GoToolkits/pkg/command"
	"github.com/EternallyAscend/GoToolkits/pkg/network/ssh"
	"log"
)

func InstallEnvironment(target *ssh.Client) bool {
	if nil == target {
		return false
	}
	res, errs, err := target.ExecuteMultiCommands(command.GenerateCommands([]string{
		"apt-get install -y golang docker.io docker-compose wget",
		"go version",
		"docker -v",
		"docker-compose -v",
		//"git clone https://github.com/hyperledger/fabric-samples.git",
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


func PullFabricBinaryFiles(target *ssh.Client) bool {
	if nil == target {
		return false
	}
	res, errs, err := target.ExecuteMultiCommands(command.GenerateCommands([]string{
		"wget https://github.com/hyperledger/fabric/releases/download/v2.2.0/hyperledger-fabric-linux-amd64-2.2.0.tar.gz",
		"wget https://github.com/hyperledger/fabric-ca/releases/download/v1.4.8/hyperledger-fabric-ca-linux-amd64-1.4.8.tar.gz",
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
