package main

import (
	"fmt"
	"strings"
	"utils"
)

type Point struct {
	x, y int
}
type PointSet map[Point]struct{}

var connections = struct{ N, E, S, W string }{
	N: "S|LJ",
	E: "S-LF",
	S: "S|7F",
	W: "S-J7",
}

func floodFill(grid [][]rune, visited PointSet, toCheck PointSet) (PointSet, PointSet) {
	next := make(PointSet)
	checkPoint := func(point Point, direction string) {
		_, checked := visited[point]
		if !checked && strings.Contains(direction, string(grid[point.y][point.x])) {
			visited[point] = struct{}{}
			next[point] = struct{}{}
		}
	}

	for curr := range toCheck {
		symbol := grid[curr.y][curr.x]
		if strings.Contains(connections.N, string(symbol)) && curr.y > 0 { // North
			checkPoint(Point{x: curr.x, y: curr.y - 1}, connections.S)
		}
		if strings.Contains(connections.E, string(symbol)) && curr.x < len(grid[0])-1 { // East
			checkPoint(Point{x: curr.x + 1, y: curr.y}, connections.W)
		}
		if strings.Contains(connections.S, string(symbol)) && curr.y < len(grid)-1 { // South
			checkPoint(Point{x: curr.x, y: curr.y + 1}, connections.N)
		}
		if strings.Contains(connections.W, string(symbol)) && curr.x > 0 { // West
			checkPoint(Point{x: curr.x - 1, y: curr.y}, connections.E)
		}
	}
	return visited, next
}

func rotate(facing rune, turn rune) int {
	if facing == turn {
		return 0 // Straight ahead
	}
	clockwise := map[rune]rune{
		'N': 'E',
		'E': 'S',
		'S': 'W',
		'W': 'N',
	}
	if clockwise[facing] == turn {
		return 1 // Turn right/clockwise
	}
	return -1 // Turn left/anti-clockwise
}

func solve() {
	lines := utils.ReadInput("input.txt", "\n")

	var source Point
	grid := make([][]rune, len(lines))
	for i, line := range lines { // Setup grid
		grid[i] = []rune(line)
		if strings.Contains(line, "S") { // Find source
			source = Point{x: strings.Index(line, "S"), y: i}
		}
	}

	next, visited := PointSet{source: struct{}{}}, PointSet{source: struct{}{}}
	steps := -1
	for len(next) > 0 {
		visited, next = floodFill(grid, visited, next)
		steps++
	}
	fmt.Println("Part 1:", steps)

	current, facing, rotation := source, 'E', 0
	finished := false
	left, right := PointSet{}, PointSet{}
	// Traverse pipe
	for !finished {
		moved := false
		goNext := func(point Point, connection string, turn rune) {
			if strings.Contains(connection, string(grid[point.y][point.x])) {
				dir := rotate(facing, turn)
				rotation += dir
				var surround []Point
				switch facing {
				case 'N':
					surround = []Point{{x: current.x - 1, y: current.y}, {x: current.x, y: current.y - 1}, {x: current.x + 1, y: current.y}}
				case 'E':
					surround = []Point{{x: current.x, y: current.y - 1}, {x: current.x + 1, y: current.y}, {x: current.x, y: current.y + 1}}
				case 'S':
					surround = []Point{{x: current.x + 1, y: current.y}, {x: current.x, y: current.y + 1}, {x: current.x - 1, y: current.y}}
				case 'W':
					surround = []Point{{x: current.x, y: current.y + 1}, {x: current.x - 1, y: current.y}, {x: current.x, y: current.y - 1}}
				}
				_, pipe0 := visited[surround[0]]
				_, pipe1 := visited[surround[1]]
				_, pipe2 := visited[surround[2]]
				if !pipe0 {
					left[surround[0]] = struct{}{}
				}
				if !pipe2 {
					right[surround[2]] = struct{}{}
				}
				if !pipe1 && dir == -1 {
					right[surround[1]] = struct{}{}
				}
				if !pipe1 && dir == 1 {
					left[surround[1]] = struct{}{}
				}
				current = point
				facing = turn
				moved = true
			}
		}

		symbol := string(grid[current.y][current.x])
		if strings.Contains(connections.N, symbol) && facing != 'S' && current.y > 0 { // North
			goNext(Point{x: current.x, y: current.y - 1}, connections.S, 'N')
		}
		if strings.Contains(connections.E, symbol) && facing != 'W' && current.x < len(grid[0])-1 && !moved { // East
			goNext(Point{x: current.x + 1, y: current.y}, connections.W, 'E')
		}
		if strings.Contains(connections.S, symbol) && facing != 'N' && current.y < len(grid)-1 && !moved { // South
			goNext(Point{x: current.x, y: current.y + 1}, connections.N, 'S')
		}
		if strings.Contains(connections.W, symbol) && facing != 'E' && current.x > 0 && !moved { // West
			goNext(Point{x: current.x - 1, y: current.y}, connections.E, 'W')
		}

		if current == source {
			finished = true
		}
	}

	inside, toCheck := PointSet{}, left
	if rotation > 0 { // If we went clockwise use the set on the right
		toCheck = right
	}
	for len(toCheck) > 0 { // Flood fill on inside set
		for point := range toCheck {
			inside[point] = struct{}{}
			delete(toCheck, point)
			adjacent := []Point{
				{x: point.x, y: point.y - 1}, // North
				{x: point.x + 1, y: point.y}, // East
				{x: point.x, y: point.y + 1}, // South
				{x: point.x - 1, y: point.y}, // West
			}
			for _, check := range adjacent {
				_, pipe := visited[check]
				_, seen := inside[check]
				if !pipe && !seen {
					toCheck[check] = struct{}{}
				}
			}
		}
	}
	fmt.Println("Part 2:", len(inside))
}

func main() {
	utils.TimeFunction(solve)
}
