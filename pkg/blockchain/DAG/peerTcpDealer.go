package DAG

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"github.com/EternallyAscend/GoToolkits/pkg/network/tcp"
	"log"
	"net"
	"time"
)

func EncodeTcpMessage(data []byte) []byte {
	length := int64(len(data))
	pkg := new(bytes.Buffer)
	err := binary.Write(pkg, binary.BigEndian, length)
	if err != nil {
		log.Println(err)
		return nil
	}
	err = binary.Write(pkg, binary.BigEndian, data)
	if err != nil {
		log.Println(err)
		return nil
	}
	return pkg.Bytes()
}

func (that *Peer) listenTcp() {
	tcp.ListenInterruptableViaTcp4(that.ctx, func(conn *net.Conn) {
		defer (*conn).Close()
		// Reader.
		reader := bufio.NewReader(*conn)
		peek, err := reader.Peek(DefaultTcpLength)
		if nil != err {
			log.Println(err)
			return
		}

		// TODO Peer Reliable Connection.
		buffer := bytes.NewBuffer(peek)
		var length int64
		binary.Read(buffer, binary.BigEndian, &length)
		for int64(reader.Buffered()) < length+DefaultTcpLength {
			// TODO Last 存在优化可能
			time.Sleep(time.Second)
		}
		packByte := make([]byte, length+DefaultTcpLength)
		_, err = reader.Read(packByte)
		if nil != err {
			log.Println(err)
			return
		}
		// Unpack Package.
		pack := UnpackPackage(packByte[DefaultTcpLength:])
		// TODO Listen for Services.
		switch pack.Type {
		case TcpMethodJoin:
			// Unpack Peer Information.
			peerInfo := UnpackPeerInfo(pack.Message)
			if nil == peerInfo {
				return
			}
			that.addNeighbor(peerInfo)
			// Send Neighbor Information Back.
			peerListByte, errIn := json.Marshal(that.Router.Neighbor)
			if nil != errIn {
				log.Println(errIn)
				return
			}
			//(*conn).Write(peerListByte)
			// Udp Transfer
			udpPack := &Package{
				Type: UdpMethodReceive,
				//Length:  uint(len(peerListByte)),
				Message: peerListByte,
			}
			udpPackByte, errIn := json.Marshal(udpPack)
			if nil != errIn {
				log.Println(errIn)
				return
			}
			go peerInfo.UdpSendToPeer(udpPackByte)
			// Add Peer to Neighbor List.
			that.addNeighbor(peerInfo)
			break
		case TcpMethodExchangeGH:
			// TODO Next.
			break
		case TcpMethodReleaseGradient: // Receive Gradient Trained by Other Process.
			// TODO File Path or Transfer Directly.
			that.readGradient("")
			// TODO Aggregate Gradient.
			that.aggregateGradient().ReleaseModel()
			break
		case TcpMethodReleaseModel:
			// TODO Verify Received Model in Other Process.
			that.receiveModel().verifyModel()
			break
		case TcpMethodGetModel:
			break
		case TcpMethodGetModelScore:
			break
		case TcpMethodCheckModel:
			// TODO Try Model Correction.
			that.Try()
			break
		}
	}, that.Info.TcpPort)

}

func (that *Peer) TcpBroadcast(method uint, data []byte) {
	for _, v := range that.Router.Neighbor {
		v.TcpCommunicateWithPeer(method, data)
	}
}

func SenderTcpFunc(conn *net.Conn, method uint, data []byte) {
	connection := *conn
	header := &PackageTcpHeader{
		Type:   method,
		Length: uint(len(data)),
	}
	headerByte, err := json.Marshal(header)
	if nil != err {
		log.Println(err)
		return
	}
	connection.Write(headerByte)
	connection.Write(data)
}

func (that *Peer) releaseModel() {}
