package valid_anagram

func IsAnagram(s, t string) bool {
	return isAnagram(s, t)
}

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var letters [26]int

	for i := range s {
		letters[s[i]-'a']++
		letters[t[i]-'a']--
	}

	for _, letter := range letters {
		if letter != 0 {
			return false
		}
	}

	return true
}
