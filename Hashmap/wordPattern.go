package main

import (
	"fmt"
)

/*
290. Word Pattern
https://leetcode.com/problems/word-pattern

I tried not to use split function
*/
func wordPattern(pattern string, s string) bool {
	itrString := 0
	current := 0
	patternMap := make(map[byte]string)
	wordMap := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		if itrString >= len(s) {
			return false
		}
		current = itrString
		for current < len(s) && s[current] != ' ' {
			current++
		}
		if word, exist := patternMap[pattern[i]]; exist {
			if word != s[itrString:current] {
				return false
			}
		} else {
			if word, exist := wordMap[s[itrString:current]]; exist {
				if word != pattern[i] {
					return false
				}
			}
			patternMap[pattern[i]] = s[itrString:current]
			wordMap[s[itrString:current]] = pattern[i]
		}
		itrString = current + 1
	}
	if itrString < len(s) {
		return false
	}
	return true
}

func findKeyByValue(inputMap map[byte]string, targetValue string) bool {
	for _, value := range inputMap {
		if value == targetValue {
			return true
		}
	}
	return false
}

func main() {
	pattern := "abc"
	s := "abc"

	k := wordPattern(pattern, s)

	fmt.Print(k)
}
