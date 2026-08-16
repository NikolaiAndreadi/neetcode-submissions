/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */  

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

    // find middle
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	
	// decouple parts
	second := slow.Next
	slow.Next = nil

	// reverse second part
	var prev *ListNode
	for second != nil {
		tmp := second.Next
		second.Next = prev
		prev = second
		second = tmp
	}

	// merge
	first := head
	second = prev
	for second != nil {
		tmp1 := first.Next
		tmp2 := second.Next
		 
		first.Next = second
		second.Next = tmp1

		first = tmp1
		second = tmp2
	}
}
