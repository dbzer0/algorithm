package main

import (
	"fmt"
)

/*
Implement pow(x, n), which calculates x raised to the power n (xn).

Example 1:

Input: 2.00000, 10
Output: 1024.00000
Example 2:

Input: 2.10000, 3
Output: 9.26100
Example 3:

Input: 2.00000, -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25
Note:

-100.0 < x < 100.0
n is a 32-bit signed integer, within the range [−231, 231 − 1]
*/

func main() {
	fmt.Println(myPow(2, 2))
	fmt.Println(myPowRecursion(2, 2))
}

func myPow(x float64, n int) float64 {
	switch {
	case x == 0 || x == 1:
		return 1
	case n == 1:
		return x
	}

	l := n
	if n < 0 {
		l = -n
	}

	result := float64(1)

	for i := 1; i <= l; i++ {
		if n < 0 {
			result *= 1 / x
		} else {
			result *= x
		}
	}

	return result
}

func myPowRecursion(x float64, n int) float64 {
	switch {
	case x == 0 || x == 1:
		return 1
	case n == 1:
		return x
	}

	if n < 0 {
		return (1 / x) * myPow(x, n+1)
	}

	return x * myPow(x, n-1)
}
