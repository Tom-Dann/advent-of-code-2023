package main

import (
	"fmt"
	"os"
	"sort"
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

func handType(card string, jokers bool) int {
	cardCount := make(map[rune]int)
	for _, s := range card {
		cardCount[s]++
	}
	jokerCount := 0
	if jokers {
		jokerCount = cardCount['J']
		cardCount['J'] = 0
	}
	counts := make([]int, 0, 5)
	for _, v := range cardCount {
		counts = append(counts, v)
	}
	sort.Slice(counts, func(i, j int) bool { // Sort in descending order
		return counts[i] > counts[j]
	})
	counts[0] += jokerCount
	if counts[0] == 5 {
		return 7 // Five-of-a-kind
	}
	if counts[0] == 4 {
		return 6 // Four-of-a-kind
	}
	if counts[0] == 3 {
		if counts[1] == 2 {
			return 5 // Full house
		}
		return 4 // Three-of-a-kind
	}
	if counts[0] == 2 {
		if counts[1] == 2 {
			return 3 // Two pair
		}
		return 2 // One pair
	}
	return 1 // High card
}

func cardVal(card byte, jokers bool) int {
	ordering := "23456789TJQKA" // Part 1
	if jokers {
		ordering = "J23456789TQKA" // Part 2
	}
	return strings.Index(ordering, string(card))
}

func compareHands(a Hand, b Hand, jokers bool) bool {
	if a.Type > b.Type { // First hand is better
		return false
	}
	if a.Type < b.Type { // Second hand is better
		return true
	}
	for i := 0; i < len(a.Cards); i++ { // Compare cards in tie-break
		valA := cardVal(a.Cards[i], jokers)
		valB := cardVal(b.Cards[i], jokers)
		if valA > valB { // First hand wins on tie-break
			return false
		}
		if valA < valB { // Second hand wins on tie-break
			return true
		}
	}
	return false // Hands tie
}

func totalWinnings(cards []Hand) int {
	sum := 0
	for rank, card := range cards {
		sum += card.Bid * (rank + 1)
	}
	return sum
}

type Hand struct {
	Cards string
	Bid   int
	Type  int
}

func solve() {
	raw, err := os.ReadFile("input.txt") // Read file
	check(err)
	input := strings.Split(strings.TrimSpace(string(raw)), "\n")

	hands1 := make([]Hand, len(input)) // Cards for part 1
	hands2 := make([]Hand, len(input)) // Cards for part 2
	for i, line := range input {       // Parse cards into a map
		s := strings.Split(line, " ")
		n, _ := strconv.Atoi(s[1])
		hands1[i] = Hand{Cards: s[0], Bid: n, Type: handType(s[0], false)}
		hands2[i] = Hand{Cards: s[0], Bid: n, Type: handType(s[0], true)}
	}

	sort.Slice(hands1, func(i, j int) bool {
		return compareHands(hands1[i], hands1[j], false)
	})
	sort.Slice(hands2, func(i, j int) bool {
		return compareHands(hands2[i], hands2[j], true)
	})

	fmt.Println("Part 1:", totalWinnings(hands1))
	fmt.Println("Part 2:", totalWinnings(hands2))
}

func main() {
	timeFunction(solve)
}
