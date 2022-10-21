package main

func permute(nums []int) [][]int {
	var results [][]int
	var candidate []int
	backtrack(nums, candidate, &results)
	return results
}

func backtrack(list []int, candidate []int, results *[][]int) {
	if len(candidate) == len(list) {
		//copy candidate and add to results
		copyCandidate := make([]int, len(candidate))
		copy(copyCandidate, candidate)
		*results = append(*results, copyCandidate)
	}

	for i := 0; i < len(list); i++ {
		//prune here for unique number
		if isIncluded(list[i], candidate) {
			continue
		}
		candidate = append(candidate, list[i])   //take
		backtrack(list, candidate, results)      //explore
		candidate = candidate[:len(candidate)-1] //clean up

	}

}

func isIncluded(num int, list []int) bool {
	for _, val := range list {
		if val == num {
			return true
		}
	}
	return false
}
