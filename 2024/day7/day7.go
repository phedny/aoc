package main

import (
	"aoc2024/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var tallyA, tallyB uint64
	for _, line := range util.ReadLines() {
		strs := strings.Split(line, " ")
		target, _ := strconv.ParseUint(strs[0][:len(strs[0])-1], 10, 64)
		ns := make([]uint64, len(strs)-1)
		for i, str := range strs[1:] {
			ns[i], _ = strconv.ParseUint(str, 10, 64)
		}
		if run(target, ns, sum, product) {
			tallyA += target
		}
		if run(target, ns, sum, product, concat) {
			tallyB += target
		}
	}
	fmt.Println(tallyA)
	fmt.Println(tallyB)
}

func run(target uint64, ns []uint64, ops ...func(a, b uint64) uint64) bool {
	carries := map[uint64]bool{ns[0]: true}
	for _, n := range ns[1:] {
		newCarries := make(map[uint64]bool)
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

func sum(a, b uint64) uint64 {
	return a + b
}

func product(a, b uint64) uint64 {
	return a * b
}

func concat(a, b uint64) uint64 {
	n, _ := strconv.ParseUint(strconv.FormatUint(a, 10)+strconv.FormatUint(b, 10), 10, 64)
	return n
}
