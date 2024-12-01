package main

import (
	"aoc2024/util"
	"fmt"
	"slices"
)

func main() {
	var as, bs []int
	for _, line := range util.ReadLines() {
		var a, b int
		fmt.Sscan(line, &a, &b)
		as = append(as, a)
		bs = append(bs, b)
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
