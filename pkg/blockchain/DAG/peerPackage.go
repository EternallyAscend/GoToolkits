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

func TransferDataToPackage(data []byte, typeCode uint) ([]byte, error) {
	p := &Package{
		Type: typeCode,
		// Length:  uint(len(data)),
		Message: data,
	}
	return json.Marshal(p)
}

const DefaultPackageTcpHeaderSize = 256

type PackageTcpHeader struct {
	Type   uint `json:"type" yaml:"type"`
	Length uint `json:"length" yaml:"length"`
}

func UnpackPackageTcpHeader(data []byte) *PackageTcpHeader {
	p := &PackageTcpHeader{}
	err := json.Unmarshal(data, p)
	if nil != err {
		log.Println("Unpack packageTcpHeader json failed,", err)
		log.Println(string(data))
		return nil
	}
	return p
}
