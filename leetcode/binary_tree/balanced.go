/*
Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:

a binary tree in which the left and right subtrees of every node differ in height by no more than 1.
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftDepth := depth(root.Left)
	rightDepth := depth(root.Right)

	return (leftDepth == rightDepth || leftDepth+1 == rightDepth || leftDepth == rightDepth+1 || leftDepth-1 == rightDepth || leftDepth == rightDepth-1) && isBalanced(root.Left) && isBalanced(root.Right)

}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(depth(root.Right), depth(root.Left)) + 1

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
