package main

import (
	"aoc2023/util"
	"fmt"
	"math"
	"regexp"
	"strconv"

	"golang.org/x/exp/slices"
)

var rNumber = regexp.MustCompile(`\d+`)

func main() {
	blocks := parseBlocks(util.ReadLines())

	supporting := make(map[int]map[int]bool)
	supportedBy := make(map[int]map[int]bool)
	columns := make(map[[2]int]columnData)
	for id, b := range blocks {
		distanceToFall := math.MaxInt
		for _, c := range b.area {
			distance := b.z - columns[c].z
			if distance < distanceToFall {
				distanceToFall = distance
			}
		}
		for _, c := range b.area {
			cd, hasCd := columns[c]
			if hasCd && cd.z == b.z-distanceToFall {
				if supporting[cd.blockId] == nil {
					supporting[cd.blockId] = make(map[int]bool)
				}
				supporting[cd.blockId][id] = true
				if supportedBy[id] == nil {
					supportedBy[id] = make(map[int]bool)
				}
				supportedBy[id][cd.blockId] = true
			}
			columns[c] = columnData{id, b.z + b.height - distanceToFall}
		}
	}

	var countA, sumB int
	for _, n := range detectFallingBlocks(blocks, supporting, supportedBy) {
		if n == 0 {
			countA++
		}
		sumB += n
	}
	fmt.Println(countA)
	fmt.Println(sumB)
}

func parseBlocks(lines []string) []block {
	blocks := make([]block, len(lines))
	for i, line := range lines {
		numbers := rNumber.FindAllString(line, 6)
		var ns [6]int
		for j, str := range numbers {
			ns[j], _ = strconv.Atoi(str)
		}
		area := make([][2]int, 0, (ns[3]-ns[0]+1)*(ns[4]-ns[1]+1))
		for x := ns[0]; x <= ns[3]; x++ {
			for y := ns[1]; y <= ns[4]; y++ {
				area = append(area, [2]int{x, y})
			}
		}
		blocks[i] = block{area, ns[2], ns[5] - ns[2] + 1}
	}
	slices.SortFunc(blocks, func(b1, b2 block) int {
		return b1.z - b2.z
	})
	return blocks
}

func detectFallingBlocks(blocks []block, supporting, supportedBy map[int]map[int]bool) []int {
	fallCount := make([]int, len(blocks))
	for i := range blocks {
		fell := make(map[int]bool)
		falling := make(map[int]bool)
		for j := range supporting[i] {
			if len(supportedBy[j]) == 1 {
				falling[j] = true
			}
		}
		for len(falling) > 0 {
			checkFalling := make(map[int]bool)
			for j := range falling {
				fell[j] = true
				for k := range supporting[j] {
					checkFalling[k] = true
				}
			}
			for j := range checkFalling {
				for k := range supportedBy[j] {
					if !fell[k] {
						delete(checkFalling, j)
						break
					}
				}
			}
			falling = checkFalling
		}
		fallCount[i] = len(fell)
	}
	return fallCount
}

type block struct {
	area      [][2]int
	z, height int
}

type columnData struct {
	blockId int
	z       int
}
