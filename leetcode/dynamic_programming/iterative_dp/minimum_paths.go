/*
Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.

Input: grid =

1,3,1			1 4 5			1  4  5
1,5,1 --> 2	_	_	-->	2	 7	6
4,2,1			6	_	_			6	 8 	7


1 2 3			1 3 6			1  3   6
4 5 6 -->	4	_	_	-->	5  10  12


Output: 7
Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.

Input: grid = [[1,2,3],[4,5,6]]
Output: 12

Algorithm:
- Iterate over first row values and replace values with current sum
- Iterate over first column values and replace values with current sum
- Fill in the rest of the grid
	- For empty squares, look at square above and to left
		- if values in current square + above square < values in current square + left square
			- change value of current square to current square + above square
		- else change current square to current square + left square
- Return right corner square: grid[len(grid)][len(grid)[0] - 1]
*/

package main

import "fmt"

func minPathSum(grid [][]int) int {

	//iterate through first row, replacing values
	for i := 1; i < len(grid[0]); i++ {
		grid[0][i] = grid[0][i-1] + grid[0][i]
	}

	//iterate through first column, replacing values
	for i := 1; i < len(grid); i++ {
		grid[i][0] = grid[i-1][0] + grid[i][0]
	}

	//Fill in rest of grid, from second row --> last row
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			currentSquare := grid[i][j]
			aboveSquare := grid[i-1][j]
			leftSquare := grid[i][j-1]

			if currentSquare+aboveSquare < currentSquare+leftSquare {
				grid[i][j] = currentSquare + aboveSquare
			} else {
				grid[i][j] = currentSquare + leftSquare
			}
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]

}

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	fmt.Println(minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}))
}
