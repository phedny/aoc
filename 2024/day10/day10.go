package main

import (
	"aoc2024/input"
	"aoc2024/util"
	"fmt"
	"iter"
	"maps"
)

func main() {
	grid := input.ReadDay10()

	fmt.Println(run(grid,
		func(c util.Coordinate) map[util.Coordinate]bool { return map[util.Coordinate]bool{c: true} },
		func(a, b map[util.Coordinate]bool) map[util.Coordinate]bool {
			return maps.Collect(concat(maps.All(a), maps.All(b)))
		},
		func(vs map[util.Coordinate]bool) int { return len(vs) },
	))

	fmt.Println(run(grid,
		func(c util.Coordinate) int { return 1 },
		func(a, b int) int { return a + b },
		func(vs int) int { return vs },
	))
}

func run[T any](grid util.Grid[byte], init func(util.Coordinate) T, step func(T, T) T, count func(T) int) int {
	front := make(map[util.Coordinate]T)
	for w := range grid.AllCells {
		if w.Get() == '9' {
			front[w.Position()] = init(w.Position())
		}
	}
	for b := byte('8'); b >= '0'; b-- {
		newFront := make(map[util.Coordinate]T)
		for w, vs := range front {
			for _, w2 := range util.WalkGrid(grid, w, util.Translation{}).OrthogonalNeighbors {
				if w2.Get() == b {
					newFront[w2.Position()] = step(newFront[w2.Position()], vs)
				}
			}
		}
		front = newFront
	}
	var tally int
	for _, vs := range front {
		tally += count(vs)
	}
	return tally
}

func concat[K comparable, V any](seqs ...iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, seq := range seqs {
			for k, v := range seq {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}
