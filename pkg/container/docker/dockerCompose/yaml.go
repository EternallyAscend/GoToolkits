package dockerCompose

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
)

type DockerYAML struct {
	Version  string              `yaml:"version"`
	Networks map[string]*Network `yaml:"networks"`
	Services map[string]*Service `yaml:"services"`
	//Volumes  []string            `yaml:"volumes"`
	//Driver   []string            `yaml:"driver"`
}

func GenerateDockerYAML(version string) *DockerYAML {
	return &DockerYAML{
		Version:  version,
		Services: map[string]*Service{},
		Networks: map[string]*Network{},
		//Volumes:  []string{},
		//Driver:   []string{},
	}
}

func (that *DockerYAML) AddNetwork(network *Network) error {
	if nil == network {
		return errors.New("Nil pointer of network struct. ")
	}
	if nil == that.Networks[network.Name] {
		that.Networks[network.Name] = network
	} else {
		return errors.New(fmt.Sprintf("Duplicate key for network %s. ", network.Name))
	}
	return nil
}

func (that *DockerYAML) AddService(service *Service) error {
	if nil == service {
		return errors.New("Nil pointer of service struct. ")
	}
	if nil == that.Networks[service.ContainerName] {
		that.Services[service.ContainerName] = service
	} else {
		return errors.New(fmt.Sprintf("Duplicate key for service %s. ", service.ContainerName))
	}
	return nil
}

func (that *DockerYAML) ExportToByteArray() ([]byte, error) {
	return yaml.Marshal(that)
}

func (that *DockerYAML) ExportToFile(path string, fileName string) error {

	return nil
}
