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
	method string
	got    interface{}
	want   float64
}

func TestCalculator(t *testing.T) {
	t.Parallel()
	testCases := []TestCase{
		{
			method: "Add",
			got:    calculator.Add(2, 2),
			want:   4,
		},
		{
			method: "Subtract",
			got:    calculator.Subtract(3, 5),
			want:   -2,
		},
		{
			method: "Multiply",
			got:    calculator.Multiply(4, 20),
			want:   80,
		},
	}

	for _, tc := range testCases {
		if tc.want != tc.got {
			t.Errorf("%s: want %f, got %f", tc.method, tc.want, tc.got)
		}
	}
}
