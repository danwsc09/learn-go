package main

import "fmt"

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !isAlphaNumberic(s[left]) {
			left++
		}

		for left < right && !isAlphaNumberic(s[right]) {
			right--
		}

		if left < right && toLower(s[left]) != toLower(s[right]) {
			return false
		}
		left++
		right--
	}

	return true
}

func isAlphaNumberic(s byte) bool {
	return (s <= 'z' && s >= 'a') || (s <= 'Z' && s >= 'A') || (s <= '9' && s >= '0')
}

func toLower(b byte) byte {
	if b <= 'Z' && b >= 'A' {
		return b + 32
	}
	return b
}
