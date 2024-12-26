package main

import (
	"aoc2024/input"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(partA())
	n, _ := partB(0, []int{2, 4, 1, 2, 7, 5, 1, 7, 4, 4, 0, 3, 5, 5, 3, 0})
	fmt.Println(n)
}

func partA() string {
	input := input.ReadDay17()
	return run(input.Program, [3]int{input.A, input.B, input.C})
}

func run(program []int, regs [3]int) string {
	var output strings.Builder
	for pc := 0; pc < len(program); {
		jump, absolute := opcodes[program[pc]](program[pc+1], regs[:], func(i int) { output.WriteByte('0' + byte(i)); output.WriteByte(',') })
		if absolute {
			pc = jump
		} else {
			pc += jump
		}
	}
	return output.String()[:output.Len()-1]
}

var opcodes = [8]func(int, []int, func(int)) (int, bool){adv, bxl, bst, jnz, bxc, out, bdv, cdv}

func combo(operand int, regs []int) int {
	if operand < 4 {
		return int(operand)
	} else {
		return regs[operand-4]
	}
}

func adv(operand int, regs []int, output func(int)) (int, bool) {
	regs[0] = regs[0] >> combo(operand, regs)
	return 2, false
}

func bxl(operand int, regs []int, output func(int)) (int, bool) {
	regs[1] ^= int(operand)
	return 2, false
}

func bst(operand int, regs []int, output func(int)) (int, bool) {
	regs[1] = combo(operand, regs) & 0x07
	return 2, false
}

func jnz(operand int, regs []int, output func(int)) (int, bool) {
	if regs[0] == 0 {
		return 2, false
	}
	return int(operand), true
}

func bxc(operand int, regs []int, output func(int)) (int, bool) {
	regs[1] ^= regs[2]
	return 2, false
}

func out(operand int, regs []int, output func(int)) (int, bool) {
	output(combo(operand, regs) & 0x07)
	return 2, false
}

func bdv(operand int, regs []int, output func(int)) (int, bool) {
	regs[1] = regs[0] >> combo(operand, regs)
	return 2, false
}

func cdv(operand int, regs []int, output func(int)) (int, bool) {
	regs[2] = regs[0] >> combo(operand, regs)
	return 2, false
}

func partB(a int, target []int) (int, bool) {
	if len(target) == 0 {
		return a, true
	}
	nextB := target[len(target)-1]
	for i := range 8 {
		tryA := a<<3 | i
		if ((tryA^5)^tryA>>((tryA&7)^2))&7 == nextB {
			a, ok := partB(tryA, target[:len(target)-1])
			if ok {
				return a, true
			}
		}
	}
	return 0, false
}
