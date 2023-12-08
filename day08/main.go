package main

import (
	"fmt"
	"os"
	"regexp"
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

type nodeMap struct {
	Left  string
	Right string
}

func part2end(locations []string) bool {
	for _, loc := range locations {
		if loc[len(loc)-1] != 'Z' {
			return false
		}
	}
	return true
}

func countSteps(lr string, paths map[string][]string, loc string, part1 bool) int {
	end := false
	count := 0
	for !end {
		if lr[(count%len(lr))] == 'L' { // Left
			loc = paths[loc][0]
		} else { // Right
			loc = paths[loc][1]
		}
		count++
		if part1 {
			end = loc == "ZZZ"
		} else {
			end = loc[len(loc)-1] == 'Z'
		}
	}
	return count
}

func gcd(a int, b int) int { // Apply Euclidean algorithm
	for b != 0 {
		r := b
		b = a % b
		a = r
	}
	return a
}

func lcm(a int, b int) int {
	return (a * b) / gcd(a, b)
}

func solve() {
	raw, err := os.ReadFile("input.txt") // Read file
	check(err)
	sections := strings.Split(strings.TrimSpace(string(raw)), "\n\n")

	lr := sections[0]
	paths := make(map[string][]string)
	re := regexp.MustCompile(`(.+) = \((.+), (.+)\)`)
	startLocs := make([]string, 0)
	for _, line := range strings.Split(sections[1], "\n") {
		m := re.FindAllStringSubmatch(line, -1)
		source := m[0][1]
		paths[source] = []string{m[0][2], m[0][3]}
		if source[len(source)-1] == 'A' { // Start location for part 2
			startLocs = append(startLocs, source)
		}
	}

	part1 := countSteps(lr, paths, "AAA", true)
	part2 := 1
	for _, start := range startLocs {
		part2 = lcm(part2, countSteps(lr, paths, start, false))
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func main() {
	timeFunction(solve)
}
