package group_anagrams

func GroupAnagrams(strs []string) [][]string {
	return groupAnagrams(strs)
}

func groupAnagrams(strs []string) [][]string {
	anagrams := [][]string{}
	primes := map[string]int{
		"a": 2,
		"b": 3,
		"c": 5,
		"d": 7,
		"e": 11,
		"f": 13,
		"g": 17,
		"h": 19,
		"i": 23,
		"j": 29,
		"k": 31,
		"l": 37,
		"m": 41,
		"n": 43,
		"o": 47,
		"p": 53,
		"q": 59,
		"r": 61,
		"s": 67,
		"t": 71,
		"u": 73,
		"v": 79,
		"w": 83,
		"x": 89,
		"y": 91,
		"z": 101,
	}

	ints := make([]int, len(strs))
	for i, str := range strs {
		product := 1
		for _, c := range str {
			product *= primes[string(c)]
		}

		ints[i] = product
	}

	intMap := make(map[int]int)
	for i, num := range ints {
		if index, ok := intMap[num]; ok {
			anagram := append(anagrams[index], strs[i])
			anagrams[index] = anagram
		} else {
			anagrams = append(anagrams, []string{strs[i]})
			intMap[num] = len(anagrams) - 1
		}
	}

	return anagrams
}
