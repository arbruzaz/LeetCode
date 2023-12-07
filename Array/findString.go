package main

import (
	"fmt"
)

/*
28. Find The Index of the First in String
https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string
*/

func strStr(haystack string, needle string) int {
	j := 0
	result := -1
	needleLen := len(needle) - 1
	for j+needleLen < len(haystack) {
		if needle == haystack[j:j+needleLen+1] {
			result = j
			break
		}
		j++
	}
	return result
}

func main() {
	var haystack = "mississippi"
	var needle = "issip"

	k := strStr(haystack, needle)

	fmt.Print(k)
}
