package calculator

import (
	"errors"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -1, -2, -3},
		{"add zero", 5, 0, 5},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Add(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Add(%v, %v) = %v; want %v", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5, 3, 2},
		{"negative result", 3, 5, -2},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Subtract(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Subtract(%v, %v) = %v; want %v", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 4, 3, 12},
		{"multiply by zero", 5, 0, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Multiply(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Multiply(%v, %v) = %v; want %v", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	t.Run("positive numbers", func(t *testing.T) {
		result, err := Divide(10, 2)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result != 5 {
			t.Errorf("Divide(10, 2) = %v; want 5", result)
		}
	})

	t.Run("divide by zero", func(t *testing.T) {
		_, err := Divide(10, 0)
		if !errors.Is(err, ErrDivisionByZero) {
			t.Errorf("expected ErrDivisionByZero, got %v", err)
		}
	})

	t.Run("decimal result", func(t *testing.T) {
		result, err := Divide(5, 2)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if result != 2.5 {
			t.Errorf("Divide(5, 2) = %v; want 2.5", result)
		}
	})
}
