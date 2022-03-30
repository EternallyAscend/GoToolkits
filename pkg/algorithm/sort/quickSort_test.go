package sort

import (
	"fmt"
	"testing"
)

func TestQuickSortInt(t *testing.T) {
	data := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println(QuickSortInt(data), data)
	input := make([]interface{}, len(data))
	for i := range data {
		input[i] = data[i]
	}
	fmt.Println(QuickSortFunc(input, true, greaterThanInt), input)
	fmt.Println(QuickSortFunc(input, true, greaterThanInt), input)
	fmt.Println(QuickSortFunc(input, false, greaterThanInt), input)
	fmt.Println(QuickSortFunc(input, false, greaterThanInt), input)
}
