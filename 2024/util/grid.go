package util

import (
	"iter"
	"slices"
)

type Grid[T any] interface {
	Get(r, c int) (T, bool)
	AllCells(yield func(Cell[T]) bool)
}

type Coordinate struct {
	row    int
	column int
}

func (c Coordinate) Row() int {
	return c.row
}

func (c Coordinate) Column() int {
	return c.column
}

type Cell[T any] struct {
	grid        Grid[T]
	row, column int
}

func CellAt[T any](g Grid[T], r, c int) Cell[T] {
	return Cell[T]{g, r, c}
}

func (c Cell[T]) Coordinate() Coordinate {
	return Coordinate{c.row, c.column}
}

func (c Cell[T]) Row() int {
	return c.row
}

func (c Cell[T]) Column() int {
	return c.column
}

func (c Cell[T]) Get() T {
	v, _ := c.grid.Get(c.row, c.column)
	return v
}

func (c Cell[T]) Valid() bool {
	_, ok := c.grid.Get(c.row, c.column)
	return ok
}

func (c Cell[T]) Move(delta Coordinate) Cell[T] {
	return Cell[T]{c.grid, c.row + delta.row, c.column + delta.column}
}

func (c Cell[T]) Up() Cell[T] {
	return c.Move(Up)
}

func (c Cell[T]) Down() Cell[T] {
	return c.Move(Down)
}

func (c Cell[T]) Left() Cell[T] {
	return c.Move(Left)
}

func (c Cell[T]) Right() Cell[T] {
	return c.Move(Right)
}

func (c Cell[T]) UpLeft() Cell[T] {
	return c.Move(UpLeft)
}

func (c Cell[T]) UpRight() Cell[T] {
	return c.Move(UpRight)
}

func (c Cell[T]) DownLeft() Cell[T] {
	return c.Move(DownLeft)
}

func (c Cell[T]) DownRight() Cell[T] {
	return c.Move(DownRight)
}

func (c Cell[T]) MoveSeq(delta Coordinate) iter.Seq2[int, Cell[T]] {
	return func(yield func(int, Cell[T]) bool) {
		for i := 0; ; i++ {
			if !c.Valid() || !yield(i, c) {
				return
			}
			c = c.Move(delta)
		}
	}
}

func (c Cell[T]) MoveAll(deltas iter.Seq[Coordinate]) iter.Seq2[Coordinate, Cell[T]] {
	return func(yield func(Coordinate, Cell[T]) bool) {
		for delta := range deltas {
			if cell2 := c.Move(delta); cell2.Valid() && !yield(delta, cell2) {
				return
			}
		}
	}
}

func (cell Cell[T]) AllNeighbors(yield func(Coordinate, Cell[T]) bool) {
	cell.MoveAll(AllNeighbors)(yield)
}

func (cell Cell[T]) OrthogonalNeighbors(yield func(Coordinate, Cell[T]) bool) {
	cell.MoveAll(OrthogonalNeighbors)(yield)
}

func (cell Cell[T]) DiagonalNeighbors(yield func(Coordinate, Cell[T]) bool) {
	cell.MoveAll(DiagonalNeighbors)(yield)
}

var (
	Up                  = Coordinate{-1, 0}
	Down                = Coordinate{1, 0}
	Left                = Coordinate{0, -1}
	Right               = Coordinate{0, 1}
	UpLeft              = Coordinate{-1, -1}
	UpRight             = Coordinate{-1, 1}
	DownLeft            = Coordinate{1, -1}
	DownRight           = Coordinate{1, 1}
	AllNeighbors        = slices.Values([]Coordinate{Up, Down, Left, Right, UpLeft, UpRight, DownLeft, DownRight})
	OrthogonalNeighbors = slices.Values([]Coordinate{Up, Down, Left, Right})
	DiagonalNeighbors   = slices.Values([]Coordinate{UpLeft, UpRight, DownLeft, DownRight})
)
