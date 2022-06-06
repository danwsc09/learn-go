package main

import (
	"fmt"
)

func main() {
	fmt.Println(len(""))
	fmt.Println(longestPalindrome("abba"))
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	n := len(s)
	res := ""

	for i := 0; i < n; i++ {
		ch := s[i]
		l, r := i, i

		for l >= 0 && s[l] == ch {
			l--
		}

		for r < n && s[r] == ch {
			r++
		}

		for l >= 0 && r < n {
			if s[l] != s[r] {
				break
			}
			l--
			r++
		}
		if r-l-1 > len(res) {
			res = s[l+1 : r]
		}
	}

	return res
}

func extendPalindrome(s string, i, j int) string {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return s[i : j+1]
}

func isPalin(s string) bool {
	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
