package solutions

import (
	"fmt"
	"github.com/PeCeSe/advent-of-code-go/utils"
	"log"
	"sort"
	"strconv"
	"strings"
)

func Day01() {
	// Read file using utils
	lines, err := utils.ReadFile("2024/input/day01.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	// Parse the lines into lists
	list1, list2, err := parseLines(lines)
	if err != nil {
		log.Fatalf("Failed to parse input: %v", err)
	}

	// Output the results of the first task
	fmt.Println("Day 1, Task 1. Total difference: ", calculateTotalDifference(list1, list2))
	// Output the results of the second task
	fmt.Println("Day 1, Task 2. Similarity score: ", calculateSimilarityScore(list1, list2))
}

func parseLines(lines []string) ([]int, []int, error) {
	var list1, list2 []int
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid input format: %s", line)
		}

		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			return nil, nil, fmt.Errorf("failed to parse numbers: %v, %v", err1, err2)
		}

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}
	return list1, list2, nil
}

func calculateTotalDifference(list1, list2 []int) int {
	// Ensure both lists are sorted
	sort.Ints(list1)
	sort.Ints(list2)

	// Calculate the total difference
	totalDifference := 0
	for i := 0; i < len(list1) && i < len(list2); i++ {
		diff := list1[i] - list2[i]
		if diff < 0 {
			diff = -diff // Absolute difference
		}
		totalDifference += diff
	}
	return totalDifference
}

func calculateSimilarityScore(list1, list2 []int) int {
	counts := make(map[int]int)
	for _, num := range list2 {
		counts[num]++
	}

	similarityScore := 0
	for _, num := range list1 {
		occurrences := counts[num]
		similarityScore += num * occurrences
	}

	return similarityScore
}
