/*You are given two sorted arrays that both only contain integers. Your task is to find a way to merge them into a single one, sorted in asc order. Complete the function mergeArrays(arr1, arr2), where arr1 and arr2 are the original sorted arrays.

You don't need to worry about validation, since arr1 and arr2 must be arrays with 0 or more Integers. If both arr1 and arr2 are empty, then just return an empty array.

Note: arr1 and arr2 may be sorted in different orders. Also arr1 and arr2 may have same integers. Remove duplicated in the returned result.

* [1, 2, 3, 4, 5], [6, 7, 8, 9, 10] -> [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

* [1, 3, 5, 7, 9], [10, 8, 6, 4, 2] -> [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

* [1, 3, 5, 7, 9, 11, 12], [1, 2, 3, 4, 5, 10, 12] -> [1, 2, 3, 4, 5, 7, 9, 10, 11, 12]
*/

package main

import (
	"fmt"
	"sort"
)

func mergeSortedArrays(arr1 []int, arr2 []int) []int {
	if len(arr1) == 0 && len(arr2) == 0 {
		return []int{}
	}

	mergedSlices := append(arr1, arr2...)
	sort.Ints(mergedSlices)

	var noDuplicates []int

	for index, value := range mergedSlices {
		if index != len(mergedSlices)-1 && mergedSlices[index+1] == value {
			continue
		}
		noDuplicates = append(noDuplicates, value)
	}

	return noDuplicates
}

func main() {
	first := []int{1, 3, 5, 7, 9, 11, 12}
	second := []int{1, 2, 3, 4, 5, 10, 12}

	fmt.Println(mergeSortedArrays(first, second))
}
