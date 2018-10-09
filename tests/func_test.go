package tests

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"testing"
)

func closure() func(i int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func Test_Func01(t *testing.T) {
	fc := closure()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + ... + %d = %d \n", i, fc(i))
	}
}

func fibonacci() func() int {

	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}

}

func Test_Fibonacci(t *testing.T) {

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}

type intGen func() int

func fibonacci2() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 1000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func Test_Fibonacci2(t *testing.T) {

	f := fibonacci2()
	printFileContents(f)

}
