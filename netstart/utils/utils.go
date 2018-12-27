package utils

import "time"

func TimeFormat(timeStamp int64) string {

	format := time.Unix(timeStamp, 0).Format("2006-01-02 15:04:05")

	return format;
}

func TimeArrays(n int32) [][]int32 {

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

func CalTimeArrays(num int32, results [][]int32) []int32 {

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
