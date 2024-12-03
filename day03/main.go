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

	first := ""
	second := ""
	state := ""
	enabled := true
	for _, c := range input {
		if state == "" && c == 'm' {
			state = "m"
		} else if state == "m" {
			if c == 'u' {
				state = "mu"
			} else {
				state = ""
			}
		} else if state == "mu" {
			if c == 'l' {
				state = "mul"
			} else {
				state = ""
			}
		} else if state == "mul" {
			if c == '(' {
				state = "mul("
			} else {
				state = ""
			}
		} else if state == "mul(" {
			if c >= '0' && c <= '9' {
				first += string(c)
			} else if c == ',' {
				state = ","
			} else {
				first = ""
				state = ""
			}
		} else if state == "," {
			if c >= '0' && c <= '9' {
				second += string(c)
			} else if c == ')' {
				state = ")"
			} else {
				first = ""
				second = ""
				state = ""
			}
		} else if handleDisabling {
			if c == 'd' {
				state = "d"
			} else if state == "d" {
				if c == 'o' {
					state = "do"
				} else {
					state = ""
				}
			} else if state == "do" {
				if c == '(' {
					state = "do("
				} else if c == 'n' {
					state = "don"
				} else {
					state = ""
				}
			} else if state == "do(" {
				if c == ')' {
					enabled = true
				}
				state = ""
			} else if state == "don" {
				if c == '\'' {
					state = "don'"
				} else {
					state = ""
				}
			} else if state == "don'" {
				if c == 't' {
					state = "don't"
				} else {
					state = ""
				}
			} else if state == "don't" {
				if c == '(' {
					state = "don't("
				} else {
					state = ""
				}
			} else if state == "don't(" {
				if c == ')' {
					enabled = false
				}
				state = ""
			}
		}

		if state == ")" {
			if enabled && first != "" && second != "" {
				firstNum := cast.ToInt(first)
				secondNum := cast.ToInt(second)
				sum += firstNum * secondNum
			}
			first = ""
			second = ""
			state = ""
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
