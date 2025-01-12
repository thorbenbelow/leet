package leet

func canBeValid(s string, locked string) bool {
	if len(s)%2 != 0 {
		return false
	}

	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || locked[i] == '0' {
			n++
		} else {
			n--
		}

		if n < 0 {
			return false
		}
	}

	n = 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' || locked[i] == '0' {
			n++
		} else {
			n--
		}

		if n < 0 {
			return false
		}
	}

	return true
}
