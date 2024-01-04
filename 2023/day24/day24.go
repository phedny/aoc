package main

import (
	"aoc2023/util"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

var rNumber = regexp.MustCompile(`-?\d+`)

func main() {
	stones := parseStones(util.ReadLines())

	lines := make([]line2, len(stones))
	for i, stone := range stones {
		lines[i] = stone.horizontalEquation()
	}

	min, max := 200000000000000., 400000000000000.
	if len(stones) == 5 {
		min, max = 7., 27.
	}

	var countA int
	for a, l1 := range lines {
		for b, l2 := range lines[a+1:] {
			intersect, at := intersectionOf(l1, l2)
			if !intersect || at[0] < min || at[0] > max || at[1] < min || at[1] > max {
				continue
			}
			at0 := int(at[0])
			sA, sB := stones[a], stones[a+b+1]
			if (sA.v[0] > 0 && at0 < sA.o[0]) || (sA.v[0] < 0 && at0 > sA.o[0]) || (sB.v[0] > 0 && at0 < sB.o[0]) || (sB.v[0] < 0 && at0 > sB.o[0]) {
				continue
			}
			countA++
		}
	}
	fmt.Println(countA)

	transform := stones[0]
	s1, s2 := stones[1], stones[2]

	l1 := line3{o: s1.o.minus(transform.o), v: s1.v.minus(transform.v)}
	l2 := line3{o: s2.o.minus(transform.o), v: s2.v.minus(transform.v)}
	l3 := line3{v: l1.o.cross(l1.v).shrink().cross(l2.o.cross(l2.v).shrink())}

	t1, t2 := l1.intersectionT(l3), l2.intersectionT(l3)
	i1, i2 := s1.atT(t1), s2.atT(t2)

	throwFrom := i1.minus(i2.minus(i1).div(t2 - t1).times(t1))
	fmt.Println(throwFrom[0] + throwFrom[1] + throwFrom[2])
}

func parseStones(lines []string) []line3 {
	stones := make([]line3, len(lines))
	for i, line := range util.ReadLines() {
		m := rNumber.FindAllString(line, 6)
		var h line3
		h.o[0], _ = strconv.Atoi(m[0])
		h.o[1], _ = strconv.Atoi(m[1])
		h.o[2], _ = strconv.Atoi(m[2])
		h.v[0], _ = strconv.Atoi(m[3])
		h.v[1], _ = strconv.Atoi(m[4])
		h.v[2], _ = strconv.Atoi(m[5])
		stones[i] = h
	}
	return stones
}

type vector [3]int

func (v vector) shrink() vector {
	g := gcd(gcd(v[0], v[1]), v[2])
	return vector{v[0] / g, v[1] / g, v[2] / g}
}

func gcd(a, b int) int {
	if a < 0 {
		a *= -1
	}
	if b < 0 {
		b *= -1
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func (v vector) times(f int) vector {
	return vector{v[0] * f, v[1] * f, v[2] * f}
}

func (v vector) div(f int) vector {
	return vector{v[0] / f, v[1] / f, v[2] / f}
}

func (v vector) minus(o vector) vector {
	return vector{v[0] - o[0], v[1] - o[1], v[2] - o[2]}
}

func (v vector) cross(o vector) vector {
	return vector{
		v[1]*o[2] - v[2]*o[1],
		v[2]*o[0] - v[0]*o[2],
		v[0]*o[1] - v[1]*o[0],
	}
}

type line3 struct {
	o, v vector
}

func (l line3) horizontalEquation() line2 {
	return line2{float64(l.o[0]) - float64(l.v[0]*l.o[1])/float64(l.v[1]), float64(l.v[0]) / float64(l.v[1])}
}

func (l line3) atT(t int) vector {
	return vector{l.o[0] + t*l.v[0], l.o[1] + t*l.v[1], l.o[2] + t*l.v[2]}
}

func (l line3) intersectionT(o line3) int {
	for _, is := range [][2]int{{0, 1}, {0, 2}, {1, 2}} {
		n := (float64(l.o[is[1]]-o.o[is[1]]) - float64(l.o[is[0]]-o.o[is[0]])*float64(o.v[is[1]])/float64(o.v[is[0]]))
		d := (float64(l.v[is[1]]) + float64(-l.v[is[0]])*float64(o.v[is[1]])/float64(o.v[is[0]]))
		if d != 0 {
			return int(math.Round(n / -d))
		}
	}
	panic("no intersection")
}

type line2 struct {
	x0, dx float64
}

func intersectionOf(l1, l2 line2) (bool, [2]float64) {
	if l1.dx == l2.dx {
		return false, [2]float64{}
	}
	y := (l1.x0 - l2.x0) / (l2.dx - l1.dx)
	return true, [2]float64{l1.x0 + y*l1.dx, y}
}
