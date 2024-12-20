package solutions

import (
	"testing"
)

func TestDay5CalculateMiddleSum(t *testing.T) {
	// Example data
	rules := [][2]int{
		{47, 53}, {97, 13}, {97, 61}, {97, 47}, {75, 29}, {61, 13},
		{75, 53}, {29, 13}, {97, 29}, {53, 29}, {61, 53}, {97, 53},
		{61, 29}, {47, 13}, {75, 47}, {97, 75}, {47, 61}, {75, 61},
		{47, 29}, {75, 13}, {53, 13},
	}
	updates := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}

	expectedSum := 143
	actualSum := calculateMiddleSum(rules, updates)

	if actualSum != expectedSum {
		t.Errorf("Expected %d, but got %d", expectedSum, actualSum)
	}
}

func TestDay5CorrectAndSumMiddlePages(t *testing.T) {
	// Example data
	rules := [][2]int{
		{47, 53}, {97, 13}, {97, 61}, {97, 47}, {75, 29}, {61, 13},
		{75, 53}, {29, 13}, {97, 29}, {53, 29}, {61, 53}, {97, 53},
		{61, 29}, {47, 13}, {75, 47}, {97, 75}, {47, 61}, {75, 61},
		{47, 29}, {75, 13}, {53, 13},
	}
	updates := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
		{75, 97, 47, 61, 53}, // Incorrect
		{61, 13, 29},         // Incorrect
		{97, 13, 75, 29, 47}, // Incorrect
	}

	expectedSum := 123
	actualSum := correctAndSumMiddlePages(rules, updates)

	if actualSum != expectedSum {
		t.Errorf("Expected %d, but got %d", expectedSum, actualSum)
	}
}
