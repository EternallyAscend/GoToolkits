package remote

import (
	"io"
	"os"
	"path"

	"github.com/EternallyAscend/GoToolkits/pkg/network/ssh"
)

func UploadFile(user string, ipv4 string, port uint, password string, publicKeyPath string, localFilePath, remoteFilePath string) error {
	cli := ssh.GenerateDefaultIPv4ClientSSH(user, ipv4, port, password, publicKeyPath)
	err := cli.Connect()
	if nil != err {
		return err
	}

	ftp, err := cli.GetSftpConnect()
	if nil != err {
		return err
	}
	defer ftp.Close()

	remoteFileName := path.Base(localFilePath)

	localFile, err := os.Open(localFilePath)
	if nil != err {
		return err
	}
	defer localFile.Close()

	remoteFile, err := ftp.Create(path.Join(remoteFilePath, remoteFileName))
	if nil != err {
		return err
	}
	defer remoteFile.Close()
	buffer := make([]byte, 1024)
	for {
		n, errIn := localFile.Read(buffer)
		if nil == errIn {
			_, errIn = remoteFile.Write(buffer[:n])
			if nil != errIn {
				return errIn
			}
			continue
		} else if errIn == io.EOF {
			break
		} else {
			return errIn
		}
	}
	return err
}

func DownloadFile(user string, ipv4 string, port uint, password string, publicKeyPath string, localFilePath, remoteFilePath string) error {
	cli := ssh.GenerateDefaultIPv4ClientSSH(user, ipv4, port, password, publicKeyPath)
	err := cli.Connect()
	if nil != err {
		return err
	}

	ftp, err := cli.GetSftpConnect()
	if nil != err {
		return err
	}
	defer ftp.Close()

	remoteFile, err := ftp.Open(remoteFilePath)
	if nil != err {
		return err
	}
	defer remoteFile.Close()
	localFilename := path.Base(remoteFilePath)
	dstFile, err := os.Create(path.Join(localFilePath, localFilename))
	if nil != err {
		return err
	}
	defer dstFile.Close()
	_, err = remoteFile.WriteTo(dstFile)
	return err
}
