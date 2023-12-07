package main

import (
	"fmt"
)

/*
14. Longest Common Prefix
https://leetcode.com/problems/longest-common-prefix
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	prefix := strs[0]
	for i := 0; i < len(prefix); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != prefix[i] {
				return prefix[:i]
			}
		}
	}
	return prefix
}

func main() {
	var str = []string{"ab", "a", "aab"}

	k := longestCommonPrefix(str)

	fmt.Print(k)
}
