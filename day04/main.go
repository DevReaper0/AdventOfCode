package main

import (
	_ "embed"
	"fmt"
	"strings"

	flag "github.com/spf13/pflag"
)

//go:embed input.dat
var input string

func part1(input string) int {
	split := strings.Split(input, "\n")

	occurances := 0
	for i := 0; i < len(split); i++ {
		for j := 0; j < len(split[i]); j++ {
			if split[i][j] == 'X' {
				if i > 2 {
					// Check up
					if split[i-1][j] == 'M' && split[i-2][j] == 'A' && split[i-3][j] == 'S' {
						occurances++
					}

					// Check diagonal up left
					if j > 2 {
						if split[i-1][j-1] == 'M' && split[i-2][j-2] == 'A' && split[i-3][j-3] == 'S' {
							occurances++
						}
					}

					// Check diagonal up right
					if j < len(split[i])-3 {
						if split[i-1][j+1] == 'M' && split[i-2][j+2] == 'A' && split[i-3][j+3] == 'S' {
							occurances++
						}
					}
				}

				// Check left
				if j > 2 {
					if split[i][j-1] == 'M' && split[i][j-2] == 'A' && split[i][j-3] == 'S' {
						occurances++
					}
				}

				// Check right
				if j < len(split[i])-3 {
					if split[i][j+1] == 'M' && split[i][j+2] == 'A' && split[i][j+3] == 'S' {
						occurances++
					}
				}

				if i < len(split)-3 {
					// Check down
					if split[i+1][j] == 'M' && split[i+2][j] == 'A' && split[i+3][j] == 'S' {
						occurances++
					}

					// Check diagonal down left
					if j > 2 {
						if split[i+1][j-1] == 'M' && split[i+2][j-2] == 'A' && split[i+3][j-3] == 'S' {
							occurances++
						}
					}

					// Check diagonal down right
					if j < len(split[i])-3 {
						if split[i+1][j+1] == 'M' && split[i+2][j+2] == 'A' && split[i+3][j+3] == 'S' {
							occurances++
						}
					}
				}
			}
		}
	}

	return occurances
}

func part2(input string) int {
	split := strings.Split(input, "\n")

	occurances := 0
	for i := 1; i < len(split)-1; i++ {
		for j := 1; j < len(split[i])-1; j++ {
			if split[i][j] == 'A' {
				if split[i-1][j-1] == 'M' && split[i+1][j+1] == 'S' {
					if split[i-1][j+1] == 'M' && split[i+1][j-1] == 'S' {
						occurances++
					} else if split[i-1][j+1] == 'S' && split[i+1][j-1] == 'M' {
						occurances++
					}
				} else if split[i-1][j-1] == 'S' && split[i+1][j+1] == 'M' {
					if split[i-1][j+1] == 'M' && split[i+1][j-1] == 'S' {
						occurances++
					} else if split[i-1][j+1] == 'S' && split[i+1][j-1] == 'M' {
						occurances++
					}
				}
			}
		}
	}

	return occurances
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
