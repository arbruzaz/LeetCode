package main

import (
	"fmt"
)

/*
26. Remove Duplicate
https://leetcode.com/problems/remove-duplicates-from-sorted-array/description
*/
func removeDeplicate(nums []int) int {
	j := 1

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			nums[j] = nums[i+1]
			j++
		}
	}
	fmt.Println(nums)
	return j
}

func main() {
	var nums = []int{0}

	k := removeDeplicate(nums)

	fmt.Print(k)
}
