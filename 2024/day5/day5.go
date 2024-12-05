package main

import (
	"aoc2024/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := util.ReadLines()
	cutAt := slices.Index(lines, "")
	order := make(map[[2]int]bool)
	for _, line := range lines[:cutAt] {
		var a, b int
		fmt.Sscanf(line, "%d|%d", &a, &b)
		order[[2]int{a, b}] = true
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
	for _, line := range lines[cutAt+1:] {
		ss := strings.Split(line, ",")
		ns := make([]int, len(ss))
		for i, s := range ss {
			n, _ := strconv.Atoi(s)
			ns[i] = n
		}
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
