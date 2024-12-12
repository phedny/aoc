package util

type MapGrid[T any] map[Coordinate]T

func (m MapGrid[T]) Get(r, c int) (T, bool) {
	v, has := m[Coordinate{r, c}]
	return v, has
}

func (m MapGrid[T]) Set(r, c int, v T) bool {
	m[Coordinate{r, c}] = v
	return true
}

func (m MapGrid[T]) Unset(r, c int) {
	delete(m, Coordinate{r, c})
}

func (m MapGrid[T]) AllCells(yield func(GridWalker[T]) bool) {
	for c := range m {
		if !yield(WalkGrid(m, c, Translation{})) {
			return
		}
	}
}
