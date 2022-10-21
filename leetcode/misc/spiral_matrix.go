/*
Input: array of arrays
Output: one array (slice) in spiral order

Explicits:
	- m is length of the outer array
	- n is lenght of inner arrays
	- outer array has to be length of at least 1
	- inner array numbers can rnage from -100-100
Implicits:
	- last 1-2 numbers are middle of rectangle
	- first numbers are the first array in order
	-  (3 x 3) --> 3, 2, 2, 1, 1 (middle)
	- (4 x 3) --> 4, 2, 3, 1, 2 middles

DS:

Algorithm:
- length of inner arrays (N) is first set of numbers
- next set of numbers is length of outer array - 1 (M  - 1); each number is +N of last
- thirst set of numbers is length N - 1; subtracting 1
- fourth set of numbers is lenfth M - 2; subtracting -N
--> until reach length 1

Algorithm:
- base case: single row, no columns --> only one slice left; outer array is length 1
	- return createdSlice + middle numbers in single slice left
- define edges and add edges to createdSlice; remove from original matrix
	- recursively call on spiralOrder (2 parameters - leftoever matrix, createdSlice)


Examples: [[1, 2, 3], [4, 5, 6], [7, 8, 9]] --> [1, 2, 3, 6, 8, 7, 4, 5]
					[[1,2,3,4],[5,6,7,8],[9,10,11,12]] --> [1,2,3,4,8,12,11,10,9,5,6,7]

*/

package main

import (
	"fmt"
)

func spiralOrder(matrix [][]int, createdSlice []int) []int {
	if len(matrix) == 1 {
		middleNumbers := matrix[0][1 : len(matrix[0])-1]

		// createdSlice = append(createdSlice, matrix[0]...)
		createdSlice = append(createdSlice, middleNumbers...)
		return createdSlice
	}

	edges := getEdges(matrix)
	createdSlice = append(createdSlice, edges...)
	leftOverMatrix := removeEdges(matrix, edges)

	return spiralOrder(leftOverMatrix, createdSlice)
}

func getEdges(matrix [][]int) (edges []int) {
	//first slice
	firstSlice := matrix[0]
	fmt.Println(firstSlice)

	//last elements
	var lastElements []int
	for _, innerMatrix := range matrix {
		lastElements = append(lastElements, innerMatrix[len(innerMatrix)-1])
	}
	lastElements = lastElements[1 : len(lastElements)-1]
	fmt.Println(lastElements)

	//last slice
	lastSlice := matrix[len(matrix)-1]
	var reversedLastSlice []int
	for index := len(lastSlice) - 1; index >= 0; index-- {
		reversedLastSlice = append(reversedLastSlice, lastSlice[index])
	}
	fmt.Println(reversedLastSlice)

	//first elements in reverse
	var firstElements []int
	for index := len(matrix) - 1; index >= 0; index-- {
		firstElements = append(firstElements, matrix[index][0])
	}
	firstElements = firstElements[1 : len(firstElements)-1]
	fmt.Println(firstElements)

	edges = append(edges, firstSlice...)
	edges = append(edges, lastElements...)
	edges = append(edges, reversedLastSlice...)
	edges = append(edges, firstElements...)

	// edges = append(edges, firstSlice..., lastElements..., reversedLastSlice..., firstElements...)

	return edges
}

func removeEdges(matrix [][]int, edges []int) (leftOverMatrix [][]int) {
	//remove first slice
	leftOverMatrix = matrix[1:]

	//remove first and last elements
	for _, innerMatrix := range leftOverMatrix {
		innerMatrix = innerMatrix[1 : len(innerMatrix)-1]
	}

	//remove last slice
	leftOverMatrix = leftOverMatrix[:len(leftOverMatrix)-1]
	fmt.Println(leftOverMatrix)

	return leftOverMatrix
}

// func spiralOrder(matrix [][]int) []int {
// 	// M := len(matrix)
// 	// N := len(matrix[0])
// 	var newArray []int

// 	for _, int := range matrix[0] {
// 		newArray = append(newArray, int)
// 	}

// 	fmt.Println(newArray)

// 	return newArray

// }

func main() {
	slice := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// anotherSlice := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	var createdSlice []int
	// fmt.Println(getEdges(slice))
	fmt.Println(spiralOrder(slice, createdSlice))
	// fmt.Println(spiralOrder(anotherSlice, createdSlice))
}
