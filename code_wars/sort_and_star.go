package main

import (
	"sort"
	"strings"
)

func TwoSort(arr []string) string {
	sort.Strings(arr)
	firstString := arr[0]
	slicedString := strings.Split(firstString, "")
	joinedString := strings.Join(slicedString, "***")
	return joinedString
}

func main() {
	slice := []string{"test", "apple", "pear"}
	TwoSort(slice)
}
