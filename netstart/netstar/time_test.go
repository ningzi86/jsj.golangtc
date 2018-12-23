package netstar

import (
	"testing"
	"time"
	"fmt"
)

func TestTime(t *testing.T) {

	now := time.Now().Unix()
	saleTime := now + 60

	for {

		n1 := time.Now().Unix()

		if (n1 >= saleTime) {
			fmt.Println("已开始")
		} else {

			if n1 > saleTime-10 {

			} else if n1 > saleTime-30 {

			} else if n1 > saleTime-60 {

			} else if n1 > saleTime-5*60 {

			} else if n1 > saleTime-30*60 {

			} else if n1 > saleTime-2*60*60 {

			} else {

			}

		}

	}

}
