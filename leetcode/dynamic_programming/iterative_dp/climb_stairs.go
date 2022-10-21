// You are climbing a staircase. It takes n steps to reach the top.

// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

// Input: n = 2
// Output: 2
// Explanation: There are two ways to climb to the top.
// 1. 1 step + 1 step
// 2. 2 steps

// Input: n = 3
// Output: 3
// Explanation: There are three ways to climb to the top.
// 1. 1 step + 1 step + 1 step
// 2. 1 step + 2 steps
// 3. 2 steps + 1 step

/*
Algorithm
- Number of ways = previous 2 number of ways added together
- Create slice with values
	- steps[0] = 1
	- steps[1] = 2
	- steps[2] = steps[0] + steps[1]...
- For loop to create this slice until reach nth step and return the computed value at that step


Algorithm2
- Create a map with int key representing step, int value representing number of ways
- Set steps[1] = 1, steps[2] = 2
- Iterate from 3 --> n, adding previous 2 steps
*/

package main

import "fmt"

// func climbStairs(n int) int {
// 	steps := []int{1, 2}

// 	for i := 2; i <= n-1; i++ {
// 		steps = append(steps, steps[i-2]+steps[i-1])
// 	}

// 	return steps[n-1]
// }

func climbStairs(n int) int {
	steps := map[int]int{}
	steps[1] = 1
	steps[2] = 2

	for i := 3; i <= n; i++ {
		steps[i] = steps[i-2] + steps[i-1]
	}

	return steps[n]
}

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
}
