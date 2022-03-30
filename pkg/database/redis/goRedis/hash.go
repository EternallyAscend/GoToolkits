package goRedis

import "errors"

func (cli *Client) CreateHashKey(maps string, key string, value interface{}) *Result {
	result := ResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	res, err := cli.GetClient().HSetNX(maps, key, value).Result()
	if nil != err {
		result.Err = err
	} else {
		result.Status = res
	}
	return result
}

func (cli *Client) UpdateHashKey(maps string, key string, value interface{}) *Result {
	result := ResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	res, err := cli.GetClient().HSet(maps, key, value).Result()
	if nil != err {
		result.Err = err
	} else {
		// TODO Maybe Here Needs Verify for `res`.
		result.Status = res
	}
	return result
}

func (cli *Client) ReadHashKey(maps string, key string) *Result {
	result := ResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	result.Info, result.Err = cli.GetClient().HGet(maps, key).Result()
	result.Status = true
	return result
}

func (cli *Client) DeleteHashKey(maps string, key string) *Result {
	result := ResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	res, err := cli.GetClient().HDel(maps, key).Result()
	if nil != err {
		result.Err = err
	} else {
		result.Status = StatusDeleted == res
	}
	return result
}

func (cli *Client) SearchHashKey(maps string, pattern string) *ArrayResult {
	result := ArrayResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	result.Info, _, result.Err = cli.GetClient().HScan(maps, 0, pattern, defaultSearchMax).Result()
	result.Status = nil == result.Err
	return result
}

func (cli *Client) ListHash(maps string) *MapResult {
	result := MapResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	result.Info, result.Err = cli.GetClient().HGetAll(maps).Result()
	result.Status = nil == result.Err
	return result
}

func (cli *Client) ListHashKey(maps string) *ArrayResult {
	result := ArrayResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	result.Info, result.Err = cli.GetClient().HKeys(maps).Result()
	result.Status = nil == result.Err
	return result
}

func (cli *Client) ListHashValue(maps string) *ArrayResult {
	result := ArrayResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	result.Info, result.Err = cli.GetClient().HVals(maps).Result()
	result.Status = nil == result.Err
	return result
}

func (cli *Client) DeleteHash(maps string) *Result {
	result := ResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	result.Err = errors.New(ErrorWIP)
	// TODO 完成删除当前哈希键。
	return result
}
