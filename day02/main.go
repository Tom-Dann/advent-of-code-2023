package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func timeFunction(function func()) {
	start := time.Now()
	function()
	fmt.Println("Time elapsed:", time.Since(start))
}

func solve() {
	// Read file
	raw, err := os.ReadFile("input.txt")
	check(err)
	lines := strings.Split(strings.TrimSpace(string(raw)), "\n")

	max := map[string]int{"red": 12, "green": 13, "blue": 14}
	part1sum, part2sum := 0, 0
	for index, line := range lines {
		re := regexp.MustCompile(`(\d+)\s(red|green|blue)`)
		cubes := re.FindAllStringSubmatch(line, -1)
		possible := true
		gameMax := map[string]int{"red": 0, "green": 0, "blue": 0}
		for _, cube := range cubes {
			n, err := strconv.Atoi(cube[1])
			check(err)
			if n > max[cube[2]] { // Part1 - Check if max is exceeded
				possible = false
			}
			if n > gameMax[cube[2]] { // Part2 - Record max for current game
				gameMax[cube[2]] = n
			}
		}
		if possible {
			part1sum += index + 1
		}
		part2sum += gameMax["red"] * gameMax["green"] * gameMax["blue"]
	}

	fmt.Println("Part 1:", part1sum)
	fmt.Println("Part 2:", part2sum)
}

func main() {
	timeFunction(solve)
}
