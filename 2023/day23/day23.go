package main

import (
	"aoc2023/util"
	"fmt"
)

var directions = map[byte][2]int{'^': {-1, 0}, 'v': {1, 0}, '<': {0, -1}, '>': {0, 1}}
var reversals = map[byte]byte{'^': 'v', 'v': '^', '<': '>', '>': '<'}

func main() {
	m := util.ReadByteMatrix()
	fmt.Println(findLongestPath(m))

	for y, line := range m {
		for x, b := range line {
			if directions[b] != [2]int{0, 0} {
				m[y][x] = '.'
			}
		}
	}
	fmt.Println(findLongestPath(m))
}

func findLongestPath(m [][]byte) int {
	graph := makeGraph(m, [2]int{0, 1}, [2]int{len(m) - 1, len(m[0]) - 2})
	_, distance := findLongestPathInGraph(graph, [2]int{0, 1}, [2]int{len(m) - 1, len(m[0]) - 2}, make(map[[2]int]bool))
	return distance
}

func makeGraph(m [][]byte, from, exit [2]int) map[[2]int]map[byte]node {
	graph := make(map[[2]int]map[byte]node)
	nodesToDo := map[[2]int]bool{{0, 1}: true}

	for len(nodesToDo) > 0 {
		for from := range nodesToDo {
			delete(nodesToDo, from)
			graph[from] = make(map[byte]node)
			for direction, delta := range directions {
				next := [2]int{from[0] + delta[0], from[1] + delta[1]}
				if next[0] < 0 || next[1] < 0 {
					continue
				}
				wouldEnter := m[next[0]][next[1]]
				if wouldEnter == '#' || wouldEnter == reversals[direction] {
					continue
				}

				to, distance := traversePath(m, from, exit, direction)
				if distance == -1 {
					continue
				}
				graph[from][direction] = node{to, distance}
				if to != exit && !nodesToDo[to] && graph[to] == nil {
					nodesToDo[to] = true
				}
			}
		}
	}

	return graph
}

func traversePath(m [][]byte, from, exit [2]int, direction byte) ([2]int, int) {
	distance := 1
	pos := [2]int{from[0] + directions[direction][0], from[1] + directions[direction][1]}
	visited := map[[2]int]bool{from: true}
	for {
		visited[pos] = true
		candidates := make(map[byte][2]int)
		for direction, delta := range directions {
			next := [2]int{pos[0] + delta[0], pos[1] + delta[1]}
			if visited[next] || next[0] < 0 || next[1] < 0 {
				continue
			}
			if next == exit {
				return next, distance + 1
			}
			wouldEnter := m[next[0]][next[1]]
			if wouldEnter == '#' || wouldEnter == reversals[direction] {
				continue
			}
			candidates[direction] = next
		}

		switch len(candidates) {
		case 0:
			return pos, -1
		case 1:
			for _, next := range candidates {
				pos = next
				distance++
			}
		default:
			return pos, distance
		}
	}
}

func findLongestPathInGraph(graph map[[2]int]map[byte]node, from, exit [2]int, seen map[[2]int]bool) (bool, int) {
	if from == exit {
		return true, 0
	}

	var max int
	seen[from] = true
	for _, to := range graph[from] {
		if !seen[to.destination] {
			finishes, distance := findLongestPathInGraph(graph, to.destination, exit, seen)
			if !finishes {
				continue
			}
			length := to.length + distance
			if length > max {
				max = length
			}
		}
	}
	delete(seen, from)

	return max > 0, max
}

type node struct {
	destination [2]int
	length      int
}
