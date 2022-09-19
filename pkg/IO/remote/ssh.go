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
	if err != nil {
		return err
	}
	defer localFile.Close()

	remoteFile, err := ftp.Create(path.Join(remoteFilePath, remoteFileName))
	if err != nil {
		return err
	}
	defer remoteFile.Close()
	buffer := make([]byte, 1024)
	for {
		n, errIn := localFile.Read(buffer)
		if errIn != nil {
			if errIn == io.EOF {
				break
			} else {
				return err
			}
		}
		remoteFile.Write(buffer[:n])
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
	if err != nil {
		return err
	}
	defer remoteFile.Close()
	localFilename := path.Base(remoteFilePath)
	dstFile, err := os.Create(path.Join(localFilePath, localFilename))
	if err != nil {
		return err
	}
	defer dstFile.Close()
	_, err = remoteFile.WriteTo(dstFile)
	return err
}
