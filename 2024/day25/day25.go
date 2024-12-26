package main

import (
	"aoc2024/input"
	"fmt"
)

func main() {
	grids := input.ReadDay25()
	var tally int
	for i, grid1 := range grids {
	GridLoop:
		for _, grid2 := range grids[i+1:] {
			for w := range grid1.AllCells {
				if b2, _ := grid2.Get(w.Position().Row(), w.Position().Column()); b2 == '#' && w.Get() == '#' {
					continue GridLoop
				}
			}
			tally++
		}
	}
	fmt.Println(tally)
}
