package environment

import "fmt"

func EnterOrCreateFolder(path string) []string {
	return []string{
		fmt.Sprintf("if [ ! -d '%s' ]; then\nmkdir %s\nfi cd %s", path, path, path),
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
