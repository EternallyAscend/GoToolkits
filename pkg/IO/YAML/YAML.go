package YAML

import (
	"fmt"
	"os"

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

func ReadStructFromFileYaml(data interface{}, path string) error {
	byteData, err := file.ReadFile(path)
	if nil != err {
		return err
	}
	err = yaml.Unmarshal(byteData, data)
	return err
}

func ExportToFolderFileYaml(data []byte, folder string, file string) error {
	f, err := os.OpenFile(fmt.Sprintf("%s/%s", folder, file), os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	if nil != err {
		return err
	}
	_, err = f.Write(data)
	if nil != err {
		return err
	}
	return f.Sync()
}
