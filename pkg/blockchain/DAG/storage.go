package DAG

import (
	"encoding/json"
	"log"

	"github.com/EternallyAscend/GoToolkits/pkg/cryptography/hash"
	"github.com/EternallyAscend/GoToolkits/pkg/cryptography/homomorphic/pedersonCommitment"
	"github.com/bwesterb/go-ristretto"
)

type Header struct {
	// Blockchain Hash.
	Merkle []byte `json:"merkle" yaml:"merkle"`
	SHA512 []byte `json:"sha512" yaml:"sha512"`
	// Dealer Shows H.E. Arguments GH and Sum of Random and Commitment.
	Dealer *pedersonCommitment.DealerUnit `json:"dealer" yaml:"dealer"`
}

// GenerateHeader Before Data Transfer, Exchange Arguments G and H.
func (that *Body) GenerateHeader(dealer *pedersonCommitment.DealerUnit) *Header {
	// TODO Start with HeaderRandom.
	header := &Header{
		Merkle: nil,
		SHA512: nil,
		Dealer: dealer,
	}
	return header
}

type Body struct {
	// H.E. Verify Information.
	Units []*pedersonCommitment.VerifiableMessageUint
	// Compression Gradients.
	Data [][]byte
}

func (that *Body) Hash() []byte {
	byteData, err := json.Marshal(that)
	if nil != err {
		log.Println(err)
		return nil
	}
	// TODO Calculate Body Byte as Merkle.
	return byteData
}

func (that *Body) HashString() string {
	return hash.SHA512String(that.Hash())
}

// Block DAG Data Structure.
type Block struct {
	Header    *Header  `json:"header" yaml:"header"`
	Body      *Body    `json:"block" yaml:"block"`
	Reference []*Block `json:"reference" yaml:"reference"`
}

// CalculateHeader After Block AppendData Finished.
func (that *Block) CalculateHeader() {
	body, err := json.Marshal(that.Body)
	if nil != err {
		log.Println(err)
		return
	}
	that.Header.Merkle = hash.SHA512(body)
	// Calculate the Sum of Random and Commit Result.
	for i := range that.Body.Units {
		that.Header.Dealer.R.Add(that.Header.Dealer.R, that.Body.Units[i].R)
		that.Header.Dealer.CommitPoint.Add(that.Header.Dealer.CommitPoint, that.Body.Units[i].CommitPoint)
	}
	var data []byte
	for i := range that.Reference {
		data = append(data, that.Reference[i].Header.SHA512...)
	}
	data = append(data, that.Header.Merkle...)
	that.Header.SHA512 = hash.SHA512(data)
}

func (that *Block) AppendData(data []byte) {
	that.Body.Data = append(that.Body.Data, data)
	// Calculate Append Data with GH.

	r := &ristretto.Scalar{}
	r.Rand()

	c := &ristretto.Point{}

	tempPoint := ristretto.Point{}
	// c = xG + rH
	x := ristretto.Scalar{}
	x.Derive(data)
	// xG
	tempPoint.ScalarMult(that.Header.Dealer.G, &x)
	// rH
	c.ScalarMult(that.Header.Dealer.H, r)
	// c = xG + rH
	c.Add(c, &tempPoint)

	unit := &pedersonCommitment.VerifiableMessageUint{
		R:           r,
		CommitPoint: c,
	}
	// Append Unit into Body.
	that.Body.Units = append(that.Body.Units, unit)
}

// Packaged Calculate Header.
func (that *Block) Packaged() {
	// Calculate Random and Commitment Summary.
	that.CalculateHeader()
	// TODO End with TailRandom Value.
}

// Verify Verify Block Correction.
func (that *Block) Verify() bool {
	if len(that.Body.Data) != len(that.Body.Units) {
		return false
	}
	if len(that.Body.Data) > 0 {
		tr := ristretto.Scalar{}
		tc := ristretto.Point{}
		tr.Set(that.Body.Units[0].R)
		tc.Set(that.Body.Units[0].CommitPoint)
		for i := 1; i < len(that.Body.Data); i++ {
			tr.Add(&tr, that.Body.Units[i].R)
			tc.Add(&tc, that.Body.Units[i].CommitPoint)
		}
		return that.Header.Dealer.CommitPoint.Equals(&tc) && that.Header.Dealer.R.Equals(&tr)
	} else {
		return true
	}
}

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
