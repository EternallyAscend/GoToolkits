package controller

import (
	"github.com/EternallyAscend/GoToolkits/pkg/command"
	"github.com/EternallyAscend/GoToolkits/pkg/network/ssh"
	"log"
)

func ExecuteControllerCommand(target *ssh.IPv4Client, commands []*command.Command) bool {
	if nil == target {
		return false
	}
	res, errs, err := target.ExecuteMultiCommands(commands)
	if nil != err {
		log.Println(err)
		return false
	}
	log.Println(len(res), res)
	log.Println(len(errs), errs)
	log.Println(err)
	return true
}

func ExecuteControllerParallelCommand(target *ssh.IPv4Client, commands []*command.Command) bool {
	if nil == target {
		return false
	}
	res, errs, err := target.ExecuteMultiParallelCommands(commands)
	if nil != err {
		log.Println(err)
		return false
	}
	log.Println(len(res), res)
	log.Println(len(errs), errs)
	log.Println(err)
	return true
}
