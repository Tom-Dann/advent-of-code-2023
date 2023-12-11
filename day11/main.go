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

type Point struct {
	x, y int
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func countElements(arr []int, a int, b int) int {
	count := 0
	min, max := a, b
	if b < a {
		min, max = b, a
	}
	for _, x := range arr {
		if min <= x && x <= max {
			count++
		}
	}
	return count
}

func solve() {
	raw, err := os.ReadFile("input.txt") // Read file
	check(err)
	lines := strings.Split(strings.TrimSpace(string(raw)), "\n")

	emptyRows, emptyCols := make([]int, 0), make([]int, 0)
	galaxies := make([]Point, 0)
	re := regexp.MustCompile(`#`)
	for i, line := range lines { // Find galaxies and empty rows
		if strings.ContainsRune(line, '#') {
			matches := re.FindAllStringIndex(line, -1)
			for _, match := range matches {
				galaxies = append(galaxies, Point{x: match[0], y: i})
			}
		} else {
			emptyRows = append(emptyRows, i)
		}
	}
	for i := 0; i < len(lines[0]); i++ { // Find empty columns
		empty := true
		for _, galaxy := range galaxies {
			if galaxy.x == i {
				empty = false
				break
			}
		}
		if empty {
			emptyCols = append(emptyCols, i)
		}
	}

	dist := func(a Point, b Point, factor int) int {
		xDelta := abs(a.x - b.x)
		yDelta := abs(a.y - b.y)
		dist := xDelta + yDelta
		if xDelta > 1 {
			dist += countElements(emptyCols, a.x, b.x) * (factor - 1)
		}
		if yDelta > 1 {
			dist += countElements(emptyRows, a.y, b.y) * (factor - 1)
		}
		return dist
	}

	part1, part2 := 0, 0
	for i, galaxyA := range galaxies { // Loop through galaxies pairwise
		for _, galaxyB := range galaxies[i+1:] {
			part1 += dist(galaxyA, galaxyB, 2)
			part2 += dist(galaxyA, galaxyB, 1000000)
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func main() {
	timeFunction(solve)
}
