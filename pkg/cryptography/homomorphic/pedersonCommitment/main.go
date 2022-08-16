package pedersonCommitment

import (
	"encoding/json"
	"github.com/bwesterb/go-ristretto"
	"gopkg.in/yaml.v2"
)

// generateParams 椭圆曲线参数，由被承诺方提供
func generateParams() (*ristretto.Point, *ristretto.Point) {
	g := ristretto.Point{}
	h := ristretto.Point{}
	return g.Rand(), h.Rand()
}

// generateRandom 随机数值，由承诺方提供
func generateRandom() *ristretto.Scalar {
	r := ristretto.Scalar{}
	return r.Rand()
}

// DealerUnit Pederson承诺处理单元
type DealerUnit struct {
	G *ristretto.Point
	H *ristretto.Point
	R *ristretto.Scalar
	CommitPoint *ristretto.Point
}

// GenerateDealer 被承诺方生成Dealer
func GenerateDealer() *DealerUnit {
	g, h := generateParams()
	return &DealerUnit{
		G: g,
		H: h,
	}
}

func (that *DealerUnit) generateRandom() *DealerUnit {
	that.R = generateRandom()
	return that
}

// Commit 承诺方提交承诺信息
func (that *DealerUnit) Commit(content []byte) *DealerUnit {
	// 承诺方提供随机数r
	that.generateRandom()
	that.CommitPoint = &ristretto.Point{}
	tempPoint := ristretto.Point{}
	// c = xG + rH
	x := ristretto.Scalar{}
	x.Derive(content)
	// xG
	tempPoint.ScalarMult(that.G, &x)
	// rH
	that.CommitPoint.ScalarMult(that.H, that.R)
	// c = xG + rH
	that.CommitPoint.Add(that.CommitPoint, &tempPoint)
	return that
}

func (that *DealerUnit) Open(content []byte) bool {
	compare := ristretto.Point{}
	tempPoint := ristretto.Point{}

	x := ristretto.Scalar{}
	x.Derive(content)

	// x'G
	tempPoint.ScalarMult(that.G, &x)
	// rH
	compare.ScalarMult(that.H, that.R)
	// c' = x'G + rH
	compare.Add(&compare, &tempPoint)
	// c' == c
	return compare.Equals(that.CommitPoint)
}

func (that *DealerUnit) CopyParams() *DealerUnit {
	// TODO Test using that.G, that.H directly.
	//g := &ristretto.Point{}
	//h := &ristretto.Point{}
	//g.Set(that.G)
	//h.Set(that.H)
	return &DealerUnit{
		G: that.G,
		H: that.H,
	}
}

func (that *DealerUnit) TransferToJsonByte() ([]byte, error) {
	return json.Marshal(that)
}

func (that *DealerUnit) TransferToYamlByte() ([]byte, error) {
	return yaml.Marshal(that)
}

// https://blog.csdn.net/qq_35739903/article/details/119453339

// https://zhuanlan.zhihu.com/p/108659500

// https://www.codeleading.com/article/59343262039/

type VerifiableMessageUint struct {
	R *ristretto.Scalar
	CommitPoint *ristretto.Point
}

type VerifiableMessage struct {
	Dealer *DealerUnit
	Data [][]byte
	Units []*VerifiableMessageUint
}

func (that *VerifiableMessage) AppendData(data []byte, encrypt bool) *VerifiableMessage {
	if encrypt {
		that.Data = append(that.Data, nil)
	} else {
		that.Data = append(that.Data, data)
	}
	return that
}

func (that *VerifiableMessage) AppendDataArray(data [][]byte, encrypt []bool) *VerifiableMessage {
	if len(data) != len(encrypt) {

	}
	for i := range data {
		that.AppendData(data[i], encrypt[i])
	}
	return that
}

func (that *VerifiableMessage) OpenLine(line uint) bool {
	return true
}

func (that *VerifiableMessage) Verify() bool {

	return true
}

func (that *VerifiableMessage) CheckAll() bool {
	for i := range that.Data {
		if nil != that.Data[i] {
			if !that.OpenLine(uint(i)) {
				return false
			}
		}
	}
	return that.Verify()
}

