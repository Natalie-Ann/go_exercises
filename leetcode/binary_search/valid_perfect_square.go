// Given a positive integer num, write a function which returns True if num is a perfect square else False.

// Follow up: Do not use any built-in library function such as sqrt.

/*
Examples:
16 --> true
14 --> false

Input: integer
Output: boolean

Explicits: - return whether input is perfect square
						- can't use libraries
Implicits: 	- what is perfect square? number multiplied by itself = given number
						- 1 * 1 = 1
						- 2 * 2 = 4
						- 3 * 3 = 9
						- go up by 2 between each number, starting at 3 --> 0, 3, 5, 7, 9, 11 ...
						-

DS: arrays

Algorithm:
- Divide input number by 2
- If middle number * itself > input number, decrease middle by 1, then divide again in 2
- if middle number * itself < input numbmer, increase number by 1 and continue adding 1 and multiplying until reach >= input number
- if reach input number exxactly then return true
- else return false
*/

package main

import "fmt"

// func isPerfectSquare(num int) bool {
// 	middleNumber := num / 2

// 	for middleNumber <= num {
// 		if middleNumber*middleNumber == num {
// 			return true
// 		} else if middleNumber*middleNumber > num {
// 			middleNumber = middleNumber/2
// 		} else {
// 			middleNumber = middleNumber + 1
// 		}
// 	}

// 	fmt.Println(middleNumber)

// 	return false
// }

func isPerfectSquare(num int) bool {
	for i := 1; i*i <= num; i += 1 {
		if i*i == num {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(isPerfectSquare(4))
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(14))
}
