package longest_substring_without_repeating_characters

func LengthOfLongestSubstring(s string) int {
	return lengthOfLongestSubstring(s)
}

func lengthOfLongestSubstring(s string) int {
	charmap := make(map[byte]int)

	maxLen := 0

	i := 0
	j := 0
	numOfDups := 0

	for j < len(s) {
		charmap[s[j]]++

		if charmap[s[j]] <= 1 {
			if j-i+1 >= maxLen {
				maxLen = j - i + 1
			}
		} else {
			numOfDups++
		}

		for numOfDups > 0 {
			if charmap[s[i]] > 1 {
				numOfDups--
			}

			charmap[s[i]]--
			i++
		}

		j++
	}

	return maxLen
}
