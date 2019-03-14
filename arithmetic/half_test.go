package arithmetic

import (
	"testing"
	"math"
	"fmt"
)

func TestHalf(t *testing.T) {

	//fmt.Println("gogo ")
	//time.Sleep(1 * time.Second)
	//
	//var list []int
	//for i := 1; i <= 100; i++ {
	//	list = append(list, i)
	//}
	//
	//for i := 1000; i <= 2000; i += 2 {
	//	list = append(list, i)
	//}
	//
	//res := Find(1201, list)
	//fmt.Println(res)

	//fmt.Println(list)
	//mp := make(map[int]map[int]int, 100)
	//
	//for i := 1; i <= 100; i++ {
	//	mp[i] = Find(i, list)
	//}
	//
	//for i := 1; i <= 100; i++ {
	//	fmt.Printf("查找值:%d ", i)
	//	for k, v := range mp[i] {
	//		fmt.Printf("索引:%d 次数:%d\n", k, v)
	//	}
	//}
	//
	//fmt.Println(mp)

	var list2 = []int{1, 2, 3, 4, 5}
	fmt.Println(Find2(5, list2))

}

func Find2(target int, list []int) int {

	start, end := 0, len(list)-1

	for start <= end {
		mid := int(math.Floor(float64((start + end) / 2)))

		if list[mid] == target {
			return mid
		}

		if list[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1

}

func Find(targetNum int, list []int) map[int]int {

	start, end := 0, len(list)-1
	time := 1

	for start <= end {
		mid := int(math.Floor(float64((start + end) / 2)))

		if targetNum == list[mid] {
			//fmt.Printf(fmt.Sprintf("目标索引：%d", mid))
			return map[int]int{mid: time}
		}

		time++
		if targetNum < list[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return map[int]int{-1: time}

}

//func init() {
//	fmt.Println("haha1")
//}
//func init() {
//	fmt.Println("haha2")
//	go func() { fmt.Println("haha3") }()
//}
