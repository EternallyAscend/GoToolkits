package DAG

import (
	"encoding/json"
	"github.com/EternallyAscend/GoToolkits/pkg/network/udp"
	"log"
)

func StartOrigin() {
	peer, err := GeneratePeer(DefaultPort, DefaultTcpPort)
	if nil != err {
		log.Println("Origin peer start failed,", err)
		return
	}
	// TODO Judge Genesis Block.
	// If genesis block is existed this process will read and continue, otherwise create.
	go OriginListenUdp4(peer)
	peer.listenTcp()
}

func OriginListenUdp4(peer *Peer) {
	go udp.ListenViaUdp4(func(data []byte) {
		p := UnpackPackage(data)
		if nil == p {
			return
		}
		switch p.Type {
		case UdpMethodRefresh:
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
					Type:    UdpMethodReceive,
					Message: d,
				})
				go peerInfo.UdpSendToPeer(pack)
				// Add Neighbor.
				//peer.Router.Neighbor = append(peer.Router.Neighbor, peerInfo)
				peer.addNeighbor(peerInfo)
				//for _, v := range peer.Router.Neighbor {
				//	log.Println(v)
				//}
				log.Println(peerInfo, "ask for blockchain network.")
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
	}, DefaultPort)
}
