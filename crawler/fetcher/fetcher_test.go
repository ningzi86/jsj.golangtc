package fetcher

import (
	"testing"
	"fmt"
)

func Test_Fetcher(t *testing.T) {

	bytes, e := Fetch("http://album.zhenai.com/u/109138413")

	if e != nil{
		panic(e)
	}

	fmt.Printf("%s", bytes)

}
