package YAML

import (
	"github.com/EternallyAscend/GoToolkits/pkg/IO/file"
	"gopkg.in/yaml.v2"
)

func ExportToFileYaml(data interface{}, path string) error {
	byteData, err := yaml.Marshal(data)
	if nil != err {
		return err
	}
	return file.CreateOrRewrite(byteData, path)
}
