/*
You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.

Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.

Questions:
Starting from index  0
- Can you get to end (last index)? yes: return true
- Can you get to next to last?
	- yes: cache and ask same question from second to last to end
	- no: Can you get to third from last? ...etc
- If go through full array and cant get to second element, then return false


Algorithm:
- Base case(s):
- Min steps to reach end: length of array - (currentIndex + 1)
- Cache: key = current index; value = furthest index the current index can reach
- Recurse:
*/

package main

import "fmt"

func canJump(nums []int) bool {
	// var cache map[int]bool
	targetIndex := len(nums) - 1
	return helper(nums, nums, targetIndex, 0)
}

func helper(nums []int, slicedNums []int, targetIndex int, currentIndex int) bool {
	currentMinSteps := minStepsNeeded(slicedNums)

	if targetIndex < 1 {
		return false
	}

	if targetIndex == len(nums)-1 && nums[currentIndex] >= currentMinSteps {
		return true
	}

	if currentMinSteps > slicedNums[currentIndex] {
		fmt.Println("checking previous reachability")
		slicedNums := nums[0:targetIndex]
		targetIndex = len(slicedNums) - 1
		return helper(nums, slicedNums, targetIndex, 0)
	}

	return helper(nums, nums[targetIndex:], targetIndex, 0)

}

func minStepsNeeded(nums []int) int {
	return len(nums) - 1
}

func main() {
	// fmt.Println(canJump([]int{3, 2, 1, 1, 4}))
	// fmt.Println(canJump([]int{1, 2, 3}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
}
