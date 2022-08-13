package gRPC

import (
	"fmt"
	"github.com/EternallyAscend/GoToolkits/pkg/command"
	"log"
)

// go get google.golang.org/grpc
// go get google.golang.org/protobuf/cmd/protoc-gen-go
// go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

func TransferProtobufToGo(goFilePath string, protoPath string, protoFilePath string) {
	totalResult := 2
	commands := []*command.Command{
		command.GenerateCommand(fmt.Sprintf("protoc --go_out=%s --proto_path=%s %s/%s", goFilePath, protoPath, protoPath, protoFilePath)),
		command.GenerateCommand(fmt.Sprintf("protoc --go-grpc_out=%s --proto_path=%s %s/%s", goFilePath, protoPath, protoPath, protoFilePath)),
	}
	res := command.ExecuteMultiCommands(commands)
	if totalResult == len(res) {
		if nil != res[0].GetErr() {
			log.Println(res[0].GetErrorAsString())
			log.Println(res[1].GetOutputAsString())
			log.Fatal(res[0].GetErr())
		}
		if nil != res[1].GetErr() {
			log.Println(res[1].GetErrorAsString())
			log.Println(res[0].GetOutputAsString())
			log.Fatal(res[1].GetErr())
		}
	} else {
		log.Fatalf("Wrong protobuf result. %d/%d.\n", len(res), totalResult)
	}
}
