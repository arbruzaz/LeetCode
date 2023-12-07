package main

import "fmt"

/*
88. Merge Sorted Array
https://leetcode.com/problems/merge-sorted-array/description/
*/

// 2nd Solution
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
	fmt.Println(nums1)
}

// 1st Solution

// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	lastIndex := m - 1
// 	current2 := 0

// 	for i := 0; i < m+n; i++ {
// 		if lastIndex < 0 || i > lastIndex {
// 			nums1[i] = nums2[current2]
// 			current2 += 1
// 		} else if current2 < n && nums1[i] > nums2[current2] {
// 			tmpNums := make([]int, m+n)
// 			copy(tmpNums[:i], nums1[:i])
// 			tmpNums[i] = nums2[current2]
// 			copy(tmpNums[i+1:], nums1[i:])
// 			copy(nums1, tmpNums)

// 			current2 += 1
// 			lastIndex += 1
// 		}

// 	}
// 	fmt.Println(nums1)
// }

func main() {

	var num1 = []int{-1, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0}
	var m = 5

	var num2 = []int{-1, -1, 0, 0, 1, 2}
	var n = 6

	merge(num1, m, num2, n)
}
