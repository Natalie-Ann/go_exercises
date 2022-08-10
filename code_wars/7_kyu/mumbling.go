// Examples:
// accum("abcd") -> "A-Bb-Ccc-Dddd"
// accum("RqaEzty") -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
// accum("cwAt") -> "C-Ww-Aaa-Tttt"
// The parameter of accum is a string which includes only letters from a..z and A..Z.

//convert string to array of strings
//change each string to first letter capital + same letter lowercase repeated index times
//convert slice to string with "-" separator
//return new string

package main

import (
	"strings"
)

func Accum(s string) string {
	stringArray := strings.Split(s, "")
	for index, value := range stringArray {
		stringArray[index] = strings.ToUpper(value) + strings.Repeat(strings.ToLower(value), index)
	}
	returnString := strings.Join(stringArray, "-")
	return returnString
}

func main() {
	Accum("abcd")
}
