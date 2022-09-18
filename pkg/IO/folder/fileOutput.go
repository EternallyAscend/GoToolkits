package IO

import (
	"bufio"
	"os"
)

/*
flag -rwxrwxrwx
os flag info list:
	os.O_CREATE: 	create if none exists 		不存在则创建
	os.O_RDONLY: 	read-only 					只读
	os.O_WRONLY: 	write-only 					只写
	os.O_RDWR: 		read-write 					可读可写
	os.O_TRUNC: 	truncate when opened 		文件长度截为0：即清空文件
	os.O_APPEND: 	append 						追加新数据到文件
*/

func BufferIOOutputByte(path string, flag int, perm os.FileMode, data []byte) error {
	file, err := os.OpenFile(path, flag, perm)
	defer file.Close()
	if nil != err {
		return err
	}
	writer := bufio.NewWriter(file)
	_, err = writer.Write(data)
	if nil != err {
		return err
	}
	err = writer.Flush()
	return err
}

func BufferOutputByte(path string, flag int, perm os.FileMode, data []byte) error {
	file, err := os.OpenFile(path, flag, perm)
	defer file.Close()
	if nil != err {
		return err
	}
	_, err = file.Write(data)
	if nil != err {
		return err
	}
	err = file.Sync()
	return err
}

func BufferOutputWriteByte(path string, perm os.FileMode, data []byte) error {
	return BufferOutputByte(path, os.O_CREATE, perm, data)
}

func BufferOutputAppendByte(path string, perm os.FileMode, data []byte) error {
	return BufferOutputByte(path, os.O_APPEND, perm, data)
}

func BufferOutputWriteOnlyByte(path string, perm os.FileMode, data []byte) error {
	return BufferOutputByte(path, os.O_WRONLY, perm, data)
}

func BufferOutputClearByte(path string, perm os.FileMode) error {
	// return BufferOutputByte(path, os.O_TRUNC, perm, []byte("\n"))
	return os.WriteFile(path, []byte(""), perm)
}

func BufferIOOutputString(path string, flag int, perm os.FileMode, data string) error {
	file, err := os.OpenFile(path, flag, perm)
	defer file.Close()
	if nil != err {
		return err
	}
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(data)
	if nil != err {
		return err
	}
	err = writer.Flush()
	return err
}

func BufferOutputString(path string, flag int, perm os.FileMode, data string) error {
	file, err := os.OpenFile(path, flag, perm)
	defer file.Close()
	if nil != err {
		return err
	}
	_, err = file.WriteString(data)
	if nil != err {
		return err
	}
	err = file.Sync()
	return err
}

func BufferOutputWriteString(path string, perm os.FileMode, data string) error {
	return BufferOutputString(path, os.O_CREATE, perm, data)
}

func BufferOutputAppendString(path string, perm os.FileMode, data string) error {
	return BufferOutputString(path, os.O_APPEND, perm, data)
}

func BufferOutputWriteOnlyString(path string, perm os.FileMode, data string) error {
	return BufferOutputString(path, os.O_WRONLY, perm, data)
}

func BufferOutputClearString(path string, perm os.FileMode) error {
	// return BufferOutputString(path, os.O_TRUNC, perm, "")
	return os.WriteFile(path, []byte(""), perm)
}
