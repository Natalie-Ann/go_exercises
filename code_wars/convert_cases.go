package main

import (
	"fmt"
	"unicode"
)

func ToAlternatingCase(str string) string {
	result := ""
	for _, ch := range str {
		if unicode.IsUpper(ch) {
			result = result + string(unicode.ToLower(ch))
		} else if unicode.IsLower(ch) {
			result = result + string(unicode.ToUpper(ch))
		} else {
			result = result + string(ch)
		}
	}

	return result
}

func main() {
	fmt.Println(ToAlternatingCase("TEST test"))
}
