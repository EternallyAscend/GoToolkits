package DAG

import (
	"encoding/json"
	"log"
)

type Package struct {
	Type    uint   `json:"type" yaml:"type"`
	Message []byte `json:"message" yaml:"message"`
}

func UnpackPackage(data []byte) *Package {
	p := &Package{}
	err := json.Unmarshal(data, p)
	if nil != err {
		log.Println("Unpack package json failed,", err)
		log.Println(string(data))
		return nil
	}
	return p
}
