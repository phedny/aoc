package input

import (
	"aoc2024/util"
)

var ReadDay1 = ReadFile(1, List("\n", Regex2[Tuple2[int, int]](`(\d+) +(\d+)`, Number, Number)))

var ReadDay2 = ReadFile(2, List("\n", List(" ", Number)))

var ReadDay3 = ReadFile(3, ByteArray)

var ReadDay4 = ReadFile(3, ByteMatrix)

type Day5 struct {
	Order       []Tuple2[int, int] `input:"1"`
	Productions [][]int            `input:"2"`
}

var ReadDay5 = ReadFile(5, Sequence2[Day5]("\n\n", List("\n", Regex2[Tuple2[int, int]](`(\d+)\|(\d+)`, Number, Number)), List("\n", List(",", Number))))

var ReadDay6 = ReadFile(6, ByteMatrix)

type Day7 struct {
	Target int   `input:"1"`
	Values []int `input:"2"`
}

var ReadDay7 = ReadFile(7, List("\n", Sequence2[Day7](": ", Number, List(" ", Number))))

var ReadDay8 = ReadFile(8, ByteMatrix)

var ReadDay9 = ReadFile(9, List("", Number))

var ReadDay10 = ReadFile(10, ByteMatrix)

var ReadDay11 = ReadFile(11, List(" ", String))

var ReadDay12 = ReadFile(12, ByteMatrix)

type Day13 struct {
	AX     int `input:"1"`
	AY     int `input:"2"`
	BX     int `input:"3"`
	BY     int `input:"4"`
	PrizeX int `input:"5"`
	PrizeY int `input:"6"`
}

var ReadDay13 = ReadFile(13, List("\n\n", Regex6[Day13]("Button A: X\\+(\\d+), Y\\+(\\d+)\nButton B: X\\+(\\d+), Y\\+(\\d+)\nPrize: X=(\\d+), Y=(\\d+)", Number, Number, Number, Number, Number, Number)))

type Day14 struct {
	PositionX int `input:"1"`
	PositionY int `input:"2"`
	VelocityX int `input:"3"`
	VelocityY int `input:"4"`
}

var ReadDay14 = ReadFile(14, List("\n", Regex4[Day14](`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`, Number, Number, Number, Number)))

type Day15 struct {
	Grid         util.ByteMatrix `input:"1"`
	Instructions []byte          `input:"2"`
}

var ReadDay15 = ReadFile(15, Sequence2[Day15]("\n\n", ByteMatrix, ByteArray))

var ReadDay16 = ReadFile(16, ByteMatrix)

type Day17 struct {
	A       int   `input:"1"`
	B       int   `input:"2"`
	C       int   `input:"3"`
	Program []int `input:"4"`
}

var ReadDay17 = ReadFile(17, Regex4[Day17]("Register A: (\\d+)\nRegister B: (\\d+)\nRegister C: (\\d+)\n\nProgram: ([0-9,]+)", Number, Number, Number, List(",", Number)))

var ReadDay18 = ReadFile(18, List("\n", Sequence2[Tuple2[int, int]](",", Number, Number)))

type Day19 struct {
	Towels  []string `input:"1"`
	Designs []string `input:"2"`
}

var ReadDay19 = ReadFile(19, Sequence2[Day19]("\n\n", List(", ", String), List("\n", String)))

var ReadDay20 = ReadFile(20, ByteMatrix)

var ReadDay21 = ReadFile(21, List("\n", String))

var ReadDay22 = ReadFile(22, List("\n", Number))

var ReadDay23 = ReadFile(23, List("\n", Sequence2[Tuple2[string, string]]("-", String, String)))

type Day24 struct {
	Wires map[string]bool `input:"1"`
	Gates []Day24Gate     `input:"2"`
}

type Day24Gate struct {
	A        string `input:"1"`
	B        string `input:"3"`
	Operator string `input:"2"`
	Out      string `input:"4"`
}

var ReadDay24 = ReadFile(24, Sequence2[Day24]("\n\n", Map("\n", ": ", String, Bool("0", "1")), List("\n", Regex4[Day24Gate](`([a-z0-9]+) (AND|OR|XOR) ([a-z0-9]+) -> ([a-z0-9]+)`, String, String, String, String))))

var ReadDay25 = ReadFile(25, List("\n\n", ByteMatrix))
