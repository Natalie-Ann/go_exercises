/* Given the head of a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list. Return the linked list sorted as well.

Input: head = [1,2,3,3,4,4,5]
Output: [1,2,5]

Input: head = [1,1,1,2,3]
Output: [2,3]


Algorithm:
- Guard clause: If head == nil or head.Next == nil return head
- Define previous node as nil node
- Define current node as head
- Define next node as head.Next

- While next node does not equal nil
	- If current node and next node values are equal
		- shift window right until reach a non-duplicate node
			- findNextNonDuplicate --> node; reassign the returned node to currentNode
			- reassign previous node.Next to new current node (the nonduplicate)
			- next node reassigned to current node.Next
	- Else
		- shift window by 1
			- previous --> current
			- current --> next
			- next --> next.Next
- Return head


Helper function findNextNonDuplicate
- currentNode is the parameter passed in that is a duplicate
- save duplicate value to a variable
- Repeatedly reassign currentnode to currentNode.Next get a new value
- return the new currentnode
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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// var previousNode *ListNode
	dummy := new(ListNode)
	previousNode := dummy
	previousNode.Next = head

	currentNode := head
	nextNode := head.Next

	for nextNode != nil {
		if currentNode.Val == nextNode.Val {
			currentNode = findNextNonDuplicate(currentNode)
			previousNode.Next = currentNode

			if currentNode != nil {
				nextNode = currentNode.Next
			}
		} else {
			previousNode = currentNode
			currentNode = nextNode
			nextNode = nextNode.Next
		}
	}

	dummy.Next.Traverse()

	return dummy.Next

}

func findNextNonDuplicate(currentNode *ListNode) *ListNode {
	duplicateValue := currentNode.Val

	for currentNode.Val == duplicateValue && currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	return currentNode
}

func main() {
	a := ListNode{1, new(ListNode)}
	b := ListNode{1, new(ListNode)}
	c := ListNode{2, new(ListNode)}
	d := ListNode{3, new(ListNode)}
	e := ListNode{3, new(ListNode)}

	a.Next = &b
	b.Next = &c
	c.Next = &d
	d.Next = &e
	e.Next = nil

	head := a

	// a2 := ListNode{1, new(ListNode)}
	// b2 := ListNode{3, new(ListNode)}
	// c2 := ListNode{4, new(ListNode)}

	// a2.Next = &b2
	// b2.Next = &c2
	// c2.Next = nil

	// example2 := a2

	// deleteDuplicates(&example2)
	// findNextNonDuplicate(&head)
	deleteDuplicates(&head)
}
