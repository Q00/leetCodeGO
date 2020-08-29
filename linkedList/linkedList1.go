package linkedlist

// ListNode daf
/*
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	for node.Next != nil && node.Next.Next != nil {
		node.Val = node.Next.Val
		node = node.Next
	}
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
