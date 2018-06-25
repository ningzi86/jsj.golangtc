package tests

import (
	"testing"
	"fmt"
	"strconv"
)

type Person struct{
	Name string
	Age int
	Sex int
}

func (this *Person) GoString() string  {
	return "&Person{Name is "+this.Name+", Age is "+strconv.Itoa(this.Age)+", Sex is "+strconv.Itoa(this.Sex)+"}"
}

func Test_Fmt1(t *testing.T)  {
	
	p := &Person{
		Name:"张宁",
		Age:22,
	}

	fmt.Printf("%#v", p)
 
}
