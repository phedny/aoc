package main

import (
	"aoc2024/input"
	"fmt"
)

func main() {
	var tally int
	totalGains := make(map[[4]int]int)
	for _, n := range input.ReadDay22() {
		var seq [4]int
		for i := range 3 {
			nextN := next(n)
			seq[i+1] = nextN%10 - n%10
			n = nextN
		}
		gains := make(map[[4]int]int)
		for range 1997 {
			nextN := next(n)
			copy(seq[:3], seq[1:])
			seq[3] = nextN%10 - n%10
			if _, has := gains[seq]; !has {
				gains[seq] = nextN % 10
			}
			n = nextN
		}
		tally += n
		for seq, gain := range gains {
			totalGains[seq] += gain
		}
	}
	fmt.Println(tally)
	var max int
	for _, gain := range totalGains {
		if gain > max {
			max = gain
		}
	}
	fmt.Println(max)
}

func next(n int) int {
	n = ((n << 6) ^ n) % 0x1000000
	n = ((n >> 5) ^ n) % 0x1000000
	n = ((n << 11) ^ n) % 0x1000000
	return n
}
