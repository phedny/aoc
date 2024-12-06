package util

type Lines []string

func (l Lines) Get(r, c int) (byte, bool) {
	if r >= 0 && r < len(l) && c >= 0 && c < len(l[r]) {
		return l[r][c], true
	}
	return 0, false
}

func (l Lines) AllCells(yield func(GridWalker[byte]) bool) {
	for r, line := range l {
		for c := range line {
			if !yield(WalkGrid(l, Coordinate{r, c}, Translation{})) {
				return
			}
		}
	}
}

type ByteMatrix [][]byte

func (m ByteMatrix) Get(r, c int) (byte, bool) {
	if r >= 0 && r < len(m) && c >= 0 && c < len(m[r]) {
		return m[r][c], true
	}
	return 0, false
}

func (m ByteMatrix) Set(r, c int, b byte) bool {
	if r >= 0 && r < len(m) && c >= 0 && c < len(m[r]) {
		m[r][c] = b
		return true
	}
	return false
}

func (m ByteMatrix) AllCells(yield func(GridWalker[byte]) bool) {
	for r, line := range m {
		for c := range line {
			if !yield(WalkGrid(m, Coordinate{r, c}, Translation{})) {
				return
			}
		}
	}
}
