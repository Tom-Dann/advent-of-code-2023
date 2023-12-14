package main

import (
	"fmt"
	"slices"
	"strings"
	"utils"
)

func getColumn(grid []string, n int) string {
	col := ""
	for _, row := range grid {
		col += string(row[n])
	}
	return col
}

func diffStrings(a string, b string) int {
	diff := 0
	max := max(len(a), len(b))
	for i := 0; i < max; i++ {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff
}

func scoreGrid(grid []string) []int {
	scores := make([]int, 2)
	for i := 0; i < len(grid)-1 && slices.Contains(scores, 0); i++ { // Check rows
		diff := 0
		for j := i; j >= 0 && len(grid)+j > 2*i+1; j-- {
			diff += diffStrings(grid[j], grid[2*i-j+1])
			if diff > 1 {
				break
			}
		}
		if diff <= 1 {
			scores[diff] = (i + 1) * 100
		}
	}

	for i := 0; i < len(grid[0])-1 && slices.Contains(scores, 0); i++ { // Check columns
		diff := 0
		for j := i; j >= 0 && len(grid[0])+j > 2*i+1; j-- {
			diff += diffStrings(getColumn(grid, j), getColumn(grid, 2*i-j+1))
			if diff > 1 {
				break
			}
		}
		if diff <= 1 {
			scores[diff] = i + 1
		}
	}
	return scores
}

func solve() {
	sections := utils.ReadInput("input.txt", "\n\n")

	part1, part2 := 0, 0
	for _, section := range sections {
		grid := strings.Split(strings.TrimSpace(section), "\n")
		scores := scoreGrid(grid)
		part1 += scores[0]
		part2 += scores[1]
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func main() {
	utils.TimeFunction(solve)
}
