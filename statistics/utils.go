package statistics

import (
	"math"
	"sort"
)

// GetMean recieves an array of integer and returns a number
func GetMean(data []int) (mean int) {
	n := len(data)
	var sum int
	for _, val := range data {
		sum = sum + val
	}
	mean = sum / n
	return
}

// GetMedian gives you the median
func GetMedian(data []int) (median float64) {
	n := len(data)
	sort.Ints(data)
	mid := int(math.Floor(float64(n) / 2))
	even := n%2 == 0
	// fmt.Printf("Even: %v \tMid: %v\n", even, mid)
	if even {
		var sum = 0.0
		for _, val := range data[mid-1 : mid+1] {
			sum += float64(val)
		}
		median = float64(sum) / 2.0
	} else {
		median = float64(data[mid])
	}
	// fmt.Println(median)
	return
}
