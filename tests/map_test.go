package tests

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func Test_Map01(t *testing.T) {

	//http://xhrwang.me/2014/12/25/golang-fundamentals-4-map-range.html
	var m1 map[string]string

	m1 = make(map[string]string)
	m1["a"] = "aa"
	m1["b"] = "bb"

	m2 := make(map[string]string)
	m2["a"] = "aa"
	m2["b"] = "bb"

	m3 := map[string]string{
		"a": "aa",
		"b": "bb",
	}

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)

	//查找键值是否存在
	if v, ok := m3["a"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Key Not Found")
	}

	//遍历
	for k, v := range m3 {
		fmt.Println(k)
		fmt.Println(v)
	}

	// 使用 make 时也可以设定预期的键值对数量，在初始化时一次性分配大量内存，
	// 从而避免使用过程中频繁动态分配
	// 这里给定的数量值不会影响初始化后 len(mapObject)
	sm := make(map[string]string, 1)
	sm["a"] = "1"
	sm["b"] = "2"
	sm["c"] = "3"

	fmt.Println(sm)
	fmt.Println(len(sm))

	// 如果尝试删除不存在的元素，对已有数据不会有影响，不会抛出异常
	fmt.Println(sm["d"])

	delete(sm, "b")
	fmt.Println(sm["b"])

}

func Test_Map02(t *testing.T) {

	var mp = make(map[string]string, 2)

	mp["1"] = "1"
	mp["2"] = "2"

	mp["3"] = "3"

	fmt.Println(mp)
	fmt.Println(len(mp))
}

func Test_Map03(t *testing.T) {

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func Test_Map04(t *testing.T) {

	mp := make(map[int]string, 10)

	mp[0] = "1"
	mp[1] = "1"
	mp[2] = "1"
	mp[3] = "1"
	mp[4] = "1"
	mp[5] = "1"

	fmt.Println(mp, len(mp))

}

func Test_Map05(t *testing.T) {

	var lk sync.Mutex
	ch := make(chan int, 2)

	mp := make(map[int]string, 10)
	go func(m map[int]string) {
		lk.Lock()
		for i := 0; i < 10; i++ {
			m[i] = strconv.Itoa(i)
		}
		lk.Unlock()
		ch <- 1
	}(mp)

	go func(m map[int]string) {

		lk.Lock()
		for i := 20; i < 30; i++ {
			m[i] = strconv.Itoa(i)
		}
		lk.Unlock()
		ch <- 1
	}(mp)

	<-ch
	<-ch

	fmt.Println(mp)

}
