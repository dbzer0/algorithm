/*
5. Longest Palindromic Substring

https://leetcode.com/problems/longest-palindromic-substring/

Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

	Input: "babad"
	Output: "bab"
	Note: "aba" is also a valid answer.

Example 2:

	Input: "cbbd"
	Output: "bb"
*/

package longest_palindromic_substring

// longestPalindrome находит самый длинный малиндром в строке.
// Для этого:
// 	1. Осуществляется основной посимвольный проход по всей строке;
//  2. Внутренним циклом просматриваются символы справа и слева, от индексного.
//     Если они одинаковые, то внутренний цикл продолжается до поиска их разности
//     или нахождения границ строки. Цикл пытается определить центр палиндрома, где
//     центром будет являться одна буква, например: aba;
//  3. Внутренним циклом просматриваются символы справа и слева от индексного до индесного+1.
//     Если они одинаковые, то внутренний цикл продолжается до поиска их разности или
//     нахождения границ строки. Цикл пыатется определить центр палиндрома, где центром
//     будет парная буква, например: abba.
func longestPalindrome(s string) string {
	max := ""
	l := len(s)

	for i := 0; i < l; i++ {
		if len(max) >= l+i {
			break
		}

		// случай, при котором центр палиндрома - это одна несимметричная буква
		left, right := i, i
		for left >= 0 && right < l && s[left] == s[right] {
			left--
			right++
		}
		if len(max) < right-left-1 {
			max = s[left+1 : right]
		}

		// случай, при котором центр палиндрома - это две повторяющиеся буквы
		left, right = i, i+1
		for left >= 0 && right < l && s[left] == s[right] {
			left--
			right++
		}
		if len(max) < right-left-1 {
			max = s[left+1 : right]
		}
	}

	return max
}
