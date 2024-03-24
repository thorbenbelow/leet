package leet

func findDuplicate(nums []int) int {
	n := nums[0]
	for {
		if nums[n] == -1 {
			return n
		}
		temp := nums[n]
		nums[n] = -1
		n = temp
	}
}
