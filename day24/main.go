package main

import (
	"fmt"
	"math"
	"utils"
)

type Point struct{ x, y, z float64 }
type Hailstone struct{ pos, vel Point }

func parseLine(line string) Hailstone {
	s := Hailstone{}
	fmt.Sscanf(line, "%f, %f, %f @ %f, %f, %f",
		&s.pos.x, &s.pos.y, &s.pos.z,
		&s.vel.x, &s.vel.y, &s.vel.z)
	return s
}

func getLine(stone Hailstone) (float64, float64) {
	m := stone.vel.y / stone.vel.x
	c := stone.pos.y - m*stone.pos.x // c=y-mx
	return m, c
}

func intersect(s Hailstone, t Hailstone) bool {
	min, max := 2e14, 4e14
	a, c := getLine(s)
	b, d := getLine(t)
	if a == b {
		return false
	}
	x := (d - c) / (a - b)
	y := a*x + c
	return x >= min && y >= min && x <= max && y <= max &&
		math.Signbit(y-s.pos.y) == math.Signbit(s.vel.y) &&
		math.Signbit(y-t.pos.y) == math.Signbit(t.vel.y)
}

func solve(lines []string) {
	stones := make([]Hailstone, len(lines))
	for i, line := range lines {
		stones[i] = parseLine(line)
	}
	count := 0
	a := Hailstone{}
	for i, s := range stones {
		for _, t := range stones[i+1:] {
			if intersect(s, t) {
				count++
			}
			if s.pos.y == t.pos.y {
				a = s
			}
		}
	}
	fmt.Println("Part 1:", count)

	throw := Hailstone{}
	throw.vel.y = a.vel.y
	throw.pos.y = a.pos.y

	b, c := stones[0], stones[1]
	t1 := (throw.pos.y - b.pos.y) / (b.vel.y - throw.vel.y)
	t2 := (throw.pos.y - c.pos.y) / (c.vel.y - throw.vel.y)
	throw.vel.x = (b.pos.x - c.pos.x + t1*b.vel.x - t2*c.vel.x) / (t1 - t2)
	throw.vel.z = (b.pos.z - c.pos.z + t1*b.vel.z - t2*c.vel.z) / (t1 - t2)
	throw.pos.x = b.pos.x + t1*(b.vel.x-throw.vel.x)
	throw.pos.z = b.pos.z + t1*(b.vel.z-throw.vel.z)

	fmt.Println("Part 2:", int(throw.pos.x+throw.pos.y+throw.pos.z))
}

func main() {
	lines := utils.ReadInput("input.txt", "\n")
	utils.TimeFunctionInput(solve, lines)
}
