package problem257

import (
	"strconv"
	"strings"

	. "github.com/mkuptsov/leetcode-in-go/utils"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	path := []string{}
	result := []string{}

	var backtrack func(*TreeNode)
	backtrack = func(node *TreeNode) {
		path = append(path, strconv.Itoa(node.Val))

		if node.Left == nil && node.Right == nil {
			result = append(result, strings.Join(path, "->"))
		}
		if node.Left != nil {
			backtrack(node.Left)
			path = path[:len(path)-1]
		}
		if node.Right != nil {
			backtrack(node.Right)
			path = path[:len(path)-1]
		}
	}

	backtrack(root)
	return result
}
