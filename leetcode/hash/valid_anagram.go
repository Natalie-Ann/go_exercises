/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Input: s = "anagram", t = "nagaram"
Output: true

Input: s = "rat", t = "car"
Output: false

Explicits: - retrun bool if anagram
Implicits: 	- different length = not anagram
							- anagrams = same count of each letter in both words

DS: hash

Algorithm:
- If lengths are different return false
- Create hash map
- Iterate through string
	- Add each letter to hash map with tally
- Iterate through second string
	- For each letter, find letter in hash map
		- if doesn't exist, return false
		- if does exist, subtract 1
- Iterate through hash map
	- If tally does not equal zero, return false
- Return true
*/

package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	tallyHash := map[rune]int{}

	for _, val := range s {
		tallyHash[val] += 1
	}

	for _, val := range t {
		if tallyHash[val] == 0 {
			return false
		} else {
			tallyHash[val] -= 1
		}
	}

	for _, element := range tallyHash {
		if element != 0 {
			return false
		}
	}

	return true

}

func main() {
	string1 := "anagram"
	string2 := "nagaram"

	fmt.Println(isAnagram(string1, string2))

}
