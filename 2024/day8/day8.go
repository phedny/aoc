package main

import (
	"aoc2024/util"
	"fmt"
)

func main() {
	grid := util.ReadByteMatrix()
	m := make(map[byte][]util.Coordinate)
	for w := range grid.AllCells {
		if b := w.Get(); b != '.' {
			m[b] = append(m[b], w.Position())
		}
	}
	antinodesA := make(map[util.Coordinate]bool)
	antinodesB := make(map[util.Coordinate]bool)
	for _, antennas := range m {
		for i, a := range antennas {
			for _, b := range antennas[:i] {
				w := util.WalkGrid(grid, a, b.DistanceTo(a))
				for i, w := range w.MoveSeq {
					if i == 1 {
						antinodesA[w.Position()] = true
					}
					antinodesB[w.Position()] = true
				}
				w = w.TurnAround().MoveForwards()
				for i, w := range w.MoveSeq {
					if i == 1 {
						antinodesA[w.Position()] = true
					}
					antinodesB[w.Position()] = true
				}
			}
		}
	}
	fmt.Println(len(antinodesA))
	fmt.Println(len(antinodesB))
}
