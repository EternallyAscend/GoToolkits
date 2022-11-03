package DAG

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/EternallyAscend/GoToolkits/pkg/cryptography/hash"
	"github.com/EternallyAscend/GoToolkits/pkg/network/ip"
	"github.com/EternallyAscend/GoToolkits/pkg/network/udp"
)

// TODO Change to Gossip Cluster https://www.jianshu.com/p/5198b869374a
// Example https://github.com/asim/kv

// DHT Discover
// https://blog.csdn.net/luoye4321/article/details/83587397
// https://blog.csdn.net/u012576116/article/details/81363829

// Timer https://seekload.blog.csdn.net/article/details/113155421

const (
	MethodJoin            = 0
	MethodRefresh         = 1
	MethodExit            = 2
	MethodReceiveGradient = 3 // Deal Local Training Result Reached.
	MethodReceiveModel    = 4 // Deal Blockchain Training Result Broadcast.
	MethodCheckModel      = 5 // Check Model Training Result.
)

const DefaultTimeGap = time.Second // time.Minute

const (
	DefaultIP      = "192.168.1.1"
	DefaultPort    = 8000
	DefaultTcpPort = 9000
)

func GetDefaultPeerInfo() *PeerInfo {
	return &PeerInfo{
		Address: DefaultIP,
		Port:    DefaultPort,
		TcpPort: DefaultTcpPort,
	}
}

type Package struct {
	Type    uint   `json:"type" yaml:"type"`
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
	Info   *PeerInfo   `json:"info" yaml:"info"`
	Router *PeerRouter `json:"router" yaml:"router"`
	Tasks  *TasksList  `json:"tasks" yaml:"tasks"`
}

type PeerInfo struct {
	Address string `json:"address" yaml:"address"`
	Port    uint   `json:"port" yaml:"port"`
	TcpPort uint   `json:"tcpPort" yaml:"tcpPort"`
}

func (that *PeerInfo) HashString() string {
	// TODO Add Random Id for Peers to Calculate Hash Value.
	return hash.SHA512String([]byte(that.Address))
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

func UnpackPeerInfoList(data []byte) []*PeerInfo {
	var pList []*PeerInfo
	err := json.Unmarshal(data, &pList)
	if nil != err {
		log.Println(err)
		return nil
	}
	return pList
}

func (that *PeerInfo) UdpListen(f func([]byte)) {
	udp.ListenViaUdp4(f, that.Port)
}

func (that *PeerInfo) UdpSendToPeer(data []byte) error {
	return udp.SendViaUdp4(that.Address, that.Port, data)
}

type PeerRouter struct {
	// Neighbor []*PeerInfo `json:"neighbor" yaml:"neighbor"`
	Neighbor map[string]*PeerInfo `json:"neighbor" yaml:"neighbor"`
}

type TasksList struct {
	Command   string      `json:"command" yaml:"command"`
	Timestamp time.Time   `json:"timestamp" yaml:"timestamp"`
	Reached   []*PeerInfo `json:"reached" yaml:"reached"`
}

type Header struct{}

// PeerStorage DAG Data Structure.
type PeerStorage struct {
	Header *Header `json:"header" yaml:"header"`
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
		case MethodJoin: // Add Peer Into Network.
			peerInfo := UnpackPeerInfo(p.Message)
			if nil == peerInfo {
				return
			} else {
				// Send back Neighbor.
				d, errIn := json.Marshal(peer.Router.Neighbor)
				if nil != errIn {
					log.Println(errIn)
					return
				}
				pack, errIn := json.Marshal(Package{
					Type:    MethodRefresh,
					Message: d,
				})
				err = peerInfo.UdpSendToPeer(pack)
				if nil != err {
					log.Println(err)
					return
				}
				// Add Neighbor.
				// peer.Router.Neighbor = append(peer.Router.Neighbor, peerInfo)
				peer.Router.Neighbor[peerInfo.HashString()] = peerInfo
			}
			break
		case MethodExit:
			neighbor := UnpackPeerInfo(p.Message)
			if nil == neighbor {
				break
			}
			delete(peer.Router.Neighbor, neighbor.HashString())
			break
		}
	}, 8000)
}

