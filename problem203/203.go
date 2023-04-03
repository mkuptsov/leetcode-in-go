package problem203

import "github.com/mkuptsov/leetcode-in-go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *utils.ListNode, val int) *utils.ListNode {
	newHead := new(utils.ListNode)
	newNode := newHead
	for node := head; node != nil; node = node.Next {
		if node.Val != val {
			newNode.Next = node
			newNode = node
		} else {
			newNode.Next = node.Next
		}
	}
	return newHead.Next
}
