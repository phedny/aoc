package main

import (
	"aoc2023/util"
	"fmt"
)

func main() {
	lines := util.ReadLines()
	emptyRows, emptyColumns := findEmptyRowsAndColumns(lines)
	fmt.Println(findDistances(findGalaxies(lines, emptyRows, emptyColumns, 2)))
	fmt.Println(findDistances(findGalaxies(lines, emptyRows, emptyColumns, 1000000)))
}

func findEmptyRowsAndColumns(lines []string) (map[int]bool, map[int]bool) {
	emptyRows := make(map[int]bool)
	emptyColumns := make(map[int]bool)
	for column := range lines[0] {
		emptyColumns[column] = true
	}
	for row, line := range lines {
		emptyRows[row] = true
		for column, b := range line {
			if b == '#' {
				delete(emptyRows, row)
				delete(emptyColumns, column)
			}
		}
	}
	return emptyRows, emptyColumns
}

func findGalaxies(lines []string, emptyRows, emptyColumns map[int]bool, rate int) map[[2]int]bool {
	galaxies := make(map[[2]int]bool)
	expandedY := 0
	for y, line := range lines {
		if emptyRows[y] {
			expandedY += rate
			continue
		}
		expandedX := 0
		for x, b := range line {
			if emptyColumns[x] {
				expandedX += rate
				continue
			}
			if b == '#' {
				galaxies[[2]int{expandedY, expandedX}] = true
			}
			expandedX++
		}
		expandedY++
	}
	return galaxies
}

func findDistances(galaxies map[[2]int]bool) int {
	sum := 0
	for g1 := range galaxies {
		for g2 := range galaxies {
			sum += util.Abs(g1[0]-g2[0]) + util.Abs(g1[1]-g2[1])
		}
	}
	return sum / 2
}
