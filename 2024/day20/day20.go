package main

import (
	"aoc2024/input"
	"aoc2024/util"
	"fmt"
)

func main() {
	var w util.GridWalker[byte]
	for w = range input.ReadDay20().AllCells {
		if w.Get() == 'S' {
			break
		}
	}
	for t, w2 := range w.OrthogonalNeighbors {
		if w2.Get() == '.' {
			w = w.OrientTowards(t)
		}
	}

	grid := make(util.MapGrid[int])
	for w.Get() != 'E' {
		grid.Set(w.Position().Row(), w.Position().Column(), len(grid))
		w = w.MoveForwards()
		if w.Get() == '#' {
			w = w.MoveBackwards().RotateRight().MoveForwards()
			if w.Get() == '#' {
				w = w.MoveBackwards().TurnAround().MoveForwards()
			}
		}
	}
	grid.Set(w.Position().Row(), w.Position().Column(), len(grid))

	fmt.Println(countCheats(grid, 2, 100))
	fmt.Println(countCheats(grid, 20, 100))
}

func countCheats(grid util.Grid[int], maxDistance int, minAdvantage int) int {
	var tally int
	for w := range grid.AllCells {
		for t, w2 := range w.AllCellsInRange(maxDistance) {
			advantage := w2.Get() - w.Get() - abs(t.Row()) - abs(t.Column())
			if advantage >= minAdvantage {
				tally++
			}
		}
	}
	return tally
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
