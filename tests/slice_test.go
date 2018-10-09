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
	arr := [...]int{1, 2, 3}
	fmt.Println(arr)
	fmt.Println(arr[2])

	slice4 := slice1[:]
	fmt.Println(slice4)
	fmt.Println(slice4[2])
}

func Test_Slice012(t *testing.T) {

	slice1 := []int{1, 2, 3}
	fc := func(s []int, v int) {
		s[1] = v
	}

	//slice 为引用类型，下面方法改变值时，对应的slice1_2值也会改变
	fc(slice1, 100)
	slice1_2 := slice1[1:2]
	fc(slice1, 101)

	fmt.Println(slice1, cap(slice1))
	fmt.Println(slice1_2, cap(slice1_2))

	//append后，slice的cap变大，指向新的内存地址，所以slice1_2值不会再改变
	slice1 = append(slice1, 4)
	fc(slice1, 101)

	fmt.Println(slice1, cap(slice1))
	fmt.Println(slice1_2, cap(slice1_2), slice1_2[0:2])
}

func Test_Slice013(t *testing.T) {

	slice1 := make([]int, 5)

	fmt.Println(slice1, cap(slice1))
	fc := func(s []int, index int, v int) {
		s[index] = v
	}

	//原始的slice
	slice1_2 := slice1[1:5]
	fmt.Println("slice1_2", slice1_2, cap(slice1_2))
	fc(slice1, 1, 100)

	//slice1的cap没变,所以slice1_2指向的是slice1的内存地址
	fmt.Println("slice1_2", slice1_2, cap(slice1_2))
	fc(slice1, 2, 101)

	//slice1的cap没变,所以slice1_2指向的是slice1的内存地址
	fmt.Println("slice1_2", slice1_2, cap(slice1_2))
	fmt.Println(slice1, cap(slice1))

	//填充元素,致使slice1的cap翻倍,指向新的内存地址
	slice1 = append(slice1, 6)
	fmt.Println(slice1, cap(slice1))

	//指向新的slice1
	slice1_3 := slice1[1:5]
	//改变新的slice1值
	fc(slice1, 3, 102)

	//原始的slice1_2值不受影响
	fmt.Println("slice1_2", slice1_2, cap(slice1_2))

	//新的slice1_3指向新的slice1地址,所以也跟着改变
	fmt.Println("slice1_3", slice1_3, cap(slice1_3))

	fmt.Println(slice1, cap(slice1))

	slice1 = append(slice1, 7)
	fmt.Println(slice1, cap(slice1), slice1[6:10])

}

func Test_Slice014(t *testing.T) {

	s := make([]int, 3, 10)

	fmt.Println(s, len(s), cap(s))

}
