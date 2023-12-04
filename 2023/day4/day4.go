package main

import (
	"aoc2023/util"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

var rLine = regexp.MustCompile(`Card +(\d+): ([0123456789 ]+) \| ([0123456789 ]+)`)
var rNumbers = regexp.MustCompile(`\d+`)

func main() {
	lines := util.ReadLines()
	cardCounts := make([]int, len(lines))

	var sumA, countB int
	for i, line := range util.ReadLines() {
		matches := matchCount(line)
		sumA += int(math.Floor(math.Pow(2, float64(matches-1))))

		thisCardCount := cardCounts[i] + 1
		countB += thisCardCount
		for j := 0; j < matches; j++ {
			cardCounts[i+j+1] += thisCardCount
		}
	}

	fmt.Println(sumA)
	fmt.Println(countB)
}

func matchCount(line string) int {
	m := rLine.FindStringSubmatch(line)
	winning := setOfNumbers(m[2])
	youHave := setOfNumbers(m[3])
	for n := range youHave {
		if !winning[n] {
			delete(youHave, n)
		}
	}
	return len(youHave)
}

func setOfNumbers(s string) map[int]bool {
	set := make(map[int]bool)
	for _, s := range rNumbers.FindAllString(s, -1) {
		n, _ := strconv.Atoi(s)
		set[n] = true
	}
	return set
}
