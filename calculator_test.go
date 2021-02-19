package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

func TestCalculator(t *testing.T) {
	t.Parallel()
	testCases := []struct{
		name string
		n    []float64
		f    func(float64, float64, ...float64) float64
		want float64
	}{
		{
			name: "Add three positive integers",
			n:    []float64{2, 2, 8},
			f:    calculator.Add,
			want: 12,
		},
		{
			name: "Add multiple negative and positive integers",
			n:    []float64{-2, -2, 23, 40, 0},
			f:    calculator.Add,
			want: 59,
		},
		{
			name: "Add a positive integer and a negative fraction ",
			n:    []float64{5, -2.48},
			f:    calculator.Add,
			want: 2.52,
		},
		{
			name: "Subtract a larger integer from a smaller integer",
			n:    []float64{3, 5},
			f:    calculator.Subtract,
			want: -2,
		},
		{
			name: "Subtract a negative integer from a negative integer",
			n:    []float64{-9, -5, 1},
			f:    calculator.Subtract,
			want: -5,
		},
		{
			name: "Multiply two positive integers",
			n:    []float64{4, 1},
			f:    calculator.Multiply,
			want: 4,
		},
		{
			name: "Multiply positive and a negative numbers",
			n:    []float64{4, -9, 2},
			f:    calculator.Multiply,
			want: -72,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.f(tc.n[0], tc.n[1], tc.n[2:]...)
			if tc.want != got {
				t.Errorf("inputs: (%f, %f, %v): want %f, got %f", tc.n[0], tc.n[1], tc.n[2:], tc.want, got)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	testCases := []struct{
		n           []float64
		want        float64
		errExpected bool
	}{
		{
			n:           []float64{5, 0, 4},
			want:        0,
			errExpected: true,
		},
		{
			n:           []float64{10, 2.5},
			want:        4,
			errExpected: false,
		},
		{
			n:           []float64{28, 2},
			want:        14,
			errExpected: false,
		},
		{
			n:           []float64{1, 1},
			want:        1,
			errExpected: false,
		},
		{
			n:           []float64{0, 2.5},
			want:        0,
			errExpected: false,
		},
		{
			n:           []float64{20, 2.5, 2},
			want:        4,
			errExpected: false,
		},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.n...)
		errorReturned := (err != nil)

		if tc.errExpected != errorReturned {
			t.Fatalf("Divide(%f): Expected error but got %v", tc.n, err)
		}
		if !tc.errExpected && !closeEnough(tc.want, got) {
			t.Errorf("Divide(%f): want %f, got %f", tc.n, tc.want, got)
		}
	}
}

func closeEnough(a, b float64) bool {
	return math.Abs(a-b) < 0.1
}

func TestSqrt(t *testing.T) {
	testCases := []struct{
		input       float64
		want        float64
		errExpected bool
	}{
		{ input: 4, want: 2, errExpected: false },
		{ input: 16, want: 4, errExpected: false },
		{ input: 4.84, want: 2.2, errExpected: false },
		{ input: -100, want: 999, errExpected: false },
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.input)
		errorReturned := (err != nil)
		if tc.errExpected != errorReturned {
			t.Fatalf("Sqrt(%f): unexpected error status %v", tc.input, err)
		}
		if !tc.errExpected && !closeEnough(tc.want, got) {
			t.Errorf("Sqrt(%f): want %f, got %f", tc.input, tc.want, got)
		}
	}
}
