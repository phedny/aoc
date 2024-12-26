package main

import (
	"aoc2024/input"
	"aoc2024/util"
	"fmt"
	"maps"
	"math"
)

const size = 70
const corruption = 1024

func main() {
	validCells := make(map[util.Coordinate]int)
	for r := range size + 1 {
		for c := range size + 1 {
			validCells[util.NewCoordinate(r, c)] = math.MaxInt
		}
	}
	for _, line := range input.ReadDay18()[:corruption] {
		delete(validCells, util.NewCoordinate(line.V1, line.V2))
	}
	cost, _ := util.AStar(Path{util.NewCoordinate(0, 0), 0}, pathCost, step(maps.Clone(validCells)))
	fmt.Println(cost)

	for _, line := range input.ReadDay18()[corruption:] {
		delete(validCells, util.NewCoordinate(line.V1, line.V2))
		_, found := util.AStar(Path{util.NewCoordinate(0, 0), 0}, pathCost, step(maps.Clone(validCells)))
		if !found {
			fmt.Printf("%d,%d\n", line.V1, line.V2)
			return
		}
	}
}

type Path struct {
	position util.Coordinate
	cost     int
}

func pathCost(p Path) int {
	return 2*size + p.cost - p.position.Row() - p.position.Column()
}

func step(validCells map[util.Coordinate]int) func(p Path) (int, []Path, bool) {
	return func(p Path) (int, []Path, bool) {
		if p.position.Row() == size && p.position.Column() == size {
			return p.cost, nil, true
		}
		nextPaths := make([]Path, 0, 4)
		for t := range util.OrthogonalNeighbors {
			p := Path{p.position.Add(t), p.cost + 1}
			if seenCost := validCells[p.position]; seenCost > p.cost {
				validCells[p.position] = p.cost
				nextPaths = append(nextPaths, p)
			}
		}
		return p.cost, nextPaths, false
	}
}
