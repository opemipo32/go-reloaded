package main

import "strings"

// ProcessText runs the full text correction pipeline in order.
func ProcessText(text string) string {
	text = applyModifiers(text)
	text = fixPunctuation(text)
	text = fixQuotes(text)
	text = fixArticles(text)
	return strings.TrimSpace(text)
}
