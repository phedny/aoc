package main

import (
	"aoc2024/input"
	"fmt"
)

func main() {
	robots := input.ReadDay14()

	// fmt.Println(partA(robots, 11, 7))
	fmt.Println(partA(robots, 101, 103))

	// partB(robots, 101, 103)

	// at frame 23 there is an interesting horizontal alignment
	// which re-occurs every frame N == 23 mod 101
	// same vertically at every frame N == 89 mod 103
	// so N == 7093 mod (101*103), let's print that one
	printFrame(robots, 101, 103, 7093)
}

func partA(robots []input.Day14, mX, mY int) int {
	var tallies [4]int
	for _, robot := range robots {
		pX := (robot.PositionX + 100*(robot.VelocityX+mX)) % mX
		pY := (robot.PositionY + 100*(robot.VelocityY+mY)) % mY
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

func printFrame(robots []input.Day14, mX, mY, frame int) {
	cs := make(map[[2]int]bool)
	for _, robot := range robots {
		x := (robot.PositionX + frame*(robot.VelocityX+mX)) % mX
		y := (robot.PositionY + frame*(robot.VelocityY+mY)) % mY
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
