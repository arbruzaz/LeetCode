package main

import (
	"fmt"
)

/*
383. isomorphic
https://leetcode.com/problems/isomorphic-strings
*/
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[byte]byte)
	tMap := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		sChar := s[i]
		tChar := t[i]

		if char, exist := sMap[sChar]; exist {
			if char != tChar {
				return false
			}

		} else {
			if _, exist := tMap[tChar]; exist {
				return false
			}
			sMap[sChar] = tChar
			tMap[tChar] = sChar
		}
	}

	return true
}

func main() {
	s := "badc"
	t := "baba"

	k := isIsomorphic(s, t)

	fmt.Print(k)
}
