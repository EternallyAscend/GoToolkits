package dockerCompose

type Network struct {
	Name string `yaml:"name"`
}

func GenerateNetwork(name string) *Network {
	return &Network{
		Name: name,
	}
}
