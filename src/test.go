package main

import (
	"fmt"
	"github.com/EternallyAscend/GoToolkits/pkg/IO/Medias/music"
	"github.com/EternallyAscend/GoToolkits/pkg/IO/YAML"
	"github.com/EternallyAscend/GoToolkits/pkg/blockchain/hyperledger/fabric/controller"
	"github.com/EternallyAscend/GoToolkits/pkg/blockchain/hyperledger/fabric/controller/docker"
	"github.com/EternallyAscend/GoToolkits/pkg/command"
	"github.com/EternallyAscend/GoToolkits/pkg/container/docker/dockerCompose"
	"github.com/EternallyAscend/GoToolkits/pkg/network/ssh"
	"os"
)

func main() {
	controller.TestFaber()
	os.Exit(0)
	scli, _ := ssh.ReadPwdClientFromYaml("./config/testServer.yaml")
	cli := scli.CreateClient()
	cli.Connect()
	fmt.Println(cli.ExecuteSingleCommand(command.GenerateCommand("which docker")))
	cli.Close()

	dealer := dockerCompose.GenerateDockerYAML("2")
	dealer.AddNetwork(dockerCompose.GenerateNetwork("testNetwork"))
	cas := docker.GenerateCaServices("2.2", "ca-org1", false, 7054, 17054, "../organizations/fabric-ca/org1", "adminpw", "test1", "test2", "test3")
	YAML.ExportToFileYaml(cas, "./cas.yaml")
	dealer.AddService(cas)

	dealer.AddNetwork(dockerCompose.GenerateNetwork("test"))
	YAML.ExportToFileYaml(dealer, "./dealer.yaml")
	data, _ := dealer.ExportToByteArray()
	fmt.Println(string(data))
	os.Exit(0)
	sshClient := ssh.GenerateDefaultIPv4ClientSSH("root", "192.168.3.21", 22, "linux", "")
	err := sshClient.Connect()
	if nil != err {
		fmt.Println(err)
	}
	bools := controller.InstallEnvironment(sshClient)
	fmt.Println(bools)
	bools = controller.CheckEnvironment(sshClient)
	fmt.Println(bools)
	//bools = controller.PullDockerImages(sshClient)
	//fmt.Println(bools)
	//bools = controller.PullFabricBinaryFiles(sshClient)
	//fmt.Println(bools)
	_ = sshClient.Close()
	music.Generate()
}
