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
	m := float64(stone.vel.y) / float64(stone.vel.x)
	c := float64(stone.pos.y) - m*float64(stone.pos.x) // c=y-mx
	return m, c
}

func intersect(s Hailstone, t Hailstone) bool {
	min, max := float64(200000000000000), float64(400000000000000)
	a, c := getLine(s)
	b, d := getLine(t)
	if a == b {
		return false
	}
	x := (d - c) / (a - b)
	y := a*x + c
	return x >= min && y >= min && x <= max && y <= max &&
		(math.Signbit(y-float64(s.pos.y)) == math.Signbit(float64(s.vel.y))) &&
		(math.Signbit(y-float64(t.pos.y)) == math.Signbit(float64(t.vel.y)))
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
	t1 := (float64(throw.pos.y) - float64(b.pos.y)) / (float64(b.vel.y) - float64(throw.vel.y))
	t2 := (float64(throw.pos.y) - float64(c.pos.y)) / (float64(c.vel.y) - float64(throw.vel.y))
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
