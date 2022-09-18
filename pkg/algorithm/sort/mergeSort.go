package sort

// MergeSortInt 归并排序，整型数组可用，直接修改原数组，稳定升序。
func MergeSortInt(data []int) []int {
	return mergeSortIntCore(data, 0, len(data))
}

// MergeSortIntCopy 归并排序，整型数组可用，深拷贝不影响原数组，稳定升序。
func MergeSortIntCopy(data []int) []int {
	result := make([]int, len(data))
	copy(result, data)
	return mergeSortIntCore(result, 0, len(data))
}

// mergeSortIntCore 归并排序核心，整型数组，直接修改数组，稳定升序。
func mergeSortIntCore(data []int, start int, end int) []int {
	if end <= start+1 {
		return data
	}
	middle := (start + end) >> 1
	//waitGroup := sync.WaitGroup{}
	//waitGroup.Add(2)
	//go func() {
	mergeSortIntCore(data, start, middle)
	//waitGroup.Done()
	//}()
	//go func() {
	mergeSortIntCore(data, middle, end)
	//waitGroup.Done()
	//}()
	leftCursor := start
	rightCursor := middle
	resultCursor := 0
	result := make([]int, end-start)
	//waitGroup.Wait()
	for leftCursor < middle && rightCursor < end {
		if data[leftCursor] > data[rightCursor] {
			result[resultCursor] = data[rightCursor]
			rightCursor++
		} else {
			result[resultCursor] = data[leftCursor]
			leftCursor++
		}
		resultCursor++
	}
	for leftCursor < middle {
		result[resultCursor] = data[leftCursor]
		leftCursor++
		resultCursor++
	}
	for i := 0; i < resultCursor; i++ {
		data[start+i] = result[i]
	}
	return data
}

// MergeSortFunc 归并排序，任意类型自定义升降序，需要传入大于判断函数实现。修改原数组，稳定排序。
func MergeSortFunc(data []interface{}, ascend bool, greaterThan func(interface{}, interface{}) bool) []interface{} {
	return mergeSortFuncCore(data, ascend, greaterThan, 0, len(data))
}

// mergeSortFuncCore 归并排序，任意类型自定义升降序，需要传入大于判断函数实现。修改原数组，稳定排序。
func mergeSortFuncCore(data []interface{}, ascend bool, greaterThan func(interface{}, interface{}) bool, start int, end int) []interface{} {
	if end <= start+1 {
		return data
	}
	middle := (start + end) >> 1
	mergeSortFuncCore(data, ascend, greaterThan, start, middle)
	mergeSortFuncCore(data, ascend, greaterThan, middle, end)
	leftCursor := start
	rightCursor := middle
	resultCursor := 0
	result := make([]interface{}, end-start)
	if ascend {
		for leftCursor < middle && rightCursor < end {
			if greaterThan(data[leftCursor], data[rightCursor]) {
				result[resultCursor] = data[rightCursor]
				rightCursor++
			} else {
				result[resultCursor] = data[leftCursor]
				leftCursor++
			}
			resultCursor++
		}
	} else {
		for leftCursor < middle && rightCursor < end {
			if greaterThan(data[rightCursor], data[leftCursor]) {
				result[resultCursor] = data[rightCursor]
				rightCursor++
			} else {
				result[resultCursor] = data[leftCursor]
				leftCursor++
			}
			resultCursor++
		}
	}
	for leftCursor < middle {
		result[resultCursor] = data[leftCursor]
		leftCursor++
		resultCursor++
	}
	for i := 0; i < resultCursor; i++ {
		data[start+i] = result[i]
	}
	return data
}
