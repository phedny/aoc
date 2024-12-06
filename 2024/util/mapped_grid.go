package util

type MappedGrid[T1 comparable, T2 any] struct {
	g Grid[T1]
	m func(T1) (T2, bool)
}

func MapGridWithFunc[T1 comparable, T2 any](g Grid[T1], m func(T1) (T2, bool)) MappedGrid[T1, T2] {
	return MappedGrid[T1, T2]{g, m}
}

func MapGridWithMap[T1 comparable, T2 any](g Grid[T1], m map[T1]T2) MappedGrid[T1, T2] {
	return MappedGrid[T1, T2]{g, func(t1 T1) (T2, bool) { t2, ok := m[t1]; return t2, ok }}
}

func (g MappedGrid[T1, T2]) Get(r, c int) (T2, bool) {
	var v2 T2
	v1, ok := g.g.Get(r, c)
	if ok {
		v2, ok = g.m(v1)
	}
	return v2, ok
}

func (g MappedGrid[T1, T2]) MustGet(r, c int, def T2) T2 {
	v, ok := g.Get(r, c)
	if ok {
		return v
	}
	return def
}

func (g MappedGrid[T1, T2]) AllCells(yield func(GridWalker[T2]) bool) {
	for cell := range g.g.AllCells {
		if _, ok := g.m(cell.Get()); ok {
			if !yield(WalkGrid(g, cell.Position(), cell.Orientation())) {
				return
			}
		}
	}
}
