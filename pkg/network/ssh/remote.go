package ssh

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

	"github.com/EternallyAscend/GoToolkits/pkg/IO/YAML"
	"github.com/EternallyAscend/GoToolkits/pkg/command"
	"golang.org/x/crypto/ssh"
)

type SavedIPv4Client struct {
	Ipv4 string `yaml:"ipv4"`
	Port uint   `yaml:"port"`
	User string `yaml:"user"`
	Pwd  string `yaml:"pwd"`
	Puk  string `yaml:"puk"`
}

func ReadPwdClientFromYaml(yamlPath string) (*SavedIPv4Client, error) {
	cli := &SavedIPv4Client{}
	//byteData, err := file.ReadFile(yamlPath)
	//if nil != err {
	//	return nil, err
	//}
	//err = yaml.Unmarshal(byteData, cli)
	//fmt.Println(string(byteData), *cli)
	err := YAML.ReadStructFromFileYaml(cli, yamlPath)
	return cli, err
}

func (that *SavedIPv4Client) CreateClient() *IPv4Client {
	return GenerateDefaultIPv4ClientSSH(that.User, that.Ipv4, that.Port, that.Pwd, that.Puk)
}

type IPv4Client struct {
	err          error
	ipv4         string
	port         uint
	user         string
	clientConfig *ssh.ClientConfig
	client       *ssh.Client
}

func GenerateDefaultIPv4ClientSSH(user string, ipv4 string, port uint, password string, publicKeyPath string) *IPv4Client {
	cli := &IPv4Client{
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
		key, err := os.ReadFile(publicKeyPath)
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

func (that *IPv4Client) GetIPv4AddressAsString() string {
	return fmt.Sprintf("%s:%d", that.ipv4, that.port)
}

func (that *IPv4Client) Connect() error {
	if nil == that.client {
		var err error
		that.client, err = ssh.Dial("tcp", that.GetIPv4AddressAsString(), that.clientConfig)
		if nil != err {
			return err
		}
	}
	return nil
}

func (that *IPv4Client) ExecuteSingleCommand(c *command.Command) (string, error, error) {
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

func (that *IPv4Client) ExecuteMultiCommands(commands []*command.Command) ([]string, []error, error) {
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

func (that *IPv4Client) ExecuteMultiParallelCommands(commands []*command.Command) ([]string, []error, error) {
	result := make([]string, len(commands))
	errorList := make([]error, len(commands))
	var err error
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(commands))
	for i := range commands {
		go func(i int) {
			var errIn error
			result[i], errorList[i], errIn = that.ExecuteSingleCommand(commands[i])
			if nil != err {
				err = errIn
			}
			log.Println(commands[i])
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()
	return result, errorList, err
}

func (that *IPv4Client) Close() error {
	if nil != that.client {
		return that.client.Close()
	}
	return nil
}
