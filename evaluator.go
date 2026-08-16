package main

import (
	"strconv"
)

// Evaluator function that takes the root of an evaluation tree and recursively computes the result of the mathematical expression represented by the tree.
func evaluator(evalTree Tree) (int, string) {
	var error string = ""

	if (evalTree.value == "+" || evalTree.value == "-" || evalTree.value == "*" || evalTree.value == "/") && evalTree.left != nil && evalTree.right != nil {
		_, leftErr := strconv.Atoi(evalTree.left.value)
		_, rightErr := strconv.Atoi(evalTree.right.value)
		if leftErr != nil {
			leftResult, _ := evaluator(*evalTree.left)
			evalTree.left.value = strconv.Itoa(leftResult)
			leftErr = nil
		}
		if rightErr != nil {
			rightResult, _ := evaluator(*evalTree.right)
			evalTree.right.value = strconv.Itoa(rightResult)
			rightErr = nil
		}
		if leftErr == nil && rightErr == nil {
			leftVal, _ := strconv.Atoi(evalTree.left.value)
			rightVal, _ := strconv.Atoi(evalTree.right.value)
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
