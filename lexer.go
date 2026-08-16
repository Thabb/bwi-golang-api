package main

import (
	"slices"
	"strings"
)

// Lexer function that takes a mathematical expression as input and returns a slice of tokens and an error message if any invalid tokens are found.
func lexer(form string) ([]string, string) {
	var tokens []string
	var error string = ""

	// Remove all white space
	form = strings.Replace(form, " ", "", -1)
	// Split form into tokens
	tokens = strings.Split(form, "")

	// Check if all tokens are valid
	for i := 0; i < len(tokens); i++ {
		if !((tokens[i] >= "0" && tokens[i] <= "9") || slices.Contains([]string{"+", "-", "*", "/", "(", ")"}, tokens[i])) {
			error = "Invalid token: " + tokens[i]
			break
		}
	}
	return tokens, error
}
