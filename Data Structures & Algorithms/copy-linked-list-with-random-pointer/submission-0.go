/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

//  3->7->4->5->N
//D  

type RandomMapValue struct {
	ogTarget *Node
	sister *Node
}

func copyRandomList(head *Node) *Node {
    randomMap := make(map[*Node]RandomMapValue)
	
	dummy := &Node{}
	curr := dummy
	
	headCopy := head

	for head != nil {
		node := &Node{Val: head.Val}
		curr.Next = node

		randomMap[head] = RandomMapValue{ogTarget: head.Random, sister: node}

		head = head.Next
		curr = curr.Next
	}

	head = headCopy
	for head != nil {
		data := randomMap[head]
		data.sister.Random = randomMap[data.ogTarget].sister
		head = head.Next
	}

	return dummy.Next
}
