package main

import (
	"aoc2024/input"
	"fmt"
	"math/big"
)

func main() {
	var tallyA, tallyB int64
	for _, line := range input.ReadDay13() {
		m := [][]*big.Rat{
			{big.NewRat(int64(line.AX), 1), big.NewRat(int64(line.BX), 1), big.NewRat(int64(line.PrizeX), 1)},
			{big.NewRat(int64(line.AY), 1), big.NewRat(int64(line.BY), 1), big.NewRat(int64(line.PrizeY), 1)},
		}
		Solve(m)
		if m[0][2].IsInt() && m[1][2].IsInt() {
			tallyA += 3*m[0][2].Num().Int64() + m[1][2].Num().Int64()
		}

		m = [][]*big.Rat{
			{big.NewRat(int64(line.AX), 1), big.NewRat(int64(line.BX), 1), big.NewRat(10000000000000+int64(line.PrizeX), 1)},
			{big.NewRat(int64(line.AY), 1), big.NewRat(int64(line.BY), 1), big.NewRat(10000000000000+int64(line.PrizeY), 1)},
		}
		Solve(m)
		if m[0][2].IsInt() && m[1][2].IsInt() {
			tallyB += 3*m[0][2].Num().Int64() + m[1][2].Num().Int64()
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
