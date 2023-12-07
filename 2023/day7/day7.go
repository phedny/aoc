package main

import (
	"aoc2023/util"
	"fmt"
	"strconv"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

var handTypes = map[string]int{"11111": 0, "1112": 1, "122": 3, "113": 4, "23": 5, "14": 6, "5": 7}
var cardValuesA = map[rune]int{'2': 0, '3': 1, '4': 2, '5': 3, '6': 4, '7': 5, '8': 6, '9': 7, 'T': 8, 'J': 9, 'Q': 10, 'K': 11, 'A': 12}
var cardValuesB = map[rune]int{'J': 0, '2': 1, '3': 2, '4': 3, '5': 4, '6': 5, '7': 6, '8': 7, '9': 8, 'T': 9, 'Q': 10, 'K': 11, 'A': 12}

func main() {
	handsA := make(map[int]int)
	handsB := make(map[int]int)
	for _, line := range util.ReadLines() {
		bid, _ := strconv.Atoi(line[6:])
		handsA[handValue(line[:5], cardValuesA, false)] = bid
		handsB[handValue(line[:5], cardValuesB, true)] = bid
	}

	fmt.Println(totalWinnings(handsA))
	fmt.Println(totalWinnings(handsB))
}

func handValue(cards string, cardValues map[rune]int, withJokers bool) int {
	labels := make(map[rune]int)
	for _, r := range cards {
		labels[r]++
	}
	jokers := 0
	if withJokers {
		jokers = labels['J']
		delete(labels, 'J')
	}
	counts := make([]byte, 0, 5)
	for _, c := range labels {
		counts = append(counts, byte(c+'0'))
	}
	slices.Sort(counts)
	if len(counts) == 0 {
		counts = []byte{'5'}
	} else {
		counts[len(counts)-1] += byte(jokers)
	}
	value, ok := handTypes[string(counts)]
	if !ok {
		panic("unexpected count: " + string(counts))
	}

	for _, r := range cards {
		value = 13*value + cardValues[r]
	}

	return value
}

func totalWinnings(hands map[int]int) int {
	handValues := maps.Keys(hands)
	slices.Sort(handValues)

	var sum int
	for rank, handValue := range handValues {
		sum += (rank + 1) * hands[handValue]
	}

	return sum
}
