// You are going to be given a word. Your job is to return the middle character of the word. If the word's length is odd, return the middle character. If the word's length is even, return the middle 2 characters.
// #Examples:
// Kata.getMiddle("test") should return "es"
// Kata.getMiddle("testing") should return "t"
// Kata.getMiddle("middle") should return "dd"
// Kata.getMiddle("A") should return "A"

package main

import (
	"fmt"
)

func GetMiddle(s string) string {
	if len(s)%2 == 0 {
		endIndex := len(s) / 2
		firstMiddleChar := string(s[endIndex-1])
		secondMiddleChar := string(s[endIndex])
		return fmt.Sprintf("%s%s", firstMiddleChar, secondMiddleChar)
	} else {
		middleIndex := len(s) / 2
		middleChar := string(s[middleIndex])
		return fmt.Sprintf(middleChar)
	}
}

func main() {
	fmt.Println(GetMiddle("test"))
	fmt.Println(GetMiddle("testing"))
	fmt.Println(GetMiddle("middle"))
	fmt.Println(GetMiddle("A"))
}
