package main

import (
	"fmt"
	"os"
	"reflect"
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

type Rock struct {
	x, y int
}
type RockSet map[Rock]struct{}

func tilt(rounds RockSet, cubes RockSet, max int) RockSet {
	for i := 0; i < max; i++ { // Loop over columns from left
		for j := 1; j < max; j++ { // Loop over rows from top
			_, roundRock := rounds[Rock{x: i, y: j}]
			if roundRock {
				for k := j - 1; k >= -1; k-- { // Loop over rows from current row back to the top
					_, cube := cubes[Rock{x: i, y: k}]
					_, round := rounds[Rock{x: i, y: k}]
					if cube || round || k == -1 {
						delete(rounds, Rock{x: i, y: j})
						rounds[Rock{x: i, y: k + 1}] = struct{}{}
						break
					}
				}
			}
		}
	}
	return rounds
}

func rotateRocks(old RockSet, max int) RockSet {
	new := make(RockSet)
	for rock := range old {
		new[Rock{x: max - rock.y - 1, y: rock.x}] = struct{}{}
	}
	return new
}

func rotate(rounds RockSet, cubes RockSet, max int) (RockSet, RockSet) {
	return rotateRocks(rounds, max), rotateRocks(cubes, max)
}

func spinCycle(rounds RockSet, cubes RockSet, size int) (RockSet, RockSet) {
	for i := 0; i < 4; i++ {
		rounds = tilt(rounds, cubes, size)
		rounds, cubes = rotate(rounds, cubes, size)
	}
	return rounds, cubes
}

func totalLoad(rocks RockSet, max int) int {
	total := 0
	for rock := range rocks {
		total += max - rock.y
	}
	return total
}

func solve() {
	raw, err := os.ReadFile("input.txt") // Read file
	check(err)
	lines := strings.Split(strings.TrimSpace(string(raw)), "\n")

	cubes, rounds := make(RockSet), make(RockSet)
	for j, line := range lines { // Parse input
		for i, char := range line {
			switch char {
			case '#':
				cubes[Rock{x: i, y: j}] = struct{}{}
			case 'O':
				rounds[Rock{x: i, y: j}] = struct{}{}
			}
		}
	}

	if len(lines) != len(lines[0]) {
		panic("Input not a square")
	}
	size := len(lines)
	part1 := totalLoad(tilt(rounds, cubes, size), size)

	repeating, settleTime := false, 1000
	history, loads := make(RockSet), make([]int, 0)
	for i := 0; !repeating; i++ {
		if i == settleTime {
			for k := range rounds {
				history[k] = struct{}{}
			}
		}
		rounds, cubes = spinCycle(rounds, cubes, size)
		if i >= settleTime {
			loads = append(loads, totalLoad(rounds, size))
			if reflect.DeepEqual(history, rounds) { // Check if equal to the 1000th rock arrangement
				repeating = true
			}
		}
	}
	index := (1000000000 - settleTime - 1) % len(loads)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", loads[index])
}

func main() {
	timeFunction(solve)
}
