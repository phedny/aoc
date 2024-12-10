package main

import (
	"aoc2024/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var safe0, safe1 int
	for _, line := range util.ReadLines() {
		var r []int
		for _, number := range strings.Split(line, " ") {
			n, _ := strconv.Atoi(number)
			r = append(r, n)
		}
		if PartA(r, safeUp) || PartA(r, safeDown) {
			safe0++
		} else if PartB(r, safeUp) || PartB(r, safeDown) {
			safe1++
		}
	}
	fmt.Println(safe0)
	fmt.Println(safe0 + safe1)
}

func PartA(r []int, f func(int, int) bool) bool {
	for i, v := range r[1:] {
		if !f(r[i], v) {
			return false
		}
	}
	return true
}

func PartB(r []int, f func(int, int) bool) bool {
	for i := range r {
		if PartA(slices.Delete(slices.Clone(r), i, i+1), f) {
			return true
		}
	}
	return false
}

func safeUp(a, b int) bool {
	return b-a >= 1 && b-a <= 3
}

func safeDown(a, b int) bool {
	return b-a >= -3 && b-a <= -1
}