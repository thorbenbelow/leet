package leet

import (
	"testing"
)

func countSubarrays(nums []int, k int) int64 {
	// find number to solve for
	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}

	// use sliding window and only keep at most k elems of max in window
	res, numberOfMaxInCurrentWindow, l := int64(0), 0, 0
	for r := range nums {
		if nums[r] == max {
			numberOfMaxInCurrentWindow += 1
		}

		for numberOfMaxInCurrentWindow >= k {
			if nums[l] == max {
				numberOfMaxInCurrentWindow -= 1
			}
			l += 1
		}

		res += int64(l)
	}
	return res
}

func Test2962Ex1(t *testing.T) {
	res := countSubarrays([]int{1, 3, 2, 3, 3}, 2)
	want := int64(6)
	if res != want {
		t.Fatalf("Got: '%v'; Want: '%v'", res, want)
	}
}
