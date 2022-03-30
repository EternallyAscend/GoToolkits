package IO

import (
	"io"
	"os"
)

func GetFileInfo(path string) os.FileInfo {
	info, err := os.Stat(path)
	if nil != err {
		return nil
	}
	return info
}

func CheckFileExist(path string) bool {
	_, err := os.Stat(path)
	if nil != err {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func DeleteFile(path string) error {
	return os.Remove(path)
}

func CopyFile(origin, target string) error {
	input, err := os.OpenFile(origin, os.O_RDONLY, 0666)
	if nil != err {
		return err
	}
	output, err := os.Create(target)
	if err != nil {
		return err
	}
	_, err = io.Copy(input, output)
	return err
}

func MoveFile(origin, target string) error {
	return RenameFile(origin, target)
}

func RenameFile(origin, target string) error {
	return os.Rename(origin, target)
}
