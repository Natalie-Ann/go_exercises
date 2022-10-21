/*
Given the head of a singly linked list, reverse the list, and return the reversed list.

Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

Input: head = [1,2]
Output: [2,1]


Algorithm:
- Guard clause: if head == nil or head.Next == nil, return head (in case of empty or one node list)
- Set current node equal to head node
- set current node's next node equal to dummy node
- While head node != nil
	- set current node's next node equal to current node
	- set current node equal to
- Return head
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Traverse() {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	current := head
	next := head.Next
	var previous *ListNode

	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}

	previous.Traverse()

	return previous
}

func main() {
	a := ListNode{1, new(ListNode)}
	b := ListNode{1, new(ListNode)}
	c := ListNode{2, new(ListNode)}
	d := ListNode{3, new(ListNode)}
	e := ListNode{4, new(ListNode)}

	a.Next = &b
	b.Next = &c
	c.Next = &d
	d.Next = &e
	e.Next = nil

	head := a

	reverseList(&head)
}
