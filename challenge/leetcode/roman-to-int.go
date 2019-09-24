/*
  Given a roman numeral, convert it to an integer.
  Input is guaranteed to be within the range from 1 to 3999.
*/
package main

import "fmt"

func c(b byte) int {
	switch b {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}

func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return c(s[0])
	}

	if c(s[0]) < c(s[1]) {
		return romanToInt(s[1:]) - c(s[0])
	} else {
		return romanToInt(s[1:]) + c(s[0])
	}
}

func main() {
	fmt.Println(romanToInt("XLIX"))
	fmt.Println(romanToInt("CXXVIII"))
	fmt.Println(romanToInt("LXXX"))
	fmt.Println(romanToInt("MDCCC"))
}
