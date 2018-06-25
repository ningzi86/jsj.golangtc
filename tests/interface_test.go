package tests


import (
    "fmt"
    "testing"
)

type IPrint interface {
    Say(s string)
}


type Print1 struct{
    Name string
}



func (this *Print1) Say(s string)  {
    fmt.Println(this.Name)
    fmt.Println(s)
}

func Test_Interface01(t *testing.T)  {
    
        var fc = func (name string, p IPrint)  {
            p.Say(name)
        }

        p := Print1{Name:"张宁"}
        
        fc("hello", &p)

}