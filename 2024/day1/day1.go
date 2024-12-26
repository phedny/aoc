package main

import (
	"aoc2024/input"
	"fmt"
	"slices"
)

func main() {
	var as, bs []int
	for _, line := range input.ReadDay1() {
		as = append(as, line.V1)
		bs = append(bs, line.V2)
	}

	fmt.Println(partA(as, bs))
	fmt.Println(partB(as, bs))
}

func partA(as, bs []int) int {
	slices.Sort(as)
	slices.Sort(bs)
	var tally int
	for i, a := range as {
		b := bs[i]
		if a > b {
			tally += a - b
		} else {
			tally += b - a
		}
	}
	return tally
}

func partB(as, bs []int) int {
	hist := make(map[int]int)
	for _, b := range bs {
		hist[b]++
	}
	var tally int
	for _, a := range as {
		tally += a * hist[a]
	}
	return tally
}
