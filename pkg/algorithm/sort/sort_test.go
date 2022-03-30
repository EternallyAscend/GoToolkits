package sort

func greaterThanFloat64(left interface{}, right interface{}) bool {
	return left.(float64) > right.(float64)
}

func greaterThanInt(left, right interface{}) bool {
	return left.(int) > right.(int)
}
