package parser

import (
	"testing"
	"io/ioutil"
	"fmt"
)

func Test_ReaderFile(t *testing.T) {

	bytes, err := ioutil.ReadFile(`crawler/zhenai/parser/user.in`)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", bytes)

}
