package leet

import (
	"reflect"
	"testing"
)

func reorderList(head *ListNode) {
	l, r := head, head
	// find mid with fast and slow pointer
	for r != nil && r.Next != nil {
		l = l.Next
		r = r.Next.Next
	}

	// reverse second half
	var prev *ListNode
	for l != nil {
		temp := l.Next
		l.Next = prev
		prev = l
		l = temp
	}

	l = head
	r = prev

	// zip first and second half
	for l != nil {
		temp := l.Next
		l.Next = r
		prev = l
		l = r
		r = temp
	}

	if prev != r.Next {
		prev.Next = r
		prev.Next.Next = nil
	}
}

func Test0143Ex1(t *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	reorderList(head)
	want := []int{1, 4, 2, 3}
	got := listToArr(head)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Got: %v ,but want: %v", got, want)
	}

}
func Test0143Ex2(t *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	reorderList(head)
	want := []int{1, 5, 2, 4, 3}
	got := listToArr(head)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Got: %v ,but want: %v", got, want)
	}

}
func Test0143Ex3(t *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}}}}
	reorderList(head)
	want := []int{1, 7, 2, 6, 3, 5, 4}
	got := listToArr(head)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Got: %v ,but want: %v", got, want)
	}

}

func listToArr(head *ListNode) []int {
	values := make([]int, 0)
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	return values
}
