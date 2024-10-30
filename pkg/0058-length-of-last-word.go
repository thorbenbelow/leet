package leet

import "strings"

func lengthOfLastWord(s string) int {
	t := strings.TrimSpace(s)
	i := len(t) - 1

	for ; i >= 0 && t[i] != ' '; i -= 1 {
	}

	return len(t) - i - 1
}
