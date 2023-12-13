package main

import (
	"aoc2023/util"
	"fmt"
)

func main() {
	maps := [][][]byte{{}}
	for _, line := range util.ReadLines() {
		if line == "" {
			maps = append(maps, [][]byte{})
		} else {
			maps[len(maps)-1] = append(maps[len(maps)-1], []byte(line))
		}
	}

	var countA, countB int
	for _, m := range maps {
		countA += mapValue(m, 0)
		countB += mapValue(m, 1)
	}
	fmt.Println(countA)
	fmt.Println(countB)
}

func mapValue(m [][]byte, t int) int {
	i, has := findReflection(convertMap(m), t)
	if has {
		return 100 * i
	}

	i, has = findReflection(convertMap(util.Transpose(m)), t)
	if has {
		return i
	}

	panic("no reflection found")
}

func convertMap(m [][]byte) []uint {
	out := make([]uint, len(m))
	for i, line := range m {
		var n uint
		for _, b := range line {
			n *= 2
			if b == '#' {
				n++
			}
		}
		out[i] = n
	}
	return out
}

func findReflection(m []uint, t int) (int, bool) {
	for i := 0; i < len(m)-1; i++ {
		if hasReflection(m, i, t) {
			return i + 1, true
		}
	}
	return 0, false
}

func hasReflection(m []uint, i int, t int) bool {
	t *= 2
	for f := 0; f < len(m); f++ {
		b := 2*i + 1 - f
		if b < 0 || b >= len(m) {
			continue
		}
		if m[f] != m[b] {
			if t == 0 {
				return false
			}
			x := m[f] ^ m[b]
			if x&(x-1) != 0 {
				return false
			}
			t--
		}
	}
	return t == 0
}
