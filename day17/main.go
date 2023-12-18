package main

import (
	"container/heap"
	"fmt"
	"utils"
)

type Position struct {
	x, y, dir, len int
}

// Priority Queue implementation https://pkg.go.dev/container/heap
type Item struct {
	position Position
	dist     int // The priority of the item in the queue.
	index    int // The index of the item in the heap.
}
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i]; pq[i].index, pq[j].index = i, j }
func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
func (pq *PriorityQueue) update(item *Item, dist int) { item.dist = dist; heap.Fix(pq, item.index) }

func move(p Position, dir int) Position {
	switch dir {
	case 0: // East
		p.x++
	case 1: // South
		p.y++
	case 2: // West
		p.x--
	case 3: // North
		p.y--
	}
	if p.dir == dir {
		p.len++
	} else {
		p.len = 1
		p.dir = dir
	}
	return p
}

func getNeightbours(pos Position, part int) []Position {
	if part == 2 && pos.len <= 3 && pos.dir != -1 { // Must go a min of 4 in same direction in part 2
		b := move(pos, pos.dir)
		return []Position{b}
	}
	var next []Position
	for dir := 0; dir <= 3; dir++ {
		if pos.dir-dir%4 != 2 { // Not allowed to doube back
			next = append(next, move(pos, dir))
		}
	}
	return next
}

func dijkstra(lines []string, part int) int {
	max := struct{ width, height int }{len(lines[0]), len(lines)}
	heat := make([][]int, max.height)
	heatLoss := map[Position]int{}
	for i, line := range lines {
		heat[i] = make([]int, max.width)
		for j, char := range line {
			heat[i][j] = int(char) - 48
		}
	}
	queue := PriorityQueue{&Item{position: Position{0, 0, -1, 0}, dist: 0, index: 0}}
	heap.Init(&queue)

	valid := func(block Position) bool {
		if block.x < 0 || block.x >= max.width || block.y < 0 || block.y >= max.height {
			return false
		}
		if part == 1 {
			return block.len <= 3
		}
		return block.len <= 10
	}

	for queue.Len() > 0 {
		next := heap.Pop(&queue).(*Item)
		if next.position.x == max.width-1 && next.position.y == max.height-1 {
			return next.dist
		}
		for _, neighbour := range getNeightbours(next.position, part) {
			if valid(neighbour) {
				dist := next.dist + heat[neighbour.y][neighbour.x]
				best, visited := heatLoss[neighbour]
				if !visited || dist < best {
					heatLoss[neighbour] = dist
					heap.Push(&queue, &Item{position: neighbour, dist: dist})
				}
			}
		}
	}
	return -1
}

func part1(lines []string) {
	fmt.Println("Part 1:", dijkstra(lines, 1))
}

func part2(lines []string) {
	fmt.Println("Part 2:", dijkstra(lines, 2))
}

func main() {
	lines := utils.ReadInput("input.txt", "\n")
	utils.TimeFunctionInput(part1, lines)
	utils.TimeFunctionInput(part2, lines)
}
