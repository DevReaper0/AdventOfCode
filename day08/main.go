package main

import (
	_ "embed"
	"fmt"
	"image"
	"strings"

	flag "github.com/spf13/pflag"
)

//go:embed input.dat
var input string

func part1(input string) int {
	bounds, frequency := parseInput(input)
	antiNodes := map[image.Point]bool{}

	for _, antennas := range frequency {
		for _, antennaOne := range antennas {
			for _, antennaTwo := range antennas {
				if antennaOne == antennaTwo {
					continue
				}

				if antenna := antennaTwo.Add(antennaTwo.Sub(antennaOne)); bounds[antenna] {
					antiNodes[antenna] = true
				}
			}
		}
	}

	return len(antiNodes)
}

func part2(input string) int {
	bounds, frequency := parseInput(input)
	antiNodes := map[image.Point]bool{}

	for _, antennas := range frequency {
		for _, antennaOne := range antennas {
			for _, antennaTwo := range antennas {
				if antennaOne == antennaTwo {
					continue
				}

				for diff := antennaTwo.Sub(antennaOne); bounds[antennaTwo]; antennaTwo = antennaTwo.Add(diff) {
					antiNodes[antennaTwo] = true
				}
			}
		}
	}

	return len(antiNodes)
}

func parseInput(input string) (map[image.Point]bool, map[rune][]image.Point) {
	bounds := map[image.Point]bool{}
	frequencies := map[rune][]image.Point{}

	for y, line := range strings.Split(input, "\n") {
		for x, char := range line {
			bounds[image.Point{x, y}] = true
			if char != '.' {
				frequencies[char] = append(frequencies[char], image.Point{x, y})
			}
		}
	}

	return bounds, frequencies
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
