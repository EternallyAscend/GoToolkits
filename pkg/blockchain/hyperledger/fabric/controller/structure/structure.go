package structure

import (
	"fmt"
	"github.com/EternallyAscend/GoToolkits/pkg/IO/JSON"
	"github.com/EternallyAscend/GoToolkits/pkg/IO/YAML"
	"gopkg.in/yaml.v2"
	"log"
)

type Config struct {
	Organizations []*Organization `yaml:"organizations" json:"organizations"`
	Channels      []*Channel      `yaml:"channels" json:"channels"`
	Applications  []*Application  `yaml:"applications" json:"applications"`
}

func (that *Config) Export(path string, name string, yamlOut bool, jsonOut bool) {
	if yamlOut {
		yamlData, err := yaml.Marshal(*that)
		if nil != err {
			log.Println(err)
		}
		err = YAML.ExportToFolderFileYaml(yamlData, path, fmt.Sprintf("%s.yaml", name))
		if nil != err {
			log.Println(err)
		}
	}
	if jsonOut {
		jsonData, err := yaml.Marshal(*that)
		if nil != err {
			log.Println(err)
		}
		err = JSON.ExportToFolderFileJson(jsonData, path, fmt.Sprintf("%s.json", name))
		if nil != err {
			log.Println(err)
		}
	}
}

func ReadConfigFromFile(path string) *Config {
	config := &Config{
		Organizations: []*Organization{},
		Channels:      []*Channel{},
		Applications:  []*Application{},
	}
	// TODO Reading Config Files.
	return config
}
func (that *Config) AddOrganization(org *Organization) {
	that.Organizations = append(that.Organizations, org)
}

func (that *Config) CreateOrganization(orgName string, id string, domainRoot string, mapPath string) *Organization {
	organization := GenerateEmptyOrganization(orgName, id, domainRoot, mapPath)
	that.AddOrganization(organization)
	return organization
}

func (that *Config) FindOrganization(orgName string) *Organization {
	for i := range that.Organizations {
		if that.Organizations[i].Name == orgName {
			return that.Organizations[i]
		}
	}
	return nil
}

func (that *Config) AddOrdererToOrg(org *Organization, peerName string, orgName string, domainRoot string, port uint) {
	org.AddOrderer(peerName, orgName, domainRoot, port)
}
func (that *Config) AddCaToOrg(org *Organization, peerName string, orgName string, domainRoot string, port uint, grpcPort uint) {
	org.AddCa(peerName, orgName, domainRoot, port, grpcPort)
}
func (that *Config) AddPeerToOrg(org *Organization, peerName string, orgName string, domainRoot string, port uint) {
	org.AddPeer(peerName, orgName, domainRoot, port)
}
