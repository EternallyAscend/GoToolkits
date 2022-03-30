package goRedis

// Result 返回结果封装。
type Result struct {
	Status bool
	Info   string
	Err    error
}

// ResultPointer 空结果指针。
func ResultPointer() *Result {
	return &Result{
		Status: false,
		Info:   "",
		Err:    nil,
	}
}

// MapResult 字符串Map结果封装。
type MapResult struct {
	Status bool
	Info   map[string]string
	Err    error
}

func MapResultPointer() *MapResult {
	return &MapResult{
		Status: false,
		Info:   nil,
		Err:    nil,
	}
}

// ArrayResult 字符串数组结果封装。
type ArrayResult struct {
	Status bool
	Info   []string
	Err    error
}

// ArrayResultPointer 字符串数组结果返回封装。
func ArrayResultPointer() *ArrayResult {
	return &ArrayResult{
		Status: false,
		Info:   nil,
		Err:    nil,
	}
}

// InterfaceResult 任意类型结果返回封装。
type InterfaceResult struct {
	Status bool
	Info   interface{}
	Err    error
}

// InterfaceResultPointer 任意类型空结果指针。
func InterfaceResultPointer() *InterfaceResult {
	return &InterfaceResult{
		Status: false,
		Info:   nil,
		Err:    nil,
	}
}
