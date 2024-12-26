package main

import (
	"aoc2024/input"
	"fmt"
	"regexp"
	"strconv"
)

var rMul = regexp.MustCompile(`(?:mul\((\d{1,3}),(\d{1,3})\))|(do\(\))|(don't\(\))`)

func main() {
	f := input.ReadDay3()
	ms := rMul.FindAllStringSubmatch(string(f), len(f))
	var tally int
	enabled := true
	for _, m := range ms {
		if m[3] != "" {
			enabled = true
		} else if m[4] != "" {
			enabled = false
		} else if enabled {
			a, err := strconv.Atoi(m[1])
			if err != nil {
				continue
			}
			b, err := strconv.Atoi(m[2])
			if err != nil {
				continue
			}
			tally += a * b
		}
	}
	fmt.Println(tally)
}
