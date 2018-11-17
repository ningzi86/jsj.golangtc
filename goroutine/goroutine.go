package main

import (
	"time"
	"fmt"
	"runtime"
)

func main() {

	j := 0
	var a [10]int

	for i := 0; i < 10; i++ {
		go func(i int, jj *int) {
			for {
				//fmt.Printf("123")

				//*jj++
				a[i]++
				runtime.Gosched()
				//fmt.Printf("current value is %d \n", i)
			}
		}(i, &j)
	}

	time.Sleep(time.Second)
	//fmt.Printf("j value is %d \n", j)
	fmt.Println(a)


}
