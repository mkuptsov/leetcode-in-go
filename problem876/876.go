package problem876

import . "github.com/mkuptsov/leetcode-in-go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	nodes := make([]*ListNode, 0)
	index := 0
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
		index++
	}
	head = nodes[index/2]
	return head
}
