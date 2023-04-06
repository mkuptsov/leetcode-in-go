package problem404

import . "github.com/mkuptsov/leetcode-in-go/utils"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	if root.Left == nil && root.Right == nil {
		return 0
	}

	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			sum += root.Left.Val
		} else {
			sum += sumOfLeftLeaves(root.Left)
		}
	}

	if root.Right != nil {
		sum += sumOfLeftLeaves(root.Right)
	}

	return sum
}
