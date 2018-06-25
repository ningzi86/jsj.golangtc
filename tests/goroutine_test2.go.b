package tests

import "testing"
import "fmt"
import "github.com/astaxie/beego/httplib"
import "sync"
import "time"

func Test_GoRoutine01(t *testing.T) {

	start := time.Now()

	fmt.Println("Begin……")

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go download1(wg)
	go download2(wg)

	fmt.Println("End……")

	wg.Wait()

	end := time.Now()
	fmt.Println(end.Sub(start).Seconds())

}

func download1(wg *sync.WaitGroup) {

	defer wg.Done()

	res := httplib.Get("http://www.baidu.com")
	res.String()

	fmt.Println("download1 is over")

}

func download2(wg *(sync.WaitGroup)) {

	defer wg.Done()

	res := httplib.Get("http://www.sina.com.cn")
	res.String()

	fmt.Println("download2 is over")

}
