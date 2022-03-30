package commonCommunication

func JoinCluster(ip string, port int, key string) (Master, Peer, error) {
	return Master{}, Peer{}, nil
}

func (this Peer) connectCluster() {
}

func (this Peer) exitCluster() error {
	return nil
}
