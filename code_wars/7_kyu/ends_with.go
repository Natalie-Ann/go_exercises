// Complete the solution so that it returns true if the first argument(string) passed in ends with the 2nd argument (also a string).

// solution("abc", "bc") // returns true
// solution("abc", "d") // returns false

package main

import (
	"fmt"
	"strings"
)

// func endsWith(str, ending string) bool {
// 	if len(str) < len(ending) {
// 		return false
// 	}

// 	lengthOfEnding := len(ending)

// 	return str[len(str)-lengthOfEnding:] == ending
// }

func endsWith(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}

func main() {
	fmt.Println(endsWith("abc", "bc"))
	fmt.Println(endsWith("abc", "d"))
}
