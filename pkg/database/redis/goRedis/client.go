package goRedis

import (
	"crypto/tls"
	"errors"
	"time"

	"github.com/go-redis/redis"
)

// Client 封装客户端所需参数的结构体。
type Client struct {
	Address  string        `json:"address"`
	Password string        `json:"password"`
	Database int           `json:"database"`
	ReadOnly bool          `json:"read_only"`
	Args     []interface{} `json:"args"`
	Client   *redis.Client `json:"client"`
	TLS      *tls.Config   `json:"tls"`
}

// Connect 客户端连接函数，大部分采取默认配置。
func (cli *Client) Connect() error {
	if cli.Client == nil {
		cli.Client = redis.NewClient(&redis.Options{
			Network:            "tcp",                  // default tcp, or unix
			Addr:               cli.Address,            // Redis数据库地址
			Dialer:             nil,                    // 自定义连接函数
			OnConnect:          nil,                    // 连接后钩子函数
			Password:           cli.Password,           // 数据库密码
			DB:                 cli.Database,           // 数据库页面
			MaxRetries:         3,                      // 执行命令失败最大重试次数，默认为0不重试。
			MinRetryBackoff:    8 * time.Millisecond,   // 重试最小间隔，-1为取消。
			MaxRetryBackoff:    512 * time.Millisecond, // 重试最大间隔，-1为取消
			DialTimeout:        5 * time.Second,        // 建立连接超时时间
			ReadTimeout:        3 * time.Second,        // 读超时，默认3秒，-1取消超时。
			WriteTimeout:       4 * time.Second,        // 写超时，默认等同于读超时。
			PoolSize:           8,                      // 连接池大小
			MinIdleConns:       2,                      // 启动时创建连接，并始终保持的连接数。在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量。
			MaxConnAge:         0,                      // 最大连接时间，从创建起计时，默认为0不关闭长时间存活的连接。
			PoolTimeout:        4 * time.Second,        // 繁忙超时，客户端等待可用连接的最大等待时长，默认读超时+1秒。
			IdleTimeout:        5 * time.Minute,        // 闲置超时，默认5分钟，-1取消闲置超时检查。
			IdleCheckFrequency: 1 * time.Minute,        // 闲置连接检查周期，默认1分钟，-1不进行周期性检查。
			TLSConfig:          cli.TLS,                // TLS配置信息
		})
	}
	_, err := cli.Client.Ping().Result()
	return err
}

// ConnectCustom 客户端连接函数，完全自定义连接方式。
func (cli *Client) ConnectCustom(connect func(client *Client, args []interface{}) error) error {
	return connect(cli, cli.Args)
}

// GetClient 返回对应的Client指针。
func (cli *Client) GetClient() *redis.Client {
	return cli.Client
}

// CheckClient 检查连接。
func (cli *Client) CheckClient() error {
	if nil == cli.GetClient() {
		defer func(cli *Client) {
			_ = cli.Close()
		}(cli)
		return errors.New(ErrorClientIsNil)
	}
	return nil
}

// Close 关闭数据库连接。
func (cli *Client) Close() error {
	if nil != cli.Client {
		return cli.Client.Close()
	}
	return nil
}

// Save 数据存储。
func (cli *Client) Save() {
	err := cli.CheckClient()
	if nil != err {
		return
	}
	cli.GetClient().Save()
}

// BGSave 数据持久化存储。
func (cli *Client) BGSave() {
	err := cli.CheckClient()
	if nil != err {
		return
	}
	cli.GetClient().BgSave()
}

// ExecuteDo 直接处理Redis命令。
func (cli *Client) ExecuteDo(args ...interface{}) *InterfaceResult {
	result := InterfaceResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	result.Info, result.Err = cli.GetClient().Do(args...).Result()
	return result
}
