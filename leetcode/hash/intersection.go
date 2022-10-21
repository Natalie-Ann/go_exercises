/*
Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must be unique and you may return the result in any order.

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Explanation: [4,9] is also accepted.

Inputs: 2 arrays
Output: 1 array

Explicits: - returned array has numbers elements that are in both input arrays
        - only unique numbers in output
        - empty array returned if no overlap
        - any order to output array
Implicits: - output array max length can only be the length of the shorter array; length of the shorter array filtered for unique values

DS: arrays

Algorithm: O(N * M)
- Filter both arrays for unique elements/numbers
- Find the shorter array = shortArray
- Iterate through the short array and if the number from longArray is also in the shortArray, then add it to output array
- return output array

Algorithm2: O(N + M)
- Declare a empty hash map
- Declare an empty slice as output array
- Iterate through one array
    - For each number, set value to 1
    - if encounter same number again, continue (do not change to 2)
- Iterate through second array
    - For each number,
        - if number in map is > 0
					- if number equals 1
            - add number to output array
            - increase value to 2
					- else continue
- Return output array
*/

package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	matches := map[int]int{}
	outputSlice := []int{}

	for _, val := range nums1 {
		matches[val] = 1
	}

	for _, val := range nums2 {
		if matches[val] > 0 {
			if matches[val] == 1 {
				outputSlice = append(outputSlice, val)
				matches[val] += 1
			}
		}
	}

	return outputSlice
}

func main() {
	example := []int{1, 2, 2, 1}
	example2 := []int{2, 2}

	example3 := []int{4, 9, 5}
	example4 := []int{9, 4, 9, 8, 4}

	fmt.Println(intersection(example, example2))
	fmt.Println(intersection(example3, example4))
}
