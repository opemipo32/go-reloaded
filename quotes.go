package main

import "regexp"

var quotesRe = regexp.MustCompile(`'\s+([^']*?)\s+'`)

// fixQuotes removes surrounding whitespace from inside single-quote pairs.
// e.g. ' awesome ' -> 'awesome'
func fixQuotes(text string) string {
	return quotesRe.ReplaceAllStringFunc(text, func(match string) string {
		sub := quotesRe.FindStringSubmatch(match)
		return "'" + sub[1] + "'"
	})
}
