package DAG

import (
	"encoding/json"
	"log"
	"time"
)

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

// Methods: Refresh and Exit. Not Reliable.

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
		case TcpMethodReceiveModel:
			// TODO Verify Received Model in Other Process.
			that.receiveModel().verifyModel()
			break
		case TcpMethodCheckModel:
			// TODO Try Model Correction.
			that.Try()
			break
		default:
			break
		}
	}, that.ctx)
}

func (that *Peer) join() {
	// Request Default Address.
	defaultPeer := GetDefaultPeerInfo()
	peerInfo, err := json.Marshal(that.Info)
	if nil != err {
		log.Println("Marshal local peerInfo for join network failed,", err)
		return
	}
	p := &Package{
		Type:    TcpMethodJoin,
		Message: peerInfo,
	}
	pack, err := json.Marshal(p)
	if nil != err {
		log.Println("Marshal local package for join network failed,", err)
		return
	}
	// TODO Change UDP to TCP Connection.
	err = defaultPeer.UdpSendToPeer(pack)
	if err != nil {
		log.Println("Send local peerInfo to origin failed,", err)
		return
	}
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

func (that *Peer) exit() {
	// Notice Neighbor.
	d, err := json.Marshal(that.Info)
	if nil != err {
		log.Println(err)
		return
	}
	p := &Package{
		Type:    UdpMethodExit,
		Message: d,
	}
	pack, err := json.Marshal(p)
	if nil != err {
		log.Println(err)
		return
	}
	that.UdpBroadcast(pack)
}
