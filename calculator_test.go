package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := calculator.Add(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(4, 2)
	if want != got {
		t.Errorf("Subtract(%f, %f): want %f, got %f", 4.0, 2.0, want, got)
	}
}

type TestCase struct {
	name string
	n    []float64
	f    func(n ...float64) float64
	want float64
}

func TestCalculator(t *testing.T) {
	t.Parallel()
	testCases := []TestCase{
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
			n:    []float64{4, 20},
			f:    calculator.Multiply,
			want: 80,
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
			got := tc.f(tc.n...)
			if tc.want != got {
				t.Errorf("want %f, got %f", tc.want, got)
			}
		})
	}
}

type TestCaseDivide struct {
	n           []float64
	want        float64
	errExpected bool
}

func TestDivide(t *testing.T) {
	testCases := []TestCaseDivide{
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
		if !tc.errExpected && tc.want != got {
			t.Errorf("Divide(%f): want %f, got %f", tc.n, tc.want, got)
		}

	}
}
