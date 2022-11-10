package DAG

import "time"

type Task struct {
	Command   string      `json:"command" yaml:"command"`
	Timestamp time.Time   `json:"timestamp" yaml:"timestamp"`
	Reached   []*PeerInfo `json:"reached" yaml:"reached"`
}
