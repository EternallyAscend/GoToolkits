package main

import (
	"fmt"
	"github.com/EternallyAscend/GoToolkits/pkg/IO/Medias/music"
	"github.com/EternallyAscend/GoToolkits/pkg/command"
	"os"
)

func main() {
	//cmd := command.GenerateCommand("go help gopath | grep help")
	//cmd := command.GenerateCommand("netstat -tdnl | grep 2")
	cmd := command.GenerateCommand("ls pkg | grep t | grep th | grep a")
	res := cmd.Execute()
	fmt.Println("error: ", res.GetErr(), "\nOutput:\n", res.GetOutputAsString(), "\nErr:\n", res.GetErrorAsString())
	os.Exit(0)
	music.Generate()
}
