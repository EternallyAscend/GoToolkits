package sort

import (
	"fmt"
	"testing"
)

func TestMergeSortInt(t *testing.T) {
	data := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println(MergeSortIntCopy(data), data)
	fmt.Println(MergeSortInt(data), data)
	fmt.Println(MergeSortInt(data), data)
	fmt.Println(MergeSortIntCopy([]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}))
}

func TestMergeSortFunc(t *testing.T) {
	data := []float64{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	input := make([]interface{}, len(data))
	for i := range data {
		input[i] = data[i]
	}
	fmt.Println(MergeSortFunc(input, true, greaterThanFloat64), input)
	fmt.Println(MergeSortFunc(input, true, greaterThanFloat64), input)
	fmt.Println(MergeSortFunc(input, false, greaterThanFloat64), input)
	fmt.Println(MergeSortFunc(input, false, greaterThanFloat64), input)
	output := make([]float64, len(input))
	for i := range input {
		output[i] = input[i].(float64)
	}
	fmt.Println(output)
}

func TestMergeSortIntCopy(t *testing.T) {
	data := make([]int, 99999999)
	MergeSortInt(data)
}
