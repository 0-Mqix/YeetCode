package link

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(node *ListNode) *ListNode {
	var prev *ListNode
	var temp *ListNode
	head := node

	for head != nil {
		temp = head
		head = head.Next
		temp.Next = prev
		prev = temp
	}

	return temp
}
