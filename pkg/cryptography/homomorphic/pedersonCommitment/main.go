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
	G           *ristretto.Point
	H           *ristretto.Point
	R           *ristretto.Scalar
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
	R           *ristretto.Scalar
	CommitPoint *ristretto.Point
}

// VerifiableMessage 可验证消息，VP中的单条Verifiable Credentials
type VerifiableMessage struct {
	Dealer *DealerUnit
	Data   [][]byte
	Units  []*VerifiableMessageUint
}

// GenerateVerifiableMessage 生成一个空的可验证消息，含有公钥
func GenerateVerifiableMessage() *VerifiableMessage {
	return &VerifiableMessage{
		Dealer: GenerateDealer(),
		Data:   [][]byte{},
		Units:  []*VerifiableMessageUint{},
	}
}

// AppendData 向单条可验证消息VC中增加一行可验证消息
func (that *VerifiableMessage) AppendData(data []byte, encrypt bool) *VerifiableMessage {
	if encrypt {
		that.Data = append(that.Data, nil)
	} else {
		that.Data = append(that.Data, data)
	}

	r := &ristretto.Scalar{}
	r.Rand()
	c := &ristretto.Point{}

	tempPoint := ristretto.Point{}
	// c = xG + rH
	x := ristretto.Scalar{}
	x.Derive(data)
	// xG
	tempPoint.ScalarMult(that.Dealer.G, &x)
	// rH
	c.ScalarMult(that.Dealer.H, r)
	// c = xG + rH
	c.Add(c, &tempPoint)

	if 0 == len(that.Units) {
		that.Dealer.R = &ristretto.Scalar{}
		that.Dealer.R.Set(r)
	} else {
		that.Dealer.R.Add(that.Dealer.R, r)
	}

	unit := &VerifiableMessageUint{
		R:           r,
		CommitPoint: c,
	}
	that.Units = append(that.Units, unit)
	return that
}

// AppendDataArray 向单条可验证消息VC中增加一组可验证消息
func (that *VerifiableMessage) AppendDataArray(data [][]byte, encrypt []bool) *VerifiableMessage {
	if len(data) != len(encrypt) {
		return nil
	}
	for i := range data {
		that.AppendData(data[i], encrypt[i])
	}
	return that
}

func (that *VerifiableMessage) ConfirmMessage(data [][]byte) *VerifiableMessage {
	if nil == that.Dealer.R {
		return nil
	}
	if len(that.Data) != len(that.Units) {
		return nil
	}
	if len(data) != len(that.Data) {
		return nil
	}
	x := ristretto.Scalar{}
	x.Derive(data[0])
	tx := ristretto.Scalar{}
	tx.Set(&x)
	for i := 1; i < len(data); i++ {
		x.Derive(data[i])
		tx.Add(&tx, &x)
	}

	tempPoint := ristretto.Point{}
	// c = xG + rH
	that.Dealer.CommitPoint = &ristretto.Point{}
	// xG
	tempPoint.ScalarMult(that.Dealer.G, &tx)
	// rH
	that.Dealer.CommitPoint.ScalarMult(that.Dealer.H, that.Dealer.R)
	// c = xG + rH
	that.Dealer.CommitPoint.Add(that.Dealer.CommitPoint, &tempPoint)

	return that
}

// OpenLine 打开提供的一行数据并验证正确性
func (that *VerifiableMessage) OpenLine(line uint) bool {
	if len(that.Units) != len(that.Data) {
		return false
	}
	if int(line) > len(that.Data) {
		return false
	}
	if nil == that.Data[line] {
		return false
	}

	tc := ristretto.Point{}

	tempPoint := ristretto.Point{}
	// c = xG + rH
	x := ristretto.Scalar{}
	x.Derive(that.Data[line])
	// xG
	tempPoint.ScalarMult(that.Dealer.G, &x)
	// rH
	tc.ScalarMult(that.Dealer.H, that.Units[line].R)
	// c = xG + rH
	tc.Add(&tc, &tempPoint)
	return tc.Equals(that.Units[line].CommitPoint)
}

// Verify 验证可验证消息的完整性
func (that *VerifiableMessage) Verify() bool {
	if len(that.Data) != len(that.Units) {
		return false
	}
	if len(that.Data) > 0 {
		tr := ristretto.Scalar{}
		tc := ristretto.Point{}
		tr.Set(that.Units[0].R)
		tc.Set(that.Units[0].CommitPoint)
		for i := 1; i < len(that.Data); i++ {
			tr.Add(&tr, that.Units[i].R)
			tc.Add(&tc, that.Units[i].CommitPoint)
		}
		return that.Dealer.CommitPoint.Equals(&tc) && that.Dealer.R.Equals(&tr)
	} else {
		return true
	}
}

// CheckAll 检查可验证消息的完整性并验证验证提供的每一条数据正确性
func (that *VerifiableMessage) CheckAll() bool {
	for i := range that.Data {
		if nil != that.Data[i] {
			if !that.OpenLine(uint(i)) {
				return false
			}
		}
	}
	// Reach here means every common lines are right. Only need to verify.
	return that.Verify()
}

func (that *VerifiableMessage) ExportMessagesAsByteArray() [][]byte {
	var data [][]byte
	for i := range that.Data {
		if nil != that.Data[i] {
			data = append(data, that.Data[i])
		}
	}
	return data
}

func (that *VerifiableMessage) ExportMessagesAsString() []string {
	var data []string
	for i := range that.Data {
		if nil != that.Data[i] {
			data = append(data, string(that.Data[i]))
		}
	}
	return data
}

type VerifiableDataUnit struct {
	Dealer  *DealerUnit
	Encrypt bool
	Data    []byte
}

type VerifiableData struct {
	Units        []*VerifiableDataUnit
	DealerUint   *DealerUnit
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
