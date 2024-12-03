package solutions

import "testing"

func TestDay02TestCountSafeReports(t *testing.T) {
	// Example data from the task
	reports := [][]int{
		{7, 6, 4, 2, 1}, // Decreasing, valid differences
		{1, 2, 7, 8, 9}, // Increasing, invalid differences
		{9, 7, 6, 2, 1}, // Decreasing, invalid differences
		{1, 3, 2, 4, 5}, // Mixed, invalid (not monotonic)
		{8, 6, 4, 4, 1}, // Mixed, invalid (difference < 1)
		{1, 3, 6, 7, 9}, // Increasing, valid differences
	}

	expected := 2 // Only 2 safe reports: {7, 6, 4, 2, 1} and {1, 3, 6, 7, 9}
	actual := countSafeReports(reports)

	if actual != expected {
		t.Errorf("Expected %d safe reports, but got %d", expected, actual)
	}
}

func TestDay02TestCountSafeReportsWithProblemDampener(t *testing.T) {
	// Example data from the task
	reports := [][]int{
		{7, 6, 4, 2, 1}, // safe
		{1, 2, 7, 8, 9}, // unsafe
		{9, 7, 6, 2, 1}, // unsafe
		{1, 3, 2, 4, 5}, // safe
		{8, 6, 4, 4, 1}, // safe
		{1, 3, 6, 7, 9}, // safe
	}

	expected := 4
	actual := countSafeReportsWithProblemDampener(reports)

	if actual != expected {
		t.Errorf("Expected %d safe reports, but got %d", expected, actual)
	}
}
