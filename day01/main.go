package main

import (
	"fmt"
	"strconv"
	"utils"
)

func checkCharIsNumber(c byte) bool {
	return '0' <= c && c <= '9'
}

func part1(lines []string) {
	sum := 0
	for _, line := range lines {
		numBytes := make([]byte, 2)
		for i := 0; i < len(line); i++ { // Check from start of line
			if checkCharIsNumber(line[i]) {
				numBytes[0] = line[i]
				break
			}
		}
		for j := len(line) - 1; j >= 0; j-- { // Check from end of line
			if checkCharIsNumber(line[j]) {
				numBytes[1] = line[j]
				break
			}
		}
		num, _ := strconv.Atoi(string(numBytes))
		sum += num
	}
	fmt.Println("Part 1:", sum)
}

func part2(lines []string) {
	// Numbers as strings :)
	numberStrings := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	sum := 0
	for _, line := range lines {
		numBytes := make([]byte, 2)
		found := false
		for i := 0; i < len(line); i++ { // Check from start of line
			// Check for numeric value
			if checkCharIsNumber(line[i]) {
				numBytes[0] = line[i]
				found = true
			}
			// Check for number as string
			for numCheck, strCheck := range numberStrings {
				if line[i:min(i+len(strCheck), len(line))] == strCheck {
					numBytes[0] = fmt.Sprint(numCheck + 1)[0]
					found = true
				}
			}
			if found { // Exit loop if number is found
				break
			}
		}
		found = false
		for j := len(line) - 1; j >= 0; j-- { // Check from end of line
			// Check for numeric value
			if checkCharIsNumber(line[j]) {
				numBytes[1] = line[j]
				found = true
			}
			// Check for number as string
			for numCheck, strCheck := range numberStrings {
				if line[j:min(j+len(strCheck), len(line))] == strCheck {
					numBytes[1] = fmt.Sprint(numCheck + 1)[0]
					found = true
				}
			}
			if found { // Exit loop if number is found
				break
			}
		}
		num, _ := strconv.Atoi(string(numBytes))
		sum += num
	}
	fmt.Println("Part 2:", sum)
}

func main() {
	input := utils.ReadInput("input.txt", "\n")
	utils.TimeFunctionInput(part1, input)
	utils.TimeFunctionInput(part2, input)
}
