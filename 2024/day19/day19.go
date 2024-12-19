package main

import (
	"aoc2024/util"
	"fmt"
	"strings"
)

func main() {
	lines := util.ReadLines()
	start := &Node{next: make(map[byte][]*Node), final: true}
	for _, m := range strings.Split(lines[0], ", ") {
		n := start
	LineLoop:
		for _, r := range m[:len(m)-1] {
			next := n.next[byte(r)]
			if len(next) == 0 {
				nn := &Node{next: make(map[byte][]*Node)}
				n.next[byte(r)] = []*Node{nn}
				n = nn
			} else {
				for _, nn := range next {
					if nn != start {
						n = nn
						continue LineLoop
					}
				}
				nn := &Node{next: make(map[byte][]*Node)}
				n.next[byte(r)] = append(next, nn)
				n = nn
			}
		}
		n.next[m[len(m)-1]] = append(n.next[m[len(m)-1]], start)
	}

	var tallyA, tallyB int
	for _, line := range lines[2:] {
		nodes := map[*Node]int{start: 1}
		for _, r := range line {
			newNodes := make(map[*Node]int)
			for node, n := range nodes {
				for _, next := range node.next[byte(r)] {
					newNodes[next] += n
				}
			}
			nodes = newNodes
		}
		var tally int
		for node, n := range nodes {
			if node.final {
				tally += n
			}
		}
		if tally > 0 {
			tallyA++
		}
		tallyB += tally
	}
	fmt.Println(tallyA)
	fmt.Println(tallyB)
}

type Node struct {
	next  map[byte][]*Node
	final bool
}
