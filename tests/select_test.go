package tests


import (
    "fmt"
    "testing"
    "time"
)

func Test_Select01(t *testing.T)  {
    
    c1 := make(chan string)
    c2 := make(chan string)
 
    go func() {
        time.Sleep(time.Second * 10)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "two"
    }()
    
    for i := 0; i < 3; i++ {

        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        default:
            fmt.Println("没有数据")
        }

        fmt.Println("over")
    }

}