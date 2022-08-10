// You are given a string of space separated numbers, and have to return the highest and lowest number.

// Examples
// HighAndLow("1 2 3 4 5")  // return "5 1"
// HighAndLow("1 2 -3 4 5") // return "5 -3"
// HighAndLow("1 9 3 4 -5") // return "9 -5"
// Notes
// All numbers are valid Int32, no need to validate them.
// There will always be at least one number in the input string.
// Output string must be two numbers separated by a single space, and highest number is first.

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func HighAndLow(inputString string) string {
	//sort ints in inputString
	splitString := strings.Split(inputString, " ")
	var numberSlice = []int{}
	for _, value := range splitString {
		newValue, _ := strconv.Atoi(value)
		numberSlice = append(numberSlice, newValue)
	}

	sort.Ints(numberSlice)

	newString := fmt.Sprintf("%d %d", numberSlice[len(numberSlice)-1], numberSlice[0])

	return newString
}

func main() {
	fmt.Println(HighAndLow("1 9 3 4 -5"))
}
