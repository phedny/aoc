package main

import (
	"aoc2023/util"
	"fmt"
)

func main() {
	lines := util.ReadByteMatrix()
	y, x := findStart(lines)
	ch := canReachIn(lines, y, x)
	discard(ch, 64)
	fmt.Println(<-ch)

	var ns [3]int
	for i := range ns {
		ns[i] = <-ch
		discard(ch, 130)
	}

	p := extendSequence(ns[:])
	discard(p, 202300)
	fmt.Println(<-p)
}

func findStart(lines [][]byte) (int, int) {
	for y, line := range lines {
		for x, b := range line {
			if b == 'S' {
				return y, x
			}
		}
	}
	panic("no start found")
}

func canReachIn(lines [][]byte, y, x int) <-chan int {
	a := make(map[[2]int]bool)
	b := make(map[[2]int]bool)
	front := make(map[[2]int]bool)
	front[[2]int{y, x}] = true

	ch := make(chan int, 1)

	go func() {
		for {
			newFront := make(map[[2]int]bool)
			for c := range front {
				a[c] = true
				for _, c := range [][2]int{{c[0] - 1, c[1]}, {c[0] + 1, c[1]}, {c[0], c[1] - 1}, {c[0], c[1] + 1}} {
					y, x := c[0]%len(lines), c[1]%len(lines[0])
					if y < 0 {
						y += len(lines)
					}
					if x < 0 {
						x += len(lines[0])
					}
					if lines[y][x] == '#' || a[c] {
						continue
					}
					newFront[c] = true
				}
			}
			front, a, b = newFront, b, a
			ch <- len(b)
		}
	}()
	return ch
}

func discard(ch <-chan int, n int) {
	for ; n > 0; n-- {
		<-ch
	}
}

func extendSequence(seq []int) <-chan int {
	out := make(chan int, 1)

	go func() {
		if isZero(seq) {
			for {
				out <- 0
			}
		}

		acc := seq[0]
		for n := range extendSequence(derive(seq)) {
			out <- acc
			acc += n
		}
	}()

	return out
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
