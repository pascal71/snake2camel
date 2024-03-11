package snake2camel

import (
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ToCamelCase converts a snake_case string to a camelCase string.
func ToCamelCase(s string) string {
	if s == "" {
		return "" // Early return for empty input
	}

	caser := cases.Title(language.English, cases.NoLower)
	words := strings.Split(s, "_")
	for i, word := range words {
		if word == "" {
			continue // Skip empty words resulting from consecutive underscores
		}

		if i > 0 {
			// Use the caser for capitalizing the first letter of each word except the first one
			words[i] = caser.String(word)
		} else {
			// Convert the first letter of the first word to lowercase if it's not already
			// Safe to access word[0] here because we know word is not empty
			words[i] = string(unicode.ToLower(rune(word[0]))) + word[1:]
		}
	}
	// Join the words back into a single string and return
	return strings.Join(words, "")
}
