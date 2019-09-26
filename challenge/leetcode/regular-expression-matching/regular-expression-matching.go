/*
10. Regular Expression Matching

Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.

'.' Matches any single character.
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).

Note:

s could be empty and contains only lowercase letters a-z.
p could be empty and contains only lowercase letters a-z, and characters like . or *.

Example 1:

	Input:
	s = "aa"
	p = "a"
	Output: false
	Explanation: "a" does not match the entire string "aa".

Example 2:

	Input:
	s = "aa"
	p = "a*"
	Output: true
	Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".

Example 3:

	Input:
	s = "ab"
	p = ".*"
	Output: true
	Explanation: ".*" means "zero or more (*) of any character (.)".

Example 4:

	Input:
	s = "aab"
	p = "c*a*b"
	Output: true
	Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore, it matches "aab".

Example 5:

	Input:
	s = "mississippi"
	p = "mis*is*p*."
	Output: false
*/

// Имплементируйте систему регулярных выражений, поддерживающую символы '.' и '*'.
//
// Где, '.' совпадает с любмым одним символом.
// Где, '*' совпадает с нулем или большим количеством предыдущих элементов.
//
// Совпадения должны покрывать строку целиком, не частично.
//
// Прототип функции должен быть:
//
//    bool isMatch(const char *s, const char *p)
//
// Примеры:
//
//    isMatch("aa","a") → false
//    isMatch("aa","aa") → true
//    isMatch("aaa","aa") → false
//    isMatch("aa", "a*") → true
//    isMatch("aa", ".*") → true
//    isMatch("ab", ".*") → true
//    isMatch("aab", "c*a*b") → true

package regular_expression_matching

import (
	"fmt"
)

type test struct {
	Text       string
	Expression string
	Expect     bool
}

func shifter(text, char string) int {
	for i, c := range text {
		if string(c) != char && char != "." {
			return i
		}
	}
	return 1
}

func isMatch2(text, expression string) bool {
	if len(text) == 0 {
		return len(expression) == 0
	}
	if len(expression) == 0 {
		return false
	}

	// если следующий элемент `*`
	if len(expression) > 1 && expression[1] == '*' {
		i := shifter(text[1:], string(expression[0]))

		// а текущий элемент не `.`
		if expression[0] != '.' {
			if expression[0] != text[0] {
				return isMatch(text, expression[2:])
			}
			if len(text[i:]) == 0 {
				return true
			}
			return isMatch(text[i+1:], expression[2:])
		}

		if len(text) == 1 || len(expression[2:]) == 0 {
			return true
		}

		return isMatch(text[i+1:], expression[2:])
	}

	// если текущий элемент `.`
	if expression[0] == '.' {
		// если это последний элменет, то выражение совпало
		if len(expression) == 1 {
			return !(len(text) > 1)
		}
		return isMatch(text[1:], expression[1:])
	}

	if expression[0] == text[0] {
		return isMatch(text[1:], expression[1:])
	}

	return false
}

func main() {
	tests := make([]test, 0)
	tests = append(tests, test{Text: "aaa", Expression: "ab*a*c*a", Expect: true})
	tests = append(tests, test{Text: "b", Expression: "b*", Expect: true})
	tests = append(tests, test{Text: "a", Expression: "c*.", Expect: true})
	tests = append(tests, test{Text: "aaa", Expression: "ab*ac*a", Expect: true})
	tests = append(tests, test{Text: "aaa", Expression: "ab*a", Expect: false})
	tests = append(tests, test{Text: "aaa", Expression: "a*a", Expect: true})
	tests = append(tests, test{Text: "aaa", Expression: "aaaa", Expect: false})
	tests = append(tests, test{Text: "ab", Expression: ".*c", Expect: false})
	tests = append(tests, test{Text: "aaaaaaaaaaaaab", Expression: "a*a*a*a*a*a*a*a*a*a*a*a*b", Expect: false})
	tests = append(tests, test{Text: "ab", Expression: ".*c", Expect: false})
	tests = append(tests, test{Text: "a", Expression: ".*", Expect: true})
	tests = append(tests, test{Text: "aaa", Expression: "a..", Expect: true})
	tests = append(tests, test{Text: "aaa", Expression: ".a.", Expect: true})
	tests = append(tests, test{Text: "aaa", Expression: "..a", Expect: true})
	tests = append(tests, test{Text: "aaa", Expression: "b.a", Expect: false})
	tests = append(tests, test{Text: "aa", Expression: "a", Expect: false})
	tests = append(tests, test{Text: "aa", Expression: "aa", Expect: true})
	tests = append(tests, test{Text: "aaa", Expression: "aa", Expect: false})
	tests = append(tests, test{Text: "aa", Expression: "a*", Expect: true})
	tests = append(tests, test{Text: "aa", Expression: "b*", Expect: false})
	tests = append(tests, test{Text: "aa", Expression: ".*", Expect: true})
	tests = append(tests, test{Text: "ab", Expression: ".*", Expect: true})
	tests = append(tests, test{Text: "aab", Expression: "c*a*b", Expect: true})
	tests = append(tests, test{Text: "aabb", Expression: "c*a*bb", Expect: true})
	tests = append(tests, test{Text: "aabb", Expression: "c*a*.b", Expect: true})
	tests = append(tests, test{Text: "aabb", Expression: "c*a*b.", Expect: true})
	tests = append(tests, test{Text: "aaa", Expression: ".*", Expect: true})
	tests = append(tests, test{Text: "", Expression: ".*", Expect: false})
	tests = append(tests, test{Text: "a", Expression: "", Expect: false})

	for _, test := range tests {
		result := isMatch(test.Text, test.Expression)
		if result == test.Expect {
			fmt.Printf("→ OK   (пришло %v)", result)
		} else {
			fmt.Printf("→ FAIL (пришло %v)", result)
		}
		fmt.Printf(" isMatch(\"%s\", \"%s\") должно быть = %v\n", test.Text, test.Expression, test.Expect)
	}
}

// lCmd находит в регулярке первую команду слева.
// Возвращает:
// 	 1. символы, которые нуждаются в проверке;
//   2. операцию, которую необходимо выполнить над возвращаемыми символами;
//   3. остаток необработанного текста регулярки.
func lCmd(p string) (string, string, string) {
	if len(p) == 0 {
		return "", "", ""
	}

	var i int
	var r rune
	var sym, cmd string

	for i, r = range p {
		if r == '.' || r == '*' {
			cmd = string(r)
			break
		}
		sym += string(r)
	}

	return sym, cmd, p[i+1:]
}

func cmdMatch(sym, cmd, cmdText string) (bool, string) {
	switch cmd {
	case ".":
	case "*":
	case "":
		// нет команды, значит мы должны проверить
		// строгое соответствие текста с sym

	}
	_ = sym
	_ = cmdText

	return
}

func isMatch(s string, p string) bool {
	if len(s) == 0 {
		return len(p) == 0
	}
	if len(p) == 0 {
		return false
	}

	sym, cmd, newP := lCmd(p)
	valid, newS := cmdMatch(sym, cmd, newP)

	return valid && isMatch(newS, newP)
}
