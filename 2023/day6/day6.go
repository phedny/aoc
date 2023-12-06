package main

import (
	"aoc2023/util"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var rNumbers = regexp.MustCompile(`\d+`)

func main() {
	lines := util.ReadLines()
	times, distances := parseNumbers(lines[0]), parseNumbers(lines[1])

	productA := 1
	for i, time := range times {
		productA *= countWinning(float64(time), float64(distances[i]))
	}
	fmt.Println(productA)

	timeB, _ := strconv.Atoi(strings.Join(strings.Split(lines[0], " "), "")[5:])
	distanceB, _ := strconv.Atoi(strings.Join(strings.Split(lines[1], " "), "")[9:])
	fmt.Println(countWinning(float64(timeB), float64(distanceB)))
}

func parseNumbers(line string) []int {
	var ns []int
	for _, s := range rNumbers.FindAllString(line, -1) {
		n, _ := strconv.Atoi(s)
		ns = append(ns, n)
	}
	return ns
}

func countWinning(time, distance float64) int {
	r := math.Sqrt(time*time - 4*distance)
	v1, v2 := (time-r)/2, (time+r)/2
	extra := 0.
	if v2-math.Floor(v2) == 0 {
		extra = 1.
	}
	return int(math.Ceil(v2) - math.Ceil(v1) - extra)
}
