/*
Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.

Input: root = [2,1,3]
Output: true

Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.

Go right in a BST, the min changes to the parent root
Go left in a BST, the max changes to the parent root

*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(root, -9999999999, 9999999999)
}

func helper(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}

	left, right := root.Left, root.Right

	return root.Val < max && root.Val > min && helper(left, min, root.Val) && helper(right, root.Val, max)

}
