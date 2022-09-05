package file

import (
	"log"
	"os"
)

func CreateOrRewrite(data []byte, path string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
	if nil != err {
		return err
	}
	defer func(file *os.File) {
		errIn := file.Close()
		if errIn != nil {
			log.Println(errIn)
		}
	}(file)
	_, err = file.Write(data)
	return err
}
