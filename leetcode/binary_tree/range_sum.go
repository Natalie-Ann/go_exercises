/*
Given the root node of a binary search tree and two integers low and high, return the sum of values of all nodes with a value in the inclusive range [low, high].

Input: root = [10,5,15,3,7,null,18], low = 7, high = 15
Output: 32
Explanation: Nodes 7, 10, and 15 are in the range [7, 15]. 7 + 10 + 15 = 32.

Input: root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
Output: 23
Explanation: Nodes 6, 7, and 10 are in the range [6, 10]. 6 + 7 + 10 = 23.
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	var sum int

	if root.Val <= high && root.Val >= low {
		sum += root.Val
	}

	if root.Val >= low {
		//check only left side
		sum += rangeSumBST(root.Left, low, high)
	}
	if root.Val <= high {
		//check only right side
		sum += rangeSumBST(root.Right, low, high)
	}

	return sum

}
