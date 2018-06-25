package tests

import (
	"fmt"
	"testing"
)

func Test_chan12(t *testing.T) {

	c1 := make(chan int)
	
	//这里打开出异常，无缓冲必须在此之前有goroutine读取
 	//c1 <- 1

	go func() {
		fmt.Println(<-c1)
	}()

	c1<-2 

}



func Test_chan13(t *testing.T) {

	c1 := make(chan int,1)
	 
 	c1 <- 1 
	//这里打开异常，缓存为1，必须等channel内值释放时才能执行
	//c1 <- 2 
}
