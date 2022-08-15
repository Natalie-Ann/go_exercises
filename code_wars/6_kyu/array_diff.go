// Your goal in this kata is to implement a difference function, which subtracts one list from another and returns the result.

// It should remove all values from list a, which are present in list b keeping their order.

package main

import (
	"fmt"
)

func ArrayDiff(a, b []int) []int {
	var newArray []int
	arrayMap := make(map[int]bool)

	for _, number := range b {
		arrayMap[number] = true
	}

	for _, number := range a {
		if !arrayMap[number] {
			newArray = append(newArray, number)
		}
	}

	return newArray
}

func main() {
	fmt.Println(ArrayDiff([]int{1, 2, 3}, []int{1, 2}))
	fmt.Println(ArrayDiff([]int{}, []int{1, 2}))
}
