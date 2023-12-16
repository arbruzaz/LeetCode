package main

import (
	"fmt"
)

/*
80. Remove Duplicates From Sorted Array II
https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
*/
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	curIndex := 2

	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[curIndex-2] {
			nums[curIndex] = nums[i]
			curIndex++
		}
	}

	return curIndex
}

func main() {
	var nums = []int{0, 0, 1, 1, 1, 1, 2, 2, 2, 4}

	k := removeDuplicates(nums)

	fmt.Print(k)
}
