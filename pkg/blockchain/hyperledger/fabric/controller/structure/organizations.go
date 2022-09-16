package structure

import "github.com/EternallyAscend/GoToolkits/pkg/blockchain/hyperledger/fabric/controller/structure/configtx"

const (
	OrdererPeer    = iota
	AnchorPeer     = iota
	CAPeer         = iota
	LeaderPeer     = iota
	CommittingPeer = iota
	EndorsingPeer  = iota
)

type Peer struct {
	Role []int  `yaml:"role" json:"role"`
	Host string `yaml:"host" json:"host"`
	Port uint   `yaml:"port" json:"port"`
}

type Organization struct {
	Name     string                         `yaml:"name" json:"name"`
	ID       string                         `yaml:"id" json:"id"`
	Domain   string                         `yaml:"domain" json:"domain"`
	MSPDir   string                         `yaml:"mspDir" json:"mspDir"`
	Peers    []*Peer                        `yaml:"peers" json:"peers"`
	Policies *configtx.OrganizationPolicies `yaml:"policies" json:"policies"`
	Channels []string                       `yaml:"channels" json:"channels"`
	channels []*Channel
}

func GenerateEmptyOrganization(name string, id string, domain string, mspDir string) *Organization {
	return &Organization{
		Name:     name,
		ID:       id,
		Domain:   domain,
		MSPDir:   mspDir,
		Peers:    []*Peer{},
		Policies: nil,
		Channels: []string{},
		channels: []*Channel{},
	}
}
