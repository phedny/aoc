package main

import (
	"aoc2023/util"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var countA, countB int
	for _, line := range util.ReadLines() {
		countA += countLine(line, parseDigit)
		countB += countLine(line, parseDigitOrText)
	}
	fmt.Println(countA)
	fmt.Println(countB)
}

func countLine(line string, parse func(string) (int, bool)) int {
	var first, last int
	for len(line) > 0 {
		f, ok := parse(line)
		if ok {
			first = f
			break
		}
		line = line[1:]
	}
	for i := len(line) - 1; i >= 0; i-- {
		f, ok := parse(line[i:])
		if ok {
			last = f
			break
		}
	}

	return 10*first + last
}

func parseDigit(line string) (int, bool) {
	if unicode.IsNumber(rune(line[0])) {
		return int(line[0] - '0'), true
	}
	return 0, false
}

func parseDigitOrText(line string) (int, bool) {
	switch {
	case unicode.IsNumber(rune(line[0])):
		return int(line[0] - '0'), true
	case strings.HasPrefix(line, "one"):
		return 1, true
	case strings.HasPrefix(line, "two"):
		return 2, true
	case strings.HasPrefix(line, "three"):
		return 3, true
	case strings.HasPrefix(line, "four"):
		return 4, true
	case strings.HasPrefix(line, "five"):
		return 5, true
	case strings.HasPrefix(line, "six"):
		return 6, true
	case strings.HasPrefix(line, "seven"):
		return 7, true
	case strings.HasPrefix(line, "eight"):
		return 8, true
	case strings.HasPrefix(line, "nine"):
		return 9, true
	default:
		return 0, false
	}
}
