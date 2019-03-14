package tests

import (
	"testing"
	"fmt"
)

func Test_Fun01(t *testing.T) {

	v := Inc()
	fmt.Println(v)
}

func Inc() (v int) {
	defer func() { v++ }()
	return 42
}
