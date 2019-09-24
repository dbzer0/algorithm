package main

import (
	"fmt"
)

func multiply2(num1 string, num2 string) string {
	if len(num1) >= 110 || len(num2) >= 110 {
		return ""
	}

	s2i := func(s string) int {
		var n int
		for _, ch := range []byte(s) {
			ch -= '0'
			n = n*10 + int(ch)
		}
		return n
	}

	i2s := func(x int) string {
		if x < 10 {
			return string('0' + x)
		}

		b := make([]byte, 0)
		for {
			b = append(b, byte(x%10+'0'))
			x = x / 10
			if x%10 == 0 && x < 10 {
				break
			}
		}
		return string(b)
	}

	return reverseString(i2s(s2i(num1) * s2i(num2)))
}

func s2i(s string) int {
	var n int
	for _, ch := range []byte(s) {
		ch -= '0'
		n = n*10 + int(ch)
	}
	return n
}

func multiply(num1 string, num2 string) string {
	if len(num1) >= 110 || len(num2) >= 110 {
		return ""
	}

	var long, short string
	var lenLong, lenShort int
	if len(num1) >= len(num2) {
		long, short = num1, num2
		lenLong, lenShort = len(num1), len(num2)
	} else {
		long, short = num2, num1
		lenLong, lenShort = len(num2), len(num1)
	}

	for i := range long {
		var s2 string
		s1 := string(long[lenLong-i-1])

		if lenShort <= i {
			s2 = "1"
		} else {
			s2 = string(short[lenShort-i-1])
		}

		fmt.Println(s2i(s1) * s2i(s2))
	}

	return ""
}

func reverseString(s string) string {
	if len(s) <= 1 {
		return s
	}
	return string(s[len(s)-1]) + reverseString(s[:len(s)-1])
}

func main() {
	fmt.Println(multiply("123", "10"))
}
