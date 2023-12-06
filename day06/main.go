package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
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

func calc(t int, d int) int {
	delta := math.Sqrt(float64((t * t) - (4 * d))) // Root of discriminant in quadratic formula
	min := int(math.Floor((float64(t)-delta)/2) + 1)
	max := int(math.Ceil((float64(t)+delta)/2) - 1)
	return max - min + 1
}

func parseNum(s string) int {
	n, _ := strconv.Atoi(strings.Split(strings.ReplaceAll(s, " ", ""), ":")[1])
	return n
}

func solve() {
	raw, err := os.ReadFile("input.txt") // Read file
	check(err)
	input := strings.Split(strings.TrimSpace(string(raw)), "\n")

	re := regexp.MustCompile(`\d+`)
	times := re.FindAllString(input[0], -1)
	distances := re.FindAllString(input[1], -1)
	time2, dist2 := parseNum(input[0]), parseNum(input[1])

	part1 := 1
	for i := 0; i < len(times); i++ {
		time, _ := strconv.Atoi(times[i])
		dist, _ := strconv.Atoi(distances[i])
		part1 *= calc(time, dist)
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", calc(time2, dist2))
}

func main() {
	timeFunction(solve)
}
