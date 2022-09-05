package main

import (
	"fmt"
	"github.com/EternallyAscend/GoToolkits/pkg/IO/Medias/music"
	"github.com/EternallyAscend/GoToolkits/pkg/IO/YAML"
	"github.com/EternallyAscend/GoToolkits/pkg/blockchain/hyperledger/fabric/controller"
	"github.com/EternallyAscend/GoToolkits/pkg/blockchain/hyperledger/fabric/controller/docker"
	"github.com/EternallyAscend/GoToolkits/pkg/command"
	"github.com/EternallyAscend/GoToolkits/pkg/container/docker/dockerCompose"
	"github.com/EternallyAscend/GoToolkits/pkg/cryptography/homomorphic/pedersonCommitment"
	"github.com/EternallyAscend/GoToolkits/pkg/network/gRPC"
	"github.com/EternallyAscend/GoToolkits/pkg/network/ssh"
	"os"
)

func main() {
	dealer := dockerCompose.GenerateDockerYAML("2")
	dealer.AddNetwork(dockerCompose.GenerateNetwork("testNetwork"))
	cas := docker.GenerateCaServices("2.2", "ca-org1", false, 7054, 17054, "../organizations/fabric-ca/org1", "adminpw", "test1", "test2", "test3")
	YAML.ExportToFileYaml(cas, "./cas.yaml")
	dealer.AddService(cas)

	dealer.AddNetwork(dockerCompose.GenerateNetwork("test"))
	YAML.ExportToFileYaml(dealer, "./dealer.yaml")
	data, _ := dealer.ExportToByteArray()
	fmt.Println(string(data))
	os.Exit(0)
	sshClient := ssh.GenerateDefaultClientSSH("root", "192.168.3.21", 22, "linux", "")
	err := sshClient.Connect()
	if nil != err {
		fmt.Println(err)
	}
	bools := controller.InstallEnvironment(sshClient)
	fmt.Println(bools)
	bools = controller.CheckEnvironment(sshClient)
	fmt.Println(bools)
	//bools = controller.PullDockerImages(sshClient)
	//fmt.Println(bools)
	//bools = controller.PullFabricBinaryFiles(sshClient)
	//fmt.Println(bools)
	_ = sshClient.Close()

	os.Exit(0)
	//vm := &pedersonCommitment.VerifiableMessage{}
	//fmt.Println(vm)
	//dealer := pedersonCommitment.GenerateDealer()
	//vd := pedersonCommitment.GenerateVerifiableData(dealer)
	//vd.AddCommonData([]byte("Common data abc."))
	//vd.AddEncryptData([]byte("Encrypt data abc."))
	//vd.AddCommonData([]byte("123."))
	//vd.AddEncryptData([]byte("Encrypt data abc."))
	//vd.SetCommonVerifyData([][]byte{
	//	[]byte("Common data abc."),
	//	[]byte("Encrypt data abc."),
	//	[]byte("123."),
	//	[]byte("Encrypt data abc."),
	//})
	//for i := range vd.Units {
	//	if !vd.Units[i].Encrypt {
	//		fmt.Println(vd.Units[i].Dealer.Open(vd.Units[i].Data))
	//		d, e := vd.Units[i].Dealer.TransferToJsonByte()
	//		fmt.Println(string(d), e)
	//	}
	//}
	//fmt.Println(vd.Verify())
	vp := pedersonCommitment.GenerateVerifiableMessage()
	vp.AppendDataArray([][]byte{
		[]byte("Common data abc."),
		[]byte("Encrypt data abc."),
		[]byte("123."),
		[]byte("Encrypt data abc."),
	}, []bool{
		false, true, false, true,
	})
	vp.ConfirmMessage([][]byte{
		[]byte("Common data abc."),
		[]byte("Encrypt data abc."),
		[]byte("123."),
		[]byte("Encrypt data abc."),
	})
	fmt.Println(vp.OpenLine(1), vp.OpenLine(2), vp.CheckAll())
	fmt.Println(vp.ExportMessagesAsString())
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
