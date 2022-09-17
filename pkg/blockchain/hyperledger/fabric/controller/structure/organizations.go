package structure

import "github.com/EternallyAscend/GoToolkits/pkg/blockchain/hyperledger/fabric/controller/config/configtx"

const (
	OrdererPeer    = iota
	AnchorPeer     = iota
	CAPeer         = iota
	LeaderPeer     = iota
	CommittingPeer = iota
	EndorsingPeer  = iota
)

type Ca struct {
	PeerName   string `yaml:"peerName" json:"peerName"`
	OrgName    string `yaml:"orgName" json:"orgName"`
	DomainRoot string `yaml:"domainRoot" json:"domainRoot"`
	Port       uint   `yaml:"port" json:"port"`
	GrpcPort   uint   `yaml:"grpcPort" json:"grpcPort"`
}

type Orderer struct {
	PeerName   string `yaml:"peerName" json:"peerName"`
	OrgName    string `yaml:"orgName" json:"orgName"`
	DomainRoot string `yaml:"domainRoot" json:"domainRoot"`
	Port       uint   `yaml:"port" json:"port"`
}

type Peer struct {
	Role       []int  `yaml:"role" json:"role"`
	PeerName   string `yaml:"peerName" json:"peerName"`
	OrgName    string `yaml:"orgName" json:"orgName"`
	DomainRoot string `yaml:"domainRoot" json:"domainRoot"`
	Port       uint   `yaml:"port" json:"port"`
}

type Organization struct {
	Name     string                         `yaml:"name" json:"name"`
	ID       string                         `yaml:"id" json:"id"`
	Domain   string                         `yaml:"domain" json:"domain"`
	MSPDir   string                         `yaml:"mspDir" json:"mspDir"`
	Cas      []*Ca                          `yaml:"Cas" json:"Cas"`
	Peers    []*Peer                        `yaml:"peers" json:"peers"`
	Orderers []*Orderer                     `yaml:"orderers" json:"orderers"`
	Policies *configtx.OrganizationPolicies `yaml:"policies" json:"policies"`
	Channels []string                       `yaml:"channels" json:"channels"`
	channels []*Channel
}

func GenerateEmptyOrganization(orgName string, id string, domainRoot string, mspDir string) *Organization {
	return &Organization{
		Name:     orgName,
		ID:       id,
		Domain:   domainRoot,
		MSPDir:   mspDir,
		Peers:    []*Peer{},
		Cas:      []*Ca{},
		Orderers: []*Orderer{},
		Policies: nil,
		Channels: []string{},
		channels: []*Channel{},
	}
}

func (that *Organization) AddOrderer(peerName string, orgName string, domainRoot string, port uint) {
	that.Orderers = append(that.Orderers, &Orderer{
		PeerName:   peerName,
		OrgName:    orgName,
		DomainRoot: domainRoot,
		Port:       port,
	})
}

func (that *Organization) AddPeer(peerName string, orgName string, domainRoot string, port uint) {
	that.Peers = append(that.Peers, &Peer{
		Role:       []int{},
		PeerName:   peerName,
		OrgName:    orgName,
		DomainRoot: domainRoot,
		Port:       port,
	})
	// TODO Set role
}

func (that *Organization) AddCa(peerName string, orgName string, domainRoot string, port uint, grpcPort uint) {
	that.Cas = append(that.Cas, &Ca{
		PeerName:   peerName,
		OrgName:    orgName,
		DomainRoot: domainRoot,
		Port:       port,
		GrpcPort:   grpcPort,
	})
}
