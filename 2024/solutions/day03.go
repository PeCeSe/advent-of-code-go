package solutions

import (
	"fmt"
	"github.com/PeCeSe/advent-of-code-go/utils"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Day03() {

	lines, err := utils.ReadFile("2024/input/day03.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	// fmt.Println("Valid muls: ", extractValidMul(lines))

	fmt.Println("Day 3, Task 1. Sum of multiplications: ", sumMultiplications(extractValidMul(lines)))
	fmt.Println("Day 3, Task 2. Sum of do() multiplications: ", sumMultiplications(extractOnlyDoMuls(lines)))

}

func extractValidMul(lines []string) []string {
	// Compile the regex to match "mul(X,Y)" with 1-3 digit numbers
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	var matches []string
	for _, line := range lines {
		// Find all matches in the current line
		matches = append(matches, re.FindAllString(line, -1)...)
	}

	return matches
}

func extractMulDoDont(lines []string) []string {
	// Compile the regex to match "mul(X,Y)" with 1-3 digit numbers
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)

	var matches []string
	for _, line := range lines {
		// Find all matches in the current line
		matches = append(matches, re.FindAllString(line, -1)...)
	}

	return matches
}

func extractOnlyDoMuls(lines []string) []string {
	doMuls := extractMulDoDont(lines)

	enabled := true // Start with mul enabled
	validMuls := []string{}

	for _, mul := range doMuls {
		if mul == "do()" {
			enabled = true
		} else if mul == "don't()" {
			enabled = false
		} else if strings.HasPrefix(mul, "mul(") && enabled {
			validMuls = append(validMuls, mul)
		}
	}

	return validMuls
}

func sumMultiplications(muls []string) int {
	// Initialize the sum
	sum := 0

	// Compile the regex to match the numbers in the "mul" expressions
	re := regexp.MustCompile(`\d{1,3}`)

	// Iterate over the matches
	for _, mul := range muls {
		// Find all numbers in the current "mul" expression
		nums := re.FindAllString(mul, -1)

		// Parse the numbers
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])

		// Add the product to the sum
		sum += num1 * num2
	}

	return sum
}
