// Package calculator provides a library for simple calculations in Go.
package calculator

import "errors"

// Add takes multiple numbers and returns the result of adding them together.
func Add(n ...float64) float64 {
	sum := 0.0
	for _, v := range n {
		sum += v
	}
	return sum
}

// Subtract takes multiple numbers deducting the second from the first and returns the result
// from the number before.
func Subtract(n ...float64) float64 {
	total := n[0]
	for i := 1; i < len(n); i++ {
		total -= n[i]
	}
	return total
}

// Multiply takes multiple numbers and returns the result of multiplying them
func Multiply(n ...float64) float64 {
	total := 1.0
	for _, v := range n {
		total *= v
	}
	return total
}

// Divide takes multiple numbers and returns the result by dividing the first by the second
func Divide(n ...float64) (float64, error) {
	total := n[0]
	for i := 1; i < len(n); i++ {
		if n[i] == 0 {
			return 0, errors.New("division by zero not allowed")
		}
		total /= n[i]
	}

	return total, nil
}
