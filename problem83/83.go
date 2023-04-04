package problem83

import . "github.com/mkuptsov/leetcode-in-go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head != nil {
		node := head
		for node.Next != nil {
			if node.Val == node.Next.Val {
				node.Next = node.Next.Next
			} else {
				node = node.Next
			}
		}
	}
	return head
}
