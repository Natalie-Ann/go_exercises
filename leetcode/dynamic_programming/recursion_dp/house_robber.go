/*
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.

Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.
Example 2:

Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
Total amount you can rob = 2 + 9 + 1 = 12.

Algorithm:
- Base case: length of 1
	- return value at index 0
- Recursively call max(rob(first house + 3rd house to end), rob(secondhouse, 4th house to end))
- Define a function (max) comparing 2 numbers and returning the max
*/

package main

func rob(nums []int) int {

	return helper(nums)

}

func helper(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	return max(nums[0]+helper(nums[2:]), nums[1]+helper(nums[3:]))
}

func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
