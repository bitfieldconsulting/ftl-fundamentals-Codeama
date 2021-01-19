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
	A    float64
	B    float64
	want float64
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []TestCase{
		{
			A:    2,
			B:    2,
			want: 4,
		},
		{
			A:    -3,
			B:    3,
			want: -9,
		},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.A, tc.B)
		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f", 2.0, 2.0, tc.want, got)
		}
	}

}
