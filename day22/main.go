package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"utils"
)

type Brick struct {
	min, max []int
	grounded bool
}
type Bricks []Brick

func parseCoord(s string) []int {
	numStrings := strings.Split(s, ",")
	nums := make([]int, 0)
	for _, str := range numStrings {
		n, _ := strconv.Atoi(str)
		nums = append(nums, n)
	}
	return nums
}

func parseBricks() Bricks {
	lines := utils.ReadInput("input.txt", "\n")
	bricks := make([]Brick, len(lines))
	for i, line := range lines {
		s := strings.Split(line, "~")
		min, max := parseCoord(s[0]), parseCoord(s[1])
		if min[2] == 1 {
			bricks[i] = Brick{min, max, true}
		} else {
			bricks[i] = Brick{min, max, false}
		}
	}
	return bricks
}

func (b Brick) fall() Brick {
	b.min[2]--
	b.max[2]--
	if b.min[2] == 1 {
		b.grounded = true
	}
	return b
}

func (a Brick) supportedBy(b Brick) bool {
	return !(a.min[0] > b.max[0] ||
		a.max[0] < b.min[0] ||
		a.min[1] > b.max[1] ||
		a.max[1] < b.min[1] ||
		a.min[2]-1 > b.max[2] ||
		a.max[2]-1 < b.min[2])
}

func supersetOf(a, b []int) bool {
	for _, n := range b {
		if !slices.Contains(a, n) {
			return false
		}
	}
	return true
}

func solve() {
	bricks := parseBricks()
	supportedBy := map[int][]int{}
	falling := true
	for falling {
		falling = false
		for i, brick := range bricks {
			if !brick.grounded {
				falling = true
				canFall := true
				for j, block := range bricks {
					if i != j {
						if brick.supportedBy(block) {
							canFall = false
							if block.grounded {
								brick.grounded = true
								bricks[i] = brick
								supportedBy[i] = append(supportedBy[i], j)
							}
						}
					}
				}
				if canFall {
					bricks[i] = brick.fall()
				}
			}
		}
	}

	part1, part2 := 0, 0
	for i := range bricks {
		fallingBricks := []int{i}
		chainReaction := true
		for chainReaction {
			chainReaction = false
			for j, supports := range supportedBy {
				if !slices.Contains(fallingBricks, j) && supersetOf(fallingBricks, supports) {
					fallingBricks = append(fallingBricks, j)
					chainReaction = true
				}
			}
		}
		if len(fallingBricks) == 1 {
			part1++
		} else {
			part2 += len(fallingBricks) - 1
		}
	}
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func main() {
	utils.TimeFunction(solve)
}
