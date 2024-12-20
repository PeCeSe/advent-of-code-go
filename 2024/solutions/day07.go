package solutions

import (
	"fmt"
	"github.com/PeCeSe/advent-of-code-go/utils"
	"log"
	"strconv"
	"strings"
)

func Day07() {

	lines, err := utils.ReadFile("2024/input/day07.txt")
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}

	targets, equations := parseInputDay7(lines)
	totalCalibration := 0

	for i, target := range targets {
		if isValid(target, equations[i]) {
			totalCalibration += target
		}
	}

	fmt.Println("Day 7, Task 1. Sum of calibration result", totalCalibration)

	precomputed := precomputeCombinations(11)

	totalCalibrationWithExtraOperator := 0
	for i, target := range targets {
		if isValidWithConcat(target, equations[i], precomputed) {
			totalCalibrationWithExtraOperator += target
		}
	}

	fmt.Println("Day 7, Task 2. Sum of calibration with extra operator", totalCalibrationWithExtraOperator)

}

func parseInputDay7(lines []string) ([]int, [][]int) {
	var targets []int
	var equations [][]int

	for _, line := range lines {
		parts := strings.Split(line, ":")
		target, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		targets = append(targets, target)

		operands := strings.Fields(parts[1])
		var nums []int
		for _, operand := range operands {
			num, _ := strconv.Atoi(operand)
			nums = append(nums, num)
		}
		equations = append(equations, nums)
	}

	return targets, equations
}

func generateCombinations(n int) [][]string {
	if n == 0 {
		return [][]string{{}} // Base case: no operators needed
	}

	var results [][]string
	subCombos := generateCombinations(n - 1)
	for _, combo := range subCombos {
		// Create new slices to avoid modifying existing ones
		newComboAdd := append([]string{}, combo...)
		newComboMul := append([]string{}, combo...)

		// Add operators
		results = append(results, append(newComboAdd, "+"))
		results = append(results, append(newComboMul, "*"))
	}

	return results
}

func evaluate(operands []int, operators []string) int {
	if len(operators) != len(operands)-1 {
		log.Fatalf("Mismatched lengths: operands=%v, operators=%v\n", operands, operators)
	}

	result := operands[0]
	for i := 1; i < len(operands); i++ {
		if operators[i-1] == "+" {
			result += operands[i]
		} else {
			result *= operands[i]
		}
	}
	return result
}

func isValid(target int, operands []int) bool {
	n := len(operands) - 1
	combinations := generateCombinations(n)

	// fmt.Printf("Operators generated for %d operands: %v\n", len(operands), combinations)

	for _, operators := range combinations {
		// fmt.Printf("Testing: %v with operators %v\n", operands, operators)
		if evaluate(operands, operators) == target {
			return true
		}
	}
	return false
}

func precomputeCombinations(maxOperators int) map[int][][]string {
	combinations := make(map[int][][]string)
	for n := 1; n <= maxOperators; n++ {
		combinations[n] = generateCombinationsWithConcat(n)
	}
	return combinations
}

// Generate combinations including "||"
func generateCombinationsWithConcat(n int) [][]string {
	if n == 0 {
		return [][]string{{}}
	}

	var results [][]string
	subCombos := generateCombinationsWithConcat(n - 1)
	for _, combo := range subCombos {
		results = append(results, append(append([]string{}, combo...), "+"))
		results = append(results, append(append([]string{}, combo...), "*"))
		results = append(results, append(append([]string{}, combo...), "||"))
	}

	return results
}

func evaluateWithConcat(operands []int, operators []string) int {
	result := operands[0]
	for i := 1; i < len(operands); i++ {
		switch operators[i-1] {
		case "+":
			result += operands[i]
		case "*":
			result *= operands[i]
		case "||":
			concatValue, _ := strconv.Atoi(fmt.Sprintf("%d%d", result, operands[i]))
			result = concatValue
		}
	}
	return result
}

func isValidWithConcat(target int, operands []int, precomputed map[int][][]string) bool {
	n := len(operands) - 1
	combinations := precomputed[n]

	for _, operators := range combinations {
		if evaluateWithConcat(operands, operators) == target {
			return true
		}
	}

	return false
}
