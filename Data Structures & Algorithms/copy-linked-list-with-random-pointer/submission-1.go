/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	memo := make(map[*Node]*Node)

	curr := head
	for curr != nil {
		memo[curr] = &Node{Val: curr.Val}
		curr = curr.Next
	}

	curr = head
	linkedCurr := memo[head]
	for linkedCurr != nil {
		linkedCurr.Next = memo[curr.Next]
		linkedCurr.Random = memo[curr.Random]
		linkedCurr = linkedCurr.Next
		curr = curr.Next
	}

	return memo[head]
}
