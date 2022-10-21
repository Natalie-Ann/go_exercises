// Given an integer array nums of unique elements, return all possible subsets (the power set).

// The solution set must not contain duplicate subsets. Return the solution in any order.

// Input: nums = [1,2,3]
// Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

// Input: nums = [0]
// Output: [[],[0]]

// candidate     -> current path of exploration
// results  -> successful finds

package main

import "fmt"

func subsets(list []int) [][]int {
	var results [][]int
	var candidate []int
	backtrack(list, candidate, &results)
	return results
}

func backtrack(list []int, candidate []int, results *[][]int) {
	copyCandidate := make([]int, len(candidate))
	copy(copyCandidate, candidate)
	*results = append(*results, copyCandidate)

	for i := 0; i < len(list); i++ {
		candidate = append(candidate, list[i])    // take
		backtrack(list[i+1:], candidate, results) // explore
		candidate = candidate[:len(candidate)-1]  // clean up
	}
}

func main() {
	result := subsets([]int{1, 2, 3})
	fmt.Println(result)
}
