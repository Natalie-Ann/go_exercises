/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false

Algorithm:
- Create empty slice (stack)
- Iterate through string
	- If character equals an opening bracket
		- append the character to the stack
	- If character equals a closing bracket
		- If stack is empty, return false (because corresponding opening bracket was not added before)
		- Pop off the last character from the stack and save to a variable
		- If the last character does not equal a matching opening bracket
			- return false
	- Return length of stack
*/

package main

func isValid(s string) bool {
	stack := []byte{}

	for i := 0; i <= len(s)-1; i++ {
		switch s[i] {
		case '{', '[', '(':
			stack = append(stack, s[i])
		case '}', ']', ')':
			if len(stack) == 0 {
				return false
			}
			lastBracket := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if s[i] == '}' && lastBracket != '{' || s[i] == ']' && lastBracket != '[' || s[i] == ')' && lastBracket != '(' {
				return false
			}
		}
	}
	return len(stack) == 0
}
