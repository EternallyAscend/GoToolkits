package daemon

import goDaemon "github.com/sevlyar/go-daemon"

// GoDaemon 使用go-daemon库实现后台进程。
func GoDaemon() {
	context := &goDaemon.Context{}
	defer func(context *goDaemon.Context) {
		err := context.Release()
		if err != nil {
		}
	}(context)
}
