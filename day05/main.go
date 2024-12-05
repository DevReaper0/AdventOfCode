package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cast"
	flag "github.com/spf13/pflag"
)

//go:embed input.dat
var input string

type Pair struct {
	first  int
	second int
}

func part1(input string) int {
	rules, updates := parseInput(input)

	var correctlyOrderedUpdates [][]int
	for _, update := range updates {
		if isCorrectlyOrdered(update, rules) {
			correctlyOrderedUpdates = append(correctlyOrderedUpdates, update)
		}
	}

	middleSum := 0
	for _, update := range correctlyOrderedUpdates {
		middleIndex := len(update) / 2
		middleSum += update[middleIndex]
	}

	return middleSum
}

func part2(input string) int {
	rules, updates := parseInput(input)

	var reorderedUpdates [][]int
	for _, update := range updates {
		if !isCorrectlyOrdered(update, rules) {
			reorderedUpdates = append(reorderedUpdates, reorderUpdate(update, rules))
		}
	}

	middleSum := 0
	for _, update := range reorderedUpdates {
		middleIndex := len(update) / 2
		middleSum += update[middleIndex]
	}

	return middleSum
}

func isCorrectlyOrdered(update []int, rules []Pair) bool {
	position := make(map[int]int)
	for i, page := range update {
		position[page] = i
	}

	for _, rule := range rules {
		if pos1, ok := position[rule.first]; ok {
			if pos2, ok := position[rule.second]; ok {
				if pos1 > pos2 {
					return false
				}
			}
		}
	}

	return true
}

func reorderUpdate(update []int, rules []Pair) []int {
	sort.Slice(update, func(i, j int) bool {
		return comesBefore(update[i], update[j], rules)
	})
	return update
}

func comesBefore(a int, b int, rules []Pair) bool {
	for _, rule := range rules {
		if rule.first == a && rule.second == b {
			return true
		}
		if rule.first == b && rule.second == a {
			return false
		}
	}

	return a < b
}

func parseInput(input string) ([]Pair, [][]int) {
	var rules []Pair
	var updates [][]int

	parts := strings.SplitN(input, "\n\n", 2)

	for _, line := range strings.Split(parts[0], "\n") {
		split := strings.SplitN(line, "|", 2)
		rules = append(rules, Pair{cast.ToInt(split[0]), cast.ToInt(split[1])})
	}

	for _, line := range strings.Split(parts[1], "\n") {
		var pages []int
		split := strings.Split(line, ",")
		for _, page := range split {
			pages = append(pages, cast.ToInt(page))
		}
		updates = append(updates, pages)
	}

	return rules, updates
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
