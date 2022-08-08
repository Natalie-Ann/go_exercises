/*Given an array of integers.

Return an array, where the first element is the count of positives numbers and the second element is sum of negative numbers. 0 is neither positive nor negative.

If the input is an empty array or is null, return an empty array.

Example
For input [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15], you should return [10, -65].
*/

package main

import "fmt"

func CountPositivesSumNegatives(numbers []int) []int {
	var res = make([]int, 2)

	if len(numbers) == 0 {
		return []int{}
	}

	for _, value := range numbers {
		if value < 0 {
			res[1] += value
		} else if value > 0 {
			res[0] += 1
		}
	}
	return res
}

func main() {
	fmt.Println(CountPositivesSumNegatives([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}))
}
