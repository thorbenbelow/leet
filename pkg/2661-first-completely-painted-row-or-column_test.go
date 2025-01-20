package leet

import (
	"testing"
)

func firstCompleteIndex(arr []int, mat [][]int) int {
	rows := make([]int, len(mat))
	cols := make([]int, len(mat[0]))

	m := make(map[int][2]int)

	for i := range mat {
		for j := range mat[0] {
			m[mat[i][j]] = [2]int{i, j}
		}
	}

	for res, n := range arr {
		i := m[n][0]
		j := m[n][1]

		rows[i]++
		cols[j]++

		if rows[i] == len(mat[0]) {
			return res
		}
		if cols[j] == len(mat) {
			return res
		}
	}
	return -1
}

func TestFirstCompleteIndexExample1(t *testing.T) {
	arr := []int{1, 4, 5, 2, 6, 3}
	mat := [][]int{{4, 3, 5}, {1, 2, 6}}

	res := firstCompleteIndex(arr, mat)

	if res != 1 {
		t.Errorf("expected 1, but got %d", res)
	}
}
