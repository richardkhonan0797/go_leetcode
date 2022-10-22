package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i := 0
	j := len(s) - 1

	for i < j {
		for i < j && !unicode.IsDigit(rune(s[i])) && !unicode.IsLetter(rune(s[i])) {
			i++
		}
		for i < j && !unicode.IsDigit(rune(s[j])) && !unicode.IsLetter(rune(s[j])) {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
}
