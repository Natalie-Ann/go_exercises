// You are given an array (which will have a length of at least 3, but could be very large) containing integers. The array is either entirely comprised of odd integers or entirely comprised of even integers except for a single integer N. Write a method that takes the array as an argument and returns this "outlier" N.

// [2, 4, 0, 100, 4, 11, 2602, 36]
// 11 (the only odd number)

// [160, 3, 1719, 19, 11, 13, -21]
// 160 (the only even number)

package main

import (
	"fmt"
)

func FindOutlier(integers []int) int {
	var specialNumber int
	dividedNumbers := make(map[string][]int)

	for _, number := range integers {
		if number%2 == 1 || number%2 == -1 {
			dividedNumbers["odd"] = append(dividedNumbers["odd"], number)
		} else if number%2 == 0 {
			dividedNumbers["even"] = append(dividedNumbers["even"], number)
		}

		if (len(dividedNumbers["odd"]) > 1 && len(dividedNumbers["even"]) == 1) || (len(dividedNumbers["even"]) > 1 && len(dividedNumbers["odd"]) == 1) {
			break
		}
	}

	if len(dividedNumbers["even"]) > 1 {
		specialNumber = dividedNumbers["odd"][0]
	} else {
		specialNumber = dividedNumbers["even"][0]
	}

	return specialNumber
}

func main() {
	fmt.Println(FindOutlier([]int{160, 3, 1719, 19, 11, 13, -21}))
	fmt.Println(FindOutlier([]int{2, 4, 0, 100, 4, 11, 2602, 36}))
	fmt.Println(FindOutlier([]int{2, 6, 8, -10, 3}))
}
