package leet

import "testing"

func canConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}
	chars := make([]rune, 26)
	for _, c := range s {
		chars[c-'a']++
	}

	var count int
	for _, n := range chars {
		if n%2 != 0 {
			count++
		}
	}
	return count <= k
}

func TestCanConstructExample1(t *testing.T) {
	s := "annabelle"
	k := 2
	result := canConstruct(s, k)
	if !result {
		t.Fatalf("Expected true, got %v", result)
	}
}

func TestCanConstructExample2(t *testing.T) {
	s := "leetcode"
	k := 3
	result := canConstruct(s, k)
	if result {
		t.Fatalf("Expected false, got %v", result)
	}
}

func TestCanConstructExample3(t *testing.T) {
	s := "true"
	k := 4
	result := canConstruct(s, k)
	if !result {
		t.Fatalf("Expected true, got %v", result)
	}
}
