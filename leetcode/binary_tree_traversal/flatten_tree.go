/*
Given the root of a binary tree, flatten the tree into a "linked list":

The "linked list" should use the same TreeNode class where the right child pointer points to the next node in the list and the left child pointer is always null.
The "linked list" should be in the same order as a pre-order traversal of the binary tree.

Input: root = [1,2,5,3,4,null,6]
Output: [1,null,2,null,3,null,4,null,5,null,6]

Input: root = []
Output: []

Input: root = [0]
Output: [0]


*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	values := []*TreeNode{}
	preorderTraversal(root, &values)
	for idx := 0; idx < len(values)-1; idx += 1 {
		values[idx].Right = values[idx+1]
		values[idx].Left = nil
	}
}

func preorderTraversal(root *TreeNode, values *[]*TreeNode) {
	if root == nil {
		return
	}

	*values = append(*values, root)
	preorderTraversal(root.Left, values)
	preorderTraversal(root.Right, values)
}
