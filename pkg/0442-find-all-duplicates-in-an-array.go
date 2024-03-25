package leet

import "math"

func findDuplicates(nums []int) []int {
	duplicates := make([]int, 0)
	for i := range nums {
		n := int(math.Abs(float64(nums[i])))
		if nums[n-1] > 0 {
			nums[n-1] = -1 * nums[n-1]
		} else {
			duplicates = append(duplicates, n)
		}

	}
	return duplicates
}

