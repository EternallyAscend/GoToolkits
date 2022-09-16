package configtx

import "fmt"

type OrganizationsPoliciesRole struct {
	Type string `yaml:"Type"` // Signature
	Rule string `yaml:"Rule"` // 'OR('org.member')' or 'OR('org.admin')'
}

type OrganizationPolicies struct {
	Readers *OrganizationsPoliciesRole `yaml:"Readers"`
	Writers *OrganizationsPoliciesRole `yaml:"Writers"`
	Admins  *OrganizationsPoliciesRole `yaml:"Admins"`
}

func GenerateDefaultPolicies(org string) *OrganizationPolicies {
	return &OrganizationPolicies{
		Readers: &OrganizationsPoliciesRole{
			Type: "Signature",
			Rule: fmt.Sprintf("OR('%s.member')", org),
		},
		Writers: &OrganizationsPoliciesRole{
			Type: "Signature",
			Rule: fmt.Sprintf("OR('%s.member')", org),
		},
		Admins: &OrganizationsPoliciesRole{
			Type: "Signature",
			Rule: fmt.Sprintf("OR('%s.admin')", org),
		},
	}
}

type OrganizationAnchorPeer struct {
	Host string `yaml:"Host"`
	Port uint   `yaml:"Port"`
}

type Organization struct {
	Name             string                    `yaml:"Name"`             // 组织名称
	ID               string                    `yaml:"ID"`               // 组织ID
	MSPDir           string                    `yaml:"MSPDir"`           // 组织MSP文件夹路径
	Policies         *OrganizationPolicies     `yaml:"Policies"`         // 组织策略
	OrdererEndpoints []string                  `yaml:"OrdererEndpoints"` // 排序节点列表
	AnchorPeers      []*OrganizationAnchorPeer `yaml:"AnchorPeers"`      // 锚节点 对外代表本组织通信
}

func GenerateEmptyOrganization(name string, msp string) *Organization {
	return &Organization{
		Name:             name,
		ID:               name,
		MSPDir:           msp,
		Policies:         GenerateDefaultPolicies(name),
		OrdererEndpoints: []string{},
		AnchorPeers:      []*OrganizationAnchorPeer{},
	}
}

func (that *Organization) AddOrderer(orderer string) {
	that.OrdererEndpoints = append(that.OrdererEndpoints, orderer)
}

func (that *Organization) AddAnchorPeer(host string, port uint) {
	that.AnchorPeers = append(that.AnchorPeers, &OrganizationAnchorPeer{
		Host: host,
		Port: port,
	})
}

func (that *Organization) SetPolicies(policies *OrganizationPolicies) {
	that.Policies = policies
}
