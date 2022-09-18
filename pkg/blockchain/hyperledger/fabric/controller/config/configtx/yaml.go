package configtx

import (
	"log"

	"github.com/EternallyAscend/GoToolkits/pkg/IO/YAML"
	"gopkg.in/yaml.v2"
)

type ConfigTx struct {
	// Organizations []map[string]Organizations `yaml:"Organizations"`
	Organizations []*Organization                 `yaml:"Organizations"`
	Capabilities  *Capabilities                   `yaml:"Capabilities"`
	Application   *Application                    `yaml:"Application"`
	Orderer       *OrdererEtcd                    `yaml:"Orderer"`
	Channel       *Channel                        `yaml:"Channel"`
	Profiles      map[string]*ProfilesChannelEtcd `yaml:"Profiles"`
}

func GenerateConfigTx() *ConfigTx {
	configtx := &ConfigTx{
		Organizations: []*Organization{},
		Capabilities:  GenerateDefaultCapabilities(),
		Profiles:      map[string]*ProfilesChannelEtcd{},
	}

	configtx.Application = GenerateDefaultApplication(configtx.Capabilities)
	configtx.Orderer = GenerateDefaultOrdererEtcd(configtx.Capabilities)
	configtx.Channel = GenerateDefaultChannel(configtx.Capabilities)

	return configtx
}

func (that *ConfigTx) Export(folder string) {
	data, err := yaml.Marshal(*that)
	if nil != err {
		log.Fatal(err)
	}
	err = YAML.ExportToFolderFileYaml(data, folder, "configtx.yaml")
	if nil != err {
		log.Fatal(err)
	}
}

func (that *ConfigTx) CreateOrganization(orgName string, mspPath string) *Organization {
	organization := GenerateEmptyOrganization(orgName, mspPath)
	that.AddOrganization(organization)
	return organization
}

func (that *ConfigTx) AddOrganization(org *Organization) {
	that.Organizations = append(that.Organizations, org)
}

func (that *ConfigTx) FindOrganization(orgName string) *Organization {
	for i := range that.Organizations {
		if that.Organizations[i].Name == orgName {
			return that.Organizations[i]
		}
	}
	return nil
}

func (that *ConfigTx) AddOrdererToOrg(org *Organization, peerName string, orgName string, domainRoot string, port uint, ClientTLSCertPath string, ServerTLSCertPath string) {
	org.AddOrderer(peerName, orgName, domainRoot, port)
	that.Orderer.AddOrdererAndConsenter(peerName, orgName, domainRoot, port, ClientTLSCertPath, ServerTLSCertPath)
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
