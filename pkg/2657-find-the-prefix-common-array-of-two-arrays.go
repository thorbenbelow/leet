package leet

func findThePrefixCommonArray(A []int, B []int) []int {
	seen := make([]bool, len(A))
	res := make([]int, len(A))
	last := 0
	for i := range A {
		if seen[A[i]-1] {
			last++
		} else {
			seen[A[i]-1] = true
		}

		if seen[B[i]-1] {
			last++
		} else {
			seen[B[i]-1] = true
		}
		res[i] = last
	}
	return res
}
