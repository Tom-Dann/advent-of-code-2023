package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"utils"
)

func parseWorkflow(s string) (string, []string) {
	line := strings.Split(s, "{")
	return line[0], strings.Split(line[1][0:len(line[1])-1], ",")
}

func getInt(s string) int {
	n, err := strconv.Atoi(s)
	utils.Check(err)
	return n
}

func parsePart(line string) []int {
	re := regexp.MustCompile(`{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}`)
	m := re.FindAllStringSubmatch(line, -1)
	return []int{getInt(m[0][1]), getInt(m[0][2]), getInt(m[0][3]), getInt(m[0][4])}
}

func eval(part []int, rule string) bool {
	if strings.Contains(rule, "<") {
		s := strings.Split(rule, "<")
		val, n := part[strings.Index("xmas", s[0])], getInt(s[1])
		return val < n
	}
	s := strings.Split(rule, ">")
	val, n := part[strings.Index("xmas", s[0])], getInt(s[1])
	return val > n
}

func process(part []int, name string, workflows map[string][]string) bool {
	switch name {
	case "A":
		return true
	case "R":
		return false
	}
	workflow := workflows[name]
	for i := 0; i < len(workflow)-1; i++ {
		instruction := strings.Split(workflow[i], ":")
		if eval(part, instruction[0]) {
			return process(part, instruction[1], workflows)
		}
	}
	return process(part, workflow[len(workflow)-1], workflows)
}

func sumCombinations(ranges [][]int) int {
	sum := 1
	for _, interval := range ranges {
		if interval[0] <= interval[1] {
			sum *= interval[1] - interval[0] + 1
		} else {
			return 0
		}
	}
	return sum
}

func deepCopy(ranges [][]int) [][]int {
	new := make([][]int, len(ranges))
	for i := range ranges {
		new[i] = make([]int, len(ranges[i]))
		copy(new[i], ranges[i])
	}
	return new
}

func combinations(name string, ranges [][]int, workflows map[string][]string) int {
	switch name {
	case "A":
		return sumCombinations(ranges)
	case "R":
		return 0
	}
	workflow := workflows[name]
	sum := 0
	for i := 0; i < len(workflow)-1; i++ {
		instruction := strings.Split(workflow[i], ":")
		rule, next := instruction[0], instruction[1]
		altPos := deepCopy(ranges)
		if strings.Contains(rule, "<") {
			s := strings.Split(rule, "<")
			index, n := strings.Index("xmas", s[0]), getInt(s[1])
			altPos[index][1] = n - 1 // Update max in alternative ranges
			ranges[index][0] = n     // Update min in current ranges
		} else {
			s := strings.Split(rule, ">")
			index, n := strings.Index("xmas", s[0]), getInt(s[1])
			altPos[index][0] = n + 1 // Update min in alternative ranges
			ranges[index][1] = n     // Update max in current ranges
		}
		sum += combinations(next, altPos, workflows)
	}
	return sum + combinations(workflow[len(workflow)-1], ranges, workflows)
}

func solve(sections []string) {
	workflows := map[string][]string{}
	for _, line := range strings.Split(sections[0], "\n") {
		name, instructions := parseWorkflow(line)
		workflows[name] = instructions
	}

	part1 := 0
	for _, line := range strings.Split(sections[1], "\n") {
		part := parsePart(line)
		if process(part, "in", workflows) {
			for _, n := range part {
				part1 += n
			}
		}
	}
	fmt.Println("Part 1:", part1)

	ranges := make([][]int, 4)
	for i := range ranges {
		ranges[i] = []int{1, 4000} // Min 1, max 4000
	}
	fmt.Println("Part 2:", combinations("in", ranges, workflows))
}

func main() {
	sections := utils.ReadInput("input.txt", "\n\n")
	utils.TimeFunctionInput(solve, sections)
}
