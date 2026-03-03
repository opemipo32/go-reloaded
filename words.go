package main

import "unicode"

// transformWords applies fn to the last `count` words in the slice.
func transformWords(words []string, count int, fn func(string) string) []string {
	if len(words) == 0 {
		return words
	}
	start := len(words) - count
	if start < 0 {
		start = 0
	}
	for i := start; i < len(words); i++ {
		words[i] = fn(words[i])
	}
	return words
}

// capitalizeWord uppercases the first rune and lowercases the rest.
func capitalizeWord(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}
