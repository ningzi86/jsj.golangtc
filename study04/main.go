package main

import (
	"fmt"
)

const (
	num      = 10
	rangeNum = 100000
)

func main() {

	var buf []int = []int{1, 8, 9, 6, 5, 4, 3, 7, 2, 0}

	//var buf []int = []int{9, 6, 5, 4, 3, 7, 2, 8}
	//randSeed := rand.New(rand.NewSource(time.Now().Unix() + time.Now().UnixNano()))
	//var buf []int
	//for i := 0; i < num; i++ {
	//	buf = append(buf, randSeed.Intn(rangeNum))
	//}
	//fmt.Println(buf)
	//maopao2(buf)
	//fmt.Println(buf)
	//maopao(buf)
	//fmt.Println(buf)

	//xuanze2(buf)
	//fmt.Println(buf)

	kuaisu(buf)
	fmt.Println(buf)

}

func maopao3(buf []int) {
	for i := 0; i < len(buf); i++ {
		for j := 0; j < len(buf)-i-1; j++ {
			if buf[j] < buf[j+1] {
				buf[j], buf[j+1] = buf[j+1], buf[j]
			}
		}
	}
}

func maopao2(buf []int) {

	for i := 0; i < len(buf); i++ {
		for j := 0; j < len(buf)-i-1; j++ {
			if buf[j] > buf[j+1] {
				buf[j], buf[j+1] = buf[j+1], buf[j]
			}
		}
	}
}

// 冒泡排序
func maopao(buf []int) {
	times := 0
	for i := 0; i < len(buf)-1; i++ {
		flag := false
		for j := 1; j < len(buf)-i; j++ {

			fmt.Println("交换前", i, j, buf[j-1], buf[j], buf)
			if buf[j-1] > buf[j] {
				times++
				tmp := buf[j-1]
				buf[j-1] = buf[j]
				buf[j] = tmp
				flag = true
			}
			fmt.Println("交换后", i, j, buf[j-1], buf[j], buf)
		}
		if !flag {
			break
		}
	}
	fmt.Println("maopao times: ", times)
}

// 选择排序
func xuanze(buf []int) {
	times := 0
	for i := 0; i < len(buf)-1; i++ {
		min := i
		for j := i; j < len(buf); j++ {
			times++
			if buf[min] > buf[j] {
				min = j
			}
		}
		if min != i {
			tmp := buf[i]
			buf[i] = buf[min]
			buf[min] = tmp
		}
	}
	fmt.Println("xuanze times: ", times)
}

func xuanze2(buf []int) {

	for i := 0; i < len(buf)-1; i ++ {

		min := i
		for j := i; j < len(buf); j++ {
			if buf[min] > buf[j] {
				min = j
			}
		}

		if min != i {
			buf[min], buf[i] = buf[i], buf[min]
		}

	}

}

// 快速排序
func kuaisu(buf []int) {
	kuai(buf, 0, len(buf)-1)
}

func kuai(a []int, l, r int) {
	if l >= r {
		return
	}
	fmt.Println("排序前", a)
	i, j, key := l, r, a[l] //选择第一个数为key
	for i < j {
		for i < j && a[j] > key { //从右向左找第一个小于key的值
			j--
		}
		if i < j {
			a[i] = a[j]
			i++
		}

		for i < j && a[i] < key { //从左向右找第一个大于key的值
			i++
		}
		if i < j {
			a[j] = a[i]
			j--
		}
	}
	//i == j
	a[i] = key
	fmt.Println("排序后", a, a[i])
	kuai(a, l, i-1)
	kuai(a, i+1, r)

}
