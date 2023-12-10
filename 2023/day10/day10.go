package main

import (
	"aoc2023/util"
	"fmt"
)

type pos struct {
	y, x int
}

type posOr struct {
	y, x        int
	orientation byte
}

type from struct {
	pipe      byte
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

var next = map[from]to{
	{'|', 'U'}: up,
	{'|', 'D'}: down,
	{'-', 'L'}: left,
	{'-', 'R'}: right,
	{'L', 'D'}: right,
	{'L', 'L'}: up,
	{'J', 'D'}: left,
	{'J', 'R'}: up,
	{'7', 'U'}: left,
	{'7', 'R'}: down,
	{'F', 'U'}: right,
	{'F', 'L'}: down,
}

func main() {
	inputMap := convertMap(util.ReadLines())
	y, x := findStart(inputMap)

	loop := findLoop(y, x, inputMap)
	cleanMap(inputMap, loop)
	fmt.Println(len(loop) / 2)

	expandedMap := expandMap(inputMap)
	fillOutside(expandedMap)
	contractedMap := contractMap(expandedMap)
	fmt.Println(countEmpty(contractedMap))
}

func convertMap(lines []string) [][]byte {
	out := make([][]byte, len(lines))
	for i, line := range lines {
		out[i] = []byte(line)
	}
	return out
}

func findStart(lines [][]byte) (int, int) {
	for y, line := range lines {
		for x, r := range line {
			if r == 'S' {
				return y, x
			}
		}
	}
	panic("no start")
}

func findLoop(y, x int, lines [][]byte) map[pos]bool {
	loop := map[pos]bool{{y, x}: true}
	for po := firstStep(y, x, lines); lines[po.y][po.x] != 'S'; {
		loop[pos{po.y, po.x}] = true
		to := next[from{lines[po.y][po.x], po.orientation}]
		po = posOr{po.y + to.dy, po.x + to.dx, to.direction}
	}
	return loop
}

func firstStep(y, x int, lines [][]byte) posOr {
	for _, po := range []posOr{{y - 1, x, 'U'}, {y + 1, x, 'D'}, {y, x - 1, 'L'}, {y, x + 1, 'R'}} {
		if po.y < 0 || po.y >= len(lines) || po.x < 0 || po.x >= len(lines[0]) {
			continue
		}
		to := next[from{lines[po.y][po.x], po.orientation}]
		if to.direction != 0 {
			return po
		}
	}
	panic("no first step")
}

func cleanMap(lines [][]byte, loop map[pos]bool) {
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			if !loop[pos{y, x}] {
				lines[y][x] = ' '
			}
		}
	}
}

func expandMap(lines [][]byte) [][]byte {
	newLines := make([][]byte, 2*len(lines)+1)
	for y := range newLines {
		newLines[y] = make([]byte, 2*len(lines[0])+1)
		for x := range newLines[y] {
			newLines[y][x] = ' '
		}
	}

	for y, line := range lines {
		for x, b := range line {
			newLines[2*y+1][2*x+1] = b
			switch b {
			case '|':
				newLines[2*y][2*x+1] = '|'
				newLines[2*y+2][2*x+1] = '|'
			case '-':
				newLines[2*y+1][2*x] = '-'
				newLines[2*y+1][2*x+2] = '-'
			case 'L':
				newLines[2*y][2*x+1] = '|'
				newLines[2*y+1][2*x+2] = '-'
			case 'J':
				newLines[2*y][2*x+1] = '|'
				newLines[2*y+1][2*x] = '-'
			case '7':
				newLines[2*y+2][2*x+1] = '|'
				newLines[2*y+1][2*x] = '-'
			case 'F':
				newLines[2*y+2][2*x+1] = '|'
				newLines[2*y+1][2*x+2] = '-'
			}
		}
	}

	return newLines
}

func contractMap(lines [][]byte) [][]byte {
	newLines := make([][]byte, len(lines)/2)
	for y, line := range lines {
		if y%2 == 0 {
			continue
		}
		newLines[y/2] = make([]byte, len(line)/2)
		for x, b := range line {
			if x%2 == 0 {
				continue
			}
			newLines[y/2][x/2] = b
		}
	}
	return newLines
}

func fillOutside(lines [][]byte) {
	ps := map[pos]bool{{0, 0}: true}
	for len(ps) > 0 {
		for p := range ps {
			lines[p.y][p.x] = 'o'
		}
		newPs := make(map[pos]bool)
		for p := range ps {
			if p.y > 0 && lines[p.y-1][p.x] == ' ' {
				newPs[pos{p.y - 1, p.x}] = true
			}
			if p.y < len(lines)-1 && lines[p.y+1][p.x] == ' ' {
				newPs[pos{p.y + 1, p.x}] = true
			}
			if p.x > 0 && lines[p.y][p.x-1] == ' ' {
				newPs[pos{p.y, p.x - 1}] = true
			}
			if p.x < len(lines[0])-1 && lines[p.y][p.x+1] == ' ' {
				newPs[pos{p.y, p.x + 1}] = true
			}
		}
		ps = newPs
	}
}

func countEmpty(lines [][]byte) int {
	count := 0
	for _, line := range lines {
		for _, b := range line {
			if b == ' ' {
				count++
			}
		}
	}
	return count
}
