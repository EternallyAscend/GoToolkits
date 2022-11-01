package DAG

import (
	"errors"
	"github.com/EternallyAscend/GoToolkits/pkg/network/ip"
	"time"
)

// TODO Change to Gossip Cluster https://www.jianshu.com/p/5198b869374a
// Example https://github.com/asim/kv

// DHT Discover
// https://blog.csdn.net/luoye4321/article/details/83587397
// https://blog.csdn.net/u012576116/article/details/81363829

// Timer https://seekload.blog.csdn.net/article/details/113155421

const DefaultOrigin = "192.168.1.1:8000"

type Peer struct {
	Info *PeerInfo `json:"info" yaml:"info"`
	Router *PeerRouter `json:"router" yaml:"router"`
	Tasks *TasksList `json:"tasks" yaml:"tasks"`
}

type PeerInfo struct {
	Address string `json:"address" yaml:"address"`
	Port uint `json:"port" yaml:"port"`
}

type PeerRouter struct {
	Neighbor []*PeerInfo `json:"neighbor" yaml:"neighbor"`
}

type TasksList struct {
	Command string `json:"command" yaml:"command"`
	Timestamp time.Time `json:"timestamp" yaml:"timestamp"`
	Reached []*PeerInfo `json:"reached" yaml:"reached"`
}

func GeneratePeer(port uint) (*Peer, error) {
	ipv4Address, err := ip.GetLocalIpv4Address()
	if nil != err {
		return nil, err
	}
	return &Peer{
		Info: &PeerInfo{
			Address: ipv4Address,
			Port:    port,
		},
	}, nil
}

func (that *Peer) Join() error {
	// TODO Request Default Address
	// TODO fetch Neighbor
	// TODO set Timer (Background Tasks)
	return errors.New("")
}

func (that *Peer) fetch() *Peer {
	// TODO Calculate by Hash
	return that
}

func (that *Peer) Refresh() *Peer {
	that.fetch()
	return that
}

func (that *Peer) ReleaseTask() *Peer {
	return that
}

func (that *Peer) Broadcast() *Peer {
	return that
}

func (that *Peer) Exit() error {
	// TODO Notice Neighbor
	return errors.New("")
}

func (that *Peer) Execute() error {
	return nil
}
