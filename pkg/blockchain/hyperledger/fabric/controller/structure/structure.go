package structure

import (
	"fmt"
	"log"

	"github.com/EternallyAscend/GoToolkits/pkg/blockchain/hyperledger/fabric/controller/config"
	"github.com/EternallyAscend/GoToolkits/pkg/blockchain/hyperledger/fabric/controller/docker"

	"github.com/EternallyAscend/GoToolkits/pkg/IO/JSON"
	"github.com/EternallyAscend/GoToolkits/pkg/IO/YAML"
	"github.com/EternallyAscend/GoToolkits/pkg/blockchain/hyperledger/fabric/controller"
	"github.com/EternallyAscend/GoToolkits/pkg/blockchain/hyperledger/fabric/controller/config/configtx"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Organizations []*Organization `yaml:"organizations" json:"organizations"`
	Channels      []*Channel      `yaml:"channels" json:"channels"`
	Applications  []*Application  `yaml:"applications" json:"applications"`
	NetworkName   string          `yaml:"networkName" json:"networkName"`
	configtx      *configtx.ConfigTx
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

// FillConfigtx 填充configtx数据
func (that *Config) FillConfigtx() {
	for i := range that.Channels {
		that.configtx.AddChannel(that.Channels[i].Name, that.Channels[i].Consortium)
	}
	for i := range that.Organizations {
		// 加入组织部分数据
		org := configtx.GenerateEmptyOrganization(that.Organizations[i].Name, controller.GenerateMspID(that.Organizations[i].Name))
		that.configtx.AddOrganization(org)

		// 为channels添加组织信息
		// TODO 为channels添加组织信息

		// 为channels加入orderer信息
		for j := range that.Organizations[i].Orderers {
			orderer := that.Organizations[i].Orderers[j]
			clientTLSCertPath := controller.GenerateClientTLSCertPath(false, orderer.PeerName, orderer.OrgName, orderer.DomainRoot)
			serverTLSCertPath := controller.GenerateServerTLSCertPath(false, orderer.PeerName, orderer.OrgName, orderer.DomainRoot)
			that.configtx.AddOrdererToOrg(org, orderer.PeerName,
				orderer.OrgName,
				orderer.DomainRoot,
				orderer.Port,
				clientTLSCertPath,
				serverTLSCertPath,
			)
		}
	}
}

// FillCryptoConfig 填充crypto-config
func (that *Config) FillCryptoConfig() {
	for i := range that.Organizations {
		user := 0
		for j := range that.Organizations[i].Peers {
			user += len(that.Organizations[i].Peers[j].PeerUser)
		}
		config.GenerateDefaultPeerCryptoConfig(that.Organizations[i].Name, that.Organizations[i].Domain, uint(len(that.Organizations[i].Peers)), uint(user))
		config.GenerateDefaultOrdererCryptoConfig(that.Organizations[i].Name, that.Organizations[i].Domain)
	}
}

// FillDockerCompose 填充docker-compose
func (that *Config) FillDockerCompose() {
	for i := range that.Organizations {
		for j := range that.Organizations[i].Peers {
			peer := that.Organizations[i].Peers[j]
			docker.GeneratePeerService(
				config.FabricVersionFull,
				peer.PeerName,
				peer.OrgName,
				peer.DomainRoot,
				that.NetworkName,
				true,
				controller.GenerateTlsCertPath(true, peer.PeerName, peer.OrgName, peer.DomainRoot),
				true,
				controller.GenerateMspID(peer.OrgName),
				controller.GenerateMspPath(true, peer.OrgName, peer.OrgName, peer.DomainRoot),
				peer.Port,
				peer.ChaincodePort,
				peer.OperationsPort,
			)
		}
	}
	for i := range that.Organizations {
		for j := range that.Organizations[i].Orderers {
			orderer := that.Organizations[i].Orderers[j]
			docker.GenerateOrdererService(
				config.FabricCaVersionFull,
				orderer.PeerName,
				orderer.OrgName,
				orderer.DomainRoot,
				that.NetworkName,
				orderer.Port,
				controller.GenerateMspID(orderer.OrgName),
				controller.GenerateMspPath(false, orderer.PeerName, orderer.OrgName, orderer.DomainRoot),
				orderer.OperationPort,
				true,
				controller.GenerateTlsCertPath(false, orderer.PeerName, orderer.OrgName, orderer.DomainRoot),
				controller.GenerateBlockPath(),
				true,
			)
		}
	}
}
