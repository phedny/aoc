package main

import (
	"aoc2024/util"
	"bytes"
	"fmt"
)

func main() {
	fileParts := bytes.SplitN(util.ReadFile(), []byte("\n\n"), 2)
	grid := util.ByteMatrix(bytes.Split(fileParts[0], []byte("\n")))

	for _, setup := range []func(util.Grid[byte]) util.GridWalker[Box]{partA, partB} {
		robot := setup(grid)
		for _, instr := range fileParts[1] {
			robot = robot.OrientTowards(orientations[instr])
			if robot.Orientation() == (util.Translation{}) {
				continue
			}
			robot = robot.MoveForwards()
			if !robot.Valid() {
				robot = robot.MoveBackwards()
				continue
			}
			if b := robot.Get(); b != nil {
				boxes := make(map[Box]bool)
				if !b.CanMove(robot.Orientation(), boxes) {
					robot = robot.MoveBackwards()
					continue
				}
				for b := range boxes {
					b.RemoveFromGrid()
				}
				for b := range boxes {
					b.Move(robot.Orientation())
				}
				for b := range boxes {
					b.AddToGrid()
				}
			}
		}

		boxes := make(map[Box]bool)
		for w := range robot.Grid().AllCells {
			boxes[w.Get()] = true
		}
		delete(boxes, nil)
		var tally int
		for b := range boxes {
			tally += 100*b.Position().Row() + b.Position().Column()
		}
		fmt.Println(tally)
	}
}

func partA(grid util.Grid[byte]) util.GridWalker[Box] {
	mGrid := make(util.MapGrid[Box])
	var robot util.GridWalker[Box]
	for w := range grid.AllCells {
		switch w.Get() {
		case '.':
			mGrid.Set(w.Position().Row(), w.Position().Column(), nil)
		case 'O':
			b := NewNarrowBox(mGrid, util.NewCoordinate(w.Position().Row(), w.Position().Column()))
			b.AddToGrid()
		case '@':
			mGrid.Set(w.Position().Row(), w.Position().Column(), nil)
			robot = util.WalkGrid(mGrid, w.Position(), util.Translation{})
		}
	}
	return robot
}

func partB(grid util.Grid[byte]) util.GridWalker[Box] {
	mGrid := make(util.MapGrid[Box])
	var robot util.GridWalker[Box]
	for w := range grid.AllCells {
		switch w.Get() {
		case '.':
			mGrid.Set(w.Position().Row(), 2*w.Position().Column(), nil)
			mGrid.Set(w.Position().Row(), 2*w.Position().Column()+1, nil)
		case 'O':
			b := NewWideBox(mGrid, util.NewCoordinate(w.Position().Row(), 2*w.Position().Column()))
			b.AddToGrid()
		case '@':
			mGrid.Set(w.Position().Row(), 2*w.Position().Column(), nil)
			mGrid.Set(w.Position().Row(), 2*w.Position().Column()+1, nil)
			robot = util.WalkGrid(mGrid, util.NewCoordinate(w.Position().Row(), 2*w.Position().Column()), util.Translation{})
		}
	}
	return robot
}

type Box interface {
	AddToGrid()
	RemoveFromGrid()
	Position() util.Coordinate
	Move(util.Translation)
	CanMove(util.Translation, map[Box]bool) bool
}

type NarrowBox struct {
	grid     util.MapGrid[Box]
	position util.Coordinate
}

func NewNarrowBox(grid util.MapGrid[Box], position util.Coordinate) Box {
	return &NarrowBox{grid: grid, position: position}
}

func (b *NarrowBox) AddToGrid() {
	b.grid.Set(b.position.Row(), b.position.Column(), b)
}

func (b *NarrowBox) RemoveFromGrid() {
	b.grid.Set(b.position.Row(), b.position.Column(), nil)
}

func (b *NarrowBox) Position() util.Coordinate {
	return b.position
}

func (b *NarrowBox) Move(t util.Translation) {
	b.position = b.position.Add(t)
}

func (b *NarrowBox) CanMove(t util.Translation, boxes map[Box]bool) bool {
	b2, valid := b.grid.Get(b.position.Row()+t.Row(), b.position.Column()+t.Column())
	if !valid {
		return false
	}
	if b2 != nil && !b2.CanMove(t, boxes) {
		return false
	}
	boxes[b] = true
	return true
}

type WideBox struct {
	grid     util.MapGrid[Box]
	position util.Coordinate
}

func NewWideBox(grid util.MapGrid[Box], position util.Coordinate) Box {
	b := &WideBox{grid: grid, position: position}
	return b
}

func (b *WideBox) AddToGrid() {
	b.grid.Set(b.position.Row(), b.position.Column(), b)
	b.grid.Set(b.position.Row(), b.position.Column()+1, b)

}

func (b *WideBox) RemoveFromGrid() {
	b.grid.Set(b.position.Row(), b.position.Column(), nil)
	b.grid.Set(b.position.Row(), b.position.Column()+1, nil)
}

func (b *WideBox) Position() util.Coordinate {
	return b.position
}

func (b *WideBox) Move(t util.Translation) {
	b.position = b.position.Add(t)
}

func (b *WideBox) CanMove(t util.Translation, boxes map[Box]bool) bool {
	b2, valid2 := b.grid.Get(b.position.Row()+t.Row(), b.position.Column()+t.Column())
	b3, valid3 := b.grid.Get(b.position.Row()+t.Row(), b.position.Column()+t.Column()+1)
	if !valid2 || !valid3 {
		return false
	}
	if b2 != nil && b2 != b && !b2.CanMove(t, boxes) {
		return false
	}
	if b3 != nil && b3 != b && !b3.CanMove(t, boxes) {
		return false
	}
	boxes[b] = true
	return true
}

var orientations = map[byte]util.Translation{'^': util.North, 'v': util.South, '<': util.West, '>': util.East}
