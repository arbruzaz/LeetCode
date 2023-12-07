package main

import (
	"fmt"
)

/*
169 Majority Element
https://leetcode.com/problems/majority-element
*/
func majorityElement(nums []int) int {
	lenNums := len(nums)
	condition := lenNums / 2
	major := -1
	count := make(map[int]int)
	for i := 0; i < lenNums; i++ {
		count[nums[i]] += 1
		if count[nums[i]] > condition {
			major = nums[i]
		}
	}
	return major
}

func main() {
	var nums = []int{2, 2, 1, 1, 1, 2, 2}

	k := majorityElement(nums)

	fmt.Print(k)
}
