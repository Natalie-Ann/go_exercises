/*
Given an integer array nums of unique elements, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.

Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Input: nums = [0]
Output: [[],[0]]

*/

package main

func combinationSum(candidates []int, target int) [][]int {
	var results [][]int
	var candidate []int
	backtrack(candidates, target, candidate, &results)
	return results
}

func backtrack(list []int, candidate []int, results *[][]int) {
	if sum(candidate) == target {
		//copy candidate and add to results
		copyCandidate := make([]int, len(candidate))
		copy(copyCandidate, candidate)
		*results = append(*results, copyCandidate)
	}

	for i := 0; i < len(list); i++ {
		//prune here for unique number
		//
		candidate = append(candidate, list[i])   //take
		backtrack(list, candidate, results)      //explore
		candidate = candidate[:len(candidate)-1] //clean up

	}

}
