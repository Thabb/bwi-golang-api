package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Tree struct {
	value string
	left  *Tree
	right *Tree
}

func main() {
	router := gin.Default()
	router.POST("/calculate", calculate)
	router.Run("localhost:8080")
}

func calculate(c *gin.Context) {
	// Call lexer
	tokens, error := lexer(c.Query("form"))
	if error != "" {
		c.JSON(400, gin.H{
			"error": error,
		})
		return
	}

	// Call parser
	evalTree, error := parser(tokens)
	if error != "" {
		c.JSON(400, gin.H{
			"error": error,
		})
		return
	}

	// Call evaluator
	result := evaluate(*evalTree)
	c.JSON(200, gin.H{
		"result": result,
	})
}

func lexer(form string) ([]string, string) {
	var tokens []string
	var error string = ""

	// remove all white space
	form = strings.Replace(form, " ", "", -1)
	// split form into tokens
	tokens = strings.Split(form, "")

	// check if all tokens are valid
	for i := 0; i < len(tokens); i++ {
		if !((tokens[i] >= "0" && tokens[i] <= "9") || slices.Contains([]string{"+", "-", "*", "/", "(", ")"}, tokens[i])) {
			error = "Invalid token: " + tokens[i]
			break
		}
	}
	return tokens, error
}

// TODO: Add error handling for invalid expressions
func parser(tokens []string) (*Tree, string) {
	var evalTree Tree
	var parensCount int = 0
	var error string = ""

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

	// Check for parentheses
	if tokens[0] == "(" && tokens[len(tokens)-1] == ")" {
		// Remove the outer parentheses and parse the inner expression
		return parser(tokens[1 : len(tokens)-1])
	}

	if (tokens[0] >= "0" && tokens[0] <= "9") && (tokens[len(tokens)-1] >= "0" && tokens[len(tokens)-1] <= "9") {
		// If the expression is a single number, return it as a leaf node
		evalTree.value = strings.Join(tokens, "")
		return &evalTree, error
	}

	return nil, error
}

// TODO: Add error handling for invalid expressions
func evaluate(evalTree Tree) int {
	if (evalTree.value == "+" || evalTree.value == "-" || evalTree.value == "*" || evalTree.value == "/") && evalTree.left != nil && evalTree.right != nil {
		_, leftErr := strconv.Atoi(evalTree.left.value)
		_, rightErr := strconv.Atoi(evalTree.right.value)
		if leftErr != nil {
			evalTree.left.value = strconv.Itoa(evaluate(*evalTree.left))
			leftErr = nil
		}
		if rightErr != nil {
			evalTree.right.value = strconv.Itoa(evaluate(*evalTree.right))
			rightErr = nil
		}
		if leftErr == nil && rightErr == nil {
			leftVal, _ := strconv.Atoi(evalTree.left.value)
			rightVal, _ := strconv.Atoi(evalTree.right.value)
			switch evalTree.value {
			case "+":
				return leftVal + rightVal
			case "-":
				return leftVal - rightVal
			case "*":
				return leftVal * rightVal
			case "/":
				return leftVal / rightVal
			}
		}
	}
	return -1
}
