package palindromic_substrings

func CountSubstrings(s string) int {
	return countSubstrings(s)
}

func countSubstrings(s string) int {
	odd := 0
	even := 0
	for k := 0; k < len(s); k++ {
		odd += count(s, k, k)
		even += count(s, k, k+1)
	}

	return odd + even
}

func count(s string, left, right int) int {
	numOfPalindromes := 0
	for left >= 0 && right < len(s) && left <= right && s[left] == s[right] {
		left--
		right++
		numOfPalindromes++
	}
	return numOfPalindromes
}
