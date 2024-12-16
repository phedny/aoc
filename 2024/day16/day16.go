package main

import (
	"aoc2024/util"
	"fmt"
	"maps"
	"math"
)

func main() {
	var start, end *Node
	nodes := make(map[util.Coordinate]*Node)
	for w := range util.ReadByteMatrix().AllCells {
		if b := w.Get(); b != '#' {
			node := &Node{w.Position(), make(map[util.Translation]Edge)}
			nodes[w.Position()] = node
			if b == 'S' {
				start = node
			} else if b == 'E' {
				end = node
			}
		}
	}
	for pos, n1 := range nodes {
		for t := range util.OrthogonalNeighbors {
			if n2, has := nodes[pos.Add(t)]; has {
				n1.edges[t] = Edge{node: n2, orientation: t, cost: 1}
			}
		}
	}
	for _, node := range nodes {
		for node != start && node != end && len(node.edges) < 2 {
			delete(nodes, node.pos)
			for _, edge := range node.edges {
				delete(edge.node.edges, edge.orientation.Negate())
				node = edge.node
			}
		}
	}
	for _, node := range nodes {
		if node != start && node != end && len(node.edges) == 2 {
			var e2, e3 Edge
			var t2, t3 util.Translation
			for t, n := range node.edges {
				if e2.node == nil {
					t2, e2 = t, n
				} else {
					t3, e3 = t, n
				}
			}
			cost := e2.cost + e3.cost
			if t2.Negate() != t3 {
				cost += 1000
			}
			skipped := map[util.Coordinate]bool{node.pos: true}
			maps.Copy(skipped, e2.skipped)
			maps.Copy(skipped, e3.skipped)
			e2.node.edges[e2.orientation.Negate()] = Edge{node: e3.node, orientation: e3.orientation, cost: cost, skipped: skipped}
			e3.node.edges[e3.orientation.Negate()] = Edge{node: e2.node, orientation: e2.orientation, cost: cost, skipped: skipped}
			delete(nodes, node.pos)
		}
	}
	reachable := map[Destination]*Path{{node: start, orientation: util.East}: &Path{cost: 0, extends: nil, visits: map[util.Coordinate]bool{start.pos: true}}}
	front := map[Destination]bool{{node: start, orientation: util.East}: true}
	for len(front) > 0 {
		newFront := make(map[Destination]bool)
		for step := range front {
			path := reachable[step]
			for t, edge := range step.node.edges {
				if step.orientation == t.Negate() {
					continue
				}
				newCost := path.cost + edge.cost
				if step.orientation != t {
					newCost += 1000
				}
				knownPath, has := reachable[Destination{edge.node, edge.orientation}]
				if !has || knownPath.cost > newCost {
					visits := maps.Clone(edge.skipped)
					visits[edge.node.pos] = true
					reachable[Destination{edge.node, edge.orientation}] = &Path{cost: newCost, extends: []*Path{path}, visits: visits}
					newFront[Destination{edge.node, edge.orientation}] = true
				} else if knownPath.cost == newCost {
					knownPath.extends = append(knownPath.extends, path)
					maps.Copy(knownPath.visits, edge.skipped)
					reachable[Destination{edge.node, edge.orientation}] = knownPath
				}
			}
		}
		front = newFront
	}
	cost := math.MaxInt
	var path *Path
	for t := range util.OrthogonalNeighbors {
		if candidate, has := reachable[Destination{end, t}]; has && candidate.cost < cost {
			cost = candidate.cost
			path = candidate
		}
	}
	fmt.Println(cost)
	visited := make(map[util.Coordinate]bool)
	paths := []*Path{path}
	for len(paths) > 0 {
		var newPaths []*Path
		for _, path := range paths {
			maps.Copy(visited, path.visits)
			newPaths = append(newPaths, path.extends...)
		}
		paths = newPaths
	}
	fmt.Println(len(visited))
}

type Node struct {
	pos   util.Coordinate
	edges map[util.Translation]Edge
}

type Edge struct {
	skipped     map[util.Coordinate]bool
	node        *Node
	orientation util.Translation
	cost        int
}

type Destination struct {
	node        *Node
	orientation util.Translation
}

type Path struct {
	visits  map[util.Coordinate]bool
	extends []*Path
	cost    int
}
