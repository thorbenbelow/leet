package leet

import (
	"strings"
	"testing"
)

func countPrefixSuffixPairs(words []string) int {
	var n int
	for i := range words {
		fix := words[i]

		for j := i + 1; j < len(words); j++ {
			if j == i {
				continue
			}
			word := words[j]
			if strings.HasPrefix(word, fix) && strings.HasSuffix(word, fix) {
				n++
			}
		}
	}

	return n
}

func TestCountPrefixSuffixPairs1(t *testing.T) {
	if countPrefixSuffixPairs([]string{"a", "aba", "ababa", "aa"}) != 4 {
		t.Fail()
	}
}
func TestCountPrefixSuffixPairs2(t *testing.T) {
	if countPrefixSuffixPairs([]string{"pa", "papa", "ma", "mama"}) != 2 {
		t.Fail()
	}
}

func TestCountPrefixSuffixPairs3(t *testing.T) {
	if countPrefixSuffixPairs([]string{"ababa", "aba"}) != 0 {
		t.Fail()
	}
}
