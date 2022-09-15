package configtx

type CapabilitiesV2 struct {
	V2_0 bool `yaml:"V2_0"`
}

type Capabilities struct {
	Channel     *CapabilitiesV2 `yaml:"Channel"`
	Orderer     *CapabilitiesV2 `yaml:"Orderer"`
	Application *CapabilitiesV2 `yaml:"Application"`
}

func GenerateDefaultCapabilities() *Capabilities {
	return &Capabilities{
		Channel: &CapabilitiesV2{
			V2_0: true,
		},
		Orderer: &CapabilitiesV2{
			V2_0: true,
		},
		Application: &CapabilitiesV2{
			V2_0: true,
		},
	}
}
