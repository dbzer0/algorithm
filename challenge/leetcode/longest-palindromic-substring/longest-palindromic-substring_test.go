package longest_palindromic_substring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestPalindrome(t *testing.T) {
	assert.Equal(t, "", longestPalindrome(""))
	assert.Equal(t, "a", longestPalindrome("a"))
	assert.Equal(t, "aa", longestPalindrome("aa"))
	assert.Equal(t, "aa", longestPalindrome("baa"))
	assert.Equal(t, "bab", longestPalindrome("babad"))
	assert.Equal(t, "bb", longestPalindrome("cbbd"))
}
