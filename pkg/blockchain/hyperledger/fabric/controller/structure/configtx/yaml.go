package configtx

import (
	"faber/pkg/file"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
)

type ConfigTx struct {
	//Organizations []map[string]Organizations `yaml:"Organizations"`
	Organizations []*Organization                 `yaml:"Organizations"`
	Capabilities  *Capabilities                   `yaml:"Capabilities"`
	Application   *Application                    `yaml:"Application"`
	Orderer       *OrdererEtcd                    `yaml:"Orderer"`
	Channel       *Channel                        `yaml:"Channel"`
	Profiles      map[string]*ProfilesChannelEtcd `yaml:"Profiles"`
}

func GenerateConfigTx() *ConfigTx {
	configtx := &ConfigTx{
		//Organizations: []map[string]Organizations{
		//	{
		//		"Name": {
		//			Name:             "Default",
		//			ID:               "Default",
		//			MSPDir:           "MSP",
		//			Policies:         *GenerateDefaultPolicies("Default"),
		//			OrdererEndpoints: []string{
		//				"orderer.default.com:7051",
		//			},
		//			AnchorPeers:      []OrganizationAnchorPeer{
		//				{
		//					Host: "peer.ddefault.com",
		//					Port: "7251",
		//				},
		//			},
		//		},
		//	},
		//},
		Organizations: []*Organization{},
		Capabilities:  GenerateDefaultCapabilities(),
		Profiles:      map[string]*ProfilesChannelEtcd{},
	}

	configtx.Application = GenerateDefaultApplication(configtx.Capabilities)
	configtx.Orderer = GenerateDefaultOrdererEtcd(configtx.Capabilities)
	configtx.Channel = GenerateDefaultChannel(configtx.Capabilities)

	configtx.Example()

	return configtx
}

func (that *ConfigTx) Export(folder string) {
	data, err := yaml.Marshal(*that)
	if nil != err {
		log.Fatal(err)
	}
	err = file.ExportYamlFile(data, folder, "configtx.yaml")
	if nil != err {
		log.Fatal(err)
	}
}

func (that *ConfigTx) CreateOrganization(org string, mspPath string) *Organization {
	organization := GenerateEmptyOrganization(org, mspPath)
	that.AddOrganization(organization)
	return organization
}

func (that *ConfigTx) AddOrganization(org *Organization) {
	that.Organizations = append(that.Organizations, org)
}

func (that *ConfigTx) FindOrganization(org string) *Organization {
	for i := range that.Organizations {
		println(that.Organizations[i].Name)
	}
	return nil
}

func (that *ConfigTx) AddOrdererToOrg(org *Organization, domain string, port uint, ClientTLSCertPath string, ServerTLSCertPath string) {
	org.AddOrderer(fmt.Sprintf("%s:%d", domain, port))
	that.Orderer.AddOrdererAndConsenter(domain, port, ClientTLSCertPath, ServerTLSCertPath)
}

func (that *ConfigTx) AddChannel(name string, consortium string) {
	if nil != that.Profiles[name] {
		return
	}
	channel := GenerateDefaultProfilesChannelWithEtcdOrderer(consortium, that.Channel,
		that.Orderer, that.Organizations, that.Capabilities,
		that.Application, that.Organizations, that.Capabilities)
	that.Profiles[name] = channel
}

func (that *ConfigTx) Example() {
	//that.Organizations = append(that.Organizations,
	//	Organization{
	//		Name:     "default",
	//		ID:       "default",
	//		MSPDir:   "defaultMSP",
	//		Policies: *GenerateDefaultPolicies("default"),
	//		OrdererEndpoints: []string{
	//			"orderer.default.com:7051",
	//		},
	//		AnchorPeers: []OrganizationAnchorPeer{
	//			{
	//				Host: "peer.default.com",
	//				Port: 7251,
	//			},
	//		},
	//	})
	defaultOrg := that.CreateOrganization("default", "defaultMSP")
	//defaultOrg := GenerateEmptyOrganization("default", "defaultMSP")
	//that.Organizations = append(that.Organizations, defaultOrg)
	that.FindOrganization("")
	defaultOrg.AddAnchorPeer("peer.default.com", 7251)
	that.AddOrdererToOrg(defaultOrg, "orderer.default.com", 7051, "", "")

	that.AddChannel("defaultChannel", "defaultConsortium")
}
