package main

import (
	"fmt"
)

/*
55. Jump Game
https://leetcode.com/problems/jump-game
*/
func jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	max := 0
	count := 0
	currentJump := 0
	for i := 0; i < n-1; i++ {
		if nums[i]+i >= max {
			max = nums[i] + i
		}
		if i == currentJump {
			count++
			currentJump = max
		}
	}

	return count
}

func main() {
	var nums = []int{1, 2}
	k := jump(nums)
	fmt.Print(k)
}