// listen Background Information.
func (that *Peer) listen() {
	that.Info.UdpListen(func(data []byte) {
		p := UnpackPackage(data)
		if nil == p {
			return
		}
		switch p.Type {
		case MethodRefresh: // Receive Neighbor PeerInfo from Origin.
			neighbor := UnpackPeerInfoList(p.Message)
			if nil == neighbor {
				break
			}
			if nil != neighbor {
				for i := 0; i < len(neighbor); i++ {
					if nil == that.Router.Neighbor[neighbor[i].HashString()] {
						that.Router.Neighbor[neighbor[i].HashString()] = neighbor[i]
					}
				}
			}
			delete(that.Router.Neighbor, that.Info.HashString())
			break
		case MethodExit: // Remove Neighbor Exited.
			neighbor := UnpackPeerInfo(p.Message)
			if nil == neighbor {
				break
			}
			delete(that.Router.Neighbor, neighbor.HashString())
			break
		}
	})
}

func (that *Peer) Join() {
	// Run Callback Function on Background.
	go that.listen()
	// Request Default Address.
	defaultPeer := GetDefaultPeerInfo()
	peerInfo, err := json.Marshal(that.Info)
	if nil != err {
		log.Println(err)
		return
	}
	err = defaultPeer.UdpSendToPeer(peerInfo)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (that *Peer) fetch() {
	length := len(that.Router.Neighbor)
	p := &Package{
		Type:    MethodRefresh,
		Message: nil,
	}
	pack, err := json.Marshal(p)
	if nil != err {
		log.Println(err)
		return
	}
	for true {
		if 0 == length {
			defaultPeer := GetDefaultPeerInfo()
			err = defaultPeer.UdpSendToPeer(pack)
			if nil != err {
				log.Println(err)
			}
		} else {
			for _, v := range that.Router.Neighbor {
				if v.Address == that.Info.Address {
					continue
				}
				err = v.UdpSendToPeer(pack)
				if nil != err {
					// TODO Timeout and Remove Neighbor.
					log.Println(err)
				}
				time.Sleep(DefaultTimeGap)
			}
		}
		time.Sleep(DefaultTimeGap)
	}
	return
}

//func (that *Peer) Refresh() *Peer {
//	that.fetch()
//	return that
//}

func (that *Peer) UdpBroadcast(data []byte) {
	for _, v := range that.Router.Neighbor {
		if v.Address == that.Info.Address {
			continue
		}
		err := v.UdpSendToPeer(data)
		if nil != err {
			log.Println(err)
		}
	}
}

// readGradient TODO Read from File or Transfer between Processes.
func (that *Peer) readGradient(path string) []float64 {
	return []float64{}
}

func (that *Peer) aggregateGradient() *Peer {
	// TODO Accumulative to Threshold. Max Circle or Gradient Size.
	return that
}

func (that *Peer) ReleaseModel() *Peer {
	// TODO Calculate Timestamp 'Start'.
	// TODO HE Calculate.
	// TODO Calculate Timestamp 'End'.
	// TODO Broadcast via TCP.
	// TODO 1.Send Id, 2. Confirm Not Reached, 3. Transfer.
	return that
}

func (that *Peer) receiveModel() *Peer {
	// TODO Calculate Timestamp 'Start'.
	// TODO HE Verify.
	// TODO Calculate Timestamp 'End'.
	return that
}

func (that *Peer) verifyModel() *Peer {
	// TODO Communicate between Processes.
	return that
}

// Try Train with Test Data and Get Result.
func (that *Peer) Try() {
	// TODO Communicate between Processes.
}

func (that *Peer) Exit() error {
	// Notice Neighbor.
	d, err := json.Marshal(that.Info)
	if nil != err {
		log.Println(err)
		return err
	}
	p := &Package{
		Type:    MethodExit,
		Message: d,
	}
	pack, err := json.Marshal(p)
	if nil != err {
		log.Println(err)
		return err
	}
	that.UdpBroadcast(pack)
	// Exit.
	os.Exit(0)
	return nil
}
