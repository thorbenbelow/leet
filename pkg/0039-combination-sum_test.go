package leet_test

import (
	"reflect"
	"sort"
	"testing"
)

func sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func combinationSum(candidates []int, target int) [][]int {
	paths := make([][]int, len(candidates))
	res := make([][]int, 0)

	for i := range candidates {
		path := make([]int, 0)
		path = append(path, candidates[i])
		if candidates[i] == target {
			res = append(res, path)
		}
		paths[i] = path
	}
	for len(paths) > 0 {
		curr := paths[len(paths)-1]
		paths = paths[:len(paths)-1]
		s := sum(curr)
		for _, n := range candidates {

			if n > curr[len(curr)-1] {
				continue
			}

			switch {
			case s+n < target:
				newPath := append([]int(nil), curr...)
				newPath = append(newPath, n)
				paths = append(paths, newPath)
			case s+n > target:
				continue
			default:
				newPath := append([]int(nil), curr...)
				newPath = append(newPath, n)
				res = append(res, newPath)

			}
		}

	}

	return res
}

func isEqual(s1, s2 [][]int) bool {
	if len(s1) != len(s2) {
		return false
	}

	sortSlices(s1)
	sortSlices(s2)

	for i := 0; i < len(s1); i++ {
		if !reflect.DeepEqual(s1[i], s2[i]) {
			return false
		}
	}

	return true
}

func sortSlices(s [][]int) {
	for i := range s {
		sort.Ints(s[i])
	}
}

func Test0039Ex1(t *testing.T) {
	res := combinationSum([]int{2, 3, 6, 7}, 7)
	want := [][]int{{7}, {2, 2, 3}}
	if !isEqual(res, want) {
		t.Fatalf("Got: '%v'; Want: '%v'", res, want)
	}
}
func Test0039Ex2(t *testing.T) {
	res := combinationSum([]int{2, 3, 5}, 8)
	want := [][]int{{5, 3}, {3, 3, 2}, {2, 2, 2, 2}}
	if !isEqual(res, want) {
		t.Fatalf("Got: '%v'; Want: '%v'", res, want)
	}
}
