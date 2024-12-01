package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/spf13/cast"
	flag "github.com/spf13/pflag"
)

//go:embed input.dat
var input string

func part1(input string) int {
	left, right := parseInput(input)
	slices.Sort(left)
	slices.Sort(right)

	sum := 0
	for i := 0; i < len(left); i++ {
		distance := math.Abs(float64(left[i] - right[i]))
		sum += int(distance)
	}

	return sum
}

func part2(input string) int {
	left, right := parseInput(input)

	similarityScore := 0
	for i := 0; i < len(left); i++ {
		occurrences := 0
		for j := 0; j < len(right); j++ {
			if left[i] == right[j] {
				occurrences++
			}
		}

		similarityScore += left[i] * occurrences
	}

	return similarityScore
}

func parseInput(input string) ([]int, []int) {
	var left []int
	var right []int

	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, "   ")
		left = append(left, cast.ToInt(split[0]))
		right = append(right, cast.ToInt(split[1]))
	}
	return left, right
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
