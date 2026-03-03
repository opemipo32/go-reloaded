package main

import "regexp"

var articleRe = regexp.MustCompile(`\b([Aa])\s+([AaEeIiOoUuHh]\w*)`)

// fixArticles replaces "a" with "an" when the following word starts with a vowel or 'h'.
// Preserves the original casing of the article.
func fixArticles(text string) string {
	return articleRe.ReplaceAllStringFunc(text, func(match string) string {
		sub := articleRe.FindStringSubmatch(match)
		article, word := sub[1], sub[2]
		if article == "A" {
			return "An " + word
		}
		return "an " + word
	})
}
