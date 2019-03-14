package arithmetic

import (
	"testing"
	"fmt"
	"math"
)

func TestRecursion(t *testing.T) {

	var arrs = []int{1, 3, 4, 6, 2, 1, 3, 1,11,33,22}

	fmt.Println(GetResults(arrs))
	fmt.Println(GetResults2(arrs))

	var arrs2 = []float64{1, 3, 4, 6, 2, 1, 3, 1,11,33,22 }
	max := *GetResults3(arrs2)
	fmt.Println(max)
}

func GetResults(arrs []int) int {

	if arrs == nil || len(arrs) == 0 {
		return 0
	}

	if len(arrs) == 1 {
		return arrs[0]
	}

	return arrs[0] + GetResults(arrs[1:])

}

func GetResults2(arrs []int) int {

	if arrs == nil || len(arrs) == 0 {
		return 0
	}

	if len(arrs) == 1 {
		return 1
	}

	return 1 + GetResults2(arrs[1:])

}

func GetResults3(arrs []float64) *float64 {

	if arrs == nil || len(arrs) == 0 {
		return nil
	}
	if len(arrs) == 1 {
		max := float64(1)
		return &max
	}

	if len(arrs) == 2 {
		max := math.Max(arrs[0], arrs[1])
		return &max
	}

	max := math.Max(arrs[0], *GetResults3(arrs[1:]))
	return &max

}
