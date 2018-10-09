package tests

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func Test_byte01(t *testing.T) {

	a := "Hello,世界"

	fmt.Println(len(a), utf8.RuneCountInString(a))

	for _, k := range a {
		fmt.Printf("%c-%X  ", k, k)
	}
	fmt.Println()

	for _, c := range []byte(a) {
		fmt.Printf("%X ", c)
	}
	fmt.Println()

	bytes := []byte(a)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}

}
