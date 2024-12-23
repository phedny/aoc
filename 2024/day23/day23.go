package main

import (
	"aoc2024/util"
	"fmt"
	"iter"
	"slices"
	"strings"
)

func main() {
	g := make(Graph)
	for _, link := range util.ReadLines() {
		g.Connect(link[0:2], link[3:5])
	}

	fmt.Println(partA(g))
	fmt.Println(partB(g))
}

func partA(g Graph) int {
	var tally int
	for clique := range g.FindCliques(3) {
		if slices.ContainsFunc(clique, func(name string) bool { return name[0] == 't' }) {
			tally++
		}
	}
	return tally
}

func partB(g Graph) string {
	var max int
	for _, nodes := range g {
		if len(nodes) > max {
			max = len(nodes)
		}
	}

	for {
		for clique := range g.FindCliques(max) {
			return strings.Join(clique, ",")
		}
		max--
	}
}

type Node map[string]bool

type Graph map[string]Node

func (g Graph) Node(name string) Node {
	if node, has := g[name]; has {
		return node
	}
	node := make(Node)
	g[name] = node
	return node
}

func (g Graph) Connect(a, b string) {
	g.Node(a)[b] = true
	g.Node(b)[a] = true
}

func (g Graph) FindCliques(n int) iter.Seq[[]string] {
	return func(yield func([]string) bool) {
		g.findCliques(nil, n, yield)
	}
}

func (g Graph) findCliques(clique []string, n int, yield func([]string) bool) bool {
	if n == 0 {
		return yield(clique)
	}

	for name, node := range g {
		if !slices.Contains(clique, name) {
			var count int
			for name2 := range node {
				if name2 < name && slices.Contains(clique, name2) {
					count++
				}
			}
			if count == len(clique) {
				if !g.findCliques(append(clique, name), n-1, yield) {
					return false
				}
			}
		}
	}
	return true
}
