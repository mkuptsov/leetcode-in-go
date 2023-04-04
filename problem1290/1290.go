package problem1290

import . "github.com/mkuptsov/leetcode-in-go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	decNum := head.Val
	for node := head; node.Next != nil; node = node.Next {
		decNum = decNum*2 + node.Next.Val
	}
	return decNum

}
