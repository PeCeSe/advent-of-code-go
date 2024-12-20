package solutions

import (
	"fmt"
	"github.com/PeCeSe/advent-of-code-go/utils"
	"log"
)

func Day04() {

	lines, err := utils.ReadFile("2024/input/day04.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	runeMatrix := convertToMatrix(lines)

	fmt.Println("Day 4, Task 1. Number of XMAS occurrences: ", countWordOccurrences(runeMatrix, "XMAS"))
	fmt.Println("Day 4, Task 2. Number of MAS crosses: ", countMasCross(runeMatrix))

}

func convertToMatrix(lines []string) [][]rune {
	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = []rune(line)
	}
	return matrix
}

func searchWord(grid [][]rune, word string, row, col int, direction [2]int) bool {
	rows, cols := len(grid), len(grid[0])
	for i, char := range word {
		r, c := row+direction[0]*i, col+direction[1]*i
		if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] != char {
			return false
		}
	}
	return true
}

func countWordOccurrences(grid [][]rune, word string) int {
	count := 0
	directions := [][2]int{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
		{1, 1}, {-1, -1}, {1, -1}, {-1, 1},
	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			for _, dir := range directions {
				if searchWord(grid, word, row, col, dir) {
					count++
				}
			}
		}
	}
	return count
}

func isMasCross(grid [][]rune, row, col int) bool {
	directions := [][2]int{
		{1, 1},  // Top-left to bottom-right
		{-1, 1}, // Bottom-left to top-right
	}

	word1 := "MAS"
	word2 := "SAM"

	// Check both diagonals
	diag1MAS := searchWord(grid, word1, row-1, col-1, directions[0])
	diag1SAM := searchWord(grid, word2, row-1, col-1, directions[0])
	diag2MAS := searchWord(grid, word1, row+1, col-1, directions[1])
	diag2SAM := searchWord(grid, word2, row+1, col-1, directions[1])

	// Combine results for X-MAS validation
	return (diag1MAS && diag2MAS) || // Both diagonals contain "MAS"
		(diag1SAM && diag2SAM) || // Both diagonals contain "SAM"
		(diag1MAS && diag2SAM) || // Mixed 1
		(diag1SAM && diag2MAS) // Mixed 2
}

func countMasCross(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			if isMasCross(grid, r, c) {
				count++
			}
		}
	}

	return count
}
