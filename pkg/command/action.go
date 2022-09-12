package command

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

type Command struct {
	str string
}

func GenerateCommand(str string) *Command {
	return &Command{
		str: strings.TrimSpace(str),
	}
}

func MergeCommandsStringArray(strs ...[]string) []string {
	var cmds []string
	for i := range strs {
		cmds = append(cmds, strs[i]...)
	}
	return cmds
}

func GenerateCommands(strs []string) []*Command {
	var cmds []*Command
	for i := range strs {
		cmds = append(cmds, GenerateCommand(strs[i]))
	}
	return cmds
}

func (that *Command) name() string {
	pos := strings.IndexAny(that.str, " ")
	if pos > 0 {
		return that.str[0:pos]
	} else {
		return that.str
	}
}

func (that *Command) args() []string {
	var args []string
	head := strings.IndexAny(that.str, " ")
	if head > 0 {
		temp := strings.TrimSpace(that.str[head:])
		pos := strings.IndexAny(temp, " ")
		for pos > 0 {
			args = append(args, temp[:pos])
			temp = strings.TrimSpace(temp[pos:])
			pos = strings.IndexAny(temp, " ")
		}
		args = append(args, temp)
	}
	return args
}

func (that *Command) GetString() string {
	return that.str
}

func (that *Command) isOutPip() bool {
	return strings.Contains(that.str, "|")
}

func (that *Command) splitOutPip() []*Command {
	strs := strings.Split(that.str, "|")
	var commands []*Command
	for i := range strs {
		commands = append(commands, GenerateCommand(strings.TrimSpace(strs[i])))
	}
	return commands
}

func (that *Command) Execute() *Result {
	result := &Result{
		err:    nil,
		stdout: bytes.Buffer{},
		stderr: bytes.Buffer{},
	}
	if that.isOutPip() {
		commands := that.splitOutPip()
		var cmds []*exec.Cmd
		for i := range commands {
			cmds = append(cmds, exec.Command(commands[i].name(), commands[i].args()...))
		}
		if len(cmds) > 1 {
			cmds[0].Stderr = &result.stderr
		} else {
			result.err = errors.New("Error: Wrong cmd pip command " + that.str)
			return result
		}
		for i := 1; i < len(cmds); i++ {
			cmds[i].Stdin, result.err = cmds[i-1].StdoutPipe()
			if nil != result.err {
				return result
			}
			cmds[i].Stderr = &result.stderr
		}
		cmds[len(cmds)-1].Stdout = &result.stdout
		for i := len(cmds) - 1; i > 0; i-- {
			result.err = cmds[i].Start()
			if nil != result.err {
				return result
			}
		}
		result.err = cmds[0].Run()
		if nil != result.err {
			return result
		}
		for i := len(cmds) - 1; i > 0; i-- {
			result.err = cmds[i].Wait()
			if nil != result.err {
				return result
			}
		}
		return result
	} else {
		cmd := exec.Command(that.name(), that.args()...)
		cmd.Stdout = &result.stdout
		cmd.Stderr = &result.stderr
		result.err = cmd.Run()
		return result
	}
}
