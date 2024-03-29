package DAG

import (
	"context"
	"github.com/EternallyAscend/GoToolkits/pkg/network/ip"
	"log"
	"time"
)

// TODO Change to Gossip Cluster https://www.jianshu.com/p/5198b869374a
// Example https://github.com/asim/kv

// DHT Discover
// https://blog.csdn.net/luoye4321/article/details/83587397
// https://blog.csdn.net/u012576116/article/details/81363829

// Timer https://seekload.blog.csdn.net/article/details/113155421

type Peer struct {
	Info   *PeerInfo   `json:"info" yaml:"info"`
	Router *PeerRouter `json:"router" yaml:"router"`
	Tasks  []*Task     `json:"tasks" yaml:"tasks"`
	ctx    context.Context
}

type PeerRouter struct {
	//Neighbor []*PeerInfo `json:"neighbor" yaml:"neighbor"`
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
		Tasks:  []*Task{},
		ctx:    context.Background(),
	}, nil
}

// setBackground [Abandoned] Setting Background Refresh Method and so on.
func (that *Peer) setBackground() {
	ticker := time.NewTicker(DefaultRefreshTime)
	ch := make(chan int)
	go func() {
	loop:
		for {
			// TODO [Unused] Background Functions.
			select {
			case <-that.ctx.Done():
				break loop
			}
		}
		ticker.Stop()
		ch <- 0
	}()
	<-ch
}

func (that *Peer) sleep(t time.Duration) {
	time.Sleep(t)
}

func (that *Peer) addNeighbor(peerInfo *PeerInfo) {
	// Verify if Key is nil before Adding.
	if nil == that.Router.Neighbor[peerInfo.HashString()] {
		that.Router.Neighbor[peerInfo.HashString()] = peerInfo
	}
}

func (that *Peer) Join() {
	defer that.Exit()
	// Run Callback Function on Background.
	go that.listenUdp()
	go that.listenTcp()
	time.Sleep(DefaultFirstJoinListenWaitingTime)
	that.join()
	// Add Timer for Refresh and so on.
	that.fetch()
	return
}

func (that *Peer) Exit() {
	log.Println(that.Info, "exit.")
	// Notice Neighbor.
	that.exit()
	// Exit.
	that.ctx.Done()
}
