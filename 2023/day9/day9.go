package main

import (
	"aoc2023/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var sumA, sumB int
	for _, line := range util.ReadLines() {
		strs := strings.Split(line, " ")
		ns := make([]int, len(strs))
		for i, str := range strs {
			ns[i], _ = strconv.Atoi(str)
		}
		sumA += predictForward(ns)
		sumB += predictBackward(ns)
	}
	fmt.Println(sumA)
	fmt.Println(sumB)
}

func predictForward(seq []int) int {
	if isZero(seq) {
		return 0
	}
	return seq[len(seq)-1] + predictForward(derive(seq))
}

func predictBackward(seq []int) int {
	if isZero(seq) {
		return 0
	}
	return seq[0] - predictBackward(derive(seq))
}

func derive(seq []int) []int {
	derivative := make([]int, len(seq)-1)
	for i, n := range seq[1:] {
		derivative[i] = n - seq[i]
	}
	return derivative
}

func isZero(seq []int) bool {
	for _, n := range seq {
		if n != 0 {
			return false
		}
	}
	return true
}
