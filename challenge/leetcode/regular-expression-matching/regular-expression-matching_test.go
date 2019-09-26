package regular_expression_matching

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lCmd(t *testing.T) {
	sym, cmd, text := lCmd("")
	assert.Equal(t, "", sym)
	assert.Equal(t, "", cmd)
	assert.Equal(t, "", text)

	sym, cmd, text = lCmd("a")
	assert.Equal(t, "a", sym)
	assert.Equal(t, "", cmd)
	assert.Equal(t, "", text)

	sym, cmd, text = lCmd(".")
	assert.Equal(t, "", sym)
	assert.Equal(t, ".", cmd)
	assert.Equal(t, "", text)

	sym, cmd, text = lCmd("a.")
	assert.Equal(t, "a", sym)
	assert.Equal(t, ".", cmd)
	assert.Equal(t, "", text)

	sym, cmd, text = lCmd("ab.")
	assert.Equal(t, "ab", sym)
	assert.Equal(t, ".", cmd)
	assert.Equal(t, "", text)

	sym, cmd, text = lCmd("ab*ac*a")
	assert.Equal(t, "ab", sym)
	assert.Equal(t, "*", cmd)
	assert.Equal(t, "ac*a", text)
}

func Test_cmdMatch(t *testing.T) {
	valid, text := cmdMatch("", "", "")
	assert.True(t, valid)
	assert.Equal(t, "", text)
}

func Test_isMatch(t *testing.T) {
	assert.False(t, isMatch("aa", "p"))
	assert.True(t, isMatch("aa", "a*"))
	assert.True(t, isMatch("ab", ".*"))
	assert.True(t, isMatch("aab", "c*a*b"))
	assert.False(t, isMatch("mississippi", "mis*is*p*."))
}
