package main

import (
	"fmt"
)

func main() {
	s := "pwwkew"
	// s := "abcabcbb"
	// s := "abba"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	seen := [256]int{}
	for i := 0; i < 256; i++ {
		seen[i] = -1
	}

	dup := -1
	res := 0
	for i := 0; i < len(s); i++ {
		if seen[s[i]] > -1 && seen[s[i]] > dup {
			dup = seen[s[i]]
		}
		seen[s[i]] = i
		if i-dup > res {
			res = i - dup
		}
	}
	return res

	/*
		lastOc := make(map[byte]int)
		n := len(s)
		res, lastDup := 0, -1

		for i := 0; i < n; i++ {
			ch := s[i]
			if val, ok := lastOc[ch]; ok {
				lastDup = max(val, lastDup)
			}
			lastOc[ch] = i
			res = max(res, i-lastDup)
			fmt.Println(lastDup, res, lastOc)
		}

		return res
	*/
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
