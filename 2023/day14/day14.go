package main

import (
	"aoc2023/util"
	"fmt"
	"strings"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type segment struct {
	ort, par int
	length   int
}

func main() {
	lines := util.ReadByteMatrix()
	horMap, horSegments := defineSegments(lines)
	verMap, verSegments := defineSegments(util.Transpose(lines))
	verMap = util.Transpose(verMap)

	horToVer := make([][]int, len(horSegments))
	for i, from := range horSegments {
		horToVer[i] = make([]int, from.length)
		for j := range horToVer[i] {
			horToVer[i][j] = verMap[from.ort][from.par+j]
		}
	}

	verToHor := make([][]int, len(verSegments))
	for i, from := range verSegments {
		verToHor[i] = make([]int, from.length)
		for j := range verToHor[i] {
			verToHor[i][j] = horMap[from.par+j][from.ort]
		}
	}

	balls := loadRoundedRocks(verMap, lines)
	fmt.Println(computeVerticalLoad(len(lines), verSegments, balls))

	balls = firstCycle(horToVer, verToHor, balls)

	seen := make(map[string]int)
	for i := 1; i < 1000000000; i++ {
		key := ballsKey(balls)
		prevI, hasSeen := seen[key]
		if hasSeen {
			cLength := i - prevI
			cSkip := (1000000000 - i) / cLength
			i += cLength * cSkip
			seen = make(map[string]int)
		}
		seen[key] = i
		balls = cycle(horToVer, verToHor, balls)
	}
	fmt.Println(computeHorizontalLoad(len(lines), horSegments, balls))
}

func defineSegments(lines [][]byte) ([][]int, []segment) {
	newSegment := true
	segmentMap := make([][]int, len(lines))
	var segments []segment
	for y, line := range lines {
		newSegment = true
		segmentMap[y] = make([]int, len(line))
		for x, b := range line {
			if b == '#' {
				newSegment = true
				segmentMap[y][x] = -1
			} else {
				if newSegment {
					segmentMap[y][x] = len(segments)
					segments = append(segments, segment{y, x, 1})
				} else {
					segmentMap[y][x] = len(segments) - 1
					segments[len(segments)-1].length++
				}
				newSegment = false
			}
		}
	}
	return segmentMap, segments
}

func loadRoundedRocks(verticalSegments [][]int, lines [][]byte) map[int]int {
	out := make(map[int]int)
	for i, line := range lines {
		for j, b := range line {
			if b == 'O' {
				out[verticalSegments[i][j]]++
			}
		}
	}
	return out
}

func computeVerticalLoad(height int, segments []segment, balls map[int]int) int {
	load := 0
	for s, count := range balls {
		segment := segments[s]
		for i := 0; i < count; i++ {
			load += height - segment.par - i
		}
	}
	return load
}

func computeHorizontalLoad(height int, segments []segment, balls map[int]int) int {
	load := 0
	for s, count := range balls {
		load += count * (height - segments[s].ort)
	}
	return load
}

func firstCycle(horToVer, verToHor [][]int, balls map[int]int) map[int]int {
	balls = mapNormal(verToHor, balls)
	balls = mapNormal(horToVer, balls)
	balls = mapReversed(verToHor, balls)
	return balls
}

func cycle(horToVer, verToHor [][]int, balls map[int]int) map[int]int {
	balls = mapReversed(horToVer, balls)
	balls = mapNormal(verToHor, balls)
	balls = mapNormal(horToVer, balls)
	balls = mapReversed(verToHor, balls)
	return balls
}

func mapNormal(mapping [][]int, balls map[int]int) map[int]int {
	out := make(map[int]int)
	for segment, count := range balls {
		for i := 0; i < count; i++ {
			out[mapping[segment][i]]++
		}
	}
	return out
}

func mapReversed(mapping [][]int, balls map[int]int) map[int]int {
	out := make(map[int]int)
	for segment, count := range balls {
		for i := 0; i < count; i++ {
			m := mapping[segment]
			out[m[len(m)-1-i]]++
		}
	}
	return out
}

func ballsKey(balls map[int]int) string {
	keys := maps.Keys(balls)
	slices.Sort(keys)
	var s strings.Builder
	for _, key := range keys {
		s.WriteString(fmt.Sprintf(",%d:%d", key, balls[key]))
	}
	return s.String()[1:]
}
