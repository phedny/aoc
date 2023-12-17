package main

import (
	"aoc2023/util"
	"fmt"

	"golang.org/x/exp/slices"
)

const (
	up = iota
	left
	down
	right
)

type move struct {
	dy, dx int
}

var moves = map[int]move{up: {-1, 0}, down: {1, 0}, left: {0, -1}, right: {0, 1}}

type todo struct {
	y, x        int
	direction   int
	onlyForward int
	maxForward  int
	doSides     bool
}

type costTodo struct {
	cost  int
	todos map[todo]bool
}

func main() {
	m := util.ReadByteMatrix()
	for _, r := range m {
		for i := range r {
			r[i] -= '0'
		}
	}

	fmt.Println(lowestCost(m, 0, 3))
	fmt.Println(lowestCost(m, 3, 10))
}

func lowestCost(m [][]byte, onlyForward, maxForward int) int {
	done := make([][][4]int, len(m))
	for y, r := range m {
		done[y] = make([][4]int, len(r))
		for x := range r {
			done[y][x] = [4]int{-1, -1, -1, -1}
		}
	}

	sortedTodos := []costTodo{{0, map[todo]bool{{0, 0, down, onlyForward, maxForward, true}: true, {0, 0, right, onlyForward, maxForward, true}: true}}}
	done[0][0] = [4]int{maxForward, maxForward, maxForward, maxForward}
	for {
		currentCost := sortedTodos[0].cost
		todos := sortedTodos[0].todos
		sortedTodos = sortedTodos[1:]

		for current := range todos {
			move := moves[current.direction]
			x, y := current.x+move.dx, current.y+move.dy
			if y < 0 || y >= len(m) || x < 0 || x >= len(m[0]) {
				continue
			}

			cost := currentCost + int(m[y][x])
			if y == len(m)-1 && x == len(m[0])-1 {
				return cost
			}

			if current.onlyForward > 0 {
				sortedTodos = addTodo(sortedTodos, cost, todo{y, x, current.direction, current.onlyForward - 1, current.maxForward - 1, done[y][x][current.direction] == -1})
				continue
			}

			if current.maxForward > 1 && current.maxForward > done[y][x][current.direction] {
				sortedTodos = addTodo(sortedTodos, cost, todo{y, x, current.direction, 0, current.maxForward - 1, done[y][x][current.direction] == -1})
				done[y][x][current.direction] = current.maxForward - 1
			}

			if current.doSides {
				for dDir := 1; dDir < 4; dDir += 2 {
					newDirection := (current.direction + dDir) % 4
					if done[y][x][newDirection] < maxForward {
						sortedTodos = addTodo(sortedTodos, cost, todo{y, x, newDirection, onlyForward, maxForward, done[y][x][newDirection] == -1})
					}
				}
			}
		}
	}
}

func addTodo(sortedTodos []costTodo, cost int, newTodo todo) []costTodo {
	pos, has := slices.BinarySearchFunc(sortedTodos, cost, func(ct costTodo, cost int) int {
		return ct.cost - cost
	})
	if !has {
		sortedTodos = slices.Insert(sortedTodos, pos, costTodo{cost, make(map[todo]bool)})
	}
	sortedTodos[pos].todos[newTodo] = true
	return sortedTodos
}
