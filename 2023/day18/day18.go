package main

import (
	"aoc2023/util"
	"fmt"
	"regexp"
	"strconv"

	"golang.org/x/exp/slices"
)

var rLine = regexp.MustCompile(`(\w) (\d+) \(#([0-9a-f]{5})([0-3])\)`)

var aNext = map[string]func(int, int, [][]trench, int) (int, int, [][]trench){
	"U": goUp,
	"D": goDown,
	"L": goLeft,
	"R": goRight,
}

var bNext = map[string]func(int, int, [][]trench, int) (int, int, [][]trench){
	"0": goRight,
	"1": goDown,
	"2": goLeft,
	"3": goUp,
}

func main() {
	var aX, aY, bX, bY int
	var aVerticalTrenches, bVerticalTrenches [][]trench

	for _, line := range util.ReadLines() {
		m := rLine.FindStringSubmatch(line)
		aN, _ := strconv.ParseInt(m[2], 10, 64)
		bN, _ := strconv.ParseInt(m[3], 16, 64)
		aX, aY, aVerticalTrenches = aNext[m[1]](aX, aY, aVerticalTrenches, int(aN))
		bX, bY, bVerticalTrenches = bNext[m[4]](bX, bY, bVerticalTrenches, int(bN))
	}

	fmt.Println(findDugOutArea(aVerticalTrenches))
	fmt.Println(findDugOutArea(bVerticalTrenches))
}

type trench struct {
	ortho      int
	start, end int
}

func goDown(x, y int, verticalTrenches [][]trench, distance int) (int, int, [][]trench) {
	xPos, xFound := slices.BinarySearchFunc(verticalTrenches, x, func(trench []trench, x int) int {
		return trench[0].ortho - x
	})
	if !xFound {
		verticalTrenches = slices.Insert(verticalTrenches, xPos, nil)
	}
	yPos, yFound := slices.BinarySearchFunc(verticalTrenches[xPos], y, func(trench trench, y int) int {
		return trench.start - y
	})
	if yFound {
		panic("already a trench")
	}
	verticalTrenches[xPos] = slices.Insert(verticalTrenches[xPos], yPos, trench{x, y, y + distance})
	return x, y + distance, verticalTrenches
}

func goUp(x, y int, verticalTrenches [][]trench, distance int) (int, int, [][]trench) {
	xPos, xFound := slices.BinarySearchFunc(verticalTrenches, x, func(trench []trench, x int) int {
		return trench[0].ortho - x
	})
	if !xFound {
		verticalTrenches = slices.Insert(verticalTrenches, xPos, nil)
	}
	yPos, yFound := slices.BinarySearchFunc(verticalTrenches[xPos], y-distance, func(trench trench, y int) int {
		return trench.start - y
	})
	if yFound {
		panic("already a trench")
	}
	verticalTrenches[xPos] = slices.Insert(verticalTrenches[xPos], yPos, trench{x, y - distance, y})
	return x, y - distance, verticalTrenches
}

func goLeft(x, y int, verticalTrenches [][]trench, distance int) (int, int, [][]trench) {
	return x - distance, y, verticalTrenches
}

func goRight(x, y int, verticalTrenches [][]trench, distance int) (int, int, [][]trench) {
	return x + distance, y, verticalTrenches
}

func findDugOutArea(verticalTrenches [][]trench) int {
	rs := verticalTrenches[0]
	rStart := rs[0].ortho
	area := 0
	for _, bs := range verticalTrenches[1:] {
		width := bs[0].ortho - rStart
		rStart = bs[0].ortho
		for _, trench := range rs {
			area += width * (trench.end - trench.start + 1)
		}

		for _, b := range bs {
			pos, found := slices.BinarySearchFunc(rs, b.start, func(r trench, b int) int {
				return r.start - b
			})
			if found {
				if rs[pos].end == b.end {
					area += b.end - b.start + 1
					rs = slices.Delete(rs, pos, pos+1)
				} else if rs[pos].end > b.end {
					area += b.end - b.start
					rs[pos].start = b.end
				} else {
					panic("crossing")
				}
			} else if pos == 0 {
				if b.end < rs[0].start {
					rs = slices.Insert(rs, 0, b)
				} else if b.end == rs[0].start {
					rs[0].start = b.start
				} else {
					panic("crossing")
				}
			} else {
				if rs[pos-1].end < b.start {
					if pos < len(rs) && rs[pos].start == b.end {
						rs[pos].start = b.start
					} else {
						rs = slices.Insert(rs, pos, b)
					}
				} else if rs[pos-1].end == b.start {
					if pos < len(rs) && rs[pos].start == b.end {
						rs[pos-1].end = rs[pos].end
						rs = slices.Delete(rs, pos, pos+1)
					} else {
						rs[pos-1].end = b.end
					}
				} else if rs[pos-1].end == b.end {
					area += rs[pos-1].end - b.start
					rs[pos-1].end = b.start
				} else {
					area += b.end - b.start - 1
					rs = slices.Insert(rs, pos, trench{0, b.end, rs[pos-1].end})
					rs[pos-1].end = b.start
				}
			}
		}
	}
	if len(rs) > 0 {
		panic("Houston, we have a problem")
	}
	return area
}
