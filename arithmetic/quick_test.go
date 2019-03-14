package arithmetic

import (
	"testing"
	"fmt"
)

//快速排序
func TestQuickSort(t *testing.T) {

	arr := []int{1, 4, 2, 5, 8, 3, 9, 7, 6, 0}

	result := quickSort(arr)
	fmt.Println(result)
}

func quickSort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}

	current := arr[0]
	var left []int
	var right []int

	for i := 0; i < len(arr); i++ {

		if arr[i] < current {
			left = append(left, arr[i])
		}

		if arr[i] > current {
			right = append(right, arr[i])
		}

	}

	leftResult := append(quickSort(left), current)
	rightResult := quickSort(right)

	return append(leftResult, rightResult...)

}
