package longest_palindromic_substring

func LongestPalindrome(s string) string {
	return longestPalindrome(s)
}

func longestPalindrome(s string) string {
	palindromeIndices := [2]int16{0, 0}

	for k := 0; k < len(s); k++ {
		oddIndices := findPalindrome(int16(k), int16(k), s)
		evenIndices := findPalindrome(int16(k), int16(k+1), s)

		if oddIndices[1]-oddIndices[0] > palindromeIndices[1]-palindromeIndices[0] {
			palindromeIndices = oddIndices
		}

		if evenIndices[1]-evenIndices[0] > palindromeIndices[1]-palindromeIndices[0] {
			palindromeIndices = evenIndices
		}
	}

	return s[palindromeIndices[0] : palindromeIndices[1]+1]
}
func findPalindrome(left, right int16, s string) [2]int16 {
	indices := [2]int16{0, 0}

	for left >= 0 && right < int16(len(s)) && s[left] == s[right] {
		if right-left > indices[1]-indices[0] {
			indices[0] = left
			indices[1] = right
		}
		left--
		right++
	}

	return indices
}
