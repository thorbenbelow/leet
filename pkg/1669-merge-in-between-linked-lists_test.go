package leet

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	cursorA := list1

	for i := 0; i < a-1; i += 1 {
		cursorA = cursorA.Next
	}

	cursorB := cursorA.Next

	for i := a; i < b+1; i += 1 {
		cursorB = cursorB.Next
	}

	cursorA.Next = list2

	for cursorA.Next != nil {
		cursorA = cursorA.Next
	}

	cursorA.Next = cursorB

	return list1
}
