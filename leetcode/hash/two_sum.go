// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

// Input: nums = [3,2,4], target = 6
// Output: [1,2]

// Input: nums = [3,3], target = 6
// Output: [0,1]

/*
Input: slice of integers, target integer
Output: slice of integers representing indices

Explicits: - return slice will be length 2
						- integers at indices will add up to target integer
						- no repeat indices
Implicits: 	-

DS: map

Algorithm:
- Declare new map with integer key and index value
- Iterate through the input slice
	- For each number
		- add number key to the map, index value
		- calculate the corresponding number for target
		- look up if corresponding number exsits in map
		- if corresponding value exists, return [currentIndex, map[correspondingKey]]
- Else return false
*/

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	correspondingNumbers := map[int]int{}

	for index, number := range nums {
		correspondingNumbers[number] = index + 1
		complementNumber := target - number
		if correspondingNumbers[complementNumber] != 0 {
			return []int{index, correspondingNumbers[complementNumber] - 1}
		}
	}

	return []int{}
}

func main() {
	example := []int{2, 7, 11, 15}
	fmt.Println(twoSum(example, 9))
}
