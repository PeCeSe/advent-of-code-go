package solutions

import (
	"fmt"
	"github.com/PeCeSe/advent-of-code-go/utils"
	"log"
)

func Day06() {

	lines, err := utils.ReadFile("2024/input/day06.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	grid, start, facing := parseMap(lines)
	visited := simulateGuard(grid, start, facing)
	fmt.Println("")

	fmt.Println("Day 6, Task 1. Number of distinct positions visited: ", len(visited))

	loopPositions := findLoopCausingPositions(grid, start, facing)
	fmt.Println("Day 6, Task 2. Number of possible loops", loopPositions)

}

func parseMap(lines []string) ([][]rune, [2]int, rune) {
	var grid [][]rune
	var start [2]int
	var facing rune

	for r, line := range lines {
		row := []rune(line)
		for c, char := range row {
			if char == '^' || char == '>' || char == 'v' || char == '<' {
				start = [2]int{r, c}
				facing = char
				row[c] = '.' // Replace the guard's initial position with open space
			}
		}
		grid = append(grid, row)
	}
	return grid, start, facing
}

func simulateGuard(grid [][]rune, start [2]int, facing rune) map[[2]int]bool {
	visited := make(map[[2]int]bool)
	directions := map[rune][2]int{
		'^': {-1, 0}, // Up
		'>': {0, 1},  // Right
		'v': {1, 0},  // Down
		'<': {0, -1}, // Left
	}
	rightTurn := map[rune]rune{
		'^': '>',
		'>': 'v',
		'v': '<',
		'<': '^',
	}

	pos := start
	visited[pos] = true

	for {
		// Calculate the position in front of the guard
		dir := directions[facing]
		next := [2]int{pos[0] + dir[0], pos[1] + dir[1]}

		// Check if the guard is leaving the map
		if next[0] < 0 || next[0] >= len(grid) || next[1] < 0 || next[1] >= len(grid[0]) {
			break
		}

		// Check if there's an obstruction
		if grid[next[0]][next[1]] == '#' {
			// Turn right
			facing = rightTurn[facing]
		} else {
			// Move forward
			pos = next
			visited[pos] = true
		}
	}

	return visited
}

func simulateGuardWithObstruction(grid [][]rune, start [2]int, facing rune, obstruction [2]int) bool {
	visited := make(map[[3]int]int) // Track visited states with step count
	directions := map[rune][2]int{
		'^': {-1, 0}, // Up
		'>': {0, 1},  // Right
		'v': {1, 0},  // Down
		'<': {0, -1}, // Left
	}
	rightTurn := map[rune]rune{
		'^': '>',
		'>': 'v',
		'v': '<',
		'<': '^',
	}

	// Place the obstruction
	grid[obstruction[0]][obstruction[1]] = '#'

	pos := start
	visited[[3]int{pos[0], pos[1], int(facing)}] = 0
	steps := 0

	for {
		steps++

		// Calculate the position in front of the guard
		dir := directions[facing]
		next := [2]int{pos[0] + dir[0], pos[1] + dir[1]}

		// Check if the guard is leaving the map
		if next[0] < 0 || next[0] >= len(grid) || next[1] < 0 || next[1] >= len(grid[0]) {
			break
		}

		// Check if there's an obstruction
		if grid[next[0]][next[1]] == '#' {
			// Turn right
			facing = rightTurn[facing]
		} else {
			// Move forward
			pos = next
			state := [3]int{pos[0], pos[1], int(facing)}
			if step, exists := visited[state]; exists {
				// A loop is detected
				if steps-step > 1 {
					grid[obstruction[0]][obstruction[1]] = '.' // Reset obstruction
					return true
				}
			}
			visited[state] = steps
		}
	}

	// Reset the obstruction
	grid[obstruction[0]][obstruction[1]] = '.'
	return false
}

func findLoopCausingPositions(grid [][]rune, start [2]int, facing rune) int {
	loopPositions := 0

	for r, row := range grid {
		for c, char := range row {
			// Skip non-empty positions and the starting position
			if char != '.' || (r == start[0] && c == start[1]) {
				continue
			}

			// Check if this position causes a loop
			if simulateGuardWithObstruction(grid, start, facing, [2]int{r, c}) {
				loopPositions++
			}
		}
	}

	return loopPositions
}
