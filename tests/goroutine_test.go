package tests

import (
	"testing"
	"fmt"
	"time"
)

func Test_Goroutine1(t *testing.T) {

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("current value is %d \n", i)
		}(i)
	}

	time.Sleep(time.Millisecond)

}


func Test_Goroutine2(t *testing.T) {

	j := 0
	var a [10]int

	for i := 0; i < 10; i++ {
		go func(i int, jj *int) {
			for {
				 //*jj++
				 a[i]++
				//fmt.Printf("current value is %d \n", i)
			}
		}(i, &j)
	}

	time.Sleep(time.Millisecond)
	fmt.Printf("j value is %d \n", j)
	fmt.Println(a)

}