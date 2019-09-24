/*
771. Jewels and Stones

https://leetcode.com/problems/jewels-and-stones/

You're given strings J representing the types of stones that are jewels, and S representing the stones you have.
Each character in S is a type of stone you have.  You want to know how many of the stones you have are also jewels.

The letters in J are guaranteed distinct, and all characters in J and S are letters. Letters are case sensitive, so "a"
is considered a different type of stone from "A".

Example 1:

	Input: J = "aA", S = "aAAbbbb"
	Output: 3

Example 2:

	Input: J = "z", S = "ZZ"
	Output: 0
*/

package jewels_and_stones

func numJewelsInStones(J string, S string) int {
	list := make(map[rune]bool, 58)
	count := 0

	for _, j := range J {
		list[j-'A'] = true
	}

	for _, s := range S {
		if list[s-'A'] {
			count++
		}
	}

	return count
}

func numJewelsInStonesStupid(J string, S string) int {
	count := 0
	for _, s := range S {
		for _, j := range J {
			if s == j {
				count++
			}
		}
	}

	return count
}
