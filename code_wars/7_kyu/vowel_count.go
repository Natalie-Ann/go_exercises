package main

import (
	"fmt"
)

// func vowelCount(str string) (count int) {
// 	vowels := "aeiou"
// 	for _, value := range str {
// 		if strings.Contains(vowels, string(value)) {
// 			count += 1
// 		}
// 	}
// 	return count
// }

func vowelCount(str string) (count int) {
	for _, character := range str {
		switch character {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(vowelCount("testing"))
}
