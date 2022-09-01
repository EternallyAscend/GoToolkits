package dockerCompose

import "gopkg.in/yaml.v2"

type DockerYAML struct {
	Version  string              `yaml:"version"`
	Networks map[string]*Network `yaml:"networks"`
	Services map[string]*Service `yaml:"services"`
	Volumes  []string            `yaml:"volumes"`
	Driver   []string            `yaml:"driver"`
}

func GenerateDockerYAML(version string) *DockerYAML {
	return &DockerYAML{
		Version:  version,
		Services: map[string]*Service{},
		Networks: map[string]*Network{},
		Volumes:  []string{},
		Driver:   []string{},
	}
}

func (that *DockerYAML) AddNetwork(key string, network *Network) error {
	// TODO Check before adding.
	that.Networks[key] = network
	return nil
}

func (that *DockerYAML) AddService(key string, service *Service) error {
	// TODO Check before adding.
	that.Services[key] = service
	return nil
}

func (that *DockerYAML) ExportToByteArray() ([]byte, error) {
	return yaml.Marshal(that)
}

func (that *DockerYAML) ExportToFile(path string, fileName string) error {

	return nil
}
