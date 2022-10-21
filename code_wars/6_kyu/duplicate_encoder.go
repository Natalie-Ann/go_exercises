// The goal of this exercise is to convert a string to a new string where each character in the new string is "(" if that character appears only once in the original string, or ")" if that character appears more than once in the original string. Ignore capitalization when determining if a character is a duplicate.

// "din"      =>  "((("
// "recede"   =>  "()()()"
// "Success"  =>  ")())())"
// "(( @"     =>  "))(("

package main

import (
	"fmt"
	"strings"
)

func DuplicateEncode(word string) string {
	word = strings.ToLower(word)
	duplicateTracker := make(map[string]int)
	var newString string

	for _, value := range word {
		duplicateTracker[string(value)] += 1
	}

	for _, value := range word {
		if duplicateTracker[string(value)] > 1 {
			newString += ")"
		} else {
			newString += "("
		}
	}

	return newString
}

func main() {
	fmt.Println(DuplicateEncode("din"))
	fmt.Println(DuplicateEncode("recede"))
	fmt.Println(DuplicateEncode("Success"))
	fmt.Println(DuplicateEncode("(( @"))
}
