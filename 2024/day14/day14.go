package main

import (
	"aoc2024/util"
	"fmt"
)

func main() {
	var robots []Robot
	for _, line := range util.ReadLines() {
		var robot Robot
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robot.pX, &robot.pY, &robot.vX, &robot.vY)
		robots = append(robots, robot)
	}

	// fmt.Println(partA(robots, 11, 7))
	fmt.Println(partA(robots, 101, 103))

	// partB(robots, 101, 103)

	// at frame 23 there is an interesting horizontal alignment
	// which re-occurs every frame N == 23 mod 101
	// same vertically at every frame N == 89 mod 103
	// so N == 7093 mod (101*103), let's print that one
	printFrame(robots, 101, 103, 7093)
}

func partA(robots []Robot, mX, mY int) int {
	var tallies [4]int
	for _, robot := range robots {
		pX := (robot.pX + 100*(robot.vX+mX)) % mX
		pY := (robot.pY + 100*(robot.vY+mY)) % mY
		if pX != mX/2 && pY != mY/2 {
			var q int
			if pX > mX/2 {
				q = 1
			}
			if pY > mY/2 {
				q += 2
			}
			tallies[q]++
		}
	}
	return tallies[0] * tallies[1] * tallies[2] * tallies[3]
}

// func partB(robots []Robot, mX, mY int) {
// 	for i := 0; ; i++ {
// 		fmt.Println(i)
// 		printFrame(robots, mX, mY, i)
// 		fmt.Println()
// 		time.Sleep(200 * time.Millisecond)
// 	}
// }

func printFrame(robots []Robot, mX, mY, frame int) {
	cs := make(map[[2]int]bool)
	for _, robot := range robots {
		x := (robot.pX + frame*(robot.vX+mX)) % mX
		y := (robot.pY + frame*(robot.vY+mY)) % mY
		cs[[2]int{y, x}] = true
	}
	for y := range mY {
		for x := range mX {
			if cs[[2]int{y, x}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

type Robot struct {
	pX, pY, vX, vY int
}
