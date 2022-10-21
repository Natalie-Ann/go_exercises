// Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prevNode := head
	nextNode := head.Next

	for nextNode != nil {
		if prevNode.Val == nextNode.Val {
			nextNode = nextNode.Next

			if nextNode == nil {
				prevNode.Next = nil
			}
		} else {
			prevNode.Next = nextNode
			prevNode = nextNode
			nextNode = prevNode.Next
		}
	}

	return head
}
