package solutions

import (
	"testing"
)

func TestDay7Examples(t *testing.T) {
	tests := []struct {
		target   int
		operands []int
		expected bool
	}{
		{190, []int{10, 19}, true},
		{3267, []int{81, 40, 27}, true},
		{292, []int{11, 6, 16, 20}, true},
		{500, []int{10, 20, 30}, false}, // Example of an invalid equation
	}

	for _, tt := range tests {
		result := isValid(tt.target, tt.operands)
		if result != tt.expected {
			t.Errorf("isValid(%d, %v) = %v; want %v", tt.target, tt.operands, result, tt.expected)
		}
	}
}

func TestDay7EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		target   int
		operands []int
		expected bool
	}{
		// Minimal input
		{"Minimal Input", 5, []int{5}, true},

		// Impossible target
		{"Impossible Target", 50, []int{2, 3, 4}, false},

		// Large numbers
		{"Large Numbers", 1000000000, []int{10000, 100000, 100}, false},

		// Multiple combinations
		{"Multiple Combinations", 15, []int{1, 2, 3, 4, 5}, true},

		// Left-to-right precedence
		{"Left-to-Right Precedence", 10, []int{2, 3, 2}, true},

		// Long operand list
		{"Long Operand List", 1000, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, false},
	}

	for _, tt := range tests {
		result := isValid(tt.target, tt.operands)
		if result != tt.expected {
			t.Errorf(
				"Test Case Failed: %s\nTarget: %d\nOperands: %v\nExpected: %v\nGot: %v\n",
				tt.name, tt.target, tt.operands, tt.expected, result,
			)
		}
	}
}
