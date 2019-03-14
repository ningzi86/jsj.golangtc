package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/astaxie/beego/httplib"
)

func _Test_Chan01(t *testing.T) {

	varChan1 := make(chan string, 1)

	go func() {
		varChan1 <- "I am a string."
	}()

	fmt.Println("Got Msg:", <-varChan1)
}

// 创建 channel 类型对象时设置了 buffer 值
func Test_Chan02(t *testing.T) {

	varChan1 := make(chan string, 3)

	go func() {
		varChan1 <- "I am a string1."
		varChan1 <- "I am a string2."
		fmt.Println("I will be before all inputs")
		varChan1 <- "I am a string3."
	}()

	fmt.Println("Got Msg:", <-varChan1)
	fmt.Println("Got Msg:", <-varChan1)
	fmt.Println("Got Msg:", <-varChan1)

}

// channel 中的同步机制
func _Test_Chan03(t *testing.T) {

	varChan3 := make(chan bool, 1)
	go func(varChan chan bool) {
		fmt.Println("begin to execute")
		// do something
		time.Sleep(time.Second * 2)
		fmt.Println("end")
		varChan <- true
	}(varChan3)
	fmt.Println("Got msg: ", <-varChan3)

}

func _Test_Chan04(t *testing.T) {

	varInChan := make(chan int, 1)
	varOutChan := make(chan int, 1)

	varInChan <- 1

	func1 := func(outParam <-chan int, inParam chan<- int) {
		inParam <- <-outParam
	}

	func1(varInChan, varOutChan)

	fmt.Println(<-varOutChan)

}

func _Test_Chan05(t *testing.T) {

	c1 := make(chan int, 1)
	c2 := make(chan int, 1)

	go func() {
		c1 <- 123
	}()

	go func() {

		time.Sleep(time.Second * 5)

		c2 <- 456
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)

		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

}

func _Test_Chan06(t *testing.T) {

	c3 := make(chan int, 1)

	go func() {
		time.Sleep(time.Second * 3)
		c3 <- 123
	}()

	select {
	case res := <-c3:
		fmt.Println(res)
	case <-time.After(time.Second * 5):
		fmt.Println("timeout")
	}

}

// https://gobyexample.com/non-blocking-channel-operations
func Test_Chan0700(t *testing.T) {

	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println(msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}

func Test_chan0800(t *testing.T) {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done

}

// https://gobyexample.com/range-over-channels
func Test_chan09(t *testing.T) {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for elem := range queue {
		fmt.Println(elem)
	}
}

func _Test_chan10(t *testing.T) {

	type Res struct {
		Url     string
		Content string
	}

	urls := []string{
		"crawler://www.baidu.com",
		"crawler://www.jsj.com.cn",
		"crawler://www.sina.com.cn",
		"crawler://www.baidu.com",
		"crawler://www.jsj.com.cn",
		"crawler://www.sina.com.cn",
	}

	varChan := make(chan Res, len(urls))

	for _, url := range urls {

		go func(u string) {

			req := httplib.Get(u)
			content, _ := req.String()

			res := Res{
				Url:     u,
				Content: content,
			}
			varChan <- res

		}(url)

	}

	for i := 0; i < len(urls); i++ {
		fmt.Println((<-varChan).Url)
	}

	fmt.Println("over")

}

func Test_chan11(t *testing.T) {

	fmt.Println("Test_chan11")

	type Res struct {
		Url     string
		Content string
	}

	urls := []string{
		"crawler://www.baidu.com",
		"crawler://www.jsj.com.cn",
		"crawler://www.sina.com.cn",
	}

	chs := make([]chan Res, len(urls))

	fun := func(url string, ch chan Res) {
		req := httplib.Get(url)
		content, _ := req.String()

		res := Res{
			Url:     url,
			Content: content,
		}

		ch <- res
	}

	for i, url := range urls {
		chs[i] = make(chan Res)
		go fun(url, chs[i])
	}

	for _, ch := range chs {
		fmt.Println((<-ch).Url)
	}
}

func Test_chan012(t *testing.T) {

	done := make(chan bool)

	go func() {
		fmt.Println("write over")
		done <- true
	}()

	fmt.Println(<-done)

}
