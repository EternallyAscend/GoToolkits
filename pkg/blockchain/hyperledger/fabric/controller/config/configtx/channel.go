package configtx

type ChannelPoliciesRole struct {
	Type string `yaml:"Type"`
	Rule string `yaml:"Rule"`
}

type ChannelPolicies struct {
	Readers *ChannelPoliciesRole `yaml:"Readers"`
	Writers *ChannelPoliciesRole `yaml:"Writers"`
	Admins  *ChannelPoliciesRole `yaml:"Admins"`
}

func GenerateDefaultChannelPolicies() *ChannelPolicies {
	return &ChannelPolicies{
		Readers: &ChannelPoliciesRole{
			Type: "ImplicitMeta",
			Rule: "ANY Readers",
		},
		Writers: &ChannelPoliciesRole{
			Type: "ImplicitMeta",
			Rule: "ANY Writers",
		},
		Admins: &ChannelPoliciesRole{
			Type: "ImplicitMeta",
			Rule: "MAJORITY Admins",
		},
	}
}

type Channel struct {
	Policies     *ChannelPolicies `yaml:"Policies"`
	Capabilities *Capabilities    `yaml:"Capabilities"`
}

func GenerateDefaultChannel(capabilities *Capabilities) *Channel {
	return &Channel{
		Policies:     GenerateDefaultChannelPolicies(),
		Capabilities: capabilities,
	}
}
