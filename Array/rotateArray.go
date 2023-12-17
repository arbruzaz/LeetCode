package main

import (
	"fmt"
)

/*
189. Rotate Array
https://leetcode.com/problems/rotate-array
*/
func rotateV1(nums []int, k int) {
	if len(nums) > 1 {
		currentIndex := 0
		realCurrent := 0
		for i := 0; i < len(nums); i++ {
			desireIndex := currentIndex + k
			for desireIndex >= len(nums) {
				desireIndex = desireIndex - len(nums)
			}

			nums[realCurrent], nums[desireIndex] = nums[desireIndex], nums[realCurrent]
			if realCurrent == desireIndex {
				realCurrent++
				currentIndex = desireIndex + 1
			} else {
				currentIndex = desireIndex
			}

		}
	}
}

func rotateV2(nums []int, k int) {
	n := len(nums)
	k %= n

	if k == 0 {
		return
	}

	reverse(&nums, 0, n-1)
	reverse(&nums, k, n-1)
	reverse(&nums, 0, k-1)

	fmt.Println(nums)
}

func reverse(nums *[]int, i, j int) {
	for i < j {
		(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
		i++
		j--
	}
}

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7}
	k := 2
	rotate(nums, k)

}
