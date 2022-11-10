package DAG

import "time"

// DefaultIP Using 127.0.0.1 for local test network.
const DefaultIP = "127.0.0.1"

const (
	DefaultPort    = 8002
	DefaultTcpPort = 9000
)

const (
	DefaultK                          = 2
	DefaultNeighborRefreshTimeGap     = time.Second // time.Minute
	DefaultFirstJoinListenWaitingTime = 3 * time.Second
	DefaultRefreshTime                = 5 * time.Second
)

const DefaultTcpLength = 8 // int64

// Methods

// const TcpMethodJoin = 0
// const UdpMethodRefresh = 1
// const UdpMethodExit = 2
// const MethodReceiveGradient = 3 // Deal Local Training Result Reached.
// const TcpMethodReceiveModel = 4    // Deal Blockchain Training Result Broadcast.
// const TcpMethodCheckModel = 5      // Check Model Training Result.

// TCP Methods

const (
	TcpMethodJoin         = iota
	MethodReceiveGradient // Deal Local Training Result Reached.
	TcpMethodReceiveModel // Deal Blockchain Training Result Broadcast.
	TcpMethodCheckModel   // Check Model Training Result.
	TcpMethodReleaseGradient
	TcpMethodReleaseModel
	TcpMethodGetModel
	TcpMethodGetModelScore
	TcpMethodExchangeGH
)

// UDP Methods

const (
	UdpMethodRefresh = iota // Request Neighbor PeerInfo.
	UdpMethodReceive        // Receive Neighbor PeerInfo from other Peers.
	UdpMethodExit           // Remove Neighbor Exited.
)
