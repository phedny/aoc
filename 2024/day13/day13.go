package main

import (
	"aoc2024/util"
	"fmt"
	"math/big"
)

func main() {
	var x, y, tallyA, tallyB int64
	mA := [][]*big.Rat{make([]*big.Rat, 3), make([]*big.Rat, 3)}
	mB := [][]*big.Rat{make([]*big.Rat, 3), make([]*big.Rat, 3)}
	for i, line := range util.ReadLines() {
		switch i % 4 {
		case 0:
			fmt.Sscanf(line, "Button A: X+%d, Y+%d", &x, &y)
			mA[0][0], mA[1][0] = big.NewRat(x, 1), big.NewRat(y, 1)
			mB[0][0], mB[1][0] = big.NewRat(x, 1), big.NewRat(y, 1)
		case 1:
			fmt.Sscanf(line, "Button B: X+%d, Y+%d", &x, &y)
			mA[0][1], mA[1][1] = big.NewRat(x, 1), big.NewRat(y, 1)
			mB[0][1], mB[1][1] = big.NewRat(x, 1), big.NewRat(y, 1)
		case 2:
			fmt.Sscanf(line, "Prize: X=%d, Y=%d", &x, &y)
			mA[0][2], mA[1][2] = big.NewRat(x, 1), big.NewRat(y, 1)
			mB[0][2], mB[1][2] = big.NewRat(10000000000000+x, 1), big.NewRat(10000000000000+y, 1)
			Solve(mA)
			Solve(mB)
			if mA[0][2].IsInt() && mA[1][2].IsInt() {
				tallyA += 3*mA[0][2].Num().Int64() + mA[1][2].Num().Int64()
			}
			if mB[0][2].IsInt() && mB[1][2].IsInt() {
				tallyB += 3*mB[0][2].Num().Int64() + mB[1][2].Num().Int64()
			}
		}
	}
	fmt.Println(tallyA)
	fmt.Println(tallyB)
}

func Solve(m [][]*big.Rat) {
	var f, t big.Rat
	for r1i, r1 := range m {
		f = *r1[r1i]
		for j := range r1 {
			r1[j].Quo(r1[j], &f)
		}
		for r2i, r2 := range m {
			if r1i != r2i {
				f.Quo(r2[r1i], r1[r1i])
				for j := range r2 {
					r2[j].Sub(r2[j], t.Mul(&f, r1[j]))
				}
			}
		}
	}
}
