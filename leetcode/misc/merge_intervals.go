// Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

// Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].

// Input: intervals = [[1,4],[4,5]]
// Output: [[1,5]]
// Explanation: Intervals [1,4] and [4,5] are considered overlapping.

// 1 <= intervals.length <= 104
// intervals[i].length == 2
// 0 <= starti <= endi <= 104

/*

when do intervals overlap?
	- inner numbers are the same [1, 4] [4, 5]
	- inner numbers are between outer numbers AND end of one interval > beginning of another [1, 3] [2, 6]
	- one interval swallows another [1, 5] [2, 3]

Algorithm:
- declare new slice (will contain the merged intervals)
- for each interval in intervals
	- iterate through other intervals
		- if current interval overlaps with another interval
			- merge intervals and append to newSlice
	- else append newSlice with current interval

*/

package main

import "fmt"

func merge(intervals [][]int) (newSlice [][]int) {
	for _, firstInterval := range intervals {
		for _, secondInterval := range intervals {
			if firstInterval[1] == secondInterval[0] {
				fmt.Println("equal inner numbers")
				mergedSlices := []int{firstInterval[0], secondInterval[1]}
				newSlice = append(newSlice, mergedSlices)
			} else if firstInterval[1] < secondInterval[1] && secondInterval[0] >= firstInterval[0] {
				fmt.Println("inner numbers in between")
				mergedSlices := []int{firstInterval[0], secondInterval[1]}
				newSlice = append(newSlice, mergedSlices)
			} else if secondInterval[0] > firstInterval[0] && secondInterval[1] < firstInterval[1] {
				fmt.Println("one interval contains another")
				newSlice = append(newSlice, firstInterval)
			}
			// } else if checkIfUniqueInterval(newSlice, firstInterval) {
			// 	fmt.Println("new interval")
			// 	newSlice = append(newSlice, firstInterval)
			// }
		}
	}
	return newSlice
}

func main() {
	example := [][]int{{1, 3}, {1, 6}}
	anotherExample := [][]int{{1, 4}, {4, 5}}
	thirdExample := [][]int{{1, 5}, {2, 3}}
	another := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(example))
	fmt.Println(merge(anotherExample))
	fmt.Println(merge(thirdExample))
	fmt.Println(merge(another))
}
