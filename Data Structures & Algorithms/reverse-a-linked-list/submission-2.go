/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	return reverseRec(head, nil)
}

func reverseRec(curr *ListNode, prev *ListNode) *ListNode {
	if curr == nil {
		return prev
	}
	tmp := curr.Next
	curr.Next = prev
	return reverseRec(tmp, curr)
}
