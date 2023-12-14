package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

func parse(line string) (string, []int) {
	a := strings.Split(line, " ")
	b := strings.Split(a[1], ",")
	nums := make([]int, 0)
	for _, s := range b {
		n, _ := strconv.Atoi(s)
		nums = append(nums, n)
	}
	return a[0], nums
}

func sum(arr []int) int {
	sum := 0
	for _, n := range arr {
		sum += n
	}
	return sum
}

var cache = make(map[string]int) // Cache of comibation counts to speed things up

func count(gears string, counts []int) int {
	key := gears + strings.Join(strings.Fields(fmt.Sprint(counts)), ",")
	n, cached := cache[key]
	if cached { // Cache hit
		return n
	}

	if len(counts) == 0 {
		if strings.Contains(gears, "#") {
			return 0
		}
		return 1
	}

	minGears := len(counts) + sum(counts) - 1
	total := 0
	for i := 0; i <= len(gears)-minGears; i++ {
		if i > 0 && gears[i-1] == '#' {
			break
		}
		if !strings.Contains(gears[i:i+counts[0]], ".") {
			if i+counts[0] == len(gears) && len(counts) == 1 {
				total++
			} else if gears[i+counts[0]] != '#' {
				total += count(gears[i+counts[0]+1:], counts[1:])
			}
		}
	}
	cache[key] = total
	return total
}

func solve() {
	lines := utils.ReadInput("input.txt", "\n")

	part1, part2 := 0, 0
	for _, line := range lines {
		gears, counts := parse(line)
		part1 += count(gears, counts)
		newGears, newCounts := gears, counts
		for i := 1; i < 5; i++ {
			newGears += "?" + gears
			newCounts = append(newCounts, counts...)
		}
		part2 += count(newGears, newCounts)
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func main() {
	utils.TimeFunction(solve)
}
