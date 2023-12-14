package main

import (
	"fmt"
	"regexp"
	"strings"
	"utils"
)

func solve() {
	cards := utils.ReadInput("input.txt", "\n")

	part1, part2 := 0, 0
	cardCount := make([]int, len(cards))
	re := regexp.MustCompile(`\d+`)
	for i, card := range cards {
		cardStrings := strings.Split(strings.Split(card, ": ")[1], " | ")
		winning := map[string]struct{}{}
		for _, s := range re.FindAllString(cardStrings[0], -1) { // Make a set of winning numbers
			winning[s] = struct{}{}
		}
		winCount := 0
		for _, s := range re.FindAllString(cardStrings[1], -1) {
			_, win := winning[s] // Check numbers against the winning set
			if win {
				winCount++
			}
		}
		part1 += (1 << winCount) >> 1
		cardCount[i]++
		part2 += cardCount[i]
		for j := 1; j <= winCount; j++ {
			cardCount[i+j] += cardCount[i]
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func main() {
	utils.TimeFunction(solve)
}
