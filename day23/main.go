package main

import (
	"fmt"
	"slices"
	"utils"
)

type Point struct{ x, y int }
type Path []Point
type Distances map[Point]int
type Graph map[Point]Distances

func (point Point) getNeighbours() []Point {
	return []Point{
		{point.x + 1, point.y}, // >
		{point.x - 1, point.y}, // <
		{point.x, point.y + 1}, // v
		{point.x, point.y - 1}, // ^
	}
}

func (point Point) valid(grid []string) bool {
	return point.x >= 0 && point.y >= 0 && point.x < len(grid[0]) && point.y < len(grid) && grid[point.y][point.x] != '#'
}

func (path Path) getNext(grid []string, part1 bool) []Path {
	head := path[len(path)-1]
	if part1 {
		symbol := rune(grid[head.y][head.x])
		index := slices.Index([]rune{'>', '<', 'v', '^'}, symbol)
		if index >= 0 {
			next := head.getNeighbours()[index]
			if next.valid(grid) && !slices.Contains(path, next) {
				return []Path{append(path, head.getNeighbours()[index])}
			}
			return []Path{}
		}
	}
	paths := []Path{}
	for _, next := range head.getNeighbours() {
		if next.valid(grid) && !slices.Contains(path, next) {
			nextPath := make(Path, len(path))
			copy(nextPath, path)
			paths = append(paths, append(nextPath, next))
		}
	}
	return paths
}

func search(grid []string, part1 bool) int {
	start, finish := Point{1, 0}, Point{len(grid) - 2, len(grid) - 1}
	graph := Graph{start: {}}
	queue := []Path{{start}}
	for len(queue) > 0 { // Create graph of distances between graph vertices
		curr := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		next := curr.getNext(grid, part1)
		vertex := curr[len(curr)-1]
		if ((vertex == start || vertex == finish) && len(curr) > 1) || len(next) > 1 { // At a new crossroads or start/end
			_, visited := graph[vertex]
			if !visited {
				graph[vertex] = Distances{}
				queue = append(queue, Path{vertex}.getNext(grid, part1)...)
			}
			graph[curr[0]][vertex] = len(curr) - 1
		} else {
			queue = append(queue, next...)
		}
	}
	max := 0
	queue = []Path{{start}}
	for len(queue) > 0 { // Depth-first search on graph of distances
		curr := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		head := curr[len(curr)-1]
		dist, last := graph[head][finish]
		if last {
			for i := 0; i < len(curr)-1; i++ {
				dist += graph[curr[i]][curr[i+1]]
			}
			if dist > max {
				max = dist
			}
		} else {
			for next := range graph[head] {
				if !slices.Contains(curr, next) {
					nextPath := make(Path, len(curr))
					copy(nextPath, curr)
					queue = append(queue, append(nextPath, next))
				}
			}
		}
	}
	return max
}

func solve(grid []string) {
	fmt.Println("Part 1:", search(grid, true))
	fmt.Println("Part 2:", search(grid, false))
}

func main() {
	grid := utils.ReadInput("input.txt", "\n")
	utils.TimeFunctionInput(solve, grid)
}
