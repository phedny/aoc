package util

type Lines []string

func (l Lines) Get(r, c int) (byte, bool) {
	if r >= 0 && r < len(l) && c >= 0 && c < len(l[r]) {
		return l[r][c], true
	}
	return 0, false
}

func (l Lines) AllCells(yield func(Cell[byte]) bool) {
	for r, line := range l {
		for c := range line {
			if !yield(Cell[byte]{l, r, c}) {
				return
			}
		}
	}
}
