package DAG

import (
	"context"
	"log"
	"time"

	"github.com/EternallyAscend/GoToolkits/pkg/network/ip"
)

// TODO Change to Gossip Cluster https://www.jianshu.com/p/5198b869374a
// Example https://github.com/asim/kv

// DHT Discover
// https://blog.csdn.net/luoye4321/article/details/83587397
// https://blog.csdn.net/u012576116/article/details/81363829

// Timer https://seekload.blog.csdn.net/article/details/113155421

type Peer struct {
	Info   *PeerInfo    `json:"info" yaml:"info"`
	Router *PeerRouter  `json:"router" yaml:"router"`
	Tasks  []*TasksList `json:"tasks" yaml:"tasks"`
	alive  bool
	ctx    *context.Context
}

type PeerRouter struct {
	// Neighbor []*PeerInfo `json:"neighbor" yaml:"neighbor"`
	Neighbor map[string]*PeerInfo `json:"neighbor" yaml:"neighbor"`
}

func GeneratePeer(port, tcpPort uint) (*Peer, error) {
	ipv4Address, err := ip.GetLocalIpv4Address()
	if nil != err {
		return nil, err
	}
	return &Peer{
		Info: &PeerInfo{
			Address: ipv4Address,
			Port:    port,
			TcpPort: tcpPort,
		},
		Router: &PeerRouter{Neighbor: map[string]*PeerInfo{}},
		Tasks:  []*TasksList{},
		alive:  true,
		ctx:    new(context.Context),
	}, nil
}

func (that *Peer) sleep(t time.Duration) {
	time.Sleep(t)
}

func (that *Peer) Join() {
	defer that.Exit()
	// Run Callback Function on Background.
	go that.listenUdp()
	go that.listenTcp()
	time.Sleep(DefaultFirstJoinListenWaitingTime)
	that.join()
	// TODO Add Timer for Refresh.
	that.sleep(time.Second * 2)
	return
}

func (that *Peer) Exit() {
	log.Println(that.Info, "exit.")
	// Notice Neighbor.
	that.exit()
	// Exit.
	that.alive = false
	c := *(that.ctx)
	c.Done()
}
