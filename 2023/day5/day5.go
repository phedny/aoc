package main

import (
	"aoc2023/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	lines := util.ReadLines()

	seedsA := make(map[int]int)
	for _, seed := range strings.Split(lines[0], " ")[1:] {
		n, _ := strconv.Atoi(seed)
		seedsA[n] = 1
	}

	seedsB := make(map[int]int)
	var start int
	for i, s := range strings.Split(lines[0], " ")[1:] {
		if i%2 == 0 {
			start, _ = strconv.Atoi(s)
		} else {
			n, _ := strconv.Atoi(s)
			seedsB[start] = n
		}
	}

	seedsA = mapSeeds(seedsA, lines[1:])
	seedsB = mapSeeds(seedsB, lines[1:])

	minA := math.MaxInt
	for seed := range seedsA {
		if seed < minA {
			minA = seed
		}
	}
	fmt.Println(minA)

	minB := math.MaxInt
	for seed := range seedsB {
		if seed < minB {
			minB = seed
		}
	}
	fmt.Println(minB)
}

func mapSeeds(seeds map[int]int, lines []string) map[int]int {
	outSeeds := make(map[int]int)
	for _, line := range lines {
		vs := strings.Split(line, " ")
		switch len(vs) {
		case 1: // empty line
		case 2: // new section
			for start, length := range seeds {
				outSeeds[start] = length
			}
			seeds = outSeeds
			outSeeds = make(map[int]int)
		case 3: // map seeds
			dstStart, _ := strconv.Atoi(vs[0])
			srcStart, _ := strconv.Atoi(vs[1])
			length, _ := strconv.Atoi(vs[2])

			newSeeds := make(map[int]int)
			for seedStart, seedLength := range seeds {
				in, out := splitRange(seedStart, seedLength, srcStart, length)
				for rStart, rLength := range in {
					outSeeds[rStart-srcStart+dstStart] = rLength
				}
				for rStart, rLength := range out {
					newSeeds[rStart] = rLength
				}
			}
			seeds = newSeeds
		}
	}
	for start, length := range seeds {
		outSeeds[start] = length
	}
	return outSeeds
}

func splitRange(seedStart, seedLength, srcStart, srcLength int) (in map[int]int, out map[int]int) {
	switch {
	case seedStart+seedLength <= srcStart:
		return nil, map[int]int{seedStart: seedLength}
	case srcStart+srcLength <= seedStart:
		return nil, map[int]int{seedStart: seedLength}
	case seedStart < srcStart && seedStart+seedLength > srcStart+srcLength:
		return map[int]int{srcStart: srcLength}, map[int]int{seedStart: srcStart - seedStart, srcStart + srcLength: seedStart + seedLength - (srcStart + srcLength)}
	case seedStart < srcStart:
		return map[int]int{srcStart: seedStart + seedLength - srcStart}, map[int]int{seedStart: srcStart - seedStart}
	case seedStart+seedLength <= srcStart+srcLength:
		return map[int]int{seedStart: seedLength}, nil
	default:
		return map[int]int{seedStart: srcStart + srcLength - seedStart}, map[int]int{srcStart + srcLength: seedStart + seedLength - (srcStart + srcLength)}
	}
}
