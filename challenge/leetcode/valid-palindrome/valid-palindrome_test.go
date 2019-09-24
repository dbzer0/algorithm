package valid_palindrome

import "testing"

func Test_isPalindrome(t *testing.T) {
	if !isPalindrome("") {
		t.Error("isPalindrome(\"\") failed, must be: true")
	}

	if !isPalindrome("A man, a plan, a canal: Panama") {
		t.Error("isPalindrome(\"A man, a plan, a canal: Panama\") failed, must be: true")
	}

	if isPalindrome("abcde") {
		t.Error("isPalindrome(\"abcde\") failed, must be: false")
	}
}
