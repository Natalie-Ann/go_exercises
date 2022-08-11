// Take 2 strings s1 and s2 including only letters from a to z. Return a new sorted string, the longest possible, containing distinct letters - each taken only once - coming from s1 or s2.

// a = "xyaabbbccccdefww"
// b = "xxxxyyyyabklmopq"
// longest(a, b) -> "abcdefklmopqwxy"

// a = "abcdefghijklmnopqrstuvwxyz"
// longest(a, a) -> "abcdefghijklmnopqrstuvwxyz"

package main

import (
	"fmt"
	"sort"
	"strings"
)

func TwoToOne(s1 string, s2 string) string {
	combinedString := s1 + s2
	splitString := strings.Split(combinedString, "")
	sort.Strings(splitString)
	var newString string

	for _, char := range splitString {
		if !(strings.Contains(newString, string(char))) {
			newString += string(char)
		}
	}

	return newString
}

func main() {
	fmt.Println(TwoToOne("abc", "abc"))
	fmt.Println(TwoToOne("aretheyhere", "yestheyarehere"))
	fmt.Println(TwoToOne("xyaabbbccccdefww", "xxxxyyyyabklmopq"))
	fmt.Println(TwoToOne("codwars", ""))
}
