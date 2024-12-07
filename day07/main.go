package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"

	"github.com/spf13/cast"
	flag "github.com/spf13/pflag"
)

//go:embed input.dat
var input string

func part1(input string) int {
	answers, allValues := parseInput(input)
	sum := 0

	for i, values := range allValues {
		if solve(values, answers[i], false) {
			sum += answers[i]
		}
	}

	return sum
}

func part2(input string) int {
	answers, allValues := parseInput(input)
	sum := 0

	for i, values := range allValues {
		if solve(values, answers[i], true) {
			sum += answers[i]
		}
	}

	return sum
}

func solve(values []int, answer int, includeConcat bool) bool {
	numOperators := len(values) - 1
	base := 2
	if includeConcat {
		base = 3
	}
	maxCombinations := int(math.Pow(float64(base), float64(numOperators)))

	for combination := 0; combination < maxCombinations; combination++ {
		operators := make([]string, numOperators)
		currentCombination := combination
		for j := 0; j < numOperators; j++ {
			if currentCombination%base == 0 {
				operators[j] = "+"
			} else if currentCombination%base == 1 {
				operators[j] = "*"
			} else if includeConcat && currentCombination%base == 2 {
				operators[j] = "||"
			}
			currentCombination /= base
		}

		result := values[0]
		for j := 0; j < numOperators; j++ {
			if operators[j] == "+" {
				result += values[j+1]
			} else if operators[j] == "*" {
				result *= values[j+1]
			} else if operators[j] == "||" {
				result = cast.ToInt(cast.ToString(result) + cast.ToString(values[j+1]))
			}
		}

		if result == answer {
			return true
		}
	}

	return false
}

func parseInput(input string) ([]int, [][]int) {
	var answers []int
	var allValues [][]int

	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, " ")
		answers = append(answers, cast.ToInt(strings.TrimRight(split[0], ":")))
		var values []int
		for _, value := range split[1:] {
			values = append(values, cast.ToInt(value))
		}
		allValues = append(allValues, values)
	}
	return answers, allValues
}

func init() {
	// This is done in init instead of main so the tests have the same input
	input = strings.TrimRight(input, "\n")
}

func main() {
	var part int

	flag.IntVarP(&part, "part", "p", 1, "part 1 or 2")
	flag.Parse()

	fmt.Println("Running part", part)
	fmt.Println()

	if part == 1 {
		fmt.Println("Output:", part1(input))
	} else {
		fmt.Println("Output:", part2(input))
	}
}
