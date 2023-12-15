package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"utils"
)

func hash(s string) int {
	val := 0
	for _, c := range s {
		val = ((val + int(c)) * 17) % 256
	}
	return val
}

type Lense struct {
	Label string
	Focus int
}

func checkLabel(label string) func(Lense) bool {
	return func(l Lense) bool { return l.Label == label }
}

func part1(input []string) {
	sum := 0
	for _, s := range input {
		sum += hash(s)
	}
	fmt.Println("Part 1:", sum)
}

func part2(input []string) {
	boxes := make([][]Lense, 256)
	for _, s := range input {
		if strings.Contains(s, "-") {
			label := strings.Split(s, "-")[0]
			hash := hash(label)
			slot := slices.IndexFunc(boxes[hash], checkLabel(label))
			if slot >= 0 {
				boxes[hash] = append(boxes[hash][:slot], boxes[hash][slot+1:]...)
			}
		} else {
			split := strings.Split(s, "=")
			label := split[0]
			focus, _ := strconv.Atoi(split[1])
			hash := hash(label)
			slot := slices.IndexFunc(boxes[hash], checkLabel(label))
			if slot >= 0 {
				boxes[hash][slot].Focus = focus
			} else {
				boxes[hash] = append(boxes[hash], Lense{label, focus})
			}
		}
	}

	power := 0
	for boxNo, box := range boxes {
		for slot, lense := range box {
			power += (boxNo + 1) * (slot + 1) * lense.Focus // Calculate focusing power
		}
	}
	fmt.Println("Part 2:", power)
}

func main() {
	strings := utils.ReadInput("input.txt", ",")
	utils.TimeFunctionInput(part1, strings)
	utils.TimeFunctionInput(part2, strings)
}
