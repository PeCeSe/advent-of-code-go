package solutions

import (
	"fmt"
	"github.com/PeCeSe/advent-of-code-go/utils"
	"log"
	"strconv"
	"strings"
)

func Day02() {

	lines, err := utils.ReadFile("2024/input/day02.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	reports, err := parseReports(lines)

	fmt.Println("Day 2, Task 1. Number of safe reports: ", countSafeReports(reports))
	fmt.Println("Day 2, Task 2. Number of safe reports with problem dampener: ", countSafeReportsWithProblemDampener(reports))

}

func parseReports(lines []string) ([][]int, error) {
	var reports [][]int
	for _, line := range lines {
		var report []int
		levels := strings.Fields(line) // Split by spaces
		for _, level := range levels {
			num, err := strconv.Atoi(level) // Convert to integer
			if err != nil {
				return nil, fmt.Errorf("failed to parse level '%s': %v", level, err)
			}
			report = append(report, num)
		}
		reports = append(reports, report)
	}
	return reports, nil
}

func isReportSafe(report []int) bool {
	asc, desc := false, false
	isSafe := true

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		if utils.Abs(diff) < 1 || utils.Abs(diff) > 3 {
			isSafe = false
			break
		}

		if diff > 0 {
			asc = true
		} else if diff < 0 {
			desc = true
		}

		if asc && desc {
			isSafe = false
			break
		}
	}
	return isSafe
}

func countSafeReports(reports [][]int) int {
	safeReports := 0

	for _, report := range reports {
		if isReportSafe(report) {
			safeReports++
		}
	}

	return safeReports
}

func countSafeReportsWithProblemDampener(reports [][]int) int {
	safeReports := 0

	for _, report := range reports {
		if isReportSafe(report) {
			safeReports++
		} else {
			// Generate all variants with one level removed
			variants := generateVariants(report)

			// Check if any variant is safe
			for _, variant := range variants {
				if isReportSafe(variant) {
					safeReports++
					break
				}
			}
		}
	}

	return safeReports
}

func generateVariants(report []int) [][]int {
	variants := [][]int{}
	for i := 0; i < len(report); i++ {
		// Create a new slice with one level removed
		variant := append([]int{}, report[:i]...)  // Copy elements before index i
		variant = append(variant, report[i+1:]...) // Append elements after index i
		variants = append(variants, variant)
	}
	return variants
}
