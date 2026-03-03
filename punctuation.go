package main

import "regexp"

var (
	punctBeforeRe = regexp.MustCompile(` +([.,!?:;]+)`)
	punctAfterRe  = regexp.MustCompile(`([.,!?:;]+)([^\s.,!?:;'")\-])`)
)

// fixPunctuation removes spaces before punctuation groups (including ... and !?)
// and ensures exactly one space follows them before the next word.
func fixPunctuation(text string) string {
	text = punctBeforeRe.ReplaceAllString(text, "$1")
	text = punctAfterRe.ReplaceAllString(text, "$1 $2")
	return text
}
