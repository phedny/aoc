package main

import (
	"aoc2023/util"
	"fmt"
	"regexp"
	"strconv"
)

var rLine = regexp.MustCompile(`Game (\d+): (.*)`)
var rSet = regexp.MustCompile(`(\d+) (red|green|blue)`)

func main() {
	var sumA, sumB int
	available := map[string]int{"red": 12, "green": 13, "blue": 14}
	for _, line := range util.ReadLines() {
		m := rLine.FindStringSubmatch(line)
		if isGamePossible(available, m[2]) {
			gameNr, _ := strconv.Atoi(m[1])
			sumA += gameNr
		}
		sumB += findGamePower(m[2])
	}
	fmt.Println(sumA)
	fmt.Println(sumB)
}

func isGamePossible(available map[string]int, counts string) bool {
	for _, m := range rSet.FindAllStringSubmatch(counts, -1) {
		n, _ := strconv.Atoi(m[1])
		if available[m[2]] < n {
			return false
		}
	}
	return true
}

func findGamePower(counts string) int {
	minimum := map[string]int{}
	for _, m := range rSet.FindAllStringSubmatch(counts, -1) {
		n, _ := strconv.Atoi(m[1])
		if minimum[m[2]] < n {
			minimum[m[2]] = n
		}
	}
	return minimum["red"] * minimum["green"] * minimum["blue"]
}
