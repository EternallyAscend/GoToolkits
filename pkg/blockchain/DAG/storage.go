package DAG

import (
	"encoding/json"
	"log"

	"github.com/EternallyAscend/GoToolkits/pkg/cryptography/hash"
	"github.com/EternallyAscend/GoToolkits/pkg/cryptography/homomorphic/pedersonCommitment"
)

type Header struct {
	// TODO Blockchain Hash.
	Merkle []byte `json:"merkle" yaml:"merkle"`
	SHA512 []byte `json:"sha512" yaml:"sha512"`
	// TODO HE GH Arguments.
	Dealer *pedersonCommitment.DealerUnit `json:"dealer" yaml:"dealer"`
}

func (that *Body) GenerateHeader(dealer *pedersonCommitment.DealerUnit, reference []*Block) *Header {
	body, err := json.Marshal(that)
	if nil != err {
		return nil
	}
	header := &Header{
		Merkle: hash.SHA512(body),
		SHA512: nil,
		Dealer: dealer,
	}
	for i := range that.Units {
		// TODO Calculate the sum of Random.
		dealer.R = that.Units[i].R
	}
	var data []byte
	for i := range reference {
		data = append(data, reference[i].Header.SHA512...)
	}
	data = append(data, header.Merkle...)
	header.SHA512 = hash.SHA512(data)
	return header
}

type Body struct {
	// TODO HE Verify Information.
	Units []*pedersonCommitment.VerifiableMessageUint
	// TODO Compression Gradients.
	Data [][]byte
}

func (that *Body) HashString() string {
	var d []byte
	// TODO Calculate Body Byte.
	return hash.SHA512String(d)
}

func (that *Body) AppendData(data []byte) {
	that.Data = append(that.Data, data)
}

// Block DAG Data Structure.
type Block struct {
	Header    *Header  `json:"header" yaml:"header"`
	Body      *Body    `json:"block" yaml:"block"`
	Reference []*Block `json:"reference" yaml:"reference"`
}

func (that *Block) Verify() {}

func (that *Block) Open() {}

func (that *Block) Check() {}

type Blockchain struct {
	Bases []*Block `json:"bases" yaml:"bases"`
}

func (that *Blockchain) Verify() bool {
	return true
}

type SingleHeader struct {
	Merkle []byte `json:"merkle" yaml:"merkle"`
	SHA512 []byte `json:"sha512" yaml:"sha512"`
}

type SingleBody struct {
	Data  []byte `json:"data" yaml:"data"`
	Nonce []byte `json:"nonce" yaml:"nonce"`
}

func (that *SingleBody) GenerateHeader(forwardHash []byte) *SingleHeader {
	body, err := json.Marshal(that)
	if nil != err {
		log.Println(err)
		return nil
	}
	header := &SingleHeader{
		Merkle: hash.SHA512(body),
	}
	if nil != forwardHash {
		header.SHA512 = hash.SHA512(append(forwardHash, body...))
	} else {
		header.SHA512 = header.Merkle
	}
	return header
}

type SingleBlock struct {
	Header *SingleHeader `json:"header" yaml:"header"`
	Body   *SingleBody   `json:"body" yaml:"body"`
	Next   *SingleBlock  `json:"next" yaml:"next"`
	Back   *SingleBlock  `json:"back" yaml:"back"`
}

func (that *SingleBlock) Verify(backward *SingleBlock) bool {
	body, err := json.Marshal(that.Body)
	if nil != err {
		log.Println(err)
		return false
	}
	bodyHash := hash.SHA512(body)
	if len(bodyHash) != len(that.Header.Merkle) {
		return false
	}
	for i := range bodyHash {
		if bodyHash[i] != that.Header.Merkle[i] {
			return false
		}
	}
	if nil != backward {
		combineHash := hash.SHA512(append(backward.Header.SHA512, bodyHash...))
		if len(combineHash) != len(that.Header.SHA512) {
			return false
		}
		for i := range combineHash {
			if combineHash[i] != that.Header.SHA512[i] {
				return false
			}
		}
	} else {
		if len(bodyHash) != len(that.Header.SHA512) {
			return false
		}
		for i := range bodyHash {
			if bodyHash[i] != that.Header.SHA512[i] {
				return false
			}
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
	head.Back = that.Latest

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
		if !block.Verify(block.Back) {
			return false
		}
	}
	return true
}
