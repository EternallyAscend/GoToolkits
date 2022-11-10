package DAG

import (
	"encoding/json"
	"log"

	"github.com/EternallyAscend/GoToolkits/pkg/cryptography/hash"
)

type Header struct {
	// TODO Blockchain Hash.
	// TODO HE GH Arguments.
}

type Body struct {
	// TODO HE Verify Information.
	// TODO Compression Gradients.
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

type SingleHeader struct {
	SHA512 []byte `json:"sha512" yaml:"sha512"`
}

type SingleBody struct{}

func (that *SingleBody) GenerateHeader() *SingleHeader {
	body, err := json.Marshal(that)
	if nil != err {
		log.Println(err)
		return nil
	}
	return &SingleHeader{
		SHA512: hash.SHA512(body),
	}
}

type SingleBlock struct {
	Header *SingleHeader `json:"header" yaml:"header"`
	Body   *SingleBody   `json:"body" yaml:"body"`
	Next   *SingleBlock  `json:"next" yaml:"next"`
}

func (that *SingleBlock) Verify() bool {
	body, err := json.Marshal(that.Body)
	if nil != err {
		log.Println(err)
		return false
	}
	bodyHash := hash.SHA512(body)
	if len(bodyHash) != len(that.Header.SHA512) {
		return false
	}
	for i := range bodyHash {
		if bodyHash[i] != that.Header.SHA512[i] {
			return false
		}
	}
	return true
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

func (that *SingleChain) Verify() bool {
	return that.VerifyFrom(that.Genesis)
}

func (that *SingleChain) VerifyFrom(block *SingleBlock) bool {
	for nil != block.Next {
		block = block.Next
		if !block.Verify() {
			return false
		}
	}
	return true
}
