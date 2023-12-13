package main

import (
	"aoc2023/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println(countOptions2("?#?#?#?#?#?#?#?", []int{1, 3, 1, 6}, ""))
	// fmt.Println(countOptions("???.##", []int{1, 1, 2}, ""))

	var sumA, sumB int
	for _, lineA := range util.ReadLines() {
		splitLine := strings.Split(lineA, " ")

		lineA = splitLine[0]
		strs := strings.Split(splitLine[1], ",")
		nsA := make([]int, len(strs))
		for i, str := range strs {
			nsA[i], _ = strconv.Atoi(str)
		}

		lineB := fmt.Sprintf("%s?%s?%s?%s?%s", lineA, lineA, lineA, lineA, lineA)
		nsB := make([]int, 0, 5*len(nsA))
		for i := 0; i < 5; i++ {
			nsB = append(nsB, nsA...)
		}

		sumA += countOptions(lineA, nsA)
		sumB += countOptions(lineB, nsB)
	}
	fmt.Println(sumA)
	fmt.Println(sumB)
}

var memo = make(map[string]int)

func countOptions(line string, ns []int) int {
	key := keyFor(line, ns)
	if r, has := memo[key]; has {
		return r
	}

	r := computeCountOptions(line, ns)
	memo[key] = r
	return r
}

func computeCountOptions(line string, ns []int) int {
	if len(line) == 0 {
		if len(ns) == 0 {
			return 1
		}
		return 0
	}

	countIfFirstIsGood := 0
	if line[0] != '#' {
		countIfFirstIsGood = countOptions(line[1:], ns)
	}

	countIfFirstIsBroken := 0
	if line[0] != '.' && len(ns) > 0 && len(line) >= ns[0] {
		hasGood := false
		for _, b := range line[:ns[0]] {
			if b == '.' {
				hasGood = true
			}
		}

		if !hasGood {
			if len(line) == ns[0] {
				countIfFirstIsBroken = countOptions("", ns[1:])
			} else if line[ns[0]] != '#' {
				countIfFirstIsBroken = countOptions(line[ns[0]+1:], ns[1:])
			}
		}
	}

	return countIfFirstIsGood + countIfFirstIsBroken
}

func keyFor(line string, ns []int) string {
	var s strings.Builder
	s.WriteString(line)
	for _, n := range ns {
		s.WriteByte(',')
		s.WriteString(strconv.Itoa(n))
	}
	return s.String()
}
