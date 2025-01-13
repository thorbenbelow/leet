package leet

func minimumLength(s string) int {
	count := make([]int, 26)

	for _, c := range s {
		count[c-'a'] += 1
	}

	sum := 0
	for _, n := range count {
		if n <= 2 {
			sum += n
			continue
		}

		if n%2 == 0 {
			sum += 2
		} else {
			sum += 1
		}
	}
	return sum
}
