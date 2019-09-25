package string_to_integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_normalize(t *testing.T) {
	assert.Equal(t, "-09", normalize("-09"))
	assert.Equal(t, "-09", normalize("      -09"))
	assert.Equal(t, "09", normalize("+09"))
	assert.Equal(t, "0", normalize("a09"))
	assert.Equal(t, "3", normalize("3.1415"))
	assert.Equal(t, "23", normalize("    23.1415"))
	assert.Equal(t, "0", normalize("+-2"))
	assert.Equal(t, "-0012", normalize("  -0012a42"))
	assert.Equal(t, "0", normalize("0-1"))
	assert.Equal(t, "21474836", normalize("21474836++"))
	assert.Equal(t, "0", normalize("  +  413"))
}

func Test_myAtoi(t *testing.T) {
	assert.Equal(t, 42, myAtoi("42"))
	assert.Equal(t, -42, myAtoi("      -42"))
	assert.Equal(t, 4193, myAtoi("4193 with words"))
	assert.Equal(t, 0, myAtoi("words and 987"))
	assert.Equal(t, -2147483648, myAtoi("-91283472332"))
	assert.Equal(t, 3, myAtoi("3.14159"))
	assert.Equal(t, 2147483647, myAtoi("9223372036854775808"))
	assert.Equal(t, 12345678, myAtoi("  0000000000012345678"))
	assert.Equal(t, -12, myAtoi("  -0012a42"))
	assert.Equal(t, 0, myAtoi("0-1"))
	assert.Equal(t, 21474836, myAtoi("21474836++"))
	assert.Equal(t, 0, myAtoi("  +  413"))
	assert.Equal(t, 2147483647, myAtoi("18446744073709551617"))
	assert.Equal(t, 2147483647, myAtoi("10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459"))
}
