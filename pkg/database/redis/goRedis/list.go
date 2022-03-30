package goRedis

import "strconv"

func (cli Client) LeftPushList(list string, key string, value interface{}) *Result {
	result := ResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	res, err := cli.GetClient().LPush(list, key, value).Result()
	if nil != err {
		result.Err = err
	} else {
		result.Status = true
		result.Info = strconv.Itoa(int(res))
	}
	return result
}

func (cli Client) RightPushList(list string, key string, value interface{}) *Result {
	result := ResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	res, err := cli.GetClient().RPush(list, key, value).Result()
	if nil != err {
		result.Err = err
	} else {
		result.Status = true
		result.Info = strconv.Itoa(int(res))
	}
	return result
}

func (cli Client) LeftPopList(list string) *Result {
	result := ResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	result.Info, result.Err = cli.GetClient().LPop(list).Result()
	result.Status = nil == result.Err
	return result
}

func (cli Client) RightPopList(list string) *Result {
	result := ResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	result.Info, result.Err = cli.GetClient().RPop(list).Result()
	result.Status = nil == result.Err
	return result
}

func (cli Client) GetListSize(list string) *Result {
	result := ResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	res, err := cli.GetClient().LLen(list).Result()
	if nil != err {
		result.Err = err
	} else {
		result.Status = true
		result.Info = strconv.Itoa(int(res))
	}
	return result
}

func (cli Client) GetAllList(list string) *ArrayResult {
	result := ArrayResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	result.Info, result.Err = cli.GetClient().LRange(list, 0, -1).Result()
	result.Status = nil == result.Err
	return result
}

func (cli Client) GetListRange(list string, start int64, stop int64) *ArrayResult {
	result := ArrayResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	result.Info, result.Err = cli.GetClient().LRange(list, start, stop).Result()
	result.Status = nil == result.Err
	return result
}

func (cli Client) GetListTrim(list string, start int64, stop int64) *Result {
	result := ResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	result.Info, result.Err = cli.GetClient().LTrim(list, start, stop).Result()
	result.Status = nil == result.Err
	return result
}

func (cli Client) RemoveListKey(list string, key string, times int64) *Result {
	result := ResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	res, err := cli.GetClient().LRem(list, times, key).Result()
	if nil != err {
		result.Err = err
	} else {
		result.Status = true
		result.Info = strconv.Itoa(int(res))
	}
	return result
}

func (cli Client) RightPopLeftPushList(list string, key string) *Result {
	result := ResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	result.Info, result.Err = cli.GetClient().RPopLPush(list, key).Result()
	result.Status = nil == result.Err
	return result
}

func (cli Client) UpdateList(list string, key int64, value interface{}) *Result {
	result := ResultPointer()
	result.Err = cli.CheckClient()
	if nil != result.Err {
		return result
	}
	result.Info, result.Err = cli.GetClient().LSet(list, key, value).Result()
	result.Status = nil == result.Err
	return result
}