type VerifiableDataUnit struct {
	Dealer *DealerUnit
	Encrypt bool
	Data []byte
}

type VerifiableData struct {
	Units []*VerifiableDataUnit
	DealerUint *DealerUnit
	VerifierData [][]byte
}

func GenerateVerifiableData(dealer *DealerUnit) *VerifiableData {
	return &VerifiableData{
		Units:        []*VerifiableDataUnit{},
		DealerUint:   dealer,
		VerifierData: nil,
	}
}

func (that *VerifiableData) AddCommonData(data []byte) *VerifiableData {
	dataUint := &VerifiableDataUnit{
		Dealer:  that.DealerUint.CopyParams(),
		Encrypt: false,
		Data:    data,
	}
	dataUint.Dealer.Commit(data)
	that.Units = append(that.Units, dataUint)
	return that
}

func (that *VerifiableData) AddEncryptData(data []byte) *VerifiableData {
	dataUint := &VerifiableDataUnit{
		Dealer:  that.DealerUint.CopyParams(),
		Encrypt: true,
		Data:    nil,
	}
	dataUint.Dealer.Commit(data)
	that.Units = append(that.Units, dataUint)
	return that
}

func (that *VerifiableData) SetCommonVerifyData(data [][]byte) *VerifiableData {
	that.SetEncryptVerifyData(data)
	that.VerifierData = data
	return that
}

func (that *VerifiableData) SetEncryptVerifyData(data [][]byte) *VerifiableData {
	if len(data) != len(that.Units) {
		return that
	}
	if len(data) > 0 {
		x := ristretto.Scalar{}
		x.Derive(data[0])
		that.DealerUint.R = &ristretto.Scalar{}
		that.DealerUint.R.Set(that.Units[0].Dealer.R)
		tempScalar := ristretto.Scalar{}
		if len(data) > 1 {
			for i := 1; i < len(data); i++ {
				if nil != data[i] {
					tempScalar.Derive(data[i])
					x.Add(&x, &tempScalar)
					that.DealerUint.R.Add(that.DealerUint.R, that.Units[i].Dealer.R)
				}
			}
		}
		that.DealerUint.CommitPoint = &ristretto.Point{}
		tempPoint := ristretto.Point{}
		tempPoint.ScalarMult(that.DealerUint.G, &x)
		that.DealerUint.CommitPoint.ScalarMult(that.DealerUint.H, that.DealerUint.R)
		that.DealerUint.CommitPoint.Add(that.DealerUint.CommitPoint, &tempPoint)
	}
	return that
}

// Verify 验证全部数据的完整性
func (that *VerifiableData) Verify() bool {
	if len(that.Units) > 1 {
		if nil == that.DealerUint {
			return false
		}
		if nil == that.Units[0].Dealer || nil == that.Units[1].Dealer {
			return false
		}
		sumR := ristretto.Scalar{}
		sumR.Add(that.Units[0].Dealer.R, that.Units[1].Dealer.R)
		sumC := ristretto.Point{}
		sumC.Add(that.Units[0].Dealer.CommitPoint, that.Units[1].Dealer.CommitPoint)
		for index := 2; index < len(that.Units); index++ {
			if nil == that.Units[index].Dealer {
				return false
			}
			sumR.Add(&sumR, that.Units[index].Dealer.R)
			sumC.Add(&sumC, that.Units[index].Dealer.CommitPoint)
		}
		return that.DealerUint.R.Equals(&sumR) && that.DealerUint.CommitPoint.Equals(&sumC)
	} else if 1 == len(that.Units) {
		if nil == that.DealerUint {
			return false
		}
		if nil == that.Units[0].Dealer {
			return false
		}
		return that.DealerUint.R.Equals(that.Units[0].Dealer.R) || that.DealerUint.CommitPoint.Equals(that.Units[0].Dealer.CommitPoint)
	} else if nil == that.VerifierData {
		return true
	} else {
		return false
	}
}
