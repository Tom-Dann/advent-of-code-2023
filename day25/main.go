package main

import (
	"fmt"
	"slices"
	"strings"
	"utils"
)

func printGraphviz(lines []string) {
	fmt.Println("graph network {")
	for _, line := range lines {
		s := strings.Split(line, ": ")
		fmt.Printf("	%s -- { %s };\n", s[0], s[1])
	}
	fmt.Println("}")
}

func cut(m map[string][]string, a string, b string) {
	i, j := slices.Index(m[a], b), slices.Index(m[b], a)
	m[a] = append(m[a][:i], m[a][i+1:]...)
	m[b] = append(m[b][:j], m[b][j+1:]...)
}

func floodFill(connections map[string][]string, start string) int {
	visited := map[string]struct{}{start: {}}
	queue := []string{start}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, dest := range connections[curr] {
			_, seen := visited[dest]
			if !seen {
				queue = append(queue, dest)
				visited[dest] = struct{}{}
			}
		}
	}
	return len(visited)
}

func solve(lines []string) {
	connections := map[string][]string{}

	for _, line := range lines {
		s := strings.Split(line, ": ")
		for _, target := range strings.Split(s[1], " ") {
			connections[s[0]] = append(connections[s[0]], target)
			connections[target] = append(connections[target], s[0])
		}
	}

	cut(connections, "zlx", "chr")
	cut(connections, "hqp", "spk")
	cut(connections, "cpq", "hlx")

	fmt.Println("Answer:", floodFill(connections, "zlx")*floodFill(connections, "hlx"))
}

func main() {
	lines := utils.ReadInput("input.txt", "\n")
	utils.TimeFunctionInput(solve, lines)
}
