// Given two integers a and b, which can be positive or negative, find the sum of all the integers between and including them and return it. If the two numbers are equal return a or b.

package main

import "fmt"

func GetSum(a, b int) int {
	var sum int

	if a == b {
		return a
	} else if a < b {
		for i := a; i <= b; i++ {
			sum += i
		}
	} else {
		for i := b; i <= a; i++ {
			sum += i
		}
	}

	fmt.Println(sum)
	return sum
}

func main() {
	GetSum(1, 0)
	GetSum(-1, 2)
	GetSum(1, 1)
}
