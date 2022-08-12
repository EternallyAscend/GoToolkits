package gRPC

import (
	"fmt"
	"github.com/EternallyAscend/GoToolkits/pkg/command"
)

// go get google.golang.org/grpc
// go get google.golang.org/protobuf/cmd/protoc-gen-go
// go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

func TransferProtobufToGo(goFilePath string, protoPath string) {
	commands := []*command.Command{
		command.GenerateCommand(fmt.Sprintf("protoc --go-grpc_out=%s %s", goFilePath, protoPath)),
		command.GenerateCommand(fmt.Sprintf("protoc --go_out=%s %s", goFilePath, protoPath)),
	}
	command.ExecuteMultiCommands(commands)
}
