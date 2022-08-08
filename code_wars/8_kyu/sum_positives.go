/*
You get an array of numbers, return the sum of all of the positives ones.

Example [1,-4,7,12] => 1 + 7 + 12 = 20

Note: if there is nothing to sum, the sum is default to 0.
*/

package main

import (
	"fmt"
)

func PositiveSum(numbers []int) int {
	var sum int
	for _, value := range numbers {
		if value > 0 {
			sum += value
		}
	}
	return sum
}

func main() {
	fmt.Println(PositiveSum([]int{1, -4, 7, 12}))
}
