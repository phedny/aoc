package main

import (
	"aoc2024/input"
	"fmt"
	"strconv"
)

func main() {
	var tallyA, tallyB int
	for _, line := range input.ReadDay7() {
		if run(line.Target, line.Values, sum, product) {
			tallyA += line.Target
		}
		if run(line.Target, line.Values, sum, product, concat) {
			tallyB += line.Target
		}
	}
	fmt.Println(tallyA)
	fmt.Println(tallyB)
}

func run(target int, ns []int, ops ...func(a, b int) int) bool {
	carries := map[int]bool{ns[0]: true}
	for _, n := range ns[1:] {
		newCarries := make(map[int]bool)
		for carry := range carries {
			for _, op := range ops {
				newCarry := op(carry, n)
				if newCarry <= target {
					newCarries[newCarry] = true
				}
			}
		}
		carries = newCarries
	}
	return carries[target]
}

func sum(a, b int) int {
	return a + b
}

func product(a, b int) int {
	return a * b
}

func concat(a, b int) int {
	n, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	return n
}
