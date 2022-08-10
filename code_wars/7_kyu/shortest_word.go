// Given a string of words, return the length of the shortest word(s).

// String will never be empty and you do not need to account for different data types.

package main

import (
	"fmt"
	// "sort"
	"strings"
)

// type byLength []string

// func (s byLength) Len() int {
// 	return len(s)
// }

// func (s byLength) Swap(i, j int) {
// 	s[i], s[j] = s[j], s[i]
// }

// func (s byLength) Less(i, j int) bool {
// 	return len(s[i]) < len(s[j])
// }

// func FindShort(s string) int {
// 	stringSlice := strings.Split(s, " ")
// 	sort.Sort(byLength(stringSlice))
// 	fmt.Println(stringSlice)
// 	return len(stringSlice[0])
// }

func FindShort(s string) (c int) {
	for _, word := range strings.Fields(s) {
		if c == 0 || len(word) < c {
			c = len(word)
		}
	}
	return c
}

func main() {
	fmt.Println(FindShort("bitcoin take over the world maybe who knows perhaps"))
}
