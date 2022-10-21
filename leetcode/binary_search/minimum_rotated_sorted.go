/*
Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:

[4,5,6,7,0,1,2] if it was rotated 4 times.
[0,1,2,4,5,6,7] if it was rotated 7 times.
Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

Given the sorted rotated array nums of unique elements, return the minimum element of this array.

You must write an algorithm that runs in O(log n) time.

Example 1:

Input: nums = [3,4,5,1,2]
Output: 1
Explanation: The original array was [1,2,3,4,5] rotated 3 times.
Example 2:

Input: nums = [4,5,6,7,0,1,2]
Output: 0
Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.
Example 3:

Input: nums = [11,13,15,17]
Output: 11
Explanation: The original array was [11,13,15,17] and it was rotated 4 times. 

Input: slice of ints
Output: int representing the lowest int in the slice

Explicits: 	- array is sorted and then rotated
						- use O(log n) time so binary search
						- all unique numbers
Implicits: 	- if rotates same number of times as length of input, it is fully ordered 
						- highest and lowest numbers next to each other, unless array is fully ordered
						- lowest number is where numbers to either side are higher OR it is at either end of a slice


DS: slice

Algorithm:
- Get start and end numbers of slice
- Declare pointer at midpoint of slice
- Compare numbers to start and end numbers
	- if midpoint is lower both sides, return that nubmer
	- if midpoint is higher than one side, change window to side that has lower number
		- get midpoint of that side and repeat until number is lower than both sides
*/

package main

import "fmt"

func findMin(nums []int) int {
	midpoint := len(nums)/2

	for nums[midpoint] > nums[midpoint + 1] || nums[midpoint] > nums[midpoint - 1] {
		if midpoint == len(nums - 1) || midpoint == nums[0] {
			return midpoint
		}

		if nums[midpoint] > nums[midpoint + 1] {
			//shift window right, from prev midpoint to end of array; need that index
			midpoint = 
		}
		//shift window left
		midpoint = 

	}
	return midpoint
}

