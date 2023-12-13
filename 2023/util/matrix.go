package util

func Transpose[T any](m [][]T) [][]T {
	out := make([][]T, len(m[0]))
	for i := range out {
		out[i] = make([]T, len(m))
	}
	for i, row := range m {
		for j, cell := range row {
			out[j][i] = cell
		}
	}
	return out
}
