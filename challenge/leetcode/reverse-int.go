package main

import (
	"fmt"
	"strconv"
)

func reverseStr2(s string) string {
	n := ""
	for i := range s {
		n = fmt.Sprintf("%s%c", n, s[len(s)-1-i])
	}
	return n
}

func reverseStr(s string) string {
	if len(s) <= 1 {
		return s
	}
	return string(s[len(s)-1]) + reverseStr(s[:len(s)-1])
}

func reverse(x int) int {
	var r int
	if x < 0 {
		r, _ = strconv.Atoi(reverseStr(strconv.Itoa(-x)))
		r = -r
	} else {
		r, _ = strconv.Atoi(reverseStr(strconv.Itoa(x)))
	}
	if r < -2147483648 || r > 2147483647 {
		return 0
	}
	return r
}

func main() {
	fmt.Printf("%d\n", reverse(123125))
}
