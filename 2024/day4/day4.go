package main

import (
	"aoc2024/util"
	"fmt"
)

func main() {
	grid := util.ReadLines()
	fmt.Println(partA(grid))
	fmt.Println(partB(grid))
}

func partA(grid util.Grid[byte]) int {
	var tally int
	for c := range grid.AllCells {
		if c.Get() == 'X' {
			for delta := range c.AllNeighbors {
				for i, c := range c.MoveSeq(delta) {
					if c.Get() != "XMAS"[i] {
						break
					} else if i == 3 {
						tally++
						break
					}
				}
			}
		}
	}
	return tally
}

func partB(grid util.Grid[byte]) int {
	var tally int
	for c := range util.MapGridWithMap(grid, map[byte]int{'M': 1, 'A': 2, 'S': 4}).AllCells {
		if c.Get() == 2 && (c.MoveNW().Get()|c.MoveSE().Get())&(c.MoveNE().Get()|c.MoveSW().Get()) == 5 {
			tally++
		}
	}
	return tally
}
