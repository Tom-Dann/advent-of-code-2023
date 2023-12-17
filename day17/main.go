package main

import (
	"fmt"
	"math"
	"utils"
)

type Position struct {
	x, y int
	dir  rune
	len  int
}

func getNeightbour(b Position, dir rune, part int) Position {
	block := b
	if map[rune]rune{'E': 'W', 'S': 'N', 'W': 'E', 'N': 'S'}[block.dir] == dir { // Not allowed to doube back
		return Position{-1, -1, '-', 0}
	}
	if part == 2 && block.len <= 3 && block.dir != dir && block.dir != 'O' { // Must go a min of 4 in same direction in part 2
		return Position{-1, -1, '-', 0}
	}
	switch dir {
	case 'E':
		block.x++
	case 'W':
		block.x--
	case 'S':
		block.y++
	case 'N':
		block.y--
	}
	if dir == block.dir {
		block.len++
	} else {
		block.dir = dir
		block.len = 1
	}
	return block
}

func dijkstra(lines []string, part int) int {
	max := struct{ width, height int }{len(lines[0]), len(lines)}
	heat := make([][]int, max.height)
	heatLoss := map[Position]int{}
	for i, line := range lines {
		heat[i] = make([]int, max.width)
		for j, char := range line {
			heat[i][j] = int(char) - 48
		}
	}
	start := Position{0, 0, 'O', 0}
	heatLoss[start] = 0
	queue := map[Position]struct{}{start: {}}

	valid := func(block Position) bool {
		if block.x < 0 || block.x >= max.width || block.y < 0 || block.y >= max.height {
			return false
		}
		if part == 1 {
			return block.len <= 3
		}
		return block.len <= 10
	}

	for len(queue) > 0 {
		var next Position
		minLoss := math.MaxInt32
		for block := range queue { // Find min
			if heatLoss[block] < minLoss {
				next = block
				minLoss = heatLoss[block]
			}
		}
		delete(queue, next)
		if next.x == max.width-1 && next.y == max.height-1 {
			return heatLoss[next]
		}
		for _, dir := range []rune{'E', 'S', 'W', 'N'} {
			neighbour := getNeightbour(next, dir, part)
			if valid(neighbour) {
				dist := heatLoss[next] + heat[neighbour.y][neighbour.x]
				best, visited := heatLoss[neighbour]
				if !visited || dist < best {
					heatLoss[neighbour] = dist
					queue[neighbour] = struct{}{}
				}
			}
		}
	}
	return -1
}

func part1(lines []string) {
	fmt.Println("Part 1:", dijkstra(lines, 1))
}

func part2(lines []string) {
	fmt.Println("Part 2:", dijkstra(lines, 2))
}

func main() {
	lines := utils.ReadInput("input.txt", "\n")
	utils.TimeFunctionInput(part1, lines)
	utils.TimeFunctionInput(part2, lines)
}
