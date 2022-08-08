/*
Complete the solution so that it reverses the string passed into it.

'world'  =>  'dlrow'
'word'   =>  'drow'

*/

package main

import (
	"fmt"
)

func reverseString(inputString string) string {
	var newString string
	for i := len(inputString) - 1; i >= 0; i-- {
		newString = newString + string((inputString[i]))
	}
	return newString
}

func main() {
	fmt.Println(reverseString("world"))
}
