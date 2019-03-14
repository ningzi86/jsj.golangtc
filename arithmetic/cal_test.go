package arithmetic

import (
	"testing"
	"fmt"
)

func TestCal(t *testing.T) {

	//arr := []int{1, 2, 3, 4}
	//
	//len := len(arr)
	//bit := 1 >> uint32(len)

	//fmt.Println(len)
	//fmt.Println(bit)
	fmt.Println(1 << 1) //2
	fmt.Println(1 << 2) //4
	fmt.Println(1 << 3) //8
	fmt.Println(1 << 4) //16
	fmt.Println(1 << 5) //32

	fmt.Println(1 >> 1) //0
	fmt.Println(1 >> 2) //0

	fmt.Println(2 >> 1) //1

	fmt.Println(10 >> 1) //5
	fmt.Println(10 >> 2) //2

}

func TestCal2(t *testing.T) {

	fmt.Println(1 & 1)
	fmt.Println(1 & 0)
	fmt.Println(0 & 0)
	fmt.Println(0 & 1)

	fmt.Println(3 & 2)
	fmt.Println(3 & 5)

	fmt.Println(3 | 2)
	fmt.Println(3 | 5)

	fmt.Println(^1)
	//0001 -> 1110 -> 1110-1=1101 -> 0010 -> 2 -> -2
	fmt.Println(^2)
	//0010 -> 1101 -> 1100 -> 0011 -> 3 -> -3
	fmt.Println(^5)
	//0101 -> 1010 -> 1001 -> 0110 -> 6 -> -6
	fmt.Println(^9)
	//1001 -> 0110 -> 0101 -> 1010 -> 10 -> -10

	fmt.Println(1 ^ 2)
	// 0001 ^ 0010 相同取0,不同取1 -> 0011

	fmt.Println(4 ^ 6)
	// 0100 ^ 0110 -> 0010 -> 2

	fmt.Println(12 & (1 << 2))
}

func TestCal3(t *testing.T) {

	for i := 1; i <= 15; i++ {
		fmt.Printf("%3d %4b% d\n", i, i, n(i))
	}
}

func n(num int) int {

	count := 0

	for num > 0 {
		num &= num - 1
		count++
	}

	return count

}
