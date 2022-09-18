package configtx

type ProfilesChannelOrdererEtcd struct {
	Orderer       *OrdererEtcd    `yaml:"Orderer"`
	Organizations []*Organization `yaml:"Organizations"`
	Capabilities  *Capabilities   `yaml:"Capabilities"`
}

func GenerateDefaultProfilesChannelOrdererEtcd(orderer *OrdererEtcd, orgs []*Organization, capabilities *Capabilities) *ProfilesChannelOrdererEtcd {
	return &ProfilesChannelOrdererEtcd{
		Orderer:       orderer,
		Organizations: orgs,
		Capabilities:  capabilities,
	}
}

type ProfilesChannelApplication struct {
	Application   *Application    `yaml:"Application"`
	Organizations []*Organization `yaml:"Organizations"`
	Capabilities  *Capabilities   `yaml:"Capabilities"`
}

func GenerateDefaultProfilesChannelApplication(application *Application, orgs []*Organization, capabilities *Capabilities) *ProfilesChannelApplication {
	return &ProfilesChannelApplication{
		Application:   application,
		Organizations: orgs,
		Capabilities:  capabilities,
	}
}

type SystemChannelConsortium struct {
	// TODO @EternallyAscend
	Organization []*Organization `yaml:"Organization"`
}

type ProfilesChannelEtcd struct {
	Consortium  string                              `yaml:"Consortium"`
	Consortiums map[string]*SystemChannelConsortium `yaml:"Consortiums"`
	Channel     *Channel                            `yaml:"Channel"`
	Orderer     *ProfilesChannelOrdererEtcd         `yaml:"Orderer"`
	Application *ProfilesChannelApplication         `yaml:"Application"`
}

func GenerateDefaultProfilesChannelWithEtcdOrderer(consortium string, channel *Channel,
	// consortium是configtx中profile中的字段，text-network中值为SampleConsortium
	orderer *OrdererEtcd, ordererOrgs []*Organization, ordererCapabilities *Capabilities,
	application *Application, applicationOrgs []*Organization, applicationCapabilities *Capabilities,
) *ProfilesChannelEtcd {
	profilesChannel := &ProfilesChannelEtcd{
		Consortium:  consortium,
		Channel:     channel,
		Orderer:     GenerateDefaultProfilesChannelOrdererEtcd(orderer, ordererOrgs, ordererCapabilities),
		Application: GenerateDefaultProfilesChannelApplication(application, applicationOrgs, applicationCapabilities),
	}
	return profilesChannel
}

func GenerateSystemProfilesChannelWithEtcdOrderer(channel *Channel, orderer *OrdererEtcd, ordererOrgs []*Organization, ordererCapabilities *Capabilities,
	application *Application, applicationOrgs []*Organization, applicationCapabilities *Capabilities,
) *ProfilesChannelEtcd {
	return &ProfilesChannelEtcd{
		Consortium:  "SystemChannel",
		Consortiums: map[string]*SystemChannelConsortium{},
		Channel:     channel,
		Orderer:     GenerateDefaultProfilesChannelOrdererEtcd(orderer, ordererOrgs, ordererCapabilities),
		Application: GenerateDefaultProfilesChannelApplication(application, applicationOrgs, applicationCapabilities),
	}
}

func AppendChannelIntoSystemChannel(configtx *ConfigTx, channelName string, organizations []*Organization) {
	configtx.Profiles["SystemChannel"].Consortiums[channelName].Organization = organizations
}
