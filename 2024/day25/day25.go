package main

import (
	"aoc2024/util"
	"bytes"
	"fmt"
)

func main() {
	blocks := bytes.Split(util.ReadFile(), []byte("\n\n"))
	grids := make([]util.ByteMatrix, len(blocks))
	for i, block := range blocks {
		grids[i] = util.ByteMatrix(bytes.Split(block, []byte("\n")))
	}
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
