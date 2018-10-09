package tests

import (
	"container/list"
	"fmt"
	"runtime"
	"sort"
	"sync"
	"testing"
	"time"
)

func Test_One(t *testing.T) {
	defer_call()
}

func defer_call() {

	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")

}

type student struct {
	Name string
	Age  int
}

func Test_Two(t *testing.T) {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
	}

	for k, v := range m {
		fmt.Printf("key=%s,value=%v \n", k, v)
	}

}

func Test_03(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(21)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("i: ", i)
			wg.Done()
		}()
	}
	for j := 0; j < 10; j++ {
		go func(x int) {
			fmt.Println("j: ", x)
			wg.Done()
		}(j)
	}
	wg.Wait()
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func Test_04(t *testing.T) {
	tea := Teacher{}
	tea.ShowA()
	tea.ShowB()
}

func Test_05(t *testing.T) {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}

type testChs struct {
	chs chan int
}

func (t *testChs) init() {

	t.chs = make(chan int, 5)

	//生产者

	go func() {

		index := 0
		for {

			t.chs <- index

			if index%5 == 0 {
				time.Sleep(1 * time.Second)
			}

			index++
		}

	}()

	for i := 0; i < 5; i++ {

		go func() {
			//消费者
			for {
				res := <-t.chs
				fmt.Println(res)
			}
		}()

	}

	ch := make(chan int, 1)
	<-ch

}

func Test_06(t *testing.T) {
	tc := &testChs{}
	tc.init()
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func Test_07(t *testing.T) {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

/*
10, 1, 2, 3
20, 0, 2, 2
2, 0, 2, 2
1, 1, 3, 4

*/

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func Test_08(t *testing.T) {

	ua := &UserAges{}
	ua.ages = make(map[string]int)

	ua.Add("cc", 10)
	ua.Add("cc1", 20)

	fmt.Println(ua.Get("cc1"))
	fmt.Println(ua.Get("cc"))
}

type People2 interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func Test_09(t *testing.T) {
	var peo People2 = &Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

type People3 interface {
	Show()
}

type Student3 struct{}

func (stu *Student3) Show() {

}

func live() People3 {
	var stu *Student3
	return stu
}

func Test_10(t *testing.T) {

	fmt.Printf("%v", live())

	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}

func TestDoit(t *testing.T) {
	doit := func(arg int) interface{} {
		var result *struct{} = nil
		if arg > 0 {
			result = &struct{}{}
		}
		return result
	}
	//输出结果。
	//-1:result: <nil>    为空的匿名结构体
	//1://result: &{}     匿名结构体的地址
	if res := doit(-1); res != nil {
		fmt.Println("result:", res)
	}
}

func TestSort(t *testing.T) {

	m := make(map[int]string)
	m[1] = "a"
	m[2] = "c"
	m[0] = "b"

	var keys []int
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println(keys)
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	fmt.Println(keys)

	for _, v := range keys {
		fmt.Println(m[v])
	}

}

func TestList(t *testing.T) {
	l := list.New() //创建一个新的list
	for i := 0; i < 5; i++ {
		l.PushBack(i)
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //输出list的值,01234
	}
	fmt.Println("")
	fmt.Println(l.Front().Value) //输出首部元素的值,0
	fmt.Println(l.Back().Value)  //输出尾部元素的值,4
	l.InsertAfter(6, l.Front())  //首部元素之后插入一个值为10的元素
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //输出list的值,061234
	}
	fmt.Println("")
	l.MoveBefore(l.Front().Next(), l.Front()) //首部两个元素位置互换
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //输出list的值,601234
	}
	fmt.Println("")
	l.MoveToFront(l.Back()) //将尾部元素移动到首部
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //输出list的值,460123
	}
	fmt.Println("")
	l2 := list.New()
	l2.PushBackList(l) //将l中元素放在l2的末尾
	for e := l2.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //输出l2的值,460123
	}
	fmt.Println("")
	fmt.Print(l.Len()) //0
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //输出list的值,无内容
	}

}

func Test_m01(t *testing.T) {

	mps := map[string]string{
		"a": "b",
		"b": "c",
		"c": "d",
		"d": "e",
		"e": "f",
	}

	fmt.Println(mps)

	for _, v := range mps {

		fmt.Println("准备删除值", v)
		delete(mps, v)
	}

	fmt.Println(mps)

}

func Test_m02(t *testing.T) {

	i := make([]int, 3)
	fmt.Println(i)

	// go func(i []int) {
	i[0] = 2
	i = append(i, 4)
	// }(i)

	fmt.Println(i)

	for _, v := range i {
		i = make([]int, 2)
		fmt.Println(v)
	}

	fmt.Println(i)
}

type AA struct {
	val1 int
	val2 int
}

type IAA interface {
	cal()
}

func (a *AA) cal() {
	a.val1 = 1
	a.val2 = 2
}

func Test_m03(t *testing.T) {

	var aa AA
	aa.cal()

	fmt.Println(aa)

	bb := &AA{}
	bb.cal()

	fmt.Println(bb)

}
