package main

import (
	"aoc2024/input"
	"aoc2024/util"
	"fmt"
)

func main() {
	grid := make(util.MapGrid[byte])
	for w := range input.ReadDay12().AllCells {
		grid.Set(w.Position().Row(), w.Position().Column(), w.Get())
	}
	var tallyA, tallyB int
	for len(grid) > 0 {
		for w := range grid.AllCells {
			region := make(map[util.Coordinate]bool)
			fences := make(map[Fence]bool)
			fill(w, region, fences)
			tallyA += len(region) * len(fences)
			tallyB += len(region) * countSides(fences)
			break
		}
	}
	fmt.Println(tallyA)
	fmt.Println(tallyB)
}

func fill(w util.GridWalker[byte], region map[util.Coordinate]bool, fences map[Fence]bool) {
	region[w.Position()] = true
	v := w.Get()
	w.Unset()
	for t := range util.OrthogonalNeighbors {
		w2 := w.Move(t)
		if _, has := region[w2.Position()]; has {
			continue
		} else if !w2.Valid() || w2.Get() != v {
			fences[Fence{w.Position(), t}] = true
		} else {
			fill(w2, region, fences)
		}
	}
}

func countSides(fences map[Fence]bool) int {
	var sides int
	for len(fences) > 0 {
		for f := range fences {
			delete(fences, f)
			sides++
			left := f.Orientation.RotateLeft()
			for p := f.Position.Add(left); ; p = p.Add(left) {
				if _, has := fences[Fence{p, f.Orientation}]; has {
					delete(fences, Fence{p, f.Orientation})
				} else {
					break
				}
			}
			right := f.Orientation.RotateRight()
			for p := f.Position.Add(right); ; p = p.Add(right) {
				if _, has := fences[Fence{p, f.Orientation}]; has {
					delete(fences, Fence{p, f.Orientation})
				} else {
					break
				}
			}
			break
		}
	}
	return sides
}

type Fence struct {
	Position    util.Coordinate
	Orientation util.Translation
}
