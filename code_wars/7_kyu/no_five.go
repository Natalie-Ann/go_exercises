// You are given the start number and the end number of a region and should return the count of all numbers except numbers with a 5 in it. The start and the end number are both inclusive.

// Examples:

// 1,9 -> 1,2,3,4,6,7,8,9 -> Result 8
// 4,17 -> 4,6,7,8,9,10,11,12,13,14,16,17 -> Result 12
// The result may contain fives.
// The start number will always be smaller than the end number. Both numbers can be also negative.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func DontGiveMeFive(start int, end int) int {
	var numbers []int

	for i := start; i <= end; i++ {
		strInt := strconv.Itoa(i)
		if strings.Contains(strInt, "5") {
			continue
		}
		numbers = append(numbers, i)
	}
	return len(numbers)
}

func main() {
	fmt.Println(DontGiveMeFive(1, 9))
	fmt.Println(DontGiveMeFive(4, 17))
}
