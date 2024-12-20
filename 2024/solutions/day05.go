package solutions

import (
	"fmt"
	"github.com/PeCeSe/advent-of-code-go/utils"
	"log"
	"strconv"
	"strings"
)

func Day05() {

	lines, err := utils.ReadFile("2024/input/day05.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	fmt.Println("Day 5, Task 1. Sum of all middle pages of valid manuals ", calculateMiddleSum(parseInput(lines)))
	fmt.Println("Day 5, Task 2. Sum of all middle pages of invalid manuals after fixing", correctAndSumMiddlePages(parseInput(lines)))

}

func parseInput(input []string) (rules [][2]int, updates [][]int) {
	// Split into rules and updates
	divider := 0
	for i, line := range input {
		if line == "" { // Blank line separates rules and updates
			divider = i
			break
		}
	}

	// Parse rules
	for _, line := range input[:divider] {
		parts := strings.Split(line, "|")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		rules = append(rules, [2]int{x, y})
	}

	// Parse updates
	for _, line := range input[divider+1:] {
		parts := strings.Split(line, ",")
		var update []int
		for _, p := range parts {
			num, _ := strconv.Atoi(p)
			update = append(update, num)
		}
		updates = append(updates, update)
	}

	return rules, updates
}

func isUpdateValid(rules [][2]int, update []int) bool {
	// Create a map to store the position of each page in the update
	position := make(map[int]int)
	for i, page := range update {
		position[page] = i
	}

	// Check all rules
	for _, rule := range rules {
		x, y := rule[0], rule[1]
		posX, existsX := position[x]
		posY, existsY := position[y]

		// If both pages are in the update, ensure the order is valid
		if existsX && existsY && posX >= posY {
			return false
		}
	}

	return true
}

func calculateMiddleSum(rules [][2]int, updates [][]int) int {
	sum := 0

	for _, update := range updates {
		if isUpdateValid(rules, update) {
			// Find the middle page
			middle := update[len(update)/2]
			sum += middle
		}
	}

	return sum
}

func correctAndSumMiddlePages(rules [][2]int, updates [][]int) int {
	sum := 0

	for _, update := range updates {
		if !isUpdateValid(rules, update) {
			fmt.Println("Incorrect update", update)
			// Correct the update
			corrected := reorderUpdate(rules, update)
			fmt.Println("Corrected update", corrected)
			// Find the middle page
			middle := corrected[len(corrected)/2]
			fmt.Println("Middle page", middle)
			sum += middle
		}
	}

	return sum
}

func reorderUpdate(rules [][2]int, update []int) []int {
	// Build a dependency graph for this update
	dependencies := make(map[int][]int)
	pageSet := make(map[int]bool)

	for _, page := range update {
		pageSet[page] = true
	}

	for _, rule := range rules {
		if pageSet[rule[0]] && pageSet[rule[1]] {
			dependencies[rule[1]] = append(dependencies[rule[1]], rule[0])
		}
	}

	// Perform topological sort
	visited := make(map[int]bool)
	stack := []int{}
	for _, page := range update {
		if !visited[page] {
			topologicalSort(page, dependencies, visited, &stack)
		}
	}

	return stack
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func topologicalSort(page int, dependencies map[int][]int, visited map[int]bool, stack *[]int) {
	visited[page] = true

	for _, dep := range dependencies[page] {
		if !visited[dep] {
			topologicalSort(dep, dependencies, visited, stack)
		}
	}

	*stack = append(*stack, page)
}
