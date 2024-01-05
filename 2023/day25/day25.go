package main

import (
	"aoc2023/util"
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	for !solve() {
	}
}

func solve() bool {
	g := make(graph)
	for _, line := range util.ReadLines() {
		l := strings.Split(line, " ")
		a := l[0][:len(l[0])-1]
		for _, b := range l[1:] {
			g.addEdge(a, b)
		}
	}

	var vertices []*vertex
	for _, v := range g {
		if len(v.edges) > 4 {
			vertices = append(vertices, v)
		}
	}

	var v1, v2 *vertex
	for {
		v1, v2 = randomItem(vertices), randomItem(vertices)
		if v1 == v2 {
			continue
		}

		ps := v1.findParallelPathsTo(v2)
		if len(ps) == 3 {
			for _, p := range ps {
				for i, v := range p[1:] {
					g.removeEdge(p[i].name, v.name)
				}
			}
			break
		}
	}

	n := len(v1.findReachableVertices())
	if n == 0 || n == len(g) {
		return false
	}

	fmt.Println(n * (len(g) - n))
	return true
}

func randomItem[T any](s []T) T {
	return s[rand.Intn(len(s))]
}

type graph map[string]*vertex

func (g graph) getVertex(n string) *vertex {
	v, ok := g[n]
	if !ok {
		v = &vertex{name: n, edges: make(map[string]*vertex)}
		g[n] = v
	}
	return v
}

func (g graph) addEdge(a, b string) {
	vA, vB := g.getVertex(a), g.getVertex(b)
	_, ok := vA.edges[b]
	if ok {
		return
	}
	vA.edges[b] = vB
	vB.edges[a] = vA
}

func (g graph) removeEdge(a, b string) {
	vA, vB := g.getVertex(a), g.getVertex(b)
	delete(vA.edges, b)
	delete(vB.edges, a)
}

type vertex struct {
	name  string
	edges map[string]*vertex
}

func (from *vertex) findParallelPathsTo(to *vertex) [][]*vertex {
	visited := map[*vertex]bool{from: true}
	finished := map[*vertex]bool{}

	var paths [][]*vertex
	partialPaths := make(map[*path]bool)
	for _, v := range from.edges {
		partialPaths[&path{p: &path{s: v, v: from}, s: v, v: v}] = true
	}

	for len(partialPaths) > 0 {
		nextPaths := make(map[*path]bool)
		for p := range partialPaths {
			for _, next := range p.v.edges {
				switch {
				case visited[next]:
				case finished[p.s]:
				case next == to:
					vs := []*vertex{to}
					for p1 := p; p1 != nil; p1 = p1.p {
						vs = append(vs, p1.v)
					}
					paths = append(paths, vs)
					finished[p.s] = true
				default:
					nextPaths[&path{p: p, s: p.s, v: next}] = true
					visited[next] = true
				}
			}
		}
		partialPaths = nextPaths
	}

	return paths
}

func (from *vertex) findReachableVertices() map[*vertex]bool {
	front := map[*vertex]bool{from: true}
	found := map[*vertex]bool{from: true}

	for len(front) > 0 {
		newFront := map[*vertex]bool{}
		for v := range front {
			for _, to := range v.edges {
				if !found[to] {
					found[to] = true
					newFront[to] = true
				}
			}
		}
		front = newFront
	}

	return found
}

type path struct {
	p    *path
	s, v *vertex
}
