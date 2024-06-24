package fuzz_demo

import "strings"

// Reverse reverses the input string.
func Reverse(s string) string {
	return strings.Join(reverseStrings(strings.Fields(s)), " ")
}

// reverseStrings reverses each word in a slice of strings.
func reverseStrings(words []string) []string {
	for i, word := range words {
		words[i] = reverseString(word)
	}
	return words
}

// reverseString reverses a single string.
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
