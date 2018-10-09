package tests

import "testing"
import "fmt"

/*
数组 Arrays
数组是内置(build-in)类型,是一组同类型数据的集合，它是值类型，
通过从0开始的下标索引访问元素值。在初始化后长度是固定的，无法修改其长度。
当作为方法的入参传入时将复制一份数组而不是引用同一指针。
数组的长度也是其类型的一部分，通过内置函数len(array)获取其长度。
*/

func Test_Array01(t *testing.T) {

	//长度为5的数组，其元素值依次为：1，2，3，4，5
	var arr1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)
	fmt.Println(len(arr1))

	//长度为5的数组，其元素值依次为：1，2，0，0，0 。
	//在初始化时没有指定初值的元素将会赋值为其元素类型int的默认值0,string的默认值是""
	var arr2 = [5]int{1, 2}
	fmt.Println(arr2)
	fmt.Println(len(arr2))

	//长度为5的数组，其长度是根据初始化时指定的元素个数决定的
	var arr3 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)
	fmt.Println(len(arr3))

	//长度为5的数组，key:value,其元素值依次为：0，0，1，2，3。
	//在初始化时指定了2，3，4索引中对应的值：1，2，3
	var arr4 = [5]int{2: 1, 3: 2, 4: 3}
	fmt.Println(arr4)
	fmt.Println(len(arr4))

	//长度为5的数组，起元素值依次为：0，0，1，0，3。
	//由于指定了最大索引4对应的值3，根据初始化的元素个数确定其长度为5
	var arr5 = [...]int{2: 1, 4: 3}
	fmt.Println(arr5)
	fmt.Println(len(arr5))

}

//数组是值类型，将一个数组赋值给另一个数组时将复制一份新的元素
//切片是引用类型
func Test_Array02(t *testing.T) {

	/*数组，值类型*/
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := arr1

	arr1[3] = 10
	fmt.Println(arr1)
	fmt.Println(arr2)

	/*切片，引用类型*/
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1

	slice1[3] = 10
	fmt.Println(slice1)
	fmt.Println(slice2)

}



func Test_Array03(t *testing.T) {

	arr1 := [5]int{1, 2, 3, 4, 5}

	fc := func(a [5]int) {
		a[0] = 100
		fmt.Println(a)
	}

	//数组拷贝,可以理解为值类型
	fc(arr1)
	fmt.Println(arr1)

}
