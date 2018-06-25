package tests

import "testing"
import "fmt"

/*

切片 Slices

数组的长度不可改变，在特定场景中这样的集合就不太适用，
Go中提供了一种灵活，功能强悍的内置类型Slices切片,与数组相比切片的长度是不固定的，
可以追加元素，在追加时可能使切片的容量增大。
切片中有两个概念：一是len长度，二是cap容量，长度是指已经被赋过值的最大下标+1，
可通过内置函数len()获得。
容量是指切片目前可容纳的最多元素个数，可通过内置函数cap()获得。
切片是引用类型，因此在当传递切片时将引用同一指针，修改值将会影响其他的对象。

*/

func Test_Slice01(t *testing.T) {

	slice1 := []int{1, 2, 3}
	var slice2 = slice1

	slice1[1] = 100

	//可以追加元素
	slice1 = append(slice1, 10, 10, 10, 10)

	//可以追加切片
	slice1 = append(slice1, slice2...)

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	//切片
	slice3 := make([]int, 3, 3)
	fmt.Println(slice3)
	fmt.Println(slice3[2])

	//数组
	arr := [...]int{1,2,3}
	fmt.Println(arr)
	fmt.Println(arr[2])

	slice4 := slice1[:]
	fmt.Println(slice4)
	fmt.Println(slice4[2])
}
