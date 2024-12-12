package util

import (
	"iter"
	"slices"
)

type Grid[T any] interface {
	Get(r, c int) (T, bool)
	AllCells(yield func(GridWalker[T]) bool)
}

type SettableGrid[T any] interface {
	Set(r, c int, v T) bool
}

type UnsettableGrid[T any] interface {
	Unset(r, c int)
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

func (c Coordinate) Add(t Translation) Coordinate {
	return Coordinate{c.row + t.row, c.column + t.column}
}

func (c Coordinate) DistanceTo(c2 Coordinate) Translation {
	return Translation{c2.row - c.row, c2.column - c.column}
}

type Translation struct {
	row    int
	column int
}

func (t Translation) Row() int {
	return t.row
}

func (t Translation) Column() int {
	return t.column
}

func (t Translation) Negate() Translation {
	return t.Factor(-1)
}

func (t Translation) RotateRight() Translation {
	return Translation{t.column, -t.row}
}

func (t Translation) RotateLeft() Translation {
	return Translation{-t.column, t.row}
}

func (t Translation) Factor(n int) Translation {
	return Translation{n * t.row, n * t.column}
}

type GridWalker[T any] struct {
	grid        Grid[T]
	position    Coordinate
	orientation Translation
}

func WalkGrid[T any](g Grid[T], position Coordinate, orientation Translation) GridWalker[T] {
	return GridWalker[T]{g, position, orientation}
}

func (w GridWalker[T]) Grid() Grid[T] {
	return w.grid
}

func (w GridWalker[T]) Position() Coordinate {
	return w.position
}

func (w GridWalker[T]) Orientation() Translation {
	return w.orientation
}

func (w GridWalker[T]) Set(v T) bool {
	if grid, isSettable := w.grid.(SettableGrid[T]); isSettable {
		return grid.Set(w.position.row, w.position.column, v)
	}
	return false
}

func (w GridWalker[T]) Unset() {
	if grid, isUnsettable := w.grid.(UnsettableGrid[T]); isUnsettable {
		grid.Unset(w.position.row, w.position.column)
	}
}

func (w GridWalker[T]) Get() T {
	v, _ := w.grid.Get(w.position.row, w.position.column)
	return v
}

func (w GridWalker[T]) Valid() bool {
	_, ok := w.grid.Get(w.position.row, w.position.column)
	return ok
}

func (w GridWalker[T]) MoveTo(position Coordinate) GridWalker[T] {
	return GridWalker[T]{w.grid, position, w.orientation}
}

func (w GridWalker[T]) Move(translation Translation) GridWalker[T] {
	return GridWalker[T]{w.grid, Coordinate{w.position.row + translation.row, w.position.column + translation.column}, w.orientation}
}

func (w GridWalker[T]) MoveN() GridWalker[T] {
	return w.Move(North)
}

func (w GridWalker[T]) MoveS() GridWalker[T] {
	return w.Move(South)
}

func (w GridWalker[T]) MoveW() GridWalker[T] {
	return w.Move(West)
}

func (w GridWalker[T]) MoveE() GridWalker[T] {
	return w.Move(East)
}

func (w GridWalker[T]) MoveNW() GridWalker[T] {
	return w.Move(NorthWest)
}

func (w GridWalker[T]) MoveNE() GridWalker[T] {
	return w.Move(NorthEast)
}

func (w GridWalker[T]) MoveSW() GridWalker[T] {
	return w.Move(SouthWest)
}

func (w GridWalker[T]) MoveSE() GridWalker[T] {
	return w.Move(SouthEast)
}

func (w GridWalker[T]) MoveForwards() GridWalker[T] {
	return w.Move(w.orientation)
}

func (w GridWalker[T]) MoveBackwards() GridWalker[T] {
	return w.Move(w.orientation.Negate())
}

func (w GridWalker[T]) MoveLeft() GridWalker[T] {
	return w.Move(w.orientation.RotateLeft())
}

func (w GridWalker[T]) MoveRight() GridWalker[T] {
	return w.Move(w.orientation.RotateRight())
}

func (w GridWalker[T]) OrientTowards(orientation Translation) GridWalker[T] {
	return GridWalker[T]{w.grid, w.position, orientation}
}

func (w GridWalker[T]) TurnAround() GridWalker[T] {
	return GridWalker[T]{w.grid, w.position, w.orientation.Negate()}
}

func (w GridWalker[T]) RotateLeft() GridWalker[T] {
	return GridWalker[T]{w.grid, w.position, w.orientation.RotateLeft()}
}

func (w GridWalker[T]) RotateRight() GridWalker[T] {
	return GridWalker[T]{w.grid, w.position, w.orientation.RotateRight()}
}

func (w GridWalker[T]) MoveSeq(yield func(int, GridWalker[T]) bool) {
	for i := 0; ; i++ {
		if !w.Valid() || !yield(i, w) {
			return
		}
		w = w.Move(w.orientation)
	}
}

func (w GridWalker[T]) MoveAll(deltas iter.Seq[Translation]) iter.Seq2[Translation, GridWalker[T]] {
	return func(yield func(Translation, GridWalker[T]) bool) {
		for delta := range deltas {
			if cell2 := w.Move(delta); cell2.Valid() && !yield(delta, cell2) {
				return
			}
		}
	}
}

func (w GridWalker[T]) AllNeighbors(yield func(Translation, GridWalker[T]) bool) {
	w.MoveAll(AllNeighbors)(yield)
}

func (w GridWalker[T]) OrthogonalNeighbors(yield func(Translation, GridWalker[T]) bool) {
	w.MoveAll(OrthogonalNeighbors)(yield)
}

func (w GridWalker[T]) DiagonalNeighbors(yield func(Translation, GridWalker[T]) bool) {
	w.MoveAll(DiagonalNeighbors)(yield)
}

var (
	North               = Translation{-1, 0}
	South               = Translation{1, 0}
	West                = Translation{0, -1}
	East                = Translation{0, 1}
	NorthWest           = Translation{-1, -1}
	NorthEast           = Translation{-1, 1}
	SouthWest           = Translation{1, -1}
	SouthEast           = Translation{1, 1}
	AllNeighbors        = slices.Values([]Translation{North, South, West, East, NorthWest, NorthEast, SouthWest, SouthEast})
	OrthogonalNeighbors = slices.Values([]Translation{North, South, West, East})
	DiagonalNeighbors   = slices.Values([]Translation{NorthWest, NorthEast, SouthWest, SouthEast})
)
