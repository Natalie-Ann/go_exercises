// Write a function that accepts an array of 10 integers (between 0 and 9), that returns a string of those numbers in the form of a phone number.

package main

import "fmt"

// func CreatePhoneNumber(numbers [10]uint) string {
// 	var firstThree string
// 	var secondThree string
// 	var finalFour string

// 	for index, digit := range numbers {
// 		if index == 0 || index == 1 || index == 2 {
// 			firstThree += fmt.Sprint(digit)
// 		} else if index == 3 || index == 4 || index == 5 {
// 			secondThree += fmt.Sprint(digit)
// 		} else {
// 			finalFour += fmt.Sprint(digit)
// 		}
// 	}

// 	firstThree = "(" + firstThree + ")"
// 	secondHalf := secondThree + "-" + finalFour

// 	return fmt.Sprint(firstThree, " ", secondHalf)
// }

func CreatePhoneNumber(n [10]uint) string {
	return fmt.Sprintf("(%d%d%d) %d%d%d-%d%d%d%d", n[0], n[1], n[2], n[3], n[4], n[5], n[6], n[7], n[8], n[9])
}

func main() {
	fmt.Println(CreatePhoneNumber([10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}
