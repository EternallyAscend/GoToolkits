package controller

import (
	"fmt"
	"github.com/EternallyAscend/GoToolkits/pkg/blockchain/hyperledger/fabric/controller/config"
	"github.com/EternallyAscend/GoToolkits/pkg/blockchain/hyperledger/fabric/controller/environment"
	"github.com/EternallyAscend/GoToolkits/pkg/command"
	"github.com/EternallyAscend/GoToolkits/pkg/network/ssh"
	"log"
)

func TestFaber() {
	// Install wget git gcc make docker docker-compose go fabric
	var cmds []string
	cmds = []string{}
	cmds = append(cmds, environment.InstallDockerCommand()...)
	cmds = append(cmds, environment.CheckDockerCommand()...)
	cmds = append(cmds, environment.CheckDockerComposeCommand()...)
	cmds = append(cmds, environment.SetDockerImageOriginCommand()...)
	cmds = append(cmds, environment.RestartDockerCommand()...)
	cmds = append(cmds, environment.PullFabricDockerImagesCommand(config.FabricVersion, config.FabricCaVersion)...)

	// Connect server for test
	cli := ssh.GenerateDefaultIPv4ClientSSH("root", "10.134.116.104", 22, "linux", "")
	err := cli.Connect()
	if err != nil {
		log.Println(err)
		return
	}
	// execute commands
	r, e, err := cli.ExecuteMultiCommands(command.GenerateCommands(cmds))
	if err != nil {
		log.Println(err)
		return
	}
	cli.Close()
	// print result
	fmt.Println(r, e)
}
