package leet_test

import "testing"

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		switch {
		case nums[m] > target:
			r = m - 1
		case nums[m] < target:
			l = m + 1
		default:
			return m
		}
	}
	return l

}

func Test0035Ex1(t *testing.T) {
	res := searchInsert([]int{1, 3, 5, 6}, 5)
	want := 2
	if res != want {
		t.Fatalf("Got: '%v'; Want: '%v'", res, want)
	}
}

func Test0035Ex2(t *testing.T) {
	res := searchInsert([]int{1, 3, 5, 6}, 2)
	want := 1
	if res != want {
		t.Fatalf("Got: '%v'; Want: '%v'", res, want)
	}
}

func Test0035Ex3(t *testing.T) {
	res := searchInsert([]int{1, 3, 5, 6}, 7)
	want := 4
	if res != want {
		t.Fatalf("Got: '%v'; Want: '%v'", res, want)
	}
}
