package minimum_window_substring

func MinWindow(s, t string) string {
	return minWindow(s, t)
}

func minWindow(s string, t string) string {
	expectedCharmap := make(map[byte]int, 52)

	for _, char := range t {
		expectedCharmap[byte(char)]++
	}

	counter := len(t)
	begin := 0
	head := 0
	end := 0
	minWinLen := len(s) + 1

	for end < len(s) {
		if expectedCharmap[s[end]] > 0 {
			counter--
		}
		expectedCharmap[s[end]]--
		end++

		for counter == 0 {
			if end-begin < minWinLen {
				head = begin
				minWinLen = end - begin
			}

			expectedCharmap[s[begin]]++
			if expectedCharmap[s[begin]] > 0 {
				counter++
			}
			begin++
		}
	}

	if minWinLen == len(s)+1 {
		return ""
	}

	return s[head : head+minWinLen]
}
