package structure

import (
	"faber/pkg/file"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
)

type Config struct {
	Organizations []*Organization `yaml:"organizations" json:"organizations"`
	Channels      []*Channel      `yaml:"channels" json:"channels"`
	Applications  []*Application  `yaml:"applications" json:"applications"`
}

func (that *Config) Export(path string, name string, yamlOut bool, jsonOut bool) {
	if yamlOut {
		yamlData, err := yaml.Marshal(*that)
		if nil != err {
			log.Println(err)
		}
		err = file.ExportYamlFile(yamlData, path, fmt.Sprintf("%s.yaml", name))
		if nil != err {
			log.Println(err)
		}
	}
	if jsonOut {
		jsonData, err := yaml.Marshal(*that)
		if nil != err {
			log.Println(err)
		}
		err = file.ExportJsonFile(jsonData, path, fmt.Sprintf("%s.json", name))
		if nil != err {
			log.Println(err)
		}
	}
}

func ReadConfigFromFile(path string) *Config {
	config := &Config{
		Organizations: []*Organization{},
		Channels:      []*Channel{},
		Applications:  []*Application{},
	}
	// TODO Reading Config Files.
	return config
}

