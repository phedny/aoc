package main

import (
	"aoc2023/util"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

var handTypes = map[string]int{"11111": 0, "1112": 1, "122": 3, "113": 4, "23": 5, "14": 6, "5": 7}

func main() {
	handsA := make(map[int]int)
	handsB := make(map[int]int)
	for _, line := range util.ReadLines() {
		bid, _ := strconv.Atoi(line[6:])
		handsA[handValue(line[:5], "23456789TJQKA", false)] = bid
		handsB[handValue(line[:5], "J23456789TQKA", true)] = bid
	}

	fmt.Println(totalWinnings(handsA))
	fmt.Println(totalWinnings(handsB))
}

func handValue(cards string, cardValues string, withJokers bool) int {
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
		value = 13*value + strings.IndexRune(cardValues, r)
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
