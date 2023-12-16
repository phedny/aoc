package main

import (
	"aoc2023/util"
	"fmt"
	"math"
)

type coor struct {
	y, x int
}

type posOr struct {
	y, x        int
	orientation byte
}

type from struct {
	ch        byte
	direction byte
}

type to struct {
	dy, dx    int
	direction byte
}

var up = to{-1, 0, 'U'}
var down = to{1, 0, 'D'}
var left = to{0, -1, 'L'}
var right = to{0, 1, 'R'}

var next = map[from][]to{
	{'.', 'U'}:  {up},
	{'.', 'D'}:  {down},
	{'.', 'L'}:  {left},
	{'.', 'R'}:  {right},
	{'|', 'U'}:  {up},
	{'|', 'D'}:  {down},
	{'|', 'L'}:  {up, down},
	{'|', 'R'}:  {up, down},
	{'-', 'L'}:  {left},
	{'-', 'R'}:  {right},
	{'-', 'U'}:  {left, right},
	{'-', 'D'}:  {left, right},
	{'/', 'U'}:  {right},
	{'/', 'D'}:  {left},
	{'/', 'L'}:  {down},
	{'/', 'R'}:  {up},
	{'\\', 'U'}: {left},
	{'\\', 'D'}: {right},
	{'\\', 'L'}: {up},
	{'\\', 'R'}: {down},
}

func main() {
	m := util.ReadByteMatrix()

	fmt.Println(countEnergized(m, posOr{0, 0, 'R'}))

	initOptions := make([]posOr, 2*len(m)+2*len(m[0]))
	for y := range m {
		initOptions = append(initOptions, posOr{y, 0, 'R'})
		initOptions = append(initOptions, posOr{y, len(m) - 1, 'L'})
	}
	for x := range m[0] {
		initOptions = append(initOptions, posOr{0, x, 'D'})
		initOptions = append(initOptions, posOr{len(m[0]) - 1, x, 'U'})
	}

	max := math.MinInt
	for _, init := range initOptions {
		n := countEnergized(m, init)
		if n > max {
			max = n
		}
	}
	fmt.Println(max)
}

func countEnergized(m [][]byte, init posOr) int {
	seen := make(map[posOr]bool)
	addSeen := map[posOr]bool{init: true}
	for len(addSeen) > 0 {
		for s := range addSeen {
			seen[s] = true
		}
		newSeen := make(map[posOr]bool)
		for s := range addSeen {
			for _, to := range next[from{m[s.y][s.x], s.orientation}] {
				po := posOr{s.y + to.dy, s.x + to.dx, to.direction}
				if po.x >= 0 && po.x < len(m) && po.y >= 0 && po.y < len(m[0]) && !seen[po] {
					newSeen[po] = true
				}
			}
		}
		addSeen = newSeen
	}

	energized := make(map[coor]bool)
	for s := range seen {
		energized[coor{s.y, s.x}] = true
	}
	return len(energized)
}
