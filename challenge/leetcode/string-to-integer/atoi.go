/*
8. String to Integer (atoi)

https://leetcode.com/problems/string-to-integer-atoi/

Implement atoi which converts a string to an integer.

The function first discards as many whitespace characters as necessary until the first
non-whitespace character is found. Then, starting from this character, takes an optional
initial plus or minus sign followed by as many numerical digits as possible, and
interprets them as a numerical value.

The string can contain additional characters after those that form the integral number,
which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid integral number,
or if no such sequence exists because either str is empty or it contains only whitespace
characters, no conversion is performed.

If no valid conversion could be performed, a zero value is returned.

Note:

Only the space character ' ' is considered as whitespace character.
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. If the numerical value is out of the range of representable values, INT_MAX (231 − 1) or INT_MIN (−231) is returned.

Example 1:

	Input: "42"
	Output: 42

Example 2:

	Input: "   -42"
	Output: -42

Explanation: The first non-whitespace character is '-', which is the minus sign.
             Then take as many numerical digits as possible, which gets 42.

Example 3:

	Input: "4193 with words"
	Output: 4193
	Explanation: Conversion stops at digit '3' as the next character is not a numerical digit.

Example 4:

	Input: "words and 987"
	Output: 0
	Explanation: The first non-whitespace character is 'w', which is not a numerical
				 digit or a +/- sign. Therefore no valid conversion could be performed.

Example 5:

	Input: "-91283472332"
	Output: -2147483648
	Explanation: The number "-91283472332" is out of the range of a 32-bit signed integer.
				 Thefore INT_MIN (−231) is returned.
*/

package string_to_integer

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func normalize(s string) string {
	res, sign, ns := "", "", ""

	trimmed := strings.TrimSpace(s)
	if len(trimmed) == 0 {
		return "0"
	}

	switch trimmed[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		ns = trimmed
	case '-':
		sign = "-"
		ns = trimmed[1:]
	case '+':
		ns = trimmed[1:]
	}

	for _, r := range ns {
		if !unicode.IsNumber(r) {
			break
		}
		res = fmt.Sprintf("%s%c", res, r)
	}

	if len(res) == 0 {
		return "0"
	}

	return sign + res
}

func myAtoi(str string) int {
	result := 0
	nStr := normalize(str)

	if len(nStr) < 1 {
		return 0
	}

	sign := 1
	if nStr[0] == '-' {
		sign = -1
		nStr = nStr[1:]
	}

	nStr = strings.TrimLeft(nStr, "0")

	for i, r := range nStr {
		result = result*10 + int(r-'0')

		if result*sign < math.MinInt32 {
			if sign < 0 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
		if result > math.MaxInt32 || i > 11 {
			if sign < 0 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}

	return result * sign
}
