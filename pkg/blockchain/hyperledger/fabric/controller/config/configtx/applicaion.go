package configtx

type ApplicationPoliciesRole struct {
	Type string `yaml:"Type"`
	Rule string `yaml:"Rule"`
}

type ApplicationPolicies struct {
	Readers              *ApplicationPoliciesRole `yaml:"Reader"`
	Writers              *ApplicationPoliciesRole `yaml:"Writers"`
	Admins               *ApplicationPoliciesRole `yaml:"Admins"`
	LifecycleEndorsement *ApplicationPoliciesRole `yaml:"LifecycleEndorsement"`
	Endorsement          *ApplicationPoliciesRole `yaml:"Endorsement"`
}

func GenerateDefaultApplicationPolicies() *ApplicationPolicies {
	return &ApplicationPolicies{
		Readers: &ApplicationPoliciesRole{
			Type: "ImplicitMeta",
			Rule: "ANY Readers",
		},
		Writers: &ApplicationPoliciesRole{
			Type: "ImplicitMeta",
			Rule: "ANY Writers",
		},
		Admins: &ApplicationPoliciesRole{
			Type: "ImplicitMeta",
			Rule: "MAJORITY Admins",
		},
		LifecycleEndorsement: &ApplicationPoliciesRole{
			Type: "ImplicitMeta",
			Rule: "MAJORITY Endorsement",
		},
		Endorsement: &ApplicationPoliciesRole{
			Type: "ImplicitMeta",
			Rule: "MAJORITY Endorsement",
		},
	}
}

type Application struct {
	// Organizations interface{}              `yaml:"Organizations"`
	Polices      *ApplicationPolicies `yaml:"Polices"`
	Capabilities *Capabilities        `yaml:"Capabilities"`
}

func GenerateDefaultApplication(capabilities *Capabilities) *Application {
	return &Application{
		Polices:      GenerateDefaultApplicationPolicies(),
		Capabilities: capabilities,
	}
}
