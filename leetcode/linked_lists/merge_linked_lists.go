/*
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

Input: list1 = [], list2 = []
Output: []

Input: list1 = [], list2 = [0]
Output: [0]


Algorithm:
- Create a dummy node with nil value
- Assign current node to dummy node
- Compare the head nodes in list1 and list2 while both != nil
	- L1 value < L2 value
		- assign current.next to list1
		- assign list1 to next node in list1
	- Else
		- assign current.next to list2
		- assign list2 to next node in list2
	- Reassign current to current.next

- If list1 == nil
	- set next node equal to list2
- If list2 == nil
	- set next node equal to list1


- Return dummyNode.Next (the head node that was first assigned)
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := new(ListNode)
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 == nil {
		current.Next = list2
	} else {
		current.Next = list1
	}

	return dummy.Next
}

func main() {
	a := ListNode{1, new(ListNode)}
	b := ListNode{2, new(ListNode)}
	c := ListNode{4, new(ListNode)}

	a.Next = &b
	b.Next = &c
	c.Next = nil

	example1 := a

	a2 := ListNode{1, new(ListNode)}
	b2 := ListNode{3, new(ListNode)}
	c2 := ListNode{4, new(ListNode)}

	a2.Next = &b2
	b2.Next = &c2
	c2.Next = nil

	example2 := a2

	fmt.Println(mergeTwoLists(&example1, &example2))

}
