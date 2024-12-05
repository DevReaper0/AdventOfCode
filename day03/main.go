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
	sum := getSum(input, false)
	return sum
}

func part2(input string) int {
	sum := getSum(input, true)
	return sum
}

func getSum(input string, handleDisabling bool) int {
	sum := 0

	enabled := true
	for i := 0; i < len(input); i++ {
		if handleDisabling && input[i] == 'd' && i < len(input)-3 && input[i+1] == 'o' && input[i+2] == '(' && input[i+3] == ')' {
			i += 3
			enabled = true
		} else if handleDisabling && input[i] == 'd' && i < len(input)-6 && input[i+1] == 'o' && input[i+2] == 'n' && input[i+3] == '\'' && input[i+4] == 't' && input[i+5] == '(' && input[i+6] == ')' {
			i += 6
			enabled = false
		} else if input[i] == 'm' && i < len(input)-3 && input[i+1] == 'u' && input[i+2] == 'l' && input[i+3] == '(' {
			i += 4

			first := ""
			second := ""

			for input[i] >= '0' && input[i] <= '9' {
				first += string(input[i])
				i++
			}
			if input[i] == ',' {
				i++
			}
			for input[i] >= '0' && input[i] <= '9' {
				second += string(input[i])
				i++
			}

			if enabled && first != "" && second != "" && input[i] == ')' {
				firstNum := cast.ToInt(first)
				secondNum := cast.ToInt(second)
				sum += firstNum * secondNum
			}
		}
	}

	return sum
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
