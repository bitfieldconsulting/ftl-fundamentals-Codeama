// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"fmt"
	"math"
)

// Add takes multiple numbers and returns the result of adding them together.
func Add(a, b float64, n ...float64) float64 {
	fmt.Println("a", a, "b", b, "n", n)
	sum := a + b
	for _, v := range n {
		sum += v
	}
	return sum
}

// Subtract takes multiple numbers deducting the second from the first and returns the result
// from the number before.
func Subtract(a, b float64, n ...float64) float64 {
	total := a - b
	for _, v := range n {
		total -= v
	}
	return total
}

// Multiply takes multiple numbers and returns the result of multiplying them
func Multiply(a, b float64, n ...float64) float64 {
	total := a * b
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

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("square root of negative number not allowed... you monster: %f", a)
	}
	result := math.Sqrt(a)
	return result, nil
}
