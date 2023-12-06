package main

import (
	"fmt"
)

/*
392. Is Subsequence
https://leetcode.com/problems/is-subsequence
*/

func isSubsequence(s string, t string) bool {
	if len(s) <= 0 {
		return true
	}

	i := 0
	for j := 0; j <= len(t)-len(s)+i; j++ {
		if t[j] == s[i] {
			i++
			if i >= len(s) {
				break
			}
		}
	}

	return i == len(s)
}

func main() {
	var s = ""
	var t = "ahbgdc"

	k := isSubsequence(s, t)

	fmt.Print(k)
}
