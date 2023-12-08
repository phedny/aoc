package main

import (
	"aoc2023/util"
	"fmt"
	"strings"
)

func main() {
	lines := util.ReadLines()

	navigations := make(map[string][2]string)
	for _, line := range lines[2:] {
		navigations[line[0:3]] = [2]string{line[7:10], line[12:15]}
	}

	fmt.Println(stepsUntilEnd(navigations, lines[0], "AAA", map[string]bool{"ZZZ": true}))

	var startB []string
	endB := make(map[string]bool)
	for position := range navigations {
		if strings.HasSuffix(position, "A") {
			startB = append(startB, position)
		} else if strings.HasSuffix(position, "Z") {
			endB[position] = true
		}
	}

	fmt.Println(parallelStepsUntilEnd(navigations, lines[0], startB, endB))
}

func stepsUntilEnd(navigations map[string][2]string, directions string, start string, end map[string]bool) int {
	steps := 0
	position := start

	for !end[position] {
		d := directions[(steps)%len(directions)]
		steps++
		if d == 'L' {
			position = navigations[position][0]
		} else {
			position = navigations[position][1]
		}
	}

	return steps
}

func parallelStepsUntilEnd(navigations map[string][2]string, directions string, starts []string, end map[string]bool) int {
	steps := make([]int, len(starts))
	for i, start := range starts {
		steps[i] = stepsUntilEnd(navigations, directions, start, end)
	}

	combinedSteps := 1
	for _, step := range steps {
		combinedSteps = combinedSteps * step / gcd(combinedSteps, step)
	}

	return combinedSteps
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
