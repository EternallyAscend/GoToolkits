package tosone

import (
	"fmt"
	"github.com/tcolgate/mp3"
	"os"
)

func Decode(path string) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	defer file.Close()
	if nil != err {
		fmt.Println(err)
		return
	}
	decoder := mp3.NewDecoder(file)
	var frame mp3.Frame
	decoder.Decode(&frame, nil)
}
