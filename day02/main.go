package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/spf13/cast"
	flag "github.com/spf13/pflag"
)

//go:embed input.dat
var input string

func part1(input string) int {
	parsed := parseInput(input)

	safe := 0
	for _, report := range parsed {
		if testSafe(report, -1) {
			safe++
		}
	}

	return safe
}

func part2(input string) int {
	parsed := parseInput(input)

	safe := 0
	for _, report := range parsed {
		if testSafe(report, -1) {
			safe++
		} else {
			for i := 0; i < len(report); i++ {
				if testSafe(report, i) {
					safe++
					break
				}
			}
		}
	}

	return safe
}

func testSafe(report []int, excludedIndex int) bool {
	startingIndex := 1
	if excludedIndex == 0 {
		startingIndex = 2
	}

	direction := ""
	for i := startingIndex; i < len(report); i++ {
		if i == excludedIndex {
			continue
		}

		previousIndex := i - 1
		if previousIndex == excludedIndex {
			previousIndex--
		}

		if report[i] > report[previousIndex] {
			if direction == "decreasing" || report[i]-report[previousIndex] < 1 || report[i]-report[previousIndex] > 3 {
				return false
			}
			direction = "increasing"
		} else if report[i] < report[previousIndex] {
			if direction == "increasing" || report[previousIndex]-report[i] < 1 || report[previousIndex]-report[i] > 3 {
				return false
			}
			direction = "decreasing"
		} else {
			return false
		}
	}
	return true
}

func parseInput(input string) [][]int {
	var answer [][]int
	for _, line := range strings.Split(input, "\n") {
		var report []int
		for _, level := range strings.Split(line, " ") {
			report = append(report, cast.ToInt(level))
		}
		answer = append(answer, report)
	}
	return answer
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
