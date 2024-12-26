package main

import (
	"aoc2024/input"
	"fmt"
	"slices"
)

func main() {
	input := input.ReadDay5()
	order := make(map[[2]int]bool)
	for _, input := range input.Order {
		order[[2]int{input.V1, input.V2}] = true
	}
	cmp := func(a, b int) int {
		if order[[2]int{a, b}] {
			return -1
		} else if order[[2]int{b, a}] {
			return 1
		} else {
			return 0
		}
	}
	var tallyA, tallyB int
	for _, ns := range input.Productions {
		if slices.IsSortedFunc(ns, cmp) {
			tallyA += ns[(len(ns)-1)/2]
		} else {
			slices.SortFunc(ns, cmp)
			tallyB += ns[(len(ns)-1)/2]
		}
	}
	fmt.Println(tallyA)
	fmt.Println(tallyB)
}
