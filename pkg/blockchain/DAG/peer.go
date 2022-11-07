package DAG

import (
	"encoding/json"
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

type Peer struct {
	Info   *PeerInfo    `json:"info" yaml:"info"`
	Router *PeerRouter  `json:"router" yaml:"router"`
	Tasks  []*TasksList `json:"tasks" yaml:"tasks"`
	alive  bool
}

type PeerRouter struct {
	//Neighbor []*PeerInfo `json:"neighbor" yaml:"neighbor"`
	Neighbor map[string]*PeerInfo `json:"neighbor" yaml:"neighbor"`
}

type TasksList struct {
	Command   string      `json:"command" yaml:"command"`
	Timestamp time.Time   `json:"timestamp" yaml:"timestamp"`
	Reached   []*PeerInfo `json:"reached" yaml:"reached"`
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
	}, nil
}

func StartOrigin() {
	peer, err := GeneratePeer(DefaultPort, DefaultTcpPort)
	if nil != err {
		log.Println("Origin peer start failed,", err)
		return
	}
	// TODO Judge Genesis Block.
	// If genesis block is existed this process will read and continue, otherwise create.
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
					log.Println("Method join send neighbor back failed,", errIn)
					return
				}
				pack, errIn := json.Marshal(Package{
					Type:    UdpMethodRefresh,
					Message: d,
				})
				err = peerInfo.UdpSendToPeer(pack)
				if nil != err {
					log.Println("Udp send failed,", err)
					return
				}
				// Add Neighbor.
				//peer.Router.Neighbor = append(peer.Router.Neighbor, peerInfo)
				peer.Router.Neighbor[peerInfo.HashString()] = peerInfo
				for _, v := range peer.Router.Neighbor {
					log.Println(v)
				}
				log.Println(peerInfo, "join blockchain network.")
			}
			break
		case UdpMethodExit:
			neighbor := UnpackPeerInfo(p.Message)
			if nil == neighbor {
				break
			}
			delete(peer.Router.Neighbor, neighbor.HashString())
			log.Println(neighbor, "exit blockchain network.")
			break
		default:
			break
		}
	}, 8000)
}

// listenUdp Background Information.
func (that *Peer) listenUdp() {
	that.Info.UdpListen(func(data []byte) {
		p := UnpackPackage(data)
		if nil == p {
			return
		}
		switch p.Type {
		case UdpMethodRefresh: // Receive Neighbor PeerInfo from Origin.
			neighbor := UnpackPeerInfoList(p.Message)
			if nil == neighbor {
				break
			}
			if nil != neighbor {
				for k, v := range neighbor {
					if nil == that.Router.Neighbor[k] {
						that.Router.Neighbor[k] = v
					}
				}
			}
			delete(that.Router.Neighbor, that.Info.HashString())
			break
		case UdpMethodExit: // Remove Neighbor Exited.
			neighbor := UnpackPeerInfo(p.Message)
			if nil == neighbor {
				break
			}
			delete(that.Router.Neighbor, neighbor.HashString())
			break
			/*
				case MethodReceiveGradient: // Receive Gradient Trained by Other Process.
					// TODO File Path or Transfer Directly.
					that.readGradient("")
					// TODO Aggregate Gradient.
					that.aggregateGradient().ReleaseModel()
					break
			*/
		case MethodReceiveModel:
			// TODO Verify Received Model in Other Process.
			that.receiveModel().verifyModel()
			break
		case MethodCheckModel:
			// TODO Try Model Correction.
			that.Try()
			break
		default:
			break
		}
	})
}

func (that *Peer) listenTcp() {

}

func (that *Peer) sleep(t time.Duration) {
	time.Sleep(t)
}

func (that *Peer) Join() {
	// Run Callback Function on Background.
	go that.listenUdp()
	time.Sleep(DefaultFirstJoinListenWaitingTime)
	// Request Default Address.
	defaultPeer := GetDefaultPeerInfo()
	peerInfo, err := json.Marshal(that.Info)
	if nil != err {
		log.Println("Marshal local peerInfo for join network failed,", err)
		return
	}
	p := &Package{
		Type:    0,
		Message: peerInfo,
	}
	pack, err := json.Marshal(p)
	if nil != err {
		log.Println("Marshal local package for join network failed,", err)
		return
	}
	err = defaultPeer.UdpSendToPeer(pack)
	if err != nil {
		log.Println("Send local peerInfo to origin failed,", err)
		return
	}
	defer that.Exit()
	that.sleep(time.Second * 2)
	that.Exit()
	return
}

func (that *Peer) fetch() {
	length := len(that.Router.Neighbor)
	p := &Package{
		Type:    UdpMethodRefresh,
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
				time.Sleep(DefaultNeighborRefreshTimeGap)
			}
		}
		time.Sleep(DefaultNeighborRefreshTimeGap)
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
	// TODO Accumulative to Threshold. Max Circle, Max Time or Gradient Size.
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
	that.verifyModel()
}

func (that *Peer) Exit() error {
	log.Println(that.Info, "exit.")
	// Notice Neighbor.
	d, err := json.Marshal(that.Info)
	if nil != err {
		log.Println(err)
		return err
	}
	p := &Package{
		Type:    UdpMethodExit,
		Message: d,
	}
	pack, err := json.Marshal(p)
	if nil != err {
		log.Println(err)
		return err
	}
	that.UdpBroadcast(pack)
	// Exit.
	//os.Exit(0)
	return nil
}
