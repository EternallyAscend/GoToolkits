package pedersonCommitment

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"
)

func TestPedersonCommitment(t *testing.T) {
	filePath := "./test.txt"
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0766)
	if nil != err {
		log.Println(err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	data, err := ioutil.ReadAll(f)
	if nil != err {
		log.Println(err)
	}
	var d [][]byte
	var c []bool
	for i := 1; i <= 8192; i *= 2 {
		for j := len(d) / 9999; j < i; j++ {
			for k := len(data) - 1; k > 0; k-- {
				d = append(d, []byte{data[k]})
				c = append(c, true)
			}
		}
		log.Println(len(d), len(c))
		start := time.Now().UnixMicro()
		// time.Sleep(time.Second)
		vp := GenerateVerifiableMessage()
		vp.AppendDataArray(d, c)
		vp.ConfirmMessage(d)
		middle := time.Now().UnixMicro()
		vp.CheckAll()
		end := time.Now().UnixMicro()
		log.Println(i, " : ", end-start, "ms ", middle-start, "ms ", end-middle, "ms")
	}
}
