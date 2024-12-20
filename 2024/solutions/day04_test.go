package solutions

import (
	"testing"
)

func TestDay04TestCountXMAS(t *testing.T) {
	// Example grid
	grid := [][]rune{
		{'.', 'M', '.', 'S', '.', '.', '.', '.', '.', '.'},
		{'.', '.', 'A', '.', '.', 'M', 'S', 'M', 'S', '.'},
		{'.', 'M', '.', 'S', '.', 'M', 'A', 'A', '.', '.'},
		{'.', '.', 'A', '.', 'A', 'S', 'M', 'S', 'M', '.'},
		{'.', 'M', '.', 'S', '.', 'M', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'S', '.', 'S', '.', 'S', '.', 'S', '.', 'S', '.'},
		{'.', 'A', '.', 'A', '.', 'A', '.', 'A', '.', '.'},
		{'M', '.', 'M', '.', 'M', '.', 'M', '.', 'M', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}

	expected := 9
	actual := countMasCross(grid)

	if actual != expected {
		t.Errorf("Expected %d X-MAS patterns, but got %d", expected, actual)
	}
}
