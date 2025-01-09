package leet

import (
	"testing"
)

type Cursor struct {
	x     int
	y     int
	count int
}

func maxMoves(grid [][]int) int {
	res := 0

	stack := make([]Cursor, 0)

	for i := range grid {
		stack = append(stack, Cursor{x: 0, y: i, count: 0})
	}

	for len(stack) > 0 {
		var next Cursor
		next, stack = stack[len(stack)-1], stack[:len(stack)-1]

		// y - 1, x + 1
		if next.y-1 >= 0 && next.x+1 < len(grid[next.y-1]) {
			if grid[next.y-1][next.x+1] > grid[next.y][next.x] {
				stack = append(stack, Cursor{x: next.x + 1, y: next.y - 1, count: next.count + 1})
				if next.count+1 > res {
					res = next.count + 1
				}
			}
		}

		// y, x + 1
		if next.x+1 < len(grid[next.y]) {
			if grid[next.y][next.x+1] > grid[next.y][next.x] {
				stack = append(stack, Cursor{x: next.x + 1, y: next.y, count: next.count + 1})
				if next.count+1 > res {
					res = next.count + 1
				}
			}
		}

		// y + 1, x + 1
		if next.y+1 < len(grid) && next.x+1 < len(grid[next.y]) {
			if grid[next.y+1][next.x+1] > grid[next.y][next.x] {
				stack = append(stack, Cursor{x: next.x + 1, y: next.y + 1, count: next.count + 1})
				if next.count+1 > res {
					res = next.count + 1
				}
			}
		}

	}

	return res
}

func TestMaxMoves(t *testing.T) {
	if maxMoves([][]int{{2, 4, 3, 5}, {5, 4, 9, 3}, {3, 4, 2, 11}, {10, 9, 13, 15}}) != 3 {
		t.Fail()
	}

	if maxMoves([][]int{{3, 2, 4}, {2, 1, 9}, {1, 1, 7}}) != 0 {
		t.Fail()
	}
}
