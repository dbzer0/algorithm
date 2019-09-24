/*
Given a string, find the length of the longest substring without repeating characters.

Examples:
	Given "abcabcbb", the answer is "abc", which the length is 3.
	Given "bbbbb", the answer is "b", with the length of 1.
	Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

package main

import "fmt"

func hashInit() [128]int {
	hash := [128]int{}
	for i := 0; i < len(hash); i++ {
		hash[i] = -1
	}
	return hash
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	maxLen, start := 0, -1
	hash := hashInit()

	for i := 0; i < len(s); i++ {
		if hash[s[i]] > start {
			start = hash[s[i]]
		}

		hash[s[i]] = i
		if maxLen < i-start {
			maxLen = i - start
		}
	}

	return maxLen
}

func main() {
	//fmt.Println(lengthOfLongestSubstring(""))
	//fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	//fmt.Println(lengthOfLongestSubstring("bbbbbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	//fmt.Println(lengthOfLongestSubstring("au"))
	//fmt.Println(lengthOfLongestSubstring("dvdf"))
}
