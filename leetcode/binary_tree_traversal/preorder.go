/*
Given the root of a binary tree, return the preorder traversal of its nodes' values.
Preorder = RLR

Input: root = [1,null,2,3]
Output: [1,2,3]

Input: root = []
Output: []

Input: root = [1]
Output: [1]

*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	preorderedNodes := []int{}

	if root == nil {
		return nil
	}

	left := preorderTraversal(root.Left)
	right := preorderTraversal(root.Right)

	preorderedNodes = append(preorderedNodes, root.Val)
	preorderedNodes = append(preorderedNodes, left...)
	preorderedNodes = append(preorderedNodes, right...)

	return preorderedNodes
}
