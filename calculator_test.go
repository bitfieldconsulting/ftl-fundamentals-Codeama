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
	a, b float64
	f    func(float64, float64) float64
	want float64
}

func TestCalculator(t *testing.T) {
	t.Parallel()
	testCases := []TestCase{
		{
			name: "Add two positive integers",
			a:    2, b: 2,
			f:    calculator.Add,
			want: 4,
		},
		{
			name: "Add two negative integers",
			a:    -2, b: -2,
			f:    calculator.Add,
			want: -4,
		},
		{
			name: "Add a positive integer and a negative fraction ",
			a:    5, b: -2.48,
			f:    calculator.Add,
			want: 2.52,
		},
		{
			name: "Subtract a larger integer from a smaller integer",
			a:    3, b: 5,
			f:    calculator.Subtract,
			want: -2,
		},
		{
			name: "Subtract a negative integer from a negative integer",
			a:    -9, b: -5,
			f:    calculator.Subtract,
			want: -4,
		},
		{
			name: "Multiply two positive integers",
			a:    4, b: 20,
			f:    calculator.Multiply,
			want: 80,
		},
		{
			name: "Multiply a positive fraction by a negative fraction",
			a:    4, b: -9,
			f:    calculator.Multiply,
			want: -36,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.f(tc.a, tc.b)
			if tc.want != got {
				t.Errorf("want %f, got %f", tc.want, got)
			}
		})
	}
}

type TestCaseDivide struct {
	a, b        float64
	f           func(float64, float64) (float64, error)
	want        float64
	errExpected bool
}

func TestDivide(t *testing.T) {
	testCases := []TestCaseDivide{
		{
			a: 5, b: 0,
			f:           calculator.Divide,
			want:        0,
			errExpected: true,
		},
		{
			a: 10, b: 2.5,
			f:           calculator.Divide,
			want:        4,
			errExpected: false,
		},
		{
			a: 28, b: 2,
			f:           calculator.Divide,
			want:        14,
			errExpected: false,
		},
		{
			a: 1, b: 1,
			f:           calculator.Divide,
			want:        1,
			errExpected: false,
		},
		{
			a: 0, b: 2.5,
			f:           calculator.Divide,
			want:        0,
			errExpected: false,
		},
	}

	for _, tc := range testCases {
		got, err := tc.f(tc.a, tc.b)
		errorReturned := (err != nil)

		if tc.errExpected != errorReturned {
			t.Fatalf("Divide(%f, %f): Expected error but got %v", tc.a, tc.b, err)
		}
		if !tc.errExpected && tc.want != got {
			t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}

	}
}
