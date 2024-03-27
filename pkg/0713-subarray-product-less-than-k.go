package leet

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k == 0 {
		return 0
	}

	res := 0

	for l := range nums {
		prod := 1
		for r := l; r < len(nums); r += 1 {
			prod *= nums[r]
			if prod >= k {
				break
			}
			res += 1
		}
	}

	return res
}
