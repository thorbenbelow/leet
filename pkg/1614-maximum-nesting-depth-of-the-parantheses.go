package leet

func maxDepth(s string) int {
	max := 0

	curr := 0
	for _, c := range s {
		if c == '(' {
			curr += 1
			if curr > max {
				max = curr
			}
		} else if c == ')' {
			curr -= 1
		}

	}

	return max
}
