// Given an array of integers, find the one that appears an odd number of times.

// There will always be only one integer that appears an odd number of times.

// Examples
// [7] should return 7, because it occurs 1 time (which is odd).
// [0] should return 0, because it occurs 1 time (which is odd).
// [1,1,2] should return 2, because it occurs 1 time (which is odd).
// [0,1,0,1,0] should return 0, because it occurs 3 times (which is odd).
// [1,2,2,3,3,3,4,3,3,3,2,2,1] should return 4, because it appears 1 time (which is odd).

package main

func FindOdd(seq []int) (oddNumber int) {
	countMap := make(map[int]int)

	for _, number := range seq {
		if _, ok := countMap[number]; ok {
			countMap[number] += 1
		} else {
			countMap[number] = 1
		}
	}

	for key, value := range countMap {
		if value%2 == 1 {
			oddNumber = key
		}
	}

	return oddNumber
}

func main() {
	FindOdd([]int{1, 1, 2})
	FindOdd([]int{0, 1, 0, 1, 0})
	FindOdd([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1})
}
