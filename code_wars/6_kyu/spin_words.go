// Write a function that takes in a string of one or more words, and returns the same string, but with all five or more letter words reversed (Just like the name of this Kata). Strings passed in will consist of only letters and spaces. Spaces will be included only when more than one word is present.

// spinWords( "Hey fellow warriors" ) => returns "Hey wollef sroirraw"
// spinWords( "This is a test") => returns "This is a test"
// spinWords( "This is another test" )=> returns "This is rehtona test"

package main

import (
	"fmt"
	"strings"
)

func Reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func SpinWords(str string) string {
	wordArray := strings.Split(str, " ")

	var wordSlice []string

	for _, word := range wordArray {
		if len(word) >= 5 {
			word = Reverse(word)
		}
		wordSlice = append(wordSlice, word)
	}
	return strings.Join(wordSlice, " ")
}

func main() {
	fmt.Println(SpinWords("Hey fellow warriors"))
	fmt.Println(SpinWords("This is a test"))
	fmt.Println(SpinWords("This is another test"))
}
