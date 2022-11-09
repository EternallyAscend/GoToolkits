package DAG

import (
	"github.com/EternallyAscend/GoToolkits/pkg/network/tcp"
	"net"
)

func (that *Peer) TcpFunc() {}

func (that *Peer) listenTcp() {
	tcp.ListenInterruptableViaTcp4(that.ctx, func(conn *net.Conn) {
		// TODO Peer Reliable Connection.
		headerByte := make([]byte, DefaultPackageTcpHeaderSize)
		(*conn).Read(headerByte)
		header := UnpackPackageTcpHeader(headerByte)
		if nil == header {
			return
		}
		// TODO Listen for Services.
		switch header.Type {
		case TcpMethodJoin:
			break
		case TcpMethodReceiveModel:
			break
		case TcpMethodCheckModel:
			break
		}

		/*
			loop:
				for {
					// Header.
					headerByte := make([]byte, DefaultPackageTcpHeaderSize)
					(*conn).Read(headerByte)
					header := UnpackPackageTcpHeader(headerByte)
					if nil == header {
						(*conn).Close()
						continue
					}
					c := *that.ctx
					select {
					case <-c.Done():
						break loop
					}
				}
		*/
	}, that.Info.TcpPort)

}

func (that *Peer) TcpBroadcast() {
	for _, v := range that.Router.Neighbor {
		v.TcpSendToPeer(nil)
	}
}

func ServerTcpFunc(conn *net.Conn) {
	connection := *conn
	connection.Read([]byte{})
}

func (that *Peer) releaseModel() {}
