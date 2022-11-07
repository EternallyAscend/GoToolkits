package DAG

import "time"

// DefaultIP Using 127.0.0.1 for local test network.
const DefaultIP = "127.0.0.1"

const (
	DefaultPort    = 8000
	DefaultTcpPort = 9000
)

const (
	DefaultK                          = 2
	DefaultNeighborRefreshTimeGap     = time.Second // time.Minute
	DefaultFirstJoinListenWaitingTime = 3 * time.Second
)

// Methods

// const MethodJoin = 0
// const UdpMethodRefresh = 1
// const UdpMethodExit = 2
// const MethodReceiveGradient = 3 // Deal Local Training Result Reached.
// const MethodReceiveModel = 4    // Deal Blockchain Training Result Broadcast.
// const MethodCheckModel = 5      // Check Model Training Result.

// TCP Methods

const (
	MethodJoin            = iota
	MethodReceiveGradient // Deal Local Training Result Reached.
	MethodReceiveModel    // Deal Blockchain Training Result Broadcast.
	MethodCheckModel      // Check Model Training Result.
)

// UDP Methods

const (
	UdpMethodRefresh = iota
	UdpMethodExit
)
