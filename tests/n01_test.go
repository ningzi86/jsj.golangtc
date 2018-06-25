package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	// "path/filepath"
)

import "flag"

import "io/ioutil"

func Test_01(t *testing.T) {

	values := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(values[:len(values)/2])
	fmt.Println(values[len(values)/2:])

	var infile *string = flag.String("i", "infile", "File contains values for sorting")
	var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
	var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

	flag.Parse()
	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =",
			*algorithm)

	}
}

func Test_02(t *testing.T) {

	path, _ := os.Getwd()
	fmt.Println(path)

	b, _ := ioutil.ReadFile(path + "/n01_test.go")
	fmt.Println(string(b))

	workpath, _ := os.Getwd()

	fmt.Println(workpath)
	workpath, _ = filepath.Abs(workpath)
	
	fmt.Println(workpath)
}
