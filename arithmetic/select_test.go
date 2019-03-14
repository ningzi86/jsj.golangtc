package arithmetic

import (
	"testing"
	"fmt"
)

//选择排序
func TestSelect(t *testing.T) {

	arr := []int{1, 4, 2, 5, 8, 3, 9, 7, 6, 0}

	for i := 0; i < len(arr); i++ {

		min := arr[i]
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			current := arr[j]
			if current < min {
				minIndex = j
				min = current
			}
		}

		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}

	}

	fmt.Println(arr)

}

func TestSelect2(t *testing.T) {

	arr := []int{1, 4, 2, 5, 8, 3, 9, 7, 6, 0}

	for i := 0; i < len(arr); i++ {

		max := arr[i]
		maxIndex := i

		for j := i + 1; j < len(arr); j++ {

			current := arr[j]
			if current > max {
				max = current
				maxIndex = j
			}
		}

		if maxIndex != i {
			arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
		}

	}

	fmt.Println(arr)
}
