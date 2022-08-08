// Write a function that accepts an integer n and a string s as parameters, and returns a string of s repeated exactly n times.

// 6, "I"     -> "IIIIII"
// 5, "Hello" -> "HelloHelloHelloHelloHello"

package main

import (
	"fmt"
)

func stringRepeat(n int, s string) string {
	var repeatString string
	for i := 0; i <= n-1; i++ {
		repeatString = repeatString + s
	}
	return repeatString
}

func main() {
	fmt.Println(stringRepeat(6, "I"))
	fmt.Println(stringRepeat(5, "Hello"))
}

// import "strings"

// func RepeatStr(repetitions int, value string) string {
//   return strings.Repeat(value, repetitions)
// }
