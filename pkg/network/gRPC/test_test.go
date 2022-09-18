package gRPC

import (
	"fmt"

	"github.com/EternallyAscend/GoToolkits/pkg/command"
)

func test() {
	// cmd := command.GenerateCommand("powershell ls ../gRPC/APIs/")
	cmd := command.GenerateCommand("powershell cat ../gRPC/APIs/test.proto")
	// gRPC.TransferProtobufToGo("D:\\PersonalFiles\\Project\\Go\\gRPC\\pkg", "D:\\PersonalFiles\\Project\\Go\\gRPC\\APIs\\test.proto")
	TransferProtobufToGo("../gRPC/pkg", "../gRPC/APIs/", "test.proto")
	// cmd := command.GenerateCommand("go help gopath | grep help")
	// cmd := command.GenerateCommand("netstat -tdnl | grep 2")
	// cmd := command.GenerateCommand("ls pkg | grep t | grep th | grep a")
	// cmd := command.GenerateCommand("go version")
	res := cmd.Execute()
	fmt.Println(res)
	// fmt.Println("error: ", res.GetErr(), "\nOutput:\n", res.GetOutputAsString(), "\nErr:\n", res.GetErrorAsString())
}
