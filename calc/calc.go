package calc

import (
	"errors"
	"math"
)

// Add returns the sum of two numbers
func Add(a, b float64) float64 {
	return a + b
}

// Subtract returns the difference of two numbers
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply returns the product of two numbers
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide returns the quotient of two numbers
// Returns an error if attempting to divide by zero
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Modulo returns the remainder of a divided by b
// Returns an error if b is zero
func Modulo(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("modulo by zero")
	}
	return float64(int(a) % int(b)), nil
}

// Power returns base raised to the power of exp
func Power(base, exp float64) float64 {
	return math.Pow(base, exp)
}
