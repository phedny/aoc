package main

import (
	"aoc2023/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

var rInstr = regexp.MustCompile(`(\w+)(-|=(\d+))`)

type labeledLens struct {
	label       string
	focalLength int
}

func main() {
	var sumA, sumB int
	var boxes [256][]labeledLens

	for _, s := range strings.Split(util.ReadLines()[0], ",") {
		sumA += int(hash(s))

		m := rInstr.FindStringSubmatch(s)
		boxNr := hash(m[1])
		if m[3] == "" {
			boxes[boxNr] = slices.DeleteFunc(boxes[boxNr], withLabel(m[1]))
		} else {
			n, _ := strconv.Atoi(m[3])
			pos := slices.IndexFunc(boxes[boxNr], withLabel(m[1]))
			if pos == -1 {
				boxes[boxNr] = append(boxes[boxNr], labeledLens{m[1], n})
			} else {
				boxes[boxNr][pos].focalLength = n
			}
		}
	}

	for boxNr, box := range boxes {
		for slotNr, lens := range box {
			sumB += (boxNr + 1) * (slotNr + 1) * lens.focalLength
		}
	}

	fmt.Println(sumA)
	fmt.Println(sumB)
}

func hash(s string) byte {
	var current byte
	for _, b := range s {
		current = 17 * (current + byte(b))
	}
	return current
}

func withLabel(label string) func(labeledLens) bool {
	return func(item labeledLens) bool {
		return item.label == label
	}
}
