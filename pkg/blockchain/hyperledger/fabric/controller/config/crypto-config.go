package config

import "fmt"

// Path: organization/cryptogen/crypto-config.yaml

type OrdererCryptoConfigSpecs struct {
	Hostname string   `yaml:"Hostname"`
	SANS     []string `yaml:"SANS"`
}

type OrdererCryptoConfig struct {
	Name          string                      `yaml:"Name"`
	Domain        string                      `yaml:"Domain"`
	EnableNodeOUs bool                        `yaml:"EnableNodeOUs"`
	Specs         []*OrdererCryptoConfigSpecs `yaml:"Specs"`
}

func GenerateDefaultOrdererCryptoConfig(name, domainRoot string) *OrdererCryptoConfig {
	return &OrdererCryptoConfig{
		Name:          name,
		Domain:        domainRoot,
		EnableNodeOUs: true,
		Specs: []*OrdererCryptoConfigSpecs{
			{
				Hostname: fmt.Sprintf("%s.%s", name, domainRoot),
				SANS: []string{
					"localhost",
				},
			},
		},
	}
}

type PeerCryptoConfigTemplate struct {
	Count uint     `yaml:"Count"`
	SANS  []string `yaml:"SANS"`
}

type PeerCryptoConfigUsers struct {
	Count uint `yaml:"Count"`
}

type PeerCryptoConfig struct {
	Name          string                    `yaml:"Name"`
	Domain        string                    `yaml:"Domain"`
	EnableNodeOUs bool                      `yaml:"EnableNodeOUs"`
	Template      *PeerCryptoConfigTemplate `yaml:"Template"`
	Users         *PeerCryptoConfigUsers    `yaml:"Users"`
}

func GenerateDefaultPeerCryptoConfig(orgName, domainRoot string, count, users uint) *PeerCryptoConfig {
	return &PeerCryptoConfig{
		Name:          orgName,
		Domain:        domainRoot,
		EnableNodeOUs: true,
		Template: &PeerCryptoConfigTemplate{
			Count: count,
			SANS: []string{
				"localhost",
			},
		},
		Users: &PeerCryptoConfigUsers{Count: users},
	}
}
