package dockerCompose

import "fmt"

type Service struct {
	Image         string   `yaml:"image"`
	Environment   []string `yaml:"environment"`
	Ports         []string `yaml:"ports"`
	Command       string   `yaml:"command"`
	Volumes       []string `yaml:"volumes"`
	ContainerName string   `yaml:"container_name"`
	Networks      []string `yaml:"networks"`
}

func GenerateService() *Service {
	return &Service{
		Image:         "",
		Environment:   []string{},
		Ports:         []string{},
		Command:       "",
		Volumes:       []string{},
		ContainerName: "",
		Networks:      []string{},
	}
}

func (that *Service) SetImage(image string) *Service {
	that.Image = image
	return that
}

func (that *Service) AddEnvironment(env string) *Service {
	that.Environment = append(that.Environment, env)
	return that
}

func (that *Service) AddEnvironments(envs []string) *Service {
	for i := range envs {
		that.AddEnvironment(envs[i])
	}
	return that
}

func (that *Service) AddPort(from uint, to uint) *Service {
	that.Ports = append(that.Ports, fmt.Sprintf("%d:%d", from, to))
	return that
}

func (that *Service) AddPorts(from []uint, to []uint) *Service {
	for i := range from {
		that.AddPort(from[i], to[i])
	}
	return that
}

func (that *Service) SetCommand(command string) *Service {
	that.Command = command
	return that
}

func (that *Service) AddVolume(from string, to string) *Service {
	that.Volumes = append(that.Volumes, fmt.Sprintf("%s:%s", from, to))
	return that
}

func (that *Service) AddVolumes(from []string, to []string) *Service {
	for i := range from {
		that.AddVolume(from[i], to[i])
	}
	return that
}

func (that *Service) SetContainerName(name string) *Service {
	that.ContainerName = name
	return that
}

func (that *Service) AddNetwork(from string, to string) *Service {
	that.Networks = append(that.Networks, fmt.Sprintf("%s:%s", from, to))
	return that
}

func (that *Service) AddNetworks(from []string, to []string) *Service {
	for i := range from {
		that.AddNetwork(from[i], to[i])
	}
	return that
}
