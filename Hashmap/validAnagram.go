package main

import (
	"fmt"
)

/*
242. Valid Anagram
https://leetcode.com/problems/valid-anagram/
*/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[rune]int)

	for _, c := range s {
		sMap[c]++
	}

	for _, c := range t {
		sMap[c]--
		if sMap[c] < 0 {
			return false
		}
	}

	return true
}

func main() {
	s := "anaagram"
	t := "nagaaram"

	k := isAnagram(s, t)

	fmt.Print(k)
}
