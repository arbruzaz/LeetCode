package main

import (
	"fmt"
	"strings"
)

/*
125. Valid Palindrome
https://leetcode.com/problems/valid-palindrome
*/
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	result := true

	i := 0
	j := len(s) - 1

	for i < len(s) && j > i && i != j {
		for (s[i] > 122 || s[i] < 97) && (s[i] > 57 || s[i] < 48) && i != j {
			i++
		}
		for (s[j] > 122 || s[j] < 97) && (s[j] > 57 || s[j] < 48) && i != j {
			j--
		}
		if s[i] != s[j] {
			result = false
			break
		}
		i++
		j--
	}

	return result
}

func main() {
	var s = ",,,,,,,,,,,,acva"

	k := isPalindrome(s)

	fmt.Print(k)
}
