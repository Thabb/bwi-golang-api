package main

import (
	"slices"
	"strings"
)

// Parser function that takes a slice of tokens and constructs an evaluation tree using a recursive descent parsing approach.
// It returns the root of the evaluation tree and an error message if any issues are encountered during parsing.
func parser(tokens []string) (*Tree, string) {
	var evalTree Tree
	var parensCount int = 0
	var error string = ""

	// Check for leading and trailing operators
	if tokens[0] == "+" || tokens[0] == "-" || tokens[0] == "*" || tokens[0] == "/" {
		error = "Leading operator"
		return nil, error
	}
	if tokens[len(tokens)-1] == "+" || tokens[len(tokens)-1] == "-" || tokens[len(tokens)-1] == "*" || tokens[len(tokens)-1] == "/" {
		error = "Trailing operator"
		return nil, error
	}

	// Add a missing * operator between a number and a parenthesis
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "(" && i > 0 {
			if tokens[i-1] >= "0" && tokens[i-1] <= "9" {
				tokens = slices.Insert(tokens, i, "*")
			}
		}
		if tokens[i] == ")" && i < len(tokens)-1 {
			if tokens[i+1] >= "0" && tokens[i+1] <= "9" {
				tokens = slices.Insert(tokens, i+1, "*")
			}
		}
	}

	// Process addition and subtraction
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "(" {
			parensCount++
		}
		if tokens[i] == ")" {
			parensCount--
		}

		if parensCount == 0 {
			if tokens[i] == "+" || tokens[i] == "-" {
				if !(tokens[i+1] >= "0" && tokens[i+1] <= "9") {
					error = "Multiple operators in a row"
					return nil, error
				}
				evalTree.value = tokens[i]
				evalTree.left, error = parser(tokens[0:i])
				evalTree.right, error = parser(tokens[i+1:])
				return &evalTree, error
			}
		}
	}

	// Check for mismatched parentheses
	if parensCount != 0 {
		error = "Mismatched parentheses"
		return nil, error
	}

	// Process multiplication and division
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "(" {
			parensCount++
		}
		if tokens[i] == ")" {
			parensCount--
		}

		if parensCount == 0 {
			if tokens[i] == "*" || tokens[i] == "/" {
				if tokens[i+1] == "+" || tokens[i+1] == "-" {
					error = "Multiple operators in a row"
					return nil, error
				}
				evalTree.value = tokens[i]
				evalTree.left, error = parser(tokens[0:i])
				evalTree.right, error = parser(tokens[i+1:])
				return &evalTree, error
			}
		}
	}

	// Check for mismatched amount of parentheses
	if parensCount != 0 {
		error = "Mismatched parentheses amount"
		return nil, error
	}

	// Process parentheses
	if tokens[0] == "(" && tokens[len(tokens)-1] == ")" {
		// Check for empty parentheses
		if len(tokens) < 3 {
			error = "Empty parentheses"
			return nil, error
		}
		// Remove the outer parentheses and parse the inner expression
		return parser(tokens[1 : len(tokens)-1])
	}

	// If the expression is just a number, return it as a leaf node
	if (tokens[0] >= "0" && tokens[0] <= "9") && (tokens[len(tokens)-1] >= "0" && tokens[len(tokens)-1] <= "9") {
		evalTree.value = strings.Join(tokens, "")
		return &evalTree, error
	}

	return nil, error
}
