package main

import (
	"fmt"
)

/*
58. Length of Last Word
https://leetcode.com/problems/length-of-last-word
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
