package main

func partition(s string) [][]string {
	var results [][]string
	var candidate []string
	backtrack(s, candidate, &results)
	return results
}

func reverse(s string) string {
	var result string

	for _, c := range s {
		result = string(c) + result
	}

	return result
}

func isPalindrome(s string) bool {
	return s == reverse(s)
}

func backtrack(s string, candidate []string, results *[][]string) {
	if len(s) == 0 {
		copyCandidate := make([]string, len(candidate))
		copy(copyCandidate, candidate)
		*results = append(*results, copyCandidate)
	}

	for i := 0; i < len(s); i++ {
		substring := s[0 : i+1]
		if isPalindrome(substring) {
			candidate = append(candidate, substring)
			backtrack(s[i+1:], candidate, results)
			candidate = candidate[:len(candidate)-1]
		}

	}
}
