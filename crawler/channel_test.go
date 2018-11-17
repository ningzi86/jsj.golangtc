package main

import (
	"testing"
	"time"
	"fmt"
)

func TestChannel(t *testing.T) {

	c1 := make(chan int)
	c2 := make(chan chan int)

	result := make(chan int)

	go func() {
		for {
			c2 <- c1
			c := <-c1

			result <- c
			fmt.Println("c", c)
			time.Sleep(time.Second * 3)
		}
	}()

	var c1s []int
	var c2s []chan int

	go func() {
		for {
			var acTc1 int
			var acTc2 chan int
			if len(c1s) > 0 && len(c2s) > 0 {
				acTc1 = c1s[0]
				acTc2 = c2s[0]
			}
			fmt.Println("c1s", len(c1s))
			fmt.Println("c2s", len(c2s))

			select {
			case c := <-c1:
				c1s = append(c1s, c)
			case c := <-c2:
				c2s = append(c2s, c)
			case acTc2 <- acTc1:
				fmt.Println("444")
				c1s = c1s[1:]
				c2s = c2s[1:]
			}
		}
	}()

	c1 <- 1

	for {

		fmt.Println(<-result)
	}
}
