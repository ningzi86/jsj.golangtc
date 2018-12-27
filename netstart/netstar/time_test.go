package netstar

import (
	"testing"
	"fmt"
	"time"
)

func TestTime(t *testing.T) {

	//sec := 200
	//results := timeArrays(int32(sec))
	//
	//
	//
	//arrays := calTimeArrays(sec, results)

	fmt.Println(time.Now().Unix(), time.Now().UnixNano()/1000000)

}

func timeArrays(n int32) [][]int32 {

	var results [][]int32

	for {
		min := n / int32(2)
		max := n

		if max <= 10 {
			results = append(results, []int32{0, 10})
			break
		}

		if min <= 10 {
			results = append(results, []int32{10, max})
			results = append(results, []int32{0, 10})
			break
		}

		var result []int32
		result = append(result, min)
		result = append(result, max)

		results = append(results, result)
		n = min
	}

	return results

}

func calTimeArrays(num int32, results [][]int32) []int32 {

	if num <= 0 {
		return []int32{0, 10}
	}

	if len(results) == 0 {
		return []int32{0, 10}
	}

	for i := 0; i < len(results); i++ {

		min := results[i][0]
		max := results[i][1]

		if num > min && num <= max {
			return results[i]
		}
	}
	return results[0]
}

func TestTimer2(t *testing.T) {

	var tick *time.Ticker

	tick.Stop()

	tick = time.NewTicker(1 * time.Second)

	tick.Stop()
	tick.Stop()
	tick = time.NewTicker(1 * time.Second)

	//此处在等待channel中的信号，执行此段代码时会阻塞两秒
	fmt.Println(<-tick.C)
	fmt.Println(<-tick.C)
	fmt.Println(<-tick.C)
	fmt.Println(<-tick.C)

	tick.Stop()
	tick = time.NewTicker(1 * time.Second)
	fmt.Println(<-tick.C)

}
