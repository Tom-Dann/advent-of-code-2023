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

func timeFunction(function func([]string), input []string) {
	start := time.Now()
	function(input)
	fmt.Println("Time elapsed:", time.Since(start))
}

func part1(sections []string) {
	var part1 int
	for _, seed := range strings.Split(sections[0], " ")[1:] { // Loop over seeds
		currVal, _ := strconv.Atoi(seed)
		for _, section := range sections[1:] { // Loop over maps
			mappings := strings.Split(section, "\n")
			for _, mapping := range mappings[1:] { // Loop over mappings in map
				vals := strings.Split(mapping, " ")
				min, _ := strconv.Atoi(vals[1])
				size, _ := strconv.Atoi(vals[2])
				if min <= currVal && currVal < min+size {
					dest, _ := strconv.Atoi(vals[0])
					currVal += dest - min
					break
				}
			}
		}
		if currVal < part1 || part1 == 0 {
			part1 = currVal
		}
	}

	fmt.Println("Part 1:", part1)
}

func part2(sections []string) {
	seedNums := strings.Split(sections[0], " ")[1:]
	intervals := make(map[int]int)
	for i := 0; i < len(seedNums)-1; i += 2 { // Loop over seed descriptions
		min, _ := strconv.Atoi(seedNums[i])
		length, _ := strconv.Atoi(seedNums[i+1])
		intervals[min] = length
	}
	for _, section := range sections[1:] { // Loop over maps
		mappings := strings.Split(section, "\n")
		mappedIntervals := make(map[int]int)
		intervalsToCheck := intervals
		for _, mapping := range mappings[1:] { // Loop over mappings in map
			vals := strings.Split(mapping, " ")
			dest, _ := strconv.Atoi(vals[0])
			sMin, _ := strconv.Atoi(vals[1])
			sLen, _ := strconv.Atoi(vals[2])
			unprocessed := make(map[int]int)
			for iMin, iLen := range intervalsToCheck {
				if sMin < iMin+iLen && sMin+sLen > iMin { // If ranges overlap
					/* Four possible scenarios
					  [------)  	<-- The interval we are checking
					[------)    	#1 Only start of interval is mapped
					[----------)	#2 All of interval is mapped
					    [------)	#3 Only end of interval is mapped
					    [--)    	#4 Only middle of interval is mapped
					*/
					if sMin <= iMin { // Start of interval is mapped - Scenarios #1 & #2
						if sMin+sLen < iMin+iLen { // End of interval is not mapped - Scenario #1
							mappedIntervals[dest+iMin-sMin] = sMin + sLen - iMin
							unprocessed[sMin+sLen] = iMin + iLen - sMin - sLen
						} else { // All of interval is mapped - Scenario #2
							mappedIntervals[dest+iMin-sMin] = iLen
						}
					} else { // Start of interval is not mapped - Scenarios #3 & #4
						unprocessed[iMin] = sMin - iMin
						if sMin+sLen >= iMin+iLen { // End of interval is mapped - Scenario #3
							mappedIntervals[dest] = iMin + iLen - sMin
						} else { // Only middle of interval is mapped - Scenario #4
							mappedIntervals[dest] = sLen
							unprocessed[sMin+sLen] = iMin + iLen - sLen - sMin
						}
					}
				} else { // Mapping does not overlap with interval
					unprocessed[iMin] = iLen
				}
			}
			intervalsToCheck = unprocessed
		}
		intervals = mappedIntervals
		for k, v := range intervalsToCheck {
			intervals[k] = v
		}
	}
	// Find min in keys of map
	var part2 int
	for k := range intervals {
		part2 = k
		break
	}
	for k := range intervals {
		if k < part2 {
			part2 = k
		}
	}
	fmt.Println("Part 2:", part2)
}

func main() {
	raw, err := os.ReadFile("input.txt") // Read file
	check(err)
	sections := strings.Split(strings.TrimSpace(string(raw)), "\n\n")

	timeFunction(part1, sections)
	timeFunction(part2, sections)
}
