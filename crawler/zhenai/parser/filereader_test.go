package parser

import (
	"testing"
	"io/ioutil"
	"fmt"
	"regexp"
)

func Test_ReaderFile(t *testing.T) {

	bytes, err := ioutil.ReadFile(`crawler/zhenai/parser/user.in`)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", bytes)

}

func Test_parseCity(t *testing.T) {

	var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/([0-9]+))"[^>]*>([^<]+)</a>`)

	url := `<a href="http://album.zhenai.com/u/109138413" >烟火</a>`

	bytes := cityRe.FindAllSubmatch([]byte(url), -1)

	fmt.Printf("%s", bytes)

}
