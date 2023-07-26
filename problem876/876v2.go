package problem876

import . "github.com/mkuptsov/leetcode-in-go/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNodeV2(head *ListNode) *ListNode {
	last := head
	middle := head
	middleIndex := 1
	len := 1
	for last.Next != nil {
		len++
		last = last.Next
		if len/2 >= middleIndex {
			middle = middle.Next
			middleIndex++
		}
	}

	return middle
}
