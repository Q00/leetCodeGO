package linkedlist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	head2 := &ListNode{Next: head}
	// front -> length -n
	front := head2
	// back -> n -1
	back := head2

	back = back.Next
	size := 0
	for head.Next != nil {
		head = head.Next
		size++
	}

	for ; size >= n; size-- {
		front = front.Next
		back = back.Next
	}

	front.Next = back.Next

	return head2.Next
}
