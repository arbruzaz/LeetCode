package main

import (
	"fmt"
)

/*
20. Valid Parentheses
https://leetcode.com/problems/valid-parentheses
*/

func isValid(s string) bool {
	charList := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	var charStack []rune

	for _, char := range s {
		if len(charStack) == 0 {
			charStack = append(charStack, char)
			continue
		}

		if charStack[len(charStack)-1] == charList[char] {
			charStack = charStack[:len(charStack)-1]
			continue
		}

		charStack = append(charStack, char)

	}

	return len(charStack) == 0
}

func main() {

	s := "()((())))"

	k := isValid(s)

	fmt.Print(k)
}
