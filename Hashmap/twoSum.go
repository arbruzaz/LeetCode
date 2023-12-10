package main

import (
	"fmt"
)

/*
1. Two Sum
https://leetcode.com/problems/two-sum
*/
func twoSum(nums []int, target int) []int {
	var mapResult = make(map[int]int)
	for i, num := range nums {
		if j, ok := mapResult[num]; ok {
			return []int{j, i}
		}
		mapResult[target-num] = i
	}
	return []int{}
}

func main() {
	nums := []int{2, 15, 7, 11, 15}
	target := 9

	k := twoSum(nums, target)

	fmt.Print(k)
}
