package tests

import "testing"
import "fmt"

func Test_fmt01(t *testing.T)  {
    var i interface{} = 23
    fmt.Printf("%d\n", i)
}