package tests

import (
	"testing"
	"io"
	"strings"
	"os"
	"fmt"
)

func TestWriter01(t *testing.T) {

	up := &UpperWriter{
		os.Stdout,
	}
	fmt.Fprintln(up, "Hello world", "haha")
}

type UpperWriter struct {
	io.Writer
}

func (w *UpperWriter) Write(p []byte) (n int, err error) {
	return w.Writer.Write([]byte(strings.ToUpper(string(p))))
}

type UpperString string

func (s UpperString) String2() string {
	return strings.ToUpper(string(s))
}

func TestWriter02(t *testing.T) {
	fmt.Fprintln(os.Stdout, UpperString("hello, world"))
	fmt.Fprintln(os.Stdout, UpperInt(123))
}

type UpperInt int32

func (i UpperInt) String() string {
	return fmt.Sprintf("%d--1", i)
}
