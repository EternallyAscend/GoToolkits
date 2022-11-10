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
		case UdpMethodRefresh:
			neighbor := UnpackPeerInfo(p.Message)
			if nil == neighbor {
				break
			}
			that.addNeighbor(neighbor)
			listByte, err := json.Marshal(that.Router.Neighbor)
			if nil != err {
				log.Println(err)
				return
			}
			pack := &Package{
				Type:    UdpMethodReceive,
				Message: listByte,
			}
			packByte, err := json.Marshal(pack)
			if nil != err {
				log.Println(err)
				return
			}
			go neighbor.UdpSendToPeer(packByte)
			break
		case UdpMethodReceive:
			neighbor := UnpackPeerInfoList(p.Message)
			if nil == neighbor {
				break
			}
			for k, v := range neighbor {
				if nil == that.Router.Neighbor[k] {
					that.Router.Neighbor[k] = v
				}
			}
			//delete(that.Router.Neighbor, that.Info.HashString())
			break
		case UdpMethodExit:
			neighbor := UnpackPeerInfo(p.Message)
			if nil == neighbor {
				break
			}
			delete(that.Router.Neighbor, neighbor.HashString())
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
	// header := PackageTcpHeader{
	// 	Type:   TcpMethodJoin,
	// 	Length: uint(len(peerInfo)),
	// }
	// pack, err := json.Marshal(header)
	// if nil != err {
	// 	log.Println("Marshal local package for join network failed,", err)
	// 	return
	// }
	// err = defaultPeer.UdpSendToPeer(pack)
	// if err != nil {
	// 	log.Println("Send local peerInfo to origin failed,", err)
	// 	return
	// }
	//
	// TCP Connection.
	defaultPeer.TcpCommunicateWithPeer(TcpMethodJoin, peerInfo)
}

func (that *Peer) fetch() {
	length := len(that.Router.Neighbor)
	peerInfoByte, err := json.Marshal(that.Info)
	if nil != err {
		log.Println(err)
		return
	}
	p := &Package{
		Type:    UdpMethodRefresh,
		Message: peerInfoByte,
	}
	pack, err := json.Marshal(p)
	if nil != err {
		log.Println(err)
		return
	}
loop:
	for {
		if 0 == length {
			log.Println("zero")
			defaultPeer := GetDefaultPeerInfo()
			err = defaultPeer.UdpSendToPeer(pack)
			if nil != err {
				log.Println(err)
				continue
			}
		} else {
			for _, v := range that.Router.Neighbor {
				if v.Address == that.Info.Address && v.Port == that.Info.Port {
					continue
				}
				log.Println(v)
				err = v.UdpSendToPeer(pack)
				if nil != err {
					// TODO Timeout and Remove Neighbor.
					log.Println(err)
					continue
				}
				// Stop this Process.
				select {
				case <-that.ctx.Done():
					break loop
				default:
					time.Sleep(DefaultNeighborRefreshTimeGap)
				}
			}
		}
		// Stop this Process.
		select {
		case <-that.ctx.Done():
			break loop
		default:
			time.Sleep(DefaultNeighborRefreshTimeGap)
			length = len(that.Router.Neighbor)
		}
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
