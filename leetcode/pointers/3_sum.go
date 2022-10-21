/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

Examples:

[-1,0,1,2,-1,-4] --> [[-1,-1,2],[-1,0,1]]
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.

[0,1,1] --> []
Explanation: The only possible triplet does not sum up to 0.

[0,0,0] --> [[0,0,0]]
Explanation: The only possible triplet sums up to 0.

Input: slice of ints
Output: slice of slices containing ints

Explicits: 	- three digits must add to 0
						- each digit must have different index
						- can't repeat triplets
Implicits: 	- if no triplets, return empty slice

DS: slices, map?

Algorithm:
	- Declare a map with integer key, slice of slices value
	- Iterate through input slice
		- For each integer, add integer to map as keys
		- For each key, call 2 sum on opposite of key --> key = 5; twoSum(inputArray, target = -5)
		- Set return of twosum as value slice in map
	--- How to track indexes? Adjust two sum with three parameters

	- Declare filteredThreeSums slice of slices
	- Iterate through map
		- Append key + non-empty slices to filteredThreeSums
	- Eliinate duplicates from filteredThreeSums -- how do you know if duplicate is different indexes though?
	- Return filteredThreeSums


*/

package main

import "fmt"

func threeSum(nums []int) [][]int {
	correspondingThreeNumbers := map[int][][]int{}

	for index, number := range nums {
		correspondingThreeNumbers[number] = twoSum(nums, number*-1, index)
	}

	filteredThreeSums := [][]int{}

	for key, value := range correspondingThreeNumbers {
		if len(value) != 0 {
			for _, double := range value {
				double = append(double, key)
				filteredThreeSums = append(filteredThreeSums, double)
			}

		}
	}

	// fmt.Println(correspondingThreeNumbers)
	fmt.Println(filteredThreeSums)

	return filteredThreeSums
}

func twoSum(nums []int, target int, notThisIndex int) (complementNumbers [][]int) {
	correspondingNumbers := map[int]int{}

	for index, number := range nums {
		correspondingNumbers[number] = index + 1
		complementNumber := target - number
		if correspondingNumbers[complementNumber] != 0 && correspondingNumbers[complementNumber] != notThisIndex+1 && index != correspondingNumbers[complementNumber]-1 {
			complementNumbers = append(complementNumbers, []int{number, nums[correspondingNumbers[complementNumber]-1]})
		}
	}
	return complementNumbers
}

func main() {
	example := []int{-1, 0, 1, 2, -1, -4}
	example2 := []int{0, 0, 0}
	threeSum(example)
	threeSum(example2)
}
