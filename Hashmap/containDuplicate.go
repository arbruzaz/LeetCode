package main

import (
	"fmt"
)

/*
219. Contains Duplicate II
https://leetcode.com/problems/contains-duplicate-ii/
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	foundNum := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		_, exist := foundNum[nums[i]]
		if !exist {
			foundNum[nums[i]] = i
			continue
		}

		if i-foundNum[nums[i]] <= k {
			return true
		}

		foundNum[nums[i]] = i
	}
	return false
}

func main() {
	nums := []int{1, 0, 1, 1}
	k := 1

	r := containsNearbyDuplicate(nums, k)

	fmt.Print(r)
}
