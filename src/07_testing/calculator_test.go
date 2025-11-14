package calculator

import (
	"testing"
)

func TestCalculator_Add(t *testing.T) {
	// Table-driven test
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed numbers", -2, 3, 1},
		{"zero", 0, 0, 0},
	}

	calc := &Calculator{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%f, %f) = %f; want %f", 
					tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculator_Divide(t *testing.T) {
	tests := []struct {
		name        string
		a, b        float64
		expected    float64
		expectError bool
	}{
		{"valid division", 6, 2, 3, false},
		{"divide by zero", 6, 0, 0, true},
		{"negative numbers", -6, -2, 3, false},
	}

	calc := &Calculator{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.Divide(tt.a, tt.b)
			
			// Check error expectation
			if tt.expectError && err == nil {
				t.Error("expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			
			// If no error expected, check result
			if !tt.expectError && result != tt.expected {
				t.Errorf("Divide(%f, %f) = %f; want %f", 
					tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestCalculator_Memory(t *testing.T) {
	calc := &Calculator{}

	// Test initial memory value
	if mem := calc.Memory(); mem != 0 {
		t.Errorf("initial memory = %f; want 0", mem)
	}

	// Test storing and retrieving value
	calc.Store(42)
	if mem := calc.Memory(); mem != 42 {
		t.Errorf("memory after Store(42) = %f; want 42", mem)
	}

	// Test clearing memory
	calc.Clear()
	if mem := calc.Memory(); mem != 0 {
		t.Errorf("memory after Clear() = %f; want 0", mem)
	}
}

func BenchmarkCalculator_Add(b *testing.B) {
	calc := &Calculator{}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		calc.Add(float64(i), float64(i))
	}
}

// Example test
func ExampleCalculator_Add() {
	calc := &Calculator{}
	result := calc.Add(2, 3)
	fmt.Printf("2 + 3 = %v\n", result)
	// Output: 2 + 3 = 5
}