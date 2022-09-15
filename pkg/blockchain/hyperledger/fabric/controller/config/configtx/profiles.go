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

type ProfilesChannelEtcd struct {
	Consortium  string                      `yaml:"Consortium"`
	Channel     *Channel                    `yaml:"Channel"`
	Orderer     *ProfilesChannelOrdererEtcd `yaml:"Orderer"`
	Application *ProfilesChannelApplication `yaml:"Application"`
}

func GenerateDefaultProfilesChannelWithEtcdOrderer(consortium string, channel *Channel,
	orderer *OrdererEtcd, ordererOrgs []*Organization, ordererCapabilities *Capabilities,
	application *Application, applicationOrgs []*Organization, applicationCapabilities *Capabilities) *ProfilesChannelEtcd {
	profilesChannel := &ProfilesChannelEtcd{
		Consortium:  consortium,
		Channel:     channel,
		Orderer:     GenerateDefaultProfilesChannelOrdererEtcd(orderer, ordererOrgs, ordererCapabilities),
		Application: GenerateDefaultProfilesChannelApplication(application, applicationOrgs, applicationCapabilities),
	}
	return profilesChannel
}

//
//type Profiles struct {
//	Channels map[string]*ProfilesChannel `yaml:"Channels"`
//}
//
//func GenerateDefaultProfiles() *Profiles {
//	return &Profiles{
//		Channels: map[string]*ProfilesChannel{},
//	}
//}
//
//func (that *Profiles) AddChannel(channelName string, channel *ProfilesChannel) {
//	that.Channels[channelName] = channel
//}
