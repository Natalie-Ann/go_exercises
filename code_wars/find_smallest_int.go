package main

func SmallestIntegerFinder(numbers []int) int {
	var smallestInt int = numbers[0]
	for _, val := range numbers {
		if smallestInt > val {
			smallestInt = val
		}
	}
	return smallestInt
}

/*
import "sort"

func SmallestIntegerFinder(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0]
}

*/
