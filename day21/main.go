package main

import (
	"fmt"
	"utils"
)

type Point struct{ x, y int }
type PointSet map[Point]struct{}

func parseGrid(grid []string) (Point, PointSet) {
	rocks := PointSet{}
	start := Point{}
	for j, line := range grid {
		for i, char := range line {
			switch char {
			case '#':
				rocks[Point{i, j}] = struct{}{}
			case 'S':
				start = Point{i, j}
			}
		}
	}
	return start, rocks
}

func (p Point) getNeighbours() PointSet {
	return PointSet{
		Point{p.x, p.y + 1}: {},
		Point{p.x, p.y - 1}: {},
		Point{p.x + 1, p.y}: {},
		Point{p.x - 1, p.y}: {},
	}
}

func mod(n int, m int) int {
	ans := n % m
	if ans < 0 {
		return ans + m
	}
	return ans
}

func countSteps(rocks PointSet, start Point, max int, steps int) int {
	curr := PointSet{start: {}}
	for step := 1; step <= steps; step++ {
		next := PointSet{}
		for plot := range curr {
			for neighbour := range plot.getNeighbours() {
				_, rock := rocks[Point{mod(neighbour.x, max), mod(neighbour.y, max)}]
				if !rock {
					next[neighbour] = struct{}{}
				}
			}
		}
		curr = next
	}
	return len(curr)
}

func part1(grid []string) {
	start, rocks := parseGrid(grid)
	fmt.Println("Part 1:", countSteps(rocks, start, len(grid), 64))
}

func part2(grid []string) {
	start, rocks := parseGrid(grid)
	max := len(grid)
	steps := 26501365
	mod := 2 * max
	remainder := steps % mod

	c := countSteps(rocks, start, max, remainder)           // f(0) = c
	f1 := countSteps(rocks, start, max, remainder+mod)      // f(1) = a + b + c
	f2 := countSteps(rocks, start, max, remainder+(mod<<1)) // f(2) = 4a + 2b + c
	a := (f2 - (f1 << 1) + c) >> 1
	b := f1 - a - c
	n := steps / mod
	fmt.Println("Part 2:", a*n*n+b*n+c) // f(x) = a*x^2 + b*x + c
}

func main() {
	grid := utils.ReadInput("input.txt", "\n")
	utils.TimeFunctionInput(part1, grid)
	utils.TimeFunctionInput(part2, grid)
}
