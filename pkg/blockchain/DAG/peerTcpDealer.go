package DAG

import (
	"github.com/EternallyAscend/GoToolkits/pkg/network/tcp"
	"net"
)

func (that *Peer) listenTcp() {
	tcp.ListenInterruptableViaTcp4(that.ctx, func(conn *net.Conn) {
		defer (*conn).Close()
		// TODO Peer Reliable Connection.
		headerByte := make([]byte, DefaultPackageTcpHeaderSize)
		(*conn).Read(headerByte)
		header := UnpackPackageTcpHeader(headerByte)
		if nil == header {
			return
		}
		pack := make([]byte, header.Length)
		(*conn).Read(pack)
		// TODO Listen for Services.
		switch header.Type {
		case TcpMethodJoin:
			// TODO Send Neighbor Information Back.
			peerInfo := UnpackPeerInfo(pack)
			if nil == peerInfo {
				return
			}
			// Send.

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

func (that *Peer) TcpBroadcast() {
	for _, v := range that.Router.Neighbor {
		v.TcpCommunicateWithPeer(nil)
	}
}

func SenderTcpFunc(conn *net.Conn, data []byte) {
	connection := *conn
	connection.Read([]byte{})
}

func (that *Peer) releaseModel() {}
