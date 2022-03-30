package goRedis

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// CreateString 创建永久字符串键值对。
func (cli *Client) CreateString(key string, value interface{}) *Result {
	return cli.CreateStringWithExpire(key, value, 0)
}

// CreateStringWithExpire 创建带过期时间的字符串键值对。
func (cli *Client) CreateStringWithExpire(key string, value interface{}, expireTime time.Duration) *Result {
	result := ResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	res, err := cli.GetClient().SetNX(key, value, expireTime).Result()
	if nil != err {
		result.Err = err
	}
	if !res {
		result.Err = errors.New(ErrorDuplicateKey)
		result.Info = ErrorDuplicateKey
	} else {
		result.Status = true
		result.Info = "OK."
	}
	return result
}

// UpdateString 更新字符串键值对。
func (cli *Client) UpdateString(key string, value interface{}) *Result {
	return cli.UpdateStringWithExpire(key, value, 0)
}

// UpdateStringWithExpire 更新带过期时间带字符串键值对。
func (cli *Client) UpdateStringWithExpire(key string, value interface{}, expireTime time.Duration) *Result {
	result := ResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	result.Info, result.Err = cli.GetClient().Set(key, value, expireTime).Result()
	result.Status = nil == result.Err
	return result
}

// ReadString 获取对应字符串键值。
func (cli *Client) ReadString(key string) *Result {
	result := ResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	result.Info, result.Err = cli.GetClient().Get(key).Result()
	if nil == result.Err {
		result.Status = true
	}
	return result
}

// ReadStringWithExpire 获取对应字符串键值。
func (cli *Client) ReadStringWithExpire() {}

// ReadStrings 获取数个键名对应对键值。
func (cli *Client) ReadStrings(keys ...string) *InterfaceResult {
	result := InterfaceResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	result.Info, result.Err = cli.GetClient().MGet(keys...).Result()
	result.Status = nil == result.Err
	return result
}

// DeleteString 删除对应字符串键值。
func (cli *Client) DeleteString(key string) *Result {
	result := ResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	res, err := cli.GetClient().Del(key).Result()
	if nil != err {
		result.Err = err
	} else {
		result.Status = true
		result.Info = strconv.Itoa(int(res))
	}
	return result
}

// DeleteStrings 删除对应的一组字符串键值。
func (cli *Client) DeleteStrings(keys ...string) *Result {
	result := ResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	res, err := cli.GetClient().Del(keys...).Result()
	if nil != err {
		result.Err = err
	} else {
		result.Status = true
		result.Info = strconv.Itoa(int(res))
	}
	return result
}

// SearchStringWithKey 字符串按键名搜索。
func (cli *Client) SearchStringWithKey(pattern string) *ArrayResult {
	result := ArrayResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	result.Info, result.Err = cli.GetClient().Keys(pattern).Result()
	result.Status = nil == result.Err
	return result
}

// GetSetString 字符串键值对读取原有值并更新为输入值。
func (cli *Client) GetSetString(key string, value interface{}) *Result {
	result := ResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	result.Info, result.Err = cli.GetClient().GetSet(key, value).Result()
	result.Status = nil == result.Err
	return result
}

// AppendString 字符串值追加。
func (cli *Client) AppendString(key string, value string) *Result {
	result := ResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	res, err := cli.GetClient().Append(key, value).Result()
	if nil != err {
		result.Err = err
	} else {
		fmt.Println(res)
		result.Status = true
	}
	return result
}

// UpdateStringExpire 更新字符串超时时间。
func (cli *Client) UpdateStringExpire(key string, expireTime time.Duration) *Result {
	result := ResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	res, err := cli.GetClient().Expire(key, expireTime).Result()
	if nil != err {
		result.Err = err
	} else {
		result.Status = res
	}
	return result
}

// IncreaseString 自增字符串键值。
func (cli *Client) IncreaseString(key string) *Result {
	result := ResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	res, err := cli.GetClient().Incr(key).Result()
	if nil != err {
		result.Err = err
	} else {
		fmt.Println(res)
		result.Status = true
	}
	return result
}

// IncreaseStringByRange 根据提供步长自增字符串键值。
func (cli *Client) IncreaseStringByRange(key string, step int64) *Result {
	result := ResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	res, err := cli.GetClient().IncrBy(key, step).Result()
	if nil != err {
		result.Err = err
	} else {
		fmt.Println(res)
		result.Status = true
	}
	return result
}

// DecreaseString 自减字符串键值。
func (cli *Client) DecreaseString(key string) *Result {
	result := ResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	res, err := cli.GetClient().Decr(key).Result()
	if nil != err {
		result.Err = err
	} else {
		fmt.Println(res)
		result.Status = true
	}
	return result
}

// DecreaseStringByRange 根据提供步长自减字符串键值。
func (cli *Client) DecreaseStringByRange(key string, step int64) *Result {
	result := ResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	res, err := cli.GetClient().DecrBy(key, step).Result()
	if nil != err {
		result.Err = err
	} else {
		fmt.Println(res)
		result.Status = true
	}
	return result
}
