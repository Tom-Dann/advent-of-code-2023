package main

import (
	"fmt"
	"slices"
	"strings"
	"utils"
)

type Module struct {
	Type        rune
	Outputs     []string
	State, Last bool            // For flip-flop
	Memory      map[string]bool // For conjunction
}
type Modules map[string]Module

func parseInput() Modules {
	lines := utils.ReadInput("input.txt", "\n")
	modules := map[string]Module{}
	for _, line := range lines {
		s := strings.Split(line, " -> ")
		outputs := strings.Split(s[1], ", ")
		id := s[0]
		if id == "broadcaster" {
			modules[id] = Module{Type: '*', Outputs: outputs}
		} else {
			modules[id[1:]] = Module{Type: rune(id[0]), Outputs: outputs, Memory: map[string]bool{}}
		}
	}
	for k, v := range modules {
		for _, output := range v.Outputs {
			m := modules[output]
			if m.Type == '&' {
				m.Memory[k] = false
				modules[output] = m
			}
		}
	}
	return modules
}

type Signal struct {
	Source, Dest string
	Pulse        bool
}

func getSignals(source string, outputs []string, pulse bool) []Signal {
	signals := make([]Signal, len(outputs))
	for i, dest := range outputs {
		signals[i] = Signal{source, dest, pulse}
	}
	return signals
}

func (modules Modules) sendSignal(signal Signal) {
	m := modules[signal.Dest]
	if m.Type == '&' {
		m.Memory[signal.Source] = signal.Pulse
	} else {
		m.Last = signal.Pulse
	}
	modules[signal.Dest] = m
}

func (modules Modules) process(name string) []Signal {
	curr := modules[name]
	switch curr.Type {
	case '%':
		if !curr.Last { // Low pulse
			curr.State = !curr.State
			modules[name] = curr
			return getSignals(name, curr.Outputs, curr.State)
		}
		return []Signal{}
	case '&':
		pulse := true
		for _, v := range curr.Memory {
			pulse = pulse && v
		}
		return getSignals(name, curr.Outputs, !pulse)
	}
	return getSignals(name, curr.Outputs, false)
}

func solve() {
	modules := parseInput()
	cycles := make([]int, 4)
	nodes := []string{"dl", "ns", "bh", "vd"}
	low, high := 0, 0
	for i := 1; i <= 1000 || slices.Contains(cycles, 0); i++ {
		queue := []Signal{{"button", "broadcaster", false}} // Initial button press
		for len(queue) > 0 {
			next := queue[0]
			queue = queue[1:]
			// Counts for part 1
			if i <= 1000 {
				if next.Pulse {
					high++
				} else {
					low++
				}
			}
			// Counts for part 2
			if slices.Contains(nodes, next.Source) && next.Pulse {
				index := slices.Index(nodes, next.Source)
				if cycles[index] == 0 {
					cycles[index] = i
				}
			}
			modules.sendSignal(next)
			queue = append(queue, modules.process(next.Dest)...)
		}
	}
	fmt.Println("Part 1:", low*high)
	fmt.Println("Part 2:", cycles[0]*cycles[1]*cycles[2]*cycles[3])
}

func printMermaid() { // Used for inspecting graph manually for part 2
	fmt.Println("```mermaid\ngraph TD;")
	modules := parseInput()
	for source, m := range modules {
		sourceName := fmt.Sprintf("%s[%s%s]", source, string(m.Type), source)
		for _, dest := range m.Outputs {
			fmt.Printf("    %s-->%s;\n", sourceName, dest)
		}
	}
	fmt.Println("```")
}

func main() {
	utils.TimeFunction(solve)
}
