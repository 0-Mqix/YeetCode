package link

import "testing"

func TestReverseLinkedList(t *testing.T) {
	output := reverseList(&ListNode{Val: 1, Next: nil})
	t.Log(output)
}
