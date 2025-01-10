package leet

import "testing"

func isUniversal(word string, maxFreq []int) bool {
	freq := make([]int, 26)

	for _, c := range word {
		freq[c-'a']++
	}
	for i := 0; i < 26; i++ {
		if freq[i] < maxFreq[i] {
			return false
		}

	}
	return true
}

func wordSubsets(words1 []string, words2 []string) []string {

	res := make([]string, 0)

	maxFreq := make([]int, 26)

	for _, word := range words2 {
		freq := make([]int, 26)

		for _, c := range word {
			freq[c-'a']++
		}

		for i := 0; i < 26; i++ {
			maxFreq[i] = max(maxFreq[i], freq[i])
		}
	}

	for _, word := range words1 {
		if isUniversal(word, maxFreq) {
			res = append(res, word)
		}
	}

	return res

}

func TestExample1(t *testing.T) {
	words1 := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	words2 := []string{"e", "o"}

	result := wordSubsets(words1, words2)
	if len(result) != 3 {
		t.Fatalf("Expected 3, got %d: %v", len(result), result)
	}
}
func TestExample2(t *testing.T) {
	words1 := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	words2 := []string{"l", "e"}

	result := wordSubsets(words1, words2)
	if len(result) != 3 {
		t.Fatalf("Expected 3, got %d: %v", len(result), result)
	}
}
