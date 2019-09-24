/*
125. Valid Palindrome

https://leetcode.com/problems/valid-palindrome/

Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:

	Input: "A man, a plan, a canal: Panama"
	Output: true

Example 2:

	Input: "race a car"
	Output: false
*/

package valid_palindrome

import "strings"

func isValidRune(r rune) bool {
	return 'z'-r < 26 || 'z'-r >= 65 && 'z'-r <= 74
}

func isPalindrome(s string) bool {
	t := strings.ToLower(s)
	si, ei := 0, len(s)-1

	for si < ei {
		if !isValidRune(rune(t[si])) {
			si++
			continue
		}
		if !isValidRune(rune(t[ei])) {
			ei--
			continue
		}
		if t[si] != t[ei] {
			return false
		}
		si++
		ei--
	}

	return true
}
