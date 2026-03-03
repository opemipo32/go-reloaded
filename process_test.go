package main

import (
	"testing"
)

func TestProcessedText(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Hex conversion",
			input: "Simply add 1E (hex) files",
			want:  "Simply add 30 files",
		},
		{
			name:  "Binary conversion",
			input: "It has been 10 (bin) years",
			want:  "It has been 2 years",
		},
		{
			name:  "Capitalize single word",
			input: "Welcome to the Brooklyn bridge (cap)",
			want:  "Welcome to the Brooklyn Bridge",
		},
		{
			name:  "Uppercase multiple words",
			input: "This is so exciting (up, 2)",
			want:  "This is SO EXCITING",
		},
		{
			name:  "Lowercase single word",
			input: "I should stop SHOUTING (low)",
			want:  "I should stop shouting",
		},
		{
			name:  "Uppercase single word",
			input: "Ready, set, go (up) !",
			want:  "Ready, set, GO!",
		},
		{
			name:  "Capitalize multiple words",
			input: "it was the age of foolishness (cap, 6)",
			want:  "It Was The Age Of Foolishness",
		},
		{
			name:  "Punctuation spacing",
			input: "I was sitting over there ,and then BAMM !!",
			want:  "I was sitting over there, and then BAMM!!",
		},
		{
			name:  "Ellipsis punctuation",
			input: "I was thinking ... You were right",
			want:  "I was thinking... You were right",
		},
		{
			name:  "Single quotes",
			input: "I am exactly how they describe me: ' awesome '",
			want:  "I am exactly how they describe me: 'awesome'",
		},
		{
			name:  "Single quotes multiple words",
			input: "As Elton John said: ' I am the most well-known homosexual in the world '",
			want:  "As Elton John said: 'I am the most well-known homosexual in the world'",
		},
		{
			name:  "Article a to an before vowel",
			input: "There it was. A amazing rock!",
			want:  "There it was. An amazing rock!",
		},
		{
			name:  "Article a to an before h",
			input: "There is no greater agony than bearing a untold story inside you.",
			want:  "There is no greater agony than bearing an untold story inside you.",
		},
		{
			name:  "Full sample 1",
			input: "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
			want:  "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.",
		},
		{
			name:  "Full sample 2",
			input: "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			want:  "Simply add 66 and 2 and you will see the result is 68.",
		},
		{
			name:  "Punctuation test",
			input: "Punctuation tests are ... kinda boring ,what do you think ?",
			want:  "Punctuation tests are... kinda boring, what do you think?",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ProcessText(tt.input)
			if got != tt.want {
				t.Errorf("\nInput: %q\nGot:   %q\nWant:  %q", tt.input, got, tt.want)
			}
		})
	}
}
