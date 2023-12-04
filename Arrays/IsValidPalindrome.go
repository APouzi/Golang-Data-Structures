package arrays

import (
	"strings"
	"unicode"
)

//A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
// Given a string s, return true if it is a palindrome, or false otherwise.

//Example 1:
// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.

// Example 2:
// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.

// Example 3:
// Input: s = " "
// Output: true
// Explanation: s is an empty string "" after removing non-alphanumeric characters.
// Since an empty string reads the same forward and backward, it is a palindrome.
func IsPalindrome(s string) bool {
	var bank []byte = []byte{}

	news := strings.ToLower(s)
	for i := 0; i < len(news); i++ {
		if unicode.IsLetter(rune(news[i])) || (news[i] >= '0' && news[i] <= '9') {
			bank = append(bank, news[i])
		}

	}
	var left, right int = 0, len(bank) - 1
	for left < right {
		if bank[left] != bank[right] {
			return false
		}
		right--
		left++
	}
	return true
}