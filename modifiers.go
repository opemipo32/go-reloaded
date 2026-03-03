package main

import (
	"regexp"
	"strconv"
	"strings"
)

// modifierRe matches tags like (hex), (up), (cap, 3), etc.
var modifierRe = regexp.MustCompile(`(?i)\((hex|bin|up|low|cap)(?:,\s*(\d+))?\)`)

// applyModifiers scans the text left-to-right, finds each modifier tag,
// and applies the corresponding transformation to the words before it.
func applyModifiers(text string) string {
	for {
		loc := modifierRe.FindStringIndex(text)
		if loc == nil {
			break
		}

		modifier, count := parseModifier(text[loc[0]:loc[1]])
		before := strings.TrimRight(text[:loc[0]], " ")
		after := text[loc[1]:]

		words := strings.Fields(before)
		words = dispatchModifier(words, modifier, count)

		text = joinSegments(strings.Join(words, " "), after)
	}
	return text
}

// parseModifier extracts the modifier name and optional word count from a tag.
func parseModifier(tag string) (modifier string, count int) {
	sub := modifierRe.FindStringSubmatch(tag)
	modifier = strings.ToLower(sub[1])
	count = 1
	if sub[2] != "" {
		if n, err := strconv.Atoi(sub[2]); err == nil && n > 0 {
			count = n
		}
	}
	return modifier, count
}

// dispatchModifier routes each modifier to its handler.
func dispatchModifier(words []string, modifier string, count int) []string {
	switch modifier {
	case "hex":
		return convertBase(words, 16)
	case "bin":
		return convertBase(words, 2)
	case "up":
		return transformWords(words, count, strings.ToUpper)
	case "low":
		return transformWords(words, count, strings.ToLower)
	case "cap":
		return transformWords(words, count, capitalizeWord)
	}
	return words
}

// joinSegments reassembles the before and after parts with correct spacing.
func joinSegments(before, after string) string {
	trimmedAfter := strings.TrimLeft(after, " ")
	if before != "" && trimmedAfter != "" {
		return before + " " + trimmedAfter
	}
	return before + trimmedAfter
}
