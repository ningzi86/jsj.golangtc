package tests

import (
    "fmt"
    "testing"
    "strconv"
)

func Test_Convert01(t *testing.T)  {
    
    /*
    int，Runes（注：Rune 是int 的别名）
    int8 ,int16 ,int32 ,int64 
    byte ,uint8 ,uint16 ,uint32 ,uint64 （注：byte是uint8 的别名）
    float32 ，float64 (没有float 类型)
    bool
    string
    complex128，complex64
    */

    var inta1 int  = 12345
    var stra1 = strconv.Itoa(inta1)
    var stra2 = fmt.Sprintf("%v", inta1)
    var stra3 = fmt.Sprintf("%d", inta1)

    inta2,_ := strconv.Atoi(stra1)

    fmt.Println(stra1)
    fmt.Println(stra2)
    fmt.Println(stra3)
    fmt.Println(inta2)

    b,_ := strconv.ParseBool("true")
    f,_ := strconv.ParseFloat("3.1415", 64)
    i,_ := strconv.ParseInt("-42", 10, 64)
    u,_ := strconv.ParseUint("42", 10, 64)

    b1 := strconv.FormatBool(b)
    f1 := strconv.FormatFloat(f, 'f', 6, 64)
    i1 := strconv.FormatInt(i, 10)
    u1 := strconv.FormatUint(u, 10)

    fmt.Println(b)
    fmt.Println(f)
    fmt.Println(i)
    fmt.Println(u)

    fmt.Println("b1", b1)
    fmt.Println("f1", f1)
    fmt.Println(i1)
    fmt.Println(u1)

    var t1 float64 = 10
    t2 := int64(t1)

    fmt.Println(t2)
    t3 := int(t2)
    fmt.Println(t3)

    t4 := float64(t2)
    fmt.Println(t4)



}