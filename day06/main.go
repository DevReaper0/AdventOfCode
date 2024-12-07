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
	routeMap, guardX, guardY, guardDir := parseInput(input)

	visited := simulate(routeMap, guardX, guardY, guardDir)
	return len(visited)
}

func part2(input string) int {
	routeMap, guardX, guardY, guardDir := parseInput(input)

	validPositions := 0

	// Try placing an obstruction at every valid position
	for y := 0; y < len(routeMap); y++ {
		for x := 0; x < len(routeMap[0]); x++ {
			if routeMap[y][x] != '.' || (x == guardX && y == guardY) {
				continue
			}

			routeMap[y] = routeMap[y][:x] + "#" + routeMap[y][x+1:]
			if causesLoop(routeMap, guardX, guardY, guardDir) {
				validPositions++
			}
			routeMap[y] = routeMap[y][:x] + "." + routeMap[y][x+1:]
		}
	}

	return validPositions
}

func simulate(routeMap []string, startX, startY int, startDir rune) map[[2]int]bool {
	visited := make(map[[2]int]bool)
	x, y, dir := startX, startY, startDir

	for {
		visited[[2]int{x, y}] = true

		nextX, nextY := nextPos(x, y, dir)

		if nextX < 0 || nextY < 0 || nextY >= len(routeMap) || nextX >= len(routeMap[0]) {
			break
		}

		if routeMap[nextY][nextX] == '#' {
			dir = nextDir(dir)
		} else {
			x, y = nextX, nextY
		}
	}

	return visited
}

func causesLoop(routeMap []string, startX, startY int, startDir rune) bool {
	x, y, dir := startX, startY, startDir
	states := make(map[[3]int]bool)

	for {
		state := [3]int{x, y, int(dir)}
		if states[state] {
			return true
		}
		states[state] = true

		nextX, nextY := nextPos(x, y, dir)

		if nextX < 0 || nextY < 0 || nextY >= len(routeMap) || nextX >= len(routeMap[0]) {
			return false
		}

		if routeMap[nextY][nextX] == '#' {
			dir = nextDir(dir)
		} else {
			x, y = nextX, nextY
		}
	}
}

func parseInput(input string) ([]string, int, int, rune) {
	routeMap := strings.Split(input, "\n")
	var guardX, guardY int
	var guardDir rune

	for y, row := range routeMap {
		for x, cell := range row {
			if cell == '^' || cell == 'v' || cell == '<' || cell == '>' {
				guardX, guardY = x, y
				guardDir = rune(cell)
				routeMap[y] = row[:x] + "." + row[x+1:]
				break
			}
		}
	}

	return routeMap, guardX, guardY, guardDir
}

func nextPos(x, y int, dir rune) (int, int) {
	switch dir {
	case '^':
		return x, y - 1
	case 'v':
		return x, y + 1
	case '<':
		return x - 1, y
	case '>':
		return x + 1, y
	}
	return x, y
}

func nextDir(dir rune) rune {
	switch dir {
	case '^':
		return '>'
	case '>':
		return 'v'
	case 'v':
		return '<'
	case '<':
		return '^'
	}
	return dir
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
