package ssh

import (
	"github.com/EternallyAscend/GoToolkits/pkg/command"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"net"
	"time"
)

type Client struct {
	err          error
	ipv4         string
	port         uint
	user         string
	clientConfig *ssh.ClientConfig
}

func GenerateDefaultClientSSH(user string, ipv4 string, port uint, password string, publicKeyPath string) *Client {
	cli := &Client{
		err:  nil,
		ipv4: ipv4,
		port: port,
		user: user,
		clientConfig: &ssh.ClientConfig{
			Config: ssh.Config{},
			User:   user,
			Auth:   []ssh.AuthMethod{},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			},
			BannerCallback: func(message string) error {
				return nil
			},
			ClientVersion: "",
			Timeout:       time.Second,
		},
	}
	if "" == publicKeyPath {
		cli.clientConfig.Auth = []ssh.AuthMethod{
			ssh.Password(password),
		}
	} else {
		key, err := ioutil.ReadFile(publicKeyPath)
		if nil != err {
			cli.err = err
			return cli
		}
		signer, err := ssh.ParsePrivateKey(key)
		if nil != err {
			cli.err = err
			return cli
		}
		cli.clientConfig.Auth = []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		}
	}
	return cli
}

func (that *Client) Connect() {}

func (that *Client) Execute(commands []*command.Command) []*command.Result {
	return nil
}

func (that *Client) Close() {}
