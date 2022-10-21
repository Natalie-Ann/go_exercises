/* Write a function that reverses a string. The input string is given as an array of characters s.

You must do this by modifying the input array in-place with O(1) extra memory.

Examples:
["h","e","l","l","o"] --> ["o","l","l","e","h"]

["H","a","n","n","a","h"]--> ["h","a","n","n","a","H"]

PEDAC
Input: slice of bytes (array of characters)
Output: none; mutates string/slice of bytes

Explicits: 	- mutate slice in place; do not create new slice
						- reverse characters/bytes
Implicits:	- maintain case

DS: slice

Algorithm:
- Declare head and tail pointer (beginning and end of slice)
- Iterate through slice until header and tail meet in the middle (headPointer < tailPointer)
	- For each iteration, swap the characters at beginning and end pointers
	- increase head pointer by 1
	- decrease tail pointer by 1
*/

package main

func reverseString(s []byte) {
	for headPointer, tailPointer := 0, len(s)-1; headPointer < tailPointer; headPointer, tailPointer = headPointer+1, tailPointer-1 {
		s[headPointer], s[tailPointer] = s[tailPointer], s[headPointer]
	}
}

func main() {
	example1 := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(example1)
}
