package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

func zeroArray(arr []int) bool {
	for _, n := range arr {
		if n != 0 {
			return false
		}
	}
	return true
}

func solve() {
	input := utils.ReadInput("input.txt", "\n")
	part1, part2 := 0, 0
	for _, line := range input {
		nums := make([]int, 0)
		for _, s := range strings.Split(line, " ") { // Parse input nums
			n, _ := strconv.Atoi(s)
			nums = append(nums, n)
		}

		for i := 0; !zeroArray(nums); i++ { // Find differences until all zero
			next := make([]int, 0)
			for j := 0; j < len(nums)-1; j++ {
				next = append(next, nums[j+1]-nums[j])
			}
			part1 += nums[len(nums)-1]
			if i%2 == 0 {
				part2 += nums[0]
			} else {
				part2 -= nums[0]
			}
			nums = next
		}
	}

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func main() {
	utils.TimeFunction(solve)
}
