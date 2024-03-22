package leet

import (
	"testing"
)

func isPalindrome(head *ListNode) bool {
	l, r := head, head
	// find mid with slow and fast pointer
	for r != nil && r.Next != nil {
		l = l.Next
		r = r.Next.Next
	}

	// reverse second halth of linked list
	prev := l
	l = l.Next
	prev.Next = nil
	for l != nil {
		temp := l.Next
		l.Next = prev
		prev = l
		l = temp
	}
	r = head
	l = prev

	// compare from start and (reversed) end
	for l != nil {
		if l.Val != r.Val {
			return false
		}
		l = l.Next
		r = r.Next

	}
	return true
}

func isPalindromeNaive(head *ListNode) bool {
	size := 0
	cursor := head
	for cursor != nil {
		cursor = cursor.Next
		size += 1
	}

	cursor = head
	values := make([]int, size/2)

	for i := 0; i < size/2; i++ {
		values[i] = cursor.Val
		cursor = cursor.Next
	}

	if size%2 != 0 {
		cursor = cursor.Next

	}

	for i := 1; i <= size/2; i++ {
		if values[size/2-i] != cursor.Val {
			return false
		}

		cursor = cursor.Next
	}
	return true
}

func Test0234Ex1Naive(t *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}
	if !isPalindromeNaive(head) {
		t.Fatal("Should be true")
	}
}
func Test0234Ex2Naive(t *testing.T) {
	head := &ListNode{1, &ListNode{0, &ListNode{1, nil}}}
	if !isPalindromeNaive(head) {
		t.Fatal("Should be true")
	}
}

func Test0234Ex3Naive(t *testing.T) {
	head := &ListNode{1, &ListNode{0, &ListNode{0, nil}}}
	if isPalindromeNaive(head) {
		t.Fatal("Should be false")
	}
}
func Test0234Ex1(t *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}
	if !isPalindrome(head) {
		t.Fatal("Should be true")
	}
}
func Test0234Ex2(t *testing.T) {
	head := &ListNode{1, &ListNode{0, &ListNode{1, nil}}}
	if !isPalindrome(head) {
		t.Fatal("Should be true")
	}
}

func Test0234Ex3(t *testing.T) {
	head := &ListNode{1, &ListNode{0, &ListNode{0, nil}}}
	if isPalindrome(head) {
		t.Fatal("Should be false")
	}
}
