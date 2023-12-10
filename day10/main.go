package main

import (
	"fmt"
	"os"
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

type Point2D struct {
	X, Y int
}

type PointSet map[Point2D]struct{}

var connections = struct{ N, E, S, W string }{
	N: "S|LJ",
	E: "S-LF",
	S: "S|7F",
	W: "S-J7",
}

func floodFill(grid [][]rune, visited PointSet, toCheck PointSet) (PointSet, PointSet) {
	next := make(PointSet)
	checkPoint := func(point Point2D, direction string) {
		_, checked := visited[point]
		if !checked && strings.Contains(direction, string(grid[point.Y][point.X])) {
			visited[point] = struct{}{}
			next[point] = struct{}{}
		}
	}

	for curr := range toCheck {
		symbol := grid[curr.Y][curr.X]
		if strings.Contains(connections.N, string(symbol)) && curr.Y > 0 { // North
			checkPoint(Point2D{X: curr.X, Y: curr.Y - 1}, connections.S)
		}
		if strings.Contains(connections.E, string(symbol)) && curr.X < len(grid[0])-1 { // East
			checkPoint(Point2D{X: curr.X + 1, Y: curr.Y}, connections.W)
		}
		if strings.Contains(connections.S, string(symbol)) && curr.Y < len(grid)-1 { // South
			checkPoint(Point2D{X: curr.X, Y: curr.Y + 1}, connections.N)
		}
		if strings.Contains(connections.W, string(symbol)) && curr.X > 0 { // West
			checkPoint(Point2D{X: curr.X - 1, Y: curr.Y}, connections.E)
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
	raw, err := os.ReadFile("input.txt") // Read file
	check(err)
	lines := strings.Split(strings.TrimSpace(string(raw)), "\n")

	var source Point2D
	grid := make([][]rune, len(lines))
	for i, line := range lines { // Setup grid
		grid[i] = []rune(line)
		if strings.Contains(line, "S") { // Find source
			source = Point2D{X: strings.Index(line, "S"), Y: i}
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
		goNext := func(point Point2D, connection string, turn rune) {
			if strings.Contains(connection, string(grid[point.Y][point.X])) {
				dir := rotate(facing, turn)
				rotation += dir
				var surround []Point2D
				switch facing {
				case 'N':
					surround = []Point2D{{X: current.X - 1, Y: current.Y}, {X: current.X, Y: current.Y - 1}, {X: current.X + 1, Y: current.Y}}
				case 'E':
					surround = []Point2D{{X: current.X, Y: current.Y - 1}, {X: current.X + 1, Y: current.Y}, {X: current.X, Y: current.Y + 1}}
				case 'S':
					surround = []Point2D{{X: current.X + 1, Y: current.Y}, {X: current.X, Y: current.Y + 1}, {X: current.X - 1, Y: current.Y}}
				case 'W':
					surround = []Point2D{{X: current.X, Y: current.Y + 1}, {X: current.X - 1, Y: current.Y}, {X: current.X, Y: current.Y - 1}}
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

		symbol := string(grid[current.Y][current.X])
		if strings.Contains(connections.N, symbol) && facing != 'S' && current.Y > 0 { // North
			goNext(Point2D{X: current.X, Y: current.Y - 1}, connections.S, 'N')
		}
		if strings.Contains(connections.E, symbol) && facing != 'W' && current.X < len(grid[0])-1 && !moved { // East
			goNext(Point2D{X: current.X + 1, Y: current.Y}, connections.W, 'E')
		}
		if strings.Contains(connections.S, symbol) && facing != 'N' && current.Y < len(grid)-1 && !moved { // South
			goNext(Point2D{X: current.X, Y: current.Y + 1}, connections.N, 'S')
		}
		if strings.Contains(connections.W, symbol) && facing != 'E' && current.X > 0 && !moved { // West
			goNext(Point2D{X: current.X - 1, Y: current.Y}, connections.E, 'W')
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
			adjacent := []Point2D{
				{X: point.X, Y: point.Y - 1}, // North
				{X: point.X + 1, Y: point.Y}, // East
				{X: point.X, Y: point.Y + 1}, // South
				{X: point.X - 1, Y: point.Y}, // West
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
	timeFunction(solve)
}
