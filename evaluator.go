package main

import (
	"strconv"
)

// Evaluator function that takes the root of an evaluation tree and recursively computes the result of the mathematical expression represented by the tree.
func evaluator(evalTree Tree) (float64, string) {
	var error string = ""

	if (evalTree.value == "+" || evalTree.value == "-" || evalTree.value == "*" || evalTree.value == "/") && evalTree.left != nil && evalTree.right != nil {
		_, leftErr := strconv.ParseFloat(evalTree.left.value, 64)
		_, rightErr := strconv.ParseFloat(evalTree.right.value, 64)

		// If either left or right child is not a number, evaluate them recursively
		if leftErr != nil {
			leftResult, _ := evaluator(*evalTree.left)
			evalTree.left.value = strconv.FormatFloat(leftResult, 'f', -1, 64)
			leftErr = nil
		}
		if rightErr != nil {
			rightResult, _ := evaluator(*evalTree.right)
			evalTree.right.value = strconv.FormatFloat(rightResult, 'f', -1, 64)
			rightErr = nil
		}

		// If both left and right children are numbers, perform the operation
		if leftErr == nil && rightErr == nil {
			leftVal, _ := strconv.ParseFloat(evalTree.left.value, 64)
			rightVal, _ := strconv.ParseFloat(evalTree.right.value, 64)
			switch evalTree.value {
			case "+":
				return leftVal + rightVal, error
			case "-":
				return leftVal - rightVal, error
			case "*":
				return leftVal * rightVal, error
			case "/":
				if rightVal == 0 {
					error = "Division by zero"
					return -1, error
				}
				return leftVal / rightVal, error
			}
		}
	}
	return -1, error
}
