package calc

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"positive numbers", 5, 3, 8},
		{"negative numbers", -5, -3, -8},
		{"mixed signs", 5, -3, 2},
		{"with zero", 0, 5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.a, tt.b); got != tt.want {
				t.Errorf("Add(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"positive numbers", 5, 3, 2},
		{"negative numbers", -5, -3, -2},
		{"mixed signs", 5, -3, 8},
		{"with zero", 5, 0, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subtract(tt.a, tt.b); got != tt.want {
				t.Errorf("Subtract(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"positive numbers", 5, 3, 15},
		{"negative numbers", -5, -3, 15},
		{"mixed signs", 5, -3, -15},
		{"with zero", 5, 0, 0},
		{"with one", 5, 1, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiply(tt.a, tt.b); got != tt.want {
				t.Errorf("Multiply(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a, b    float64
		want    float64
		wantErr bool
	}{
		{"positive numbers", 6, 3, 2, false},
		{"negative numbers", -6, -3, 2, false},
		{"mixed signs", 6, -3, -2, false},
		{"with one", 5, 1, 5, false},
		{"division by zero", 5, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Divide(%v, %v) error = %v, wantErr %v", tt.a, tt.b, err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Divide(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestModulo(t *testing.T) {
	tests := []struct {
		name    string
		a, b    float64
		want    float64
		wantErr bool
	}{
		{"positive numbers", 10, 3, 1, false},
		{"exact division", 9, 3, 0, false},
		{"negative dividend", -10, 3, -1, false},
		{"negative divisor", 10, -3, 1, false},
		{"both negative", -10, -3, -1, false},
		{"modulo by zero", 5, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Modulo(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Modulo(%v, %v) error = %v, wantErr %v", tt.a, tt.b, err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Modulo(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		name     string
		base, exp float64
		want     float64
	}{
		{"positive base and exponent", 2, 3, 8},
		{"base to power of zero", 5, 0, 1},
		{"base to power of one", 5, 1, 5},
		{"negative exponent", 2, -2, 0.25},
		{"fractional exponent", 4, 0.5, 2},
		{"zero to positive power", 0, 5, 0},
		{"negative base and even exponent", -2, 2, 4},
		{"negative base and odd exponent", -2, 3, -8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Power(tt.base, tt.exp); got != tt.want {
				t.Errorf("Power(%v, %v) = %v, want %v", tt.base, tt.exp, got, tt.want)
			}
		})
	}
}
