package leet

import (
	"testing"
)

func maxSubarrayLength(nums []int, k int) int {
	l, r := 0, 0
	max := 1
	nmap := make(map[int]int)

	for l < len(nums) {
		nmap[nums[l]] += 1
		for nmap[nums[l]] > k {
			nmap[nums[r]] -= 1
			r += 1
		}
		if l-r+1 > max {
			max = l - r + 1
		}
		l++
	}
	return max
}

func Test2958Ex1(t *testing.T) {
	res := maxSubarrayLength([]int{1, 11}, 2)
	want := 2
	if res != want {
		t.Fatalf("Got: '%v'; Want: '%v'", res, want)
	}
}
func Test2958Ex2(t *testing.T) {
	res := maxSubarrayLength([]int{1, 2, 3, 1, 2, 3, 1, 2}, 2)
	want := 6
	if res != want {
		t.Fatalf("Got: '%v'; Want: '%v'", res, want)
	}
}
func Test2958Ex3(t *testing.T) {
	res := maxSubarrayLength([]int{1}, 2)
	want := 1
	if res != want {
		t.Fatalf("Got: '%v'; Want: '%v'", res, want)
	}
}
func Test2958Ex4(t *testing.T) {
	res := maxSubarrayLength([]int{2, 2, 3}, 1)
	want := 2
	if res != want {
		t.Fatalf("Got: '%v'; Want: '%v'", res, want)
	}
}
