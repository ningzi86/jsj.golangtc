package tests

import (
	"fmt"
	"testing"

	. "github.com/ahmetalpbalkan/go-linq"
)

type foo struct {
	f1 int
}

func (f foo) CompareTo(c Comparable) int {
	a, b := f.f1, c.(foo).f1

	if a < b {
		return 1
	} else {
		return -1
	}
}

func Test01(t *testing.T) {
    
	var f1 = foo{f1: 1}
	var f2 = foo{f1: 2}

	var v = f1.CompareTo(f2)
	fmt.Println(v)

}
