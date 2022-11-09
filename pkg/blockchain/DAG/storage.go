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

func (that *Blockchain) Verify() bool {
	return true
}

type SingleHeader struct{}

type SingleBody struct{}

type SingleBlock struct {
	Header *SingleHeader `json:"header" yaml:"header"`
	Body   *SingleBody   `json:"body" yaml:"body"`
	Next   *SingleBlock  `json:"next" yaml:"next"`
}

type SingleChain struct {
	Genesis *SingleBlock `json:"genesis" yaml:"genesis"`
	Latest  *SingleBlock `json:"latest" yaml:"latest"`
}

func (that *SingleChain) Append(head *SingleBlock) {
	if nil == head {
		return
	}
	p := head
	for nil != p.Next {
		p = p.Next
	}
	that.Latest = p
}
