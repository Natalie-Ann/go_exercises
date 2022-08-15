// Given n, take the sum of the digits of n. If that value has more than one digit, continue reducing in this way until a single-digit number is produced. The input will be a non-negative integer.

// 942  -->  9 + 4 + 2 = 15  -->  1 + 5 = 6

package main

import (
	"fmt"
)

func splitInt(n int) []int {
	digitSlice := []int{}

	for n > 0 {
		digitSlice = append(digitSlice, n%10)
		n = n / 10
	}

	return digitSlice
}

func DigitalRoot(n int) int {
	if n < 10 {
		return n
	}

	count := 0
	digitSlice := splitInt(n)
	fmt.Println(digitSlice)

	for _, digit := range digitSlice {
		count += digit
	}

	return 0 + DigitalRoot(count)

}

func main() {
	fmt.Println(DigitalRoot(942))
}
