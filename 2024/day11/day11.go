package main

import (
	"aoc2024/input"
	"fmt"
	"strconv"
)

func main() {
	m := make(map[string]int)
	for _, s := range input.ReadDay11() {
		m[s]++
	}
	m, partA := stepsAndTally(m, 25)
	fmt.Println(partA)
	_, partB := stepsAndTally(m, 50)
	fmt.Println(partB)
}

func stepsAndTally(m map[string]int, n int) (map[string]int, int) {
	for range n {
		m = step(m)
	}
	var tally int
	for _, n := range m {
		tally += n
	}
	return m, tally
}

func step(in map[string]int) map[string]int {
	out := make(map[string]int)
	for s, n := range in {
		if s == "0" {
			out["1"] += n
		} else if i := len(s); i%2 == 0 {
			out[s[:i/2]] += n
			s = s[i/2:]
			for len(s) > 1 && s[0] == '0' {
				s = s[1:]
			}
			out[s] += n
		} else {
			num, _ := strconv.Atoi(s)
			out[strconv.Itoa(2024*num)] += n
		}
	}
	return out
}
