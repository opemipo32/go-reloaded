package main

import "strconv"

// convertBase converts the last word in words from the given base to decimal.
func convertBase(words []string, base int) []string {
	if len(words) == 0 {
		return words
	}
	idx := len(words) - 1
	n, err := strconv.ParseInt(words[idx], base, 64)
	if err == nil {
		words[idx] = strconv.FormatInt(n, 10)
	}
	return words
}
