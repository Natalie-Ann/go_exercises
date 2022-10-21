/*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

A subarray is a contiguous part of an array.

Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Example 2:

Input: nums = [1]
Output: 1
Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23
*/

package main

import "fmt"

func maxSubarray(nums []int) int {
	right := len(nums) - 1

	return findMaxSubArray(nums, 0, right)
}

func findMaxSubArray(nums []int, left int, right int) int {
	if left == right {
		return nums[left]
	} else {
		mid := (left + right) / 2

		//find max subarray right
		//find max subarray left
		//find max subarray intersection
		maxLeft := findMaxSubArray(nums, left, mid)
		maxRight := findMaxSubArray(nums, mid+1, right)
		maxIntersection := maxCrossing(nums, left, right, mid)

		//find maximum of left, right, and intersection
		return max(maxLeft, maxRight, maxIntersection)
	}
}

func max(values ...int) int {
	maxVal := values[0]
	for _, v := range values {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func maxCrossing(nums []int, left int, right int, mid int) int {
	//array must pass midpoint;  find max of both halfs of array and add up
	leftSum := -9999999
	rightSum := -9999999
	sum := 0

	for i := mid; i >= left; i-- {
		sum += nums[i]
		leftSum = max(leftSum, sum)
	}

	sum = 0

	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		rightSum = max(rightSum, sum)
	}

	return leftSum + rightSum
}

func main() {
	example := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Println(maxSubarray(example))
}
