package main

import (
	"fmt"
	"utils"
)

type Point struct {
	x, y      int
	direction rune
}

func move(p Point) Point {
	switch p.direction {
	case 'E':
		p.x++
	case 'W':
		p.x--
	case 'S':
		p.y++
	case 'N':
		p.y--
	}
	return p
}

func directions(p Point, grid []string) []rune {
	tile := rune(grid[p.y][p.x])
	switch tile {
	case '-':
		switch p.direction {
		case 'E', 'W':
			return []rune{p.direction}
		case 'S', 'N':
			return []rune{'E', 'W'}
		}
	case '|':
		switch p.direction {
		case 'S', 'N':
			return []rune{p.direction}
		case 'E', 'W':
			return []rune{'S', 'N'}
		}
	case '/':
		return []rune{map[rune]rune{
			'E': 'N',
			'S': 'W',
			'W': 'S',
			'N': 'E',
		}[p.direction]}
	case '\\':
		return []rune{map[rune]rune{
			'E': 'S',
			'S': 'E',
			'W': 'N',
			'N': 'W',
		}[p.direction]}
	case '.':
		return []rune{p.direction}
	}
	panic(fmt.Sprint("Unexpected grid tile or direction ", tile, p.direction))
}

func inBounds(p Point, grid []string) bool {
	return p.x >= 0 && p.y >= 0 && p.x < len(grid[0]) && p.y < len(grid)
}

func energise(start Point, grid []string) int {
	visited := map[Point]struct{}{}
	toCheck := map[Point]struct{}{start: {}}
	for len(toCheck) > 0 {
		newCheck := map[Point]struct{}{}
		for check := range toCheck {
			point := move(check)
			if inBounds(point, grid) {
				for _, dir := range directions(point, grid) {
					point.direction = dir
					_, done := visited[point]
					if !done {
						visited[point] = struct{}{}
						newCheck[point] = struct{}{}
					}
				}
			}
		}
		toCheck = newCheck
	}
	energised := map[Point]struct{}{}
	for point := range visited {
		point.direction = 'O'
		energised[point] = struct{}{}
	}
	return len(energised)
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve() {
	grid := utils.ReadInput("input.txt", "\n")
	fmt.Println("Part 1:", energise(Point{-1, 0, 'E'}, grid))

	max := 0
	for i := 0; i < len(grid[0]); i++ { // Iterate over rows
		max = maxInt(max, energise(Point{-1, i, 'E'}, grid))
		max = maxInt(max, energise(Point{len(grid[0]), i, 'W'}, grid))
	}
	for i := 0; i < len(grid); i++ { // Iterate over columns
		max = maxInt(max, energise(Point{i, -1, 'S'}, grid))
		max = maxInt(max, energise(Point{i, len(grid), 'N'}, grid))
	}
	fmt.Println("Part 2:", max)
}

func main() {
	utils.TimeFunction(solve)
}
