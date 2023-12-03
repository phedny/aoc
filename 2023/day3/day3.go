package main

import (
	"aoc2023/util"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var sumA, sumB int
	lines := extendInput(util.ReadLines())
	partNrs, gears := extractPartNumbersAndGears(lines)

	for _, partNr := range partNrs {
		sumA += partNr
	}

	for _, touchingPartNrs := range gears {
		if len(touchingPartNrs) == 2 {
			sumB += touchingPartNrs[0] * touchingPartNrs[1]
		}
	}

	fmt.Println(sumA)
	fmt.Println(sumB)
}

func extendInput(lines []string) []string {
	width := len(lines[0]) + 2
	firstAndLast := strings.Repeat(".", width)
	out := make([]string, 1, len(lines)+2)
	out[0] = firstAndLast
	for _, line := range lines {
		out = append(out, "."+line+".")
	}
	out = append(out, firstAndLast)
	return out
}

func extractPartNumbersAndGears(lines []string) ([]int, map[[2]int][]int) {
	var partNrs []int
	gearsTouchingPartNrs := make(map[[2]int][]int)

	for lineNr, line := range lines {
		if lineNr == 0 || lineNr == len(lines)-1 {
			continue
		}

		partNr, adjacent, touchingGears := 0, false, make(map[[2]int]bool)
		for charNr, char := range line {
			switch charNr {
			case 0:
				continue
			case len(line) - 1:
				if partNr != 0 && adjacent {
					partNrs = append(partNrs, partNr)
					for gear := range touchingGears {
						gearsTouchingPartNrs[gear] = append(gearsTouchingPartNrs[gear], partNr)
					}
					partNr, adjacent, touchingGears = 0, false, make(map[[2]int]bool)
				}
			default:
				if unicode.IsNumber(char) {
					partNr = 10*partNr + int(char-'0')
					newAdjacent, newGears := isAdjacentAndGetGears(lines, lineNr, charNr)
					adjacent = adjacent || newAdjacent
					for _, gear := range newGears {
						touchingGears[gear] = true
					}
				} else {
					if partNr != 0 {
						if adjacent {
							partNrs = append(partNrs, partNr)
							for gear := range touchingGears {
								gearsTouchingPartNrs[gear] = append(gearsTouchingPartNrs[gear], partNr)
							}
						}
						partNr, adjacent, touchingGears = 0, false, make(map[[2]int]bool)
					}
				}
			}
		}
	}

	return partNrs, gearsTouchingPartNrs
}

func isAdjacentAndGetGears(lines []string, line, char int) (adjacent bool, gears [][2]int) {
	for _, lineNr := range []int{line - 1, line, line + 1} {
		for _, charNr := range []int{char - 1, char, char + 1} {
			if !strings.ContainsRune(".0123456789", rune(lines[lineNr][charNr])) {
				adjacent = true
				if lines[lineNr][charNr] == '*' {
					gears = append(gears, [2]int{lineNr, charNr})
				}
			}
		}
	}
	return
}
