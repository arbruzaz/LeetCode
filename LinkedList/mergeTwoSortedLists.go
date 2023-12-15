package main

import (
	"fmt"
)

/*
141. Linked List Cycle
https://leetcode.com/problems/linked-list-cycle
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}
	var newHead = list1

	if list1.Val > list2.Val {
		newHead = list2
		list2 = list2.Next
	} else {
		newHead = list1
		list1 = list1.Next
	}
	newLinkedList := newHead
	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			newHead.Next = list2
			newHead = newHead.Next
			list2 = list2.Next
		} else {
			newHead.Next = list1
			newHead = newHead.Next
			list1 = list1.Next
		}
	}

	for list1 != nil {
		newHead.Next = list1
		newHead = newHead.Next
		list1 = list1.Next
	}

	for list2 != nil {
		newHead.Next = list2
		newHead = newHead.Next
		list2 = list2.Next
	}

	return newLinkedList
}

func main() {

	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 4}

	node1.Next = node2
	node2.Next = node3
	node3.Next = nil

	node4 := &ListNode{Val: 1}
	node5 := &ListNode{Val: 3}
	node6 := &ListNode{Val: 4}

	node4.Next = node5
	node5.Next = node6
	node6.Next = nil

	k := mergeTwoLists(node1, node4)

	fmt.Print(k)
}
