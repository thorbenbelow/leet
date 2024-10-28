package leet

import (
	"math"
	"sort"
)

func longestSquareStreak(nums []int) int {
	sort.Ints(nums)
	res := -1
	m := make(map[int]int)

	for _, n := range nums {
		sqrt := int(math.Sqrt(float64(n)))

		count, ok := m[sqrt]
		if sqrt*sqrt == n && ok {
			m[n] = count + 1
			if m[n] > res {
				res = m[n]
			}
		} else {
			m[n] = 1
		}
	}
	return res
}
