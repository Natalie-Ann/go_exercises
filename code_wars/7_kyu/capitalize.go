//Capitalize first letter of every word in string

package main

import (
	"fmt"
	"strings"
)

// func Capitalize(str string) string {
// 	splitString := strings.Split(str, " ")
// 	var capitalizedString string
// 	for _, word := range splitString {
// 		word = strings.ToUpper(string(word[0])) + word[1:]
// 		capitalizedString += word + " "
// 	}

// 	return strings.Trim(capitalizedString, " ")

// }

func Capitalize(str string) string {
	return strings.Title(str)
}

func main() {
	fmt.Println(Capitalize("How can mirrors be real if our eyes aren't real"))
}
