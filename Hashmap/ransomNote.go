package main

import (
	"fmt"
)

/*
383. Ramson Note
https://leetcode.com/problems/ransom-note
*/
func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	if len(magazine) <= 0 {
		return false
	}

	var result = true
	magMap := make(map[rune]int)

	for _, c := range magazine {
		magMap[c]++
	}
	for _, c := range ransomNote {
		magMap[c]--
		if magMap[c] < 0 {
			result = false
			break
		}
	}

	return result
}

func main() {
	ransomNote := "aa"
	magazine := "aab"

	k := canConstruct(ransomNote, magazine)

	fmt.Print(k)
}
