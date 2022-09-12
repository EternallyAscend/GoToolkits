package environment

import (
	"fmt"
	"github.com/EternallyAscend/GoToolkits/pkg/blockchain/hyperledger/fabric/controller"
)

func DownloadGoCommand(addr, version, os, arch string) []string {
	tarFile := fmt.Sprintf("go%s.%s-%s.tar.gz", version, os, arch)
	return []string{
		fmt.Sprintf("cd %s && if [ ! -f '%s' ]; then\n wget %s%s\nfi", controller.AssertsPath, tarFile, addr, tarFile),
	}
}

func OpenGoWithTarCommand(version, os, arch string) []string {
	tarFile := fmt.Sprintf("go%s.%s-%s.tar.gz", version, os, arch)
	return []string{
		fmt.Sprintf("cd %s && if [ ! -d 'go' ]; then\ntar -zxvf %s\nfi", controller.FaberRootPath, tarFile),
	}
}

// ExportGoEnvironmentCommand Append Env info into envFilePath.
func ExportGoEnvironmentCommand() []string {
	return []string{
		fmt.Sprintf("echo 'export PATH=%sbin:$PATH' >> %s", controller.Go, controller.EnvironmentFilePath),
		fmt.Sprintf("source /etc/profile && go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct"),
	}
}
