// Package calculator provides a library for simple calculations in Go.
package calculator

import "errors"

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and returns the result by dividing the first by the second
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero not allowed")
	}

	return a / b, nil
}
