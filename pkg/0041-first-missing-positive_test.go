package leet

import (
	"testing"
)

func firstMissingPositive(nums []int) int {
	// The first missing positive has to be within len(nums) + 1
	// so we will sort the values within this range by assigning for each n the index n-1

	for i := range nums {
		// we only care about numbers that can be within this range
		// and we only have to swap if n-1 does not already have the value n
		for nums[i] > 0 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}

	}
	// now that all values are sorted we just need to find the first value where n != i+1
	for i, n := range nums {
		if n != i+1 {
			return i + 1
		}
	}

	// or if the array only had its own indices as elements the first missing will be len + 1
	return len(nums) + 1
}

func Test0042Ex1(t *testing.T) {
	got := firstMissingPositive([]int{1, 1})
	want := 2
	if got != want {
		t.Fatalf("Want %v, but got %v", want, got)
	}
}
func Test0042Ex2(t *testing.T) {
	got := firstMissingPositive([]int{2, 1})
	want := 3
	if got != want {
		t.Fatalf("Want %v, but got %v", want, got)
	}
}
func Test0042Ex3(t *testing.T) {
	got := firstMissingPositive([]int{0, 1, 2})
	want := 3
	if got != want {
		t.Fatalf("Want %v, but got %v", want, got)
	}
}
