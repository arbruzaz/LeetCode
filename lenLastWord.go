package main

import (
	"fmt"
)

/*
58. Length of Last Word
https://leetcode.com/problems/length-of-last-word
*/

func lengthOfLastWord(s string) int {
	wordLen := len(s) - 1
	flag := false
	result := 0
	for i := wordLen; i >= 0; i-- {
		if s[i] != ' ' {
			flag = true
			result++
		} else if flag == true {
			break
		}
	}
	return result
}

func main() {
	var s = "World"

	k := lengthOfLastWord(s)

	fmt.Print(k)
}
