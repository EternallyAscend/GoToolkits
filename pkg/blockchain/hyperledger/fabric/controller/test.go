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
	// other
	cmds = append(cmds, environment.InstallGitCommand()...)
	cmds = append(cmds, environment.InstallWgetCommand()...)
	cmds = append(cmds, environment.InstallBuildEssentialCommand()...)

	// docker
	cmds = append(cmds, environment.InstallDockerCommand()...)
	cmds = append(cmds, environment.CheckDockerCommand()...)
	cmds = append(cmds, environment.CheckDockerComposeCommand()...)
	cmds = append(cmds, environment.SetDockerImageOriginCommand()...)
	cmds = append(cmds, environment.RestartDockerCommand()...)
	cmds = append(cmds, environment.PullFabricDockerImagesCommand(config.FabricVersion, config.FabricCaVersion)...)

	cmds = []string{} // skip

	// golang
	cmds = append(cmds, environment.DownloadGoCommand(config.GoDownloadPath, config.GoVersion, config.OS, config.Arch)...)
	cmds = append(cmds, environment.OpenGoWithTarCommand(config.GoVersion, config.OS, config.Arch)...)
	cmds = append(cmds, environment.ExportGoEnvironmentCommand()...)

	cmds = []string{} // skip

	// fabric
	cmds = append(cmds, environment.CloneFabricRepositoriesCommand(config.FabricVersion, config.FabricCaVersion)...)
	makeBin := false
	if makeBin {
		cmds = append(cmds, environment.CompileFabricBinaryFilesCommand(config.FabricVersion, config.FabricCaVersion)...)
	} else {
		cmds = append(cmds, environment.DownloadFabricBinaryFilesCommand(config.FabricVersionFull, config.FabricCaVersionFull)...)
		cmds = append(cmds, environment.OpenFabricBinaryFilesWithTarCommand(config.FabricVersionFull, config.FabricCaVersionFull)...)
	}
	cmds = append(cmds, environment.ExportFabricBinPath()...)

	cmds = []string{} // skip

	// Test Network

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
