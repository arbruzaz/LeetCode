package main

import (
	"fmt"
)

/*
27. Remove Element
https://leetcode.com/problems/remove-element/description
*/
func removeElement(nums []int, val int) int {
	var j int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	fmt.Print(nums)
	return j
}

func main() {
	var nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	var val = 2

	k := removeElement(nums, val)

	fmt.Print(k)
}
