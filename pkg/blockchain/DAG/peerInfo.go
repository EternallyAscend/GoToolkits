package DAG

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"strconv"

	"github.com/EternallyAscend/GoToolkits/pkg/cryptography/hash"
	"github.com/EternallyAscend/GoToolkits/pkg/network/tcp"
	"github.com/EternallyAscend/GoToolkits/pkg/network/udp"
)

func GetDefaultPeerInfo() *PeerInfo {
	return &PeerInfo{
		Address: DefaultIP,
		Port:    DefaultPort,
		TcpPort: DefaultTcpPort,
	}
}

type PeerInfo struct {
	Address string `json:"address" yaml:"address"`
	Port    uint   `json:"port" yaml:"port"`
	TcpPort uint   `json:"tcpPort" yaml:"tcpPort"`
}

func (that *PeerInfo) HashString() string {
	// TODO Add Random Id for Peers to Calculate Hash Value.
	return hash.SHA512String([]byte(that.Address + strconv.Itoa(int(that.Port))))
	// return hash.MD5String([]byte(that.Address + strconv.Itoa(int(that.Port))))
}

func UnpackPeerInfo(data []byte) *PeerInfo {
	p := &PeerInfo{}
	err := json.Unmarshal(data, p)
	if nil != err {
		log.Println("Unmarshal peerInfo failed,", err)
		return nil
	}
	return p
}

func UnpackPeerInfoList(data []byte) map[string]*PeerInfo {
	var pList map[string]*PeerInfo
	err := json.Unmarshal(data, &pList)
	if nil != err {
		log.Println(string(data))
		log.Println("Unmarshal peerInfo list failed,", err)
		return nil
	}
	return pList
}

func (that *PeerInfo) UdpListen(f func([]byte), ctx *context.Context) {
	// TODO Using context for interrupting.
	udp.ListenViaUdp4(f, that.Port)
}

func (that *PeerInfo) UdpSendToPeer(data []byte) error {
	return udp.SendViaUdp4(that.Address, that.Port, data)
}

func (that *PeerInfo) TcpCommunicateWithPeer(data []byte) {
	tcp.RequestViaTcp4(that.Address, that.Port, SenderTcpFunc, data)
}

func (that *PeerInfo) TcpListen(f func(conn *net.Conn)) {
	// TODO Using context for interrupting.
	tcp.ListenViaTcp4(f, that.TcpPort)
}
