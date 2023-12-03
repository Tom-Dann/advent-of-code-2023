package main

import (
	"fmt"
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

func solve() {
	// Read file
	raw, err := os.ReadFile("input.txt")
	check(err)
	lines := strings.Split(strings.TrimSpace(string(raw)), "\n")

	part1 := 0
	re := regexp.MustCompile(`(\d+)`)
	gears := make(map[string][]int)
	for index, line := range lines {
		for _, match := range re.FindAllStringIndex(line, -1) {
			n, _ := strconv.Atoi(line[match[0]:match[1]])
			isPartNo := false
			for i := max(index-1, 0); i <= min(index+1, len(lines)-1); i++ { // Loop vertically
				for j := max(match[0]-1, 0); j <= min(match[1], len(line)-1); j++ { // Loop horizontally
					symbol := string(lines[i][j])
					if !strings.Contains("0123456789.", symbol) { // Part 1 check for any symbols
						isPartNo = true
						if symbol == "*" { // Gear for part 2
							gearName := fmt.Sprint(i, "-", j)
							if len(gears[gearName]) == 0 {
								gears[gearName] = []int{1, n} // [count, product]
							} else {
								gears[gearName][0]++    // Increment count
								gears[gearName][1] *= n // Update product
							}
						}
					}
				}
			}
			if isPartNo {
				part1 += n
			}
		}
	}

	part2 := 0
	for _, gear := range gears {
		if gear[0] == 2 {
			part2 += gear[1]
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func main() {
	timeFunction(solve)
}
