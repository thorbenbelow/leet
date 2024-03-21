package leet

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := head
	cursor := head.Next
	prev.Next = nil

	for cursor.Next != nil {
		temp := cursor.Next
		cursor.Next = prev
		prev = cursor
		cursor = temp
	}

	cursor.Next = prev

	return cursor
}
