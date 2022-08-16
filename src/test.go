package main

import (
	"fmt"
	"github.com/EternallyAscend/GoToolkits/pkg/IO/Medias/music"
	"github.com/EternallyAscend/GoToolkits/pkg/command"
	"github.com/EternallyAscend/GoToolkits/pkg/cryptography/homomorphic/pedersonCommitment"
	"github.com/EternallyAscend/GoToolkits/pkg/network/gRPC"
	"os"
)

func main() {
	dealer := pedersonCommitment.GenerateDealer()
	vd := pedersonCommitment.GenerateVerifiableData(dealer)
	vd.AddCommonData([]byte("Common data abc."))
	vd.AddEncryptData([]byte("Encrypt data abc."))
	vd.AddCommonData([]byte("123."))
	vd.AddEncryptData([]byte("Encrypt data abc."))
	vd.SetCommonVerifyData([][]byte{
		[]byte("Common data abc."),
		[]byte("Encrypt data abc."),
		[]byte("123."),
		[]byte("Encrypt data abc."),
	})
	for i := range vd.Units {
		if !vd.Units[i].Encrypt {
			fmt.Println(vd.Units[i].Dealer.Open(vd.Units[i].Data))
			d, e := vd.Units[i].Dealer.TransferToJsonByte()
			fmt.Println(string(d), e)
		}
	}
	fmt.Println(vd.Verify())
	//cmd := command.GenerateCommand("powershell ls ../gRPC/APIs/")
	cmd := command.GenerateCommand("powershell cat ../gRPC/APIs/test.proto")
	//gRPC.TransferProtobufToGo("D:\\PersonalFiles\\Project\\Go\\gRPC\\pkg", "D:\\PersonalFiles\\Project\\Go\\gRPC\\APIs\\test.proto")
	gRPC.TransferProtobufToGo("../gRPC/pkg", "../gRPC/APIs/", "test.proto")
	//cmd := command.GenerateCommand("go help gopath | grep help")
	//cmd := command.GenerateCommand("netstat -tdnl | grep 2")
	//cmd := command.GenerateCommand("ls pkg | grep t | grep th | grep a")
	//cmd := command.GenerateCommand("go version")
	res := cmd.Execute()
	fmt.Println(res)
	//fmt.Println("error: ", res.GetErr(), "\nOutput:\n", res.GetOutputAsString(), "\nErr:\n", res.GetErrorAsString())
	os.Exit(0)
	music.Generate()
}
