package DAG

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestStartOrigin(t *testing.T) {
	StartOrigin()
}

func RunTestPeerLocal(t *testing.T, port, tcpPort uint) {
	peer, err := GeneratePeer(port, tcpPort)
	if nil != err {
		log.Println(err)
		t.Fail()
		return
	}
	peer.Join()
	// time.Sleep(time.Second * 5)
	for k, v := range peer.Router.Neighbor {
		fmt.Println(k, v)
	}
	return
}

func TestPeerCluster_Join(t *testing.T) {
	udpBase := 8000
	tcpBase := 9000
	for i := 1; i < 10; i++ {
		RunTestPeerLocal(t, uint(udpBase+i), uint(tcpBase+i))
	}
	time.Sleep(time.Second * 5)
	RunTestPeerLocal(t, 8010, 9010)
}

func TestPeer1_Join(t *testing.T) {
	RunTestPeerLocal(t, 8001, 9001)
}

func TestPeer2_Join(t *testing.T) {
	RunTestPeerLocal(t, 8002, 9002)
}

func TestPeer3_Join(t *testing.T) {
	RunTestPeerLocal(t, 8003, 9003)
}

func TestPeer4_Join(t *testing.T) {
	RunTestPeerLocal(t, 8004, 9004)
}

func TestPeer5_Join(t *testing.T) {
	RunTestPeerLocal(t, 8005, 9005)
}
