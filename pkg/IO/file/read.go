package file

import (
	"io"
	"log"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0766)
	if nil != err {
		return nil, err
	}
	defer func(file *os.File) {
		errIn := file.Close()
		if errIn != nil {
			log.Println(errIn)
		}
	}(file)
	data, err := io.ReadAll(file)
	return data, err
}
