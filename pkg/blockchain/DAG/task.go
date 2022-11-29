package DAG

import (
	"time"

	"github.com/EternallyAscend/GoToolkits/pkg/command"
)

type Task struct {
	Command   string      `json:"command" yaml:"command"`
	Timestamp time.Time   `json:"timestamp" yaml:"timestamp"`
	Reached   []*PeerInfo `json:"reached" yaml:"reached"`
}

func GenerateTask(command string, broadcast bool) *Task {
	t := &Task{
		Command:   command,
		Timestamp: time.Now(),
		Reached:   nil,
	}
	if broadcast {
		t.Reached = []*PeerInfo{}
	}
	return t
}

func (that *Task) Execute() *command.Result {
	return command.GenerateCommand(that.Command).Execute()
}
