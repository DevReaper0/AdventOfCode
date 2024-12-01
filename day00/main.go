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
	_ = parseInput(input)

	return 0
}

func part2(input string) int {
	_ = parseInput(input)

	return 0
}

func parseInput(input string) []int {
	var answer []int
	for _, line := range strings.Split(input, "\n") {
		answer = append(answer, cast.ToInt(line))
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
