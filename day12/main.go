package main

import (
	"fmt"
	"os"
	"strconv"
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

func countArragements(gears string, counts []int) int {
	check := gears + strings.Join(strings.Fields(fmt.Sprint(counts)), ",")
	n, cached := cache[check]
	if cached { // Cache hit
		return n
	}
	n = count(gears, counts)
	cache[check] = n
	return n
}

func count(gears string, counts []int) int {
	if len(counts) == 0 {
		if strings.Contains(gears, "#") {
			return 0
		}
		return 1
	}

	minGears := len(counts) + sum(counts) - 1
	total := 0
	for i := 0; len(gears)-i >= minGears; i++ {
		if i > 0 && gears[i-1] == '#' {
			break
		}
		if !strings.Contains(gears[i:i+counts[0]], ".") {
			if i+counts[0] == len(gears) && len(counts) == 1 {
				total++
			} else if gears[i+counts[0]] != '#' {
				total += countArragements(gears[i+counts[0]+1:], counts[1:])
			}
		}
	}
	return total
}

func solve() {
	raw, err := os.ReadFile("input.txt") // Read file
	check(err)
	lines := strings.Split(strings.TrimSpace(string(raw)), "\n")

	part1, part2 := 0, 0
	for _, line := range lines {
		gears, counts := parse(line)
		part1 += countArragements(gears, counts)
		newGears, newCounts := gears, counts
		for i := 1; i < 5; i++ {
			newGears += "?" + gears
			newCounts = append(newCounts, counts...)
		}
		part2 += countArragements(newGears, newCounts)
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func main() {
	timeFunction(solve)
}
