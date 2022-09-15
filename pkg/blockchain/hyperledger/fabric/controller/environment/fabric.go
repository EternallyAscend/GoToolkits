package environment

import (
	"fmt"
	"github.com/EternallyAscend/GoToolkits/pkg/blockchain/hyperledger/fabric/controller/config"
)

func DownloadFabricBinaryFilesCommand(version, versionCA string) []string {
	common := "wget https://github.com/hyperledger/fabric/releases/download/v%s/hyperledger-fabric-linux-amd64-%s.tar.gz"
	ca := "wget https://github.com/hyperledger/fabric-ca/releases/download/v%s/hyperledger-fabric-ca-linux-amd64-%s.tar.gz"
	return []string{
		fmt.Sprintf(common, version, version),
		fmt.Sprintf(ca, versionCA, versionCA),
	}
}

func OpenFabricBinaryFilesWithTarCommand(version, versionCA string) []string {
	common := fmt.Sprintf("hyperledger-fabric-linux-amd64-%s.tar.gz)", version)
	ca := fmt.Sprintf("hyperledger-fabric-ca-linux-amd64-%s.tar.gz", versionCA)
	return []string{
		fmt.Sprintf("cd %s && tar -zxvf %s && mv ./bin/* %s", config.AssertsPath, common, config.BinaryPath),
		fmt.Sprintf("cd %s && tar -zxvf %s && mv ./bin/* %s", config.AssertsPath, ca, config.BinaryPath),
	}
}

func CloneFabricRepositoriesCommand(version, versionCA string) []string {
	return []string{
		fmt.Sprintf("cd %s && git clone https://github.com/hyperledger/fabric", config.AssertsPath),
		fmt.Sprintf("cd %s && git clone https://github.com/hyperledger/fabric-ca", config.AssertsPath),
		fmt.Sprintf("cd %s && git clone https://github.com/hyperledger/fabric-samples", config.AssertsPath),
	}
}

func CompileFabricBinaryFilesCommand(version, versionCA string) []string {
	return []string{
		fmt.Sprintf("mkdir %s", config.BinaryPath),
		fmt.Sprintf("cd %sfabric && git checkout release-%s && make release && mv ./release/* %s", config.AssertsPath, version, config.BinaryPath),
		fmt.Sprintf("cd %sfabric-ca && git checkout release-%s && make release && mv ./release/* %s", config.AssertsPath, versionCA, config.BinaryPath),
	}
}

func ExportFabricBinPath() []string {
	return []string{
		fmt.Sprintf("echo \"export PATH=$PATH:%s\" >> %s", config.BinaryPath, config.EnvironmentFilePath),
	}
}
