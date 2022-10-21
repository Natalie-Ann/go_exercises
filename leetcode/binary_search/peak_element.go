// A peak element is an element that is strictly greater than its neighbors.

// Given a 0-indexed integer array nums, find a peak element, and return its index. If the array contains multiple peaks, return the index to any of the peaks.

// You may imagine that nums[-1] = nums[n] = -∞. In other words, an element is always considered to be strictly greater than a neighbor that is outside the array.

// You must write an algorithm that runs in O(log n) time.

/*
Example 1:

Input: nums = [1,2,3,1]
Output: 2
Explanation: 3 is a peak element and your function should return the index number 2.
Example 2:

Input: nums = [1,2,1,3,5,6,4]
Output: 5
Explanation: Your function can return either index number 1 where the peak element is 2, or index number 5 where the peak element is 6.


Input: slice of ints
Output: int representing index

Explicits: - peak is number that has 2 surrounding numbers that are less than it
						- must be O(log n) - binary search
						- elements outside of array are considered to be lower
Implicits: 	- can early return; any peak counts
						- no repeat numbers next to each other


DS: arrays

Algorithm1:
- Anchor at 0, runner at 1
- Iterate through array
	- runner moves with every iteration
	- if runner number is lower than anchor number and anchor index is 0, return index (early return)
	- else if runner number is higher than anchor number
		- move anchor to right 1
		- if runner is at the end of the array return runner index
	- else if runner number is lower than anchor number
		- check prior anchor number
			- if prior number is lower, return anchor index
			- if prior number is higher, move anchor right 1

Algorithm2:
-
*/

package main

import "fmt"

func findPeakElement(nums []int) int {
	anchor, runner := 0, 1

	if len(nums) == 1 {
		return 0
	}

	for index, _ := range nums {
		if nums[runner] < nums[anchor] && anchor == 0 {
			return index
		} else if nums[runner] > nums[anchor] {
			anchor++
			if runner == len(nums)-1 {
				return runner
			}
		} else if nums[runner] < nums[anchor] {
			if nums[anchor-1] < nums[anchor] {
				return anchor
			} else {
				anchor++
			}
		}
		runner++
	}

	return 0

}

func main() {
	example := []int{1, 2, 3, 1}
	example2 := []int{1, 2, 1, 3, 5, 6, 4}
	fmt.Println(findPeakElement(example))
	fmt.Println(findPeakElement(example2))
}
