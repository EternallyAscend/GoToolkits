package environment

import "fmt"

func EnterWithCreateFolder(path string) []string {
	return []string{
		fmt.Sprintf("if [ ! -d '%s' ]; then\nmkdir %s\nfi cd %s", path, path, path),
	}
}

func InstallJqCommand() []string {
	return []string{
		"apt-get -y install jq",
	}
}

func InstallGitCommand() []string {
	return []string{
		"apt-get -y install git",
	}
}

func InstallWgetCommand() []string {
	return []string{
		"apt-get -y install wget",
	}
}

func InstallBuildEssentialCommand() []string {
	return []string{
		"apt-get -y install build-essential",
	}
}

// SetHosts For test using
func SetHosts() []string {
	var cmds []string
	// TODO Set Hyperledger Hosts.
	return cmds
}
