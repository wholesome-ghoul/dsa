package valid_palindrome

import (
	"strings"
)

func IsPalindrome(s string) bool {
	return isPalindrome(s)
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left <= right {
		if !isAlphaNumeric(s[left]) {
			left++
			continue
		}

		if !isAlphaNumeric(s[right]) {
			right--
			continue
		}

		if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphaNumeric(letter byte) bool {
	return (letter >= 48 && letter <= 57) || (letter >= 65 && letter <= 90) || (letter >= 97 && letter <= 122)
}
