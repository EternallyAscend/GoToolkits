package IO

import (
	"bufio"
	"io"
	"os"
)

func BufferInputByte(path string, size uint) ([]byte, error) {
	file, err := os.Open(path)
	defer file.Close()
	if nil != err {
		return nil, err
	}
	buffer := make([]byte, size)
	var data []byte
	for {
		count, err := file.Read(buffer)
		if nil != err {
			if io.EOF == err {
				break
			}
		}
		currentBytes := buffer[:count]
		data = append(data, currentBytes...)
	}
	return data, nil
}

func BufferInputString(path string, size uint) (string, error) {
	data, err := BufferInputByte(path, size)
	return string(data), err
}

func LineInputString(path string) ([]string, error) {
	file, err := os.Open(path)
	defer file.Close()
	if nil != err {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data, nil
}
