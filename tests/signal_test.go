package tests

import (
	"fmt"
	"os"
	"os/signal"
	"testing"
	"time"

	l4g "github.com/alecthomas/log4go"
)

func init() {
	l4g.LoadConfiguration("log.xml") //使用加载配置文件,类似与java的log4j.propertites
	// l4g.Debug("the time is now :%s -- %s", "213", "sad")
	// l4g.Error("this Error")
}

// func Test_Signal(t *testing.T) {

// 	c := make(chan os.Signal, 0)
// 	signal.Notify(c)

// 	// Block until a signal is received.
// 	s := <-c
// 	fmt.Println("Got signal:", s) //Got signal: terminated

// 	// l4g.Debug("WriteMessage")

// 	// c := make(chan os.Signal, 0)
// 	// signal.Notify(c)

// 	// // signal.Stop(c) //不允许继续往c中存入内容
// 	// s := <-c //c无内容，此处阻塞，所以不会执行下面的语句，也就没有输出
// 	// fmt.Println("Got signal:", s)
// 	// l4g.Debug("Got signal:", s)

// }

func Test_Signal(t *testing.T) {

	l4g.Debug("the time is now :%s -- %s", "213", "sad")

	go signalListen()
	time.Sleep(time.Hour)
}

func signalListen() {

	l4g.Debug("get signal1:")

	c := make(chan os.Signal)
	signal.Notify(c)
	for {
		s := <-c

		l4g.Debug("get signal2:")
		l4g.Debug("get signal:", s)
		
		//收到信号后的处理，这里只是输出信号内容，可以做一些更有意思的事
		fmt.Println("get signal:", s)

		os.Exit(1)

	}
}
