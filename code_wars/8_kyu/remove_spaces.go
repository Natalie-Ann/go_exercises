// Remove the spaces from the string, then return the resultant string.

package main

import (
	"fmt"
	"strings"
)

// func NoSpace(word string) string {
// 	var newWord string
// 	for _, char := range word {
// 		if string(char) == " " {
// 			continue
// 		}
// 		newWord += string(char)
// 	}
// 	return newWord
// }

func NoSpace(word string) string {
	return strings.ReplaceAll(word, " ", "")
}

func main() {
	fmt.Println(NoSpace("   S P A C E S    "))
}
