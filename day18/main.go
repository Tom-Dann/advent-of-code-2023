package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

func parseLine(s string, part1 bool) (rune, int) {
	line := strings.Split(s, " ")
	if part1 {
		n, _ := strconv.Atoi(line[1])
		return rune(line[0][0]), n
	}
	dist, _ := strconv.ParseInt(line[2][2:7], 16, 32)
	return rune(line[2][7]), int(dist)
}

type Point struct{ x, y int }

func move(p Point, dir rune, dist int) Point {
	switch dir {
	case 'R', '0':
		p.x += dist
	case 'D', '1':
		p.y += dist
	case 'L', '2':
		p.x -= dist
	case 'U', '3':
		p.y -= dist
	}
	return p
}

func findArea(lines []string, part1 bool) int {
	curr := Point{0, 0}
	shoelace, boundary := 0, 0
	for _, line := range lines {
		dir, dist := parseLine(line, part1)
		new := move(curr, dir, dist)
		boundary += dist
		shoelace += curr.x*new.y - curr.y*new.x // Shoelace formula
		curr = new
	}
	return (shoelace+boundary)>>1 + 1 // Pick's formula
}

func solve() {
	lines := utils.ReadInput("input.txt", "\n")
	fmt.Println("Part 1:", findArea(lines, true))
	fmt.Println("Part 2:", findArea(lines, false))
}

func main() {
	utils.TimeFunction(solve)
}
