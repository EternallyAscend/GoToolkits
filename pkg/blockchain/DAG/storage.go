package DAG

import "github.com/EternallyAscend/GoToolkits/pkg/cryptography/hash"

type Header struct{}

type Body struct {
	// TODO HE Verify Information.
}

func (that *Body) HashString() string {
	var d []byte
	// TODO Calculate Body Byte.
	return hash.SHA512String(d)
}

// Block DAG Data Structure.
type Block struct {
	Header    *Header  `json:"header" yaml:"header"`
	Body      *Body    `json:"block" yaml:"block"`
	Reference []*Block `json:"reference" yaml:"reference"`
}

type Blockchain struct {
	Bases []*Block `json:"bases" yaml:"bases"`
}

func (that *Blockchain) Fetch() {
}

type SingleChain struct {
	Genesis *Block `json:"genesis" yaml:"genesis"`
	Latest  *Block `json:"latest" yaml:"latest"`
}
