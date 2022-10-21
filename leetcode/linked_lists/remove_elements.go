// Given the head of a linked list and an integer val, remove all the nodes of the linked list that has Node.val == val, and return the new head.

// Input: head = [1,2,6,3,4,5,6], val = 6
// Output: [1,2,3,4,5]
// Example 2:

// Input: head = [], val = 1
// Output: []
// Example 3:

// Input: head = [7,7,7,7], val = 7
// Output: []

/*
Algorithm:
- Guard clause: if head == nil
	- return head
- Set prevNode to nil
- Current node = head
- nextNode = head.Next
- While nextNode != nil
	- if currentNode.Val = val
		- if next node is nil
			- prev.Next = nil
		- prevNode.Next = nextNode
		- current = next
		- next = current.Next

	- proceed by 1
			- prevNode = currnetNode
			- currentNode = nextNode
			- nextNode = nextNode.Next
- Return head

*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	currentNode := head
	nextNode := head.Next

	for nextNode != nil {
		if currentNode.Val == val {
			previousNode.Next = nextNode
			currentNode = nextNode
			nextNode = currentNode.Next

			if nextNode == nil {
				previousNode.Next = nil
			}
		}

		previousNode = currentNode
		currentNode = nextNode
		nextNode = nextNode.Next
	}

	return head
}

func main() {

}
