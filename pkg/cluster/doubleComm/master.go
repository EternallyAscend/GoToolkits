package commonCommunication

type Master struct {
	IP    string `json:"ip"`
	Port  int    `json:"port"`
	Info  string `json:"info"`
	Peers []Peer `json:"peers"`
	Puk   string `json:"puk"`
	Prk   string `json:"prk"`
}

type Peer struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
	Info string `json:"info"`
	Key  string `json:"key"`
}

func CreateCluster(port int, limit int, crypto bool) (Master, error) {
	return Master{}, nil
}

func (this Master) Listen() error {
	defer func(this Master) {
		_ = this.CloseCluster()
	}(this)
	return nil
}

func (this Master) CloseCluster() error {
	return nil
}

func (this Master) GetPublicKey() string {
	return this.Puk
}

func (this Master) GetPrivateKey() string {
	return this.Prk
}

func (this Master) FetchPeerList() []Peer {
	return this.Peers
}

func (this Master) Broadcast(info string) error {
	return nil
}

func (this Master) SendInfoToPeer(peer Peer, info string) error {
	return nil
}
