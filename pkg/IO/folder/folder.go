package IO

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileNode struct {
	Sub  []FileNode `json:"sub"`
	Path string     `json:"path"`
	Name string     `json:"name"`
	File bool       `json:"file"`
}

type FileTree struct {
	Root FileNode `json:"root"`
	Err  error    `json:"err"`
}

func GetCurrentPath() string {
	path, _ := os.Getwd()
	return path
}

func GetFolderInfo(path string) ([]string, error) {
	return filepath.Glob(filepath.Join(path, "*"))
}

func GetFolderTree(path string) FileTree {
	tree := FileTree{Root: FileNode{
		Sub:  nil,
		Path: path,
		Name: "",
		File: false,
	}}
	// temp := &tree.Root
	tree.Err = filepath.Walk(path, func(paths string, info os.FileInfo, err error) error {
		fmt.Println(paths, info)
		// TODO Make Right FileTree with Right Path.
		return nil
	})
	return tree
}

func CheckFolderExist(path string) bool {
	return CheckFileExist(path)
}

func CreateFolder(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func DeleteFolder(path string) error {
	return os.RemoveAll(path)
}

func CopyFolder(origin, target string) error {
	return filepath.Walk(origin, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path, info)
		if nil != err {
			// TODO Make Right Error Handler.
			return err
		}
		// TODO Modify Scan Logic with Right Path.
		if info.IsDir() {
			err = CreateFolder(target)
		} else {
			err = CopyFile(origin, target)
		}
		return err
	})
}

func MoveFolder(origin, target string) error {
	return RenameFolder(origin, target)
}

func RenameFolder(origin, target string) error {
	return RenameFile(origin, target)
}
