package main

import "fmt"

func clurse() func(i int) int {

	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}

}

func main() {

	fc := clurse()

	v1 := fc(1)
	v2 := fc(10)
	v3 := fc(2)

	fmt.Println(v1, v2, v3)
}
