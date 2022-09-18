package configtx

import "fmt"

type OrdererEtcdRaftConsenters struct {
	Host          string `yaml:"Host"`
	Port          uint   `yaml:"Port"`
	ClientTLSCert string `yaml:"ClientTlsCert"`
	ServerTLSCert string `yaml:"ServerTlsCert"`
}

type OrdererBatchSize struct {
	MaxMessageCount   uint   `yaml:"MaxMessageCount"`
	AbsoluteMaxBytes  string `yaml:"AbsoluteMaxBytes"`
	PreferredMaxBytes string `yaml:"PreferredMaxBytes"`
}

func GenerateDefaultOrdererBatchSize() *OrdererBatchSize {
	return &OrdererBatchSize{
		MaxMessageCount:   32,
		AbsoluteMaxBytes:  "64 MB",
		PreferredMaxBytes: "4096 KB",
	}
}

type OrdererPoliciesRole struct {
	Type string `yaml:"Type"`
	Rule string `yaml:"Rule"`
}

type OrdererPolicies struct {
	Readers         *OrdererPoliciesRole `yaml:"Readers"`
	Writers         *OrdererPoliciesRole `yaml:"Writers"`
	Admins          *OrdererPoliciesRole `yaml:"Admins"`
	BlockValidation *OrdererPoliciesRole `yaml:"BlockValidation"`
}

func GenerateDefaultOrdererPolicies() *OrdererPolicies {
	return &OrdererPolicies{
		Readers: &OrdererPoliciesRole{
			Type: "ImplicitMeta",
			Rule: "ANY Readers",
		},
		Writers: &OrdererPoliciesRole{
			Type: "ImplicitMeta",
			Rule: "ANY Writers",
		},
		Admins: &OrdererPoliciesRole{
			Type: "ImplicitMeta",
			Rule: "MAJORITY Admins",
		},
		BlockValidation: &OrdererPoliciesRole{
			Type: "ImplicitMeta",
			Rule: "ANY Writers",
		},
	}
}

type OrdererEtcdRaft struct {
	Consenters []*OrdererEtcdRaftConsenters `yaml:"Consenters"`
}

type OrdererEtcd struct {
	OrdererType  string            `yaml:"OrdererType"`
	Addresses    []string          `yaml:"Addresses"`
	EtcdRaft     *OrdererEtcdRaft  `yaml:"EtcdRaft"`
	BatchTimeout string            `yaml:"BatchTimeout"`
	BatchSize    *OrdererBatchSize `yaml:"BatchSize"`
	Policies     *OrdererPolicies  `yaml:"Policies"`
	Capabilities *Capabilities     `yaml:"Capabilities"`
}

func GenerateDefaultOrdererEtcd(capabilities *Capabilities) *OrdererEtcd {
	return &OrdererEtcd{
		OrdererType:  "etcdraft",
		Addresses:    []string{},
		EtcdRaft:     &OrdererEtcdRaft{},
		BatchTimeout: "2s",
		BatchSize:    GenerateDefaultOrdererBatchSize(),
		Policies:     GenerateDefaultOrdererPolicies(),
		Capabilities: capabilities,
	}
}

func (that *OrdererEtcd) AddConsenter(peerName string, orgName string, domainRoot string, port uint, ClientTLSCertPath string, ServerTLSCertPath string) {
	that.EtcdRaft.Consenters = append(that.EtcdRaft.Consenters, &OrdererEtcdRaftConsenters{
		Host:          fmt.Sprintf("%s.%s.%s", peerName, orgName, domainRoot),
		Port:          port,
		ClientTLSCert: ClientTLSCertPath,
		ServerTLSCert: ServerTLSCertPath,
	})
}

func (that *OrdererEtcd) AddOrdererAndConsenter(peerName string, orgName string, domainRoot string, port uint, ClientTLSCertPath string, ServerTLSCertPath string) {
	that.AddOrderer(peerName, orgName, domainRoot, port)
	that.AddConsenter(peerName, orgName, domainRoot, port, ClientTLSCertPath, ServerTLSCertPath)
}

func (that *OrdererEtcd) AddOrderer(peerName string, orgName string, domainRoot string, port uint) {
	domain := fmt.Sprintf("%s.%s.%s", peerName, orgName, domainRoot)
	that.Addresses = append(that.Addresses, fmt.Sprintf("%s:%d", domain, port))
}
