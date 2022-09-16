package pedersonCommitment

import "fmt"

func test() {
	//vm := &pedersonCommitment.VerifiableMessage{}
	//fmt.Println(vm)
	//dealer := pedersonCommitment.GenerateDealer()
	//vd := pedersonCommitment.GenerateVerifiableData(dealer)
	//vd.AddCommonData([]byte("Common data abc."))
	//vd.AddEncryptData([]byte("Encrypt data abc."))
	//vd.AddCommonData([]byte("123."))
	//vd.AddEncryptData([]byte("Encrypt data abc."))
	//vd.SetCommonVerifyData([][]byte{
	//	[]byte("Common data abc."),
	//	[]byte("Encrypt data abc."),
	//	[]byte("123."),
	//	[]byte("Encrypt data abc."),
	//})
	//for i := range vd.Units {
	//	if !vd.Units[i].Encrypt {
	//		fmt.Println(vd.Units[i].Dealer.Open(vd.Units[i].Data))
	//		d, e := vd.Units[i].Dealer.TransferToJsonByte()
	//		fmt.Println(string(d), e)
	//	}
	//}
	//fmt.Println(vd.Verify())
	vp := GenerateVerifiableMessage()
	vp.AppendDataArray([][]byte{
		[]byte("Common data abc."),
		[]byte("Encrypt data abc."),
		[]byte("123."),
		[]byte("Encrypt data abc."),
	}, []bool{
		false, true, false, true,
	})
	vp.ConfirmMessage([][]byte{
		[]byte("Common data abc."),
		[]byte("Encrypt data abc."),
		[]byte("123."),
		[]byte("Encrypt data abc."),
	})
	fmt.Println(vp.OpenLine(1), vp.OpenLine(2), vp.CheckAll())
	fmt.Println(vp.ExportMessagesAsString())
}
