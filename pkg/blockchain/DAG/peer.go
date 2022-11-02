package DAG

import (
	"encoding/json"
	"errors"
	"github.com/EternallyAscend/GoToolkits/pkg/network/ip"
	"github.com/EternallyAscend/GoToolkits/pkg/network/udp"
	"log"
	"time"
)

// TODO Change to Gossip Cluster https://www.jianshu.com/p/5198b869374a
// Example https://github.com/asim/kv

// DHT Discover
// https://blog.csdn.net/luoye4321/article/details/83587397
// https://blog.csdn.net/u012576116/article/details/81363829

// Timer https://seekload.blog.csdn.net/article/details/113155421

const DefaultIP = "192.168.1.1"
const DefaultPort = 8000
const DefaultTcpPort = 9000

type Package struct {
	Type int `json:"type" yaml:"type"`
	Message []byte `json:"message" yaml:"message"`
}

func UnpackPackage(data []byte) *Package {
	p := &Package{}
	err := json.Unmarshal(data, p)
	if nil != err {
		log.Println(err)
		return nil
	}
	return p
}

type Peer struct {
	Info *PeerInfo `json:"info" yaml:"info"`
	Router *PeerRouter `json:"router" yaml:"router"`
	Tasks *TasksList `json:"tasks" yaml:"tasks"`
}

type PeerInfo struct {
	Address string `json:"address" yaml:"address"`
	Port uint `json:"port" yaml:"port"`
	TcpPort uint `json:"tcpPort" yaml:"tcpPort"`
}

func UnpackPeerInfo(data []byte) *PeerInfo {
	p := &PeerInfo{}
	err := json.Unmarshal(data, p)
	if nil != err {
		log.Println(err)
		return nil
	}
	return p
}

func (that *PeerInfo) SendToPeerByUdp(data []byte) error {
	return udp.SendViaUdp4(that.Address, that.Port, data)
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

func StartOrigin() {
	peer, err := GeneratePeer(DefaultPort)
	if nil != err {
		log.Println(err)
		return
	}
	udp.ListenViaUdp4(func(data []byte) {
		p := UnpackPackage(data)
		if nil == p {
			return
		}
		switch p.Type {
		case 0: // Add Peer Into Network.
			peerInfo := UnpackPeerInfo(p.Message)
			if nil == peerInfo {
				return
			} else {
				// Send back Neighbor.

				// Add Neighbor.
				peer.Router.Neighbor = append(peer.Router.Neighbor, peerInfo)
			}
			break
		}
	}, 8000)
}

func (that *Peer) listen() {

}

func (that *Peer) Join() {
	// Request Default Address
	defaultPeer := PeerInfo{
		Address: DefaultIP,
		Port:    DefaultPort,
		TcpPort: DefaultTcpPort,
	}
	peerInfo, err := json.Marshal(that.Info)
	if nil != err {
		log.Println(err)
		return
	}
	err = defaultPeer.SendToPeerByUdp(peerInfo)
	if err != nil {
		log.Println(err)
		return
	}
	// TODO fetch Neighbor
	// TODO set Timer (Background Tasks)
	return
}

func (that *Peer) fetch() *Peer {
	// TODO Calculate by Hash
	return that
}

func (that *Peer) Refresh() *Peer {
	that.fetch()
	return that
}

//func (that *Peer) ReleaseTask() *Peer {
//	return that
//}
//
//func (that *Peer) BroadcastTaskResult() *Peer {
//	return that
//}

func (that *Peer) ReleaseModel() *Peer {
	return that
}

func (that *Peer) VerifyModel() *Peer {
	return that
}

func (that *Peer) Exit() error {
	// TODO Notice Neighbor
	return errors.New("")
}

func (that *Peer) Execute() error {
	return nil
}

