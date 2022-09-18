package sort

// QuickSortInt 快速排序，整型数组升序修改原素组。
func QuickSortInt(data []int) []int {
	return quickSortIntCore(data, 0, len(data)-1)
}

// QuickSortIntCopy 快速排序，整型数组升序，深拷贝。
func QuickSortIntCopy(data []int) []int {
	result := make([]int, len(data))
	copy(result, data)
	return quickSortIntCore(result, 0, len(data)-1)
}

// quickSortIntCore 快速排序核心，整型数组默认升序修改原数组。
func quickSortIntCore(data []int, left int, right int) []int {
	if right <= left {
		return data
	}
	flag := data[left]
	i, j := left, right
	for j > i {
		for j > i && data[j] >= flag {
			j--
		}
		if j > i {
			data[i] = data[j]
			i++
		}
		for j > i && data[i] <= flag {
			i++
		}
		if j > i {
			data[j] = data[i]
			j--
		}
	}
	data[i] = flag
	quickSortIntCore(data, left, i-1)
	quickSortIntCore(data, j+1, right)
	return data
}

// QuickSortFunc 快速排序，任意类型自定升降序，需要传入大于函数实现，修改至原数组。
func QuickSortFunc(data []interface{}, ascend bool, greaterThan func(interface{}, interface{}) bool) []interface{} {
	return quickSortFuncCore(data, ascend, greaterThan, 0, len(data)-1)
}

// quickSortFuncCore 快速排序核心，任意类型自定升降序，需要传入大于函数实现，修改至原数组。
func quickSortFuncCore(data []interface{}, ascend bool, greaterThan func(interface{}, interface{}) bool, left int, right int) []interface{} {
	if right <= left {
		return data
	}
	flag := data[left]
	i, j := left, right
	if ascend {
		for j > i {
			for j > i && !greaterThan(flag, data[j]) {
				j--
			}
			if j > i {
				data[i] = data[j]
				i++
			}
			for j > i && !greaterThan(data[i], flag) {
				i++
			}
			if j > i {
				data[j] = data[i]
				j--
			}
		}
	} else {
		for j > i {
			for j > i && greaterThan(flag, data[j]) {
				j--
			}
			if j > i {
				data[i] = data[j]
				i++
			}
			for j > i && greaterThan(data[i], flag) {
				i++
			}
			if j > i {
				data[j] = data[i]
				j--
			}
		}
	}
	data[i] = flag
	quickSortFuncCore(data, ascend, greaterThan, left, i-1)
	quickSortFuncCore(data, ascend, greaterThan, j+1, right)
	return data
}
