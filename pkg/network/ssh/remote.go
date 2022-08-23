package ssh

import (
	"errors"
	"fmt"
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
	client *ssh.Client
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

func (that *Client) GetIPv4AddressAsString() string {
	return fmt.Sprintf("%s:%d", that.ipv4, that.port)
}

func (that *Client) Connect() error {
	if nil == that.client {
		var err error
		that.client, err = ssh.Dial("tcp", that.GetIPv4AddressAsString(), that.clientConfig)
		if nil != err {
			return err
		}
	}
	return nil
}

func (that *Client) ExecuteSingleCommand(c *command.Command) (string, error, error) {
	if nil != that.client {
		session, err := that.client.NewSession()
		defer session.Close()
		if nil != err {
			return "", nil, err
		}
		output, err := session.CombinedOutput(c.GetString())
		return string(output), err, nil
	}
	return "", nil, errors.New("SSH: Must connect before execute commands (client is nil). ")
}

func (that *Client) ExecuteMultiCommands(commands []*command.Command) ([]string, []error, error) {
	if nil != that.client {
		var result []string
		var errorList []error
		for i := range commands {
			//session, err := that.client.NewSession()
			//if nil != err {
			//	return nil, nil, err
			//}
			//output, errIn := session.CombinedOutput(commands[i])
			//result = append(result, string(output))
			//errorList = append(errorList, errIn)
			//session.Close()
			output, errIn, errOut := that.ExecuteSingleCommand(commands[i])
			if nil != errOut {
				return nil, nil, errOut
			}
			result = append(result, string(output))
			errorList = append(errorList, errIn)
		}
		return result, errorList, nil
	}
	return nil, nil, errors.New("SSH: Must connect before execute commands (client is nil). ")
}

func (that *Client) Close() error {
	if nil != that.client {
		return that.client.Close()
	}
	return nil
}
