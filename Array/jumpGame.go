package main

import (
	"fmt"
)

/*
55. Jump Game
https://leetcode.com/problems/jump-game
*/
func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	final := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i]+i > final {
			final = nums[i] + i
		}

		if nums[i] == 0 && i == final && i != n-1 {
			return false
		}
	}

	return true
}

func main() {
	var nums = []int{2, 0, 0}
	k := canJump(nums)
	fmt.Print(k)
}
