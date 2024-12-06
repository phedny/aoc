package main

import (
	"aoc2024/util"
	"fmt"
)

func main() {
	grid := util.ReadByteMatrix()
	var w util.GridWalker[byte]
	for w = range grid.AllCells {
		if w.Get() == '^' {
			break
		}
	}
	w = w.OrientTowards(util.North)
	path := path(w)
	fmt.Println(len(path))

	var tally int
	for pos := range path {
		obstacle := util.WalkGrid(grid, pos, util.Translation{})
		if obstacle.Get() == '.' {
			obstacle.Set('#')
			if hasLoop(w) {
				tally++
			}
			obstacle.Set('.')
		}
	}
	fmt.Println(tally)
}

func path(w util.GridWalker[byte]) map[util.Coordinate]bool {
	seenPositions := make(map[util.Coordinate]bool)
	for {
		switch w.Get() {
		case '#':
			w = w.MoveBackwards().RotateRight()
		case 0:
			return seenPositions
		default:
			seenPositions[w.Position()] = true
			w = w.MoveForwards()
		}
	}
}

func hasLoop(w util.GridWalker[byte]) bool {
	seenWalkers := make(map[PositionAndOrientation]bool)
	for {
		switch w.Get() {
		case '#':
			w = w.MoveBackwards().RotateRight()
		case 0:
			return false
		default:
			pao := PositionAndOrientation{w.Position(), w.Orientation()}
			if seenWalkers[pao] {
				return true
			}
			seenWalkers[pao] = true
			w = w.MoveForwards()
		}
	}
}

type PositionAndOrientation struct {
	position    util.Coordinate
	orientation util.Translation
}
