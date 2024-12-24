package main

import (
	"aoc2024/util"
	"fmt"
	"maps"
	"slices"
	"strings"
)

func main() {
	lines := util.ReadLines()

	i := slices.Index(lines, "")
	faults := make(map[string]bool)
	wires, mux := parseNetwork(lines[i+1:])

	for _, line := range lines[:i] {
		s := strings.SplitN(line, " ", 2)
		wires.Get(s[0][:len(s[0])-1]).SetBit(s[1] == "1")
	}
	fmt.Println(mux.Get())

	z00 := detectHalfAdder(wires["x00"], wires["y00"])
	if z00.name != "z00" {
		panic("first bit wrong")
	}

	for i := 1; ; i++ {
		x := wires[fmt.Sprintf("x%02d", i)]
		y := wires[fmt.Sprintf("y%02d", i)]
		z := wires[fmt.Sprintf("z%02d", i)]
		if x == nil || y == nil {
			break
		}
		out, carryOut, fault1, fault2 := detectFullAdder(x, y)
		if fault1 == nil && fault2 == nil {
			if out != z && carryOut == z {
				out, carryOut = carryOut, out
				faults[out.name] = true
				faults[carryOut.name] = true
			}
		} else {
			faults[fault1.name] = true
			faults[fault2.name] = true
		}
	}

	s := slices.Collect(maps.Keys(faults))
	slices.Sort(s)
	fmt.Println(strings.Join(s, ","))
}

func parseNetwork(lines []string) (Wires, *Mux) {
	wires := make(Wires)
	for _, line := range lines {
		s := strings.SplitN(line, " ", 5)
		var gate BitGate
		switch s[1] {
		case "AND":
			gate = &And{InputBitWires: InputBitWires{in: make(map[int]*Wire), bits: make(map[int]bool)}, OutputBitWires: OutputBitWires{out: make(map[string]*Wire)}}
		case "OR":
			gate = &Or{InputBitWires: InputBitWires{in: make(map[int]*Wire), bits: make(map[int]bool)}, OutputBitWires: OutputBitWires{out: make(map[string]*Wire)}}
		case "XOR":
			gate = &Xor{InputBitWires: InputBitWires{in: make(map[int]*Wire), bits: make(map[int]bool)}, OutputBitWires: OutputBitWires{out: make(map[string]*Wire)}}
		}
		wires.Get(s[0]).ConnectToGateInput(gate, 0)
		wires.Get(s[2]).ConnectToGateInput(gate, 1)
		wires.Get(s[4]).ConnectToGateOutput(gate, "out")
	}

	mux := &Mux{InputBitWires{in: make(map[int]*Wire), bits: make(map[int]bool)}}
	for i := 0; ; i++ {
		wire, has := wires[fmt.Sprintf("z%02d", i)]
		if !has {
			return wires, mux
		}
		wire.ConnectToGateInput(mux, i)
	}
}

func detectHalfAdder(a, b *Wire) *Wire {
	xorA, andA := getXorAnd(a)
	xorB, andB := getXorAnd(b)
	if xorA == nil || andA == nil || xorB == nil || andB == nil || xorA != xorB || andA != andB {
		panic("XOR/AND error")
	}
	return xorA.out["out"]
}

func detectFullAdder(a, b *Wire) (out, carryOut, fault1, fault2 *Wire) {
	xorA, andA := getXorAnd(a)
	xorB, andB := getXorAnd(b)
	if xorA == nil || andA == nil || xorB == nil || andB == nil || xorA != xorB || andA != andB {
		return nil, nil, a, b
	}
	andAOut := andA.out["out"]
	xor2, and2 := getXorAnd(xorA.out["out"])
	if xor2 == nil || and2 == nil {
		xor2, and2 = getXorAnd(andA.out["out"])
		fault1 = xorA.out["out"]
		fault2 = andAOut
		andAOut = xorA.out["out"]
	}
	out = xor2.out["out"]
	or1, isOr1 := andAOut.out[0].gate.(*Or)
	or2, isOr2 := and2.out["out"].out[0].gate.(*Or)
	if !isOr1 && isOr2 {
		fault1 = andAOut
		fault2 = or2.in[1-and2.out["out"].out[0].index]
		carryOut = or2.out["out"]
	} else if isOr1 && !isOr2 {
		fault1 = and2.out["out"]
		fault2 = or1.in[1-andAOut.out[0].index]
		carryOut = or1.out["out"]
	} else {
		carryOut = or1.out["out"]
	}
	return
}

func getXorAnd(wire *Wire) (xor *Xor, and *And) {
	if len(wire.out) != 2 {
		return nil, nil
	}
	for _, out := range wire.out {
		switch gate := out.gate.(type) {
		case *And:
			and = gate
		case *Xor:
			xor = gate
		}
	}
	return
}

type Wires map[string]*Wire

func (wires Wires) Get(z string) *Wire {
	if wire, has := wires[z]; has {
		return wire
	}
	wire := &Wire{name: z}
	wires[z] = wire
	return wire
}

func (wires Wires) ConvertToHalfAdder(xor *Xor, and *And) {
	a := xor.in[0]
	b := xor.in[1]
	out := xor.out["out"]
	carry := and.out["out"]
	a.out = slices.DeleteFunc(a.out, func(input GateInput) bool { return input.gate == xor || input.gate == and })
	b.out = slices.DeleteFunc(b.out, func(input GateInput) bool { return input.gate == xor || input.gate == and })
	halfAdder := &HalfAdder{InputBitWires: InputBitWires{in: make(map[int]*Wire), bits: make(map[int]bool)}, OutputBitWires: OutputBitWires{out: make(map[string]*Wire)}}
	a.ConnectToGateInput(halfAdder, 0)
	b.ConnectToGateInput(halfAdder, 1)
	out.ConnectToGateOutput(halfAdder, "out")
	carry.ConnectToGateOutput(halfAdder, "carry")
}

func (wires Wires) ConvertToFullAdder(abXor, cXor *Xor, abAnd, cAnd *And, or *Or, carryIn *Wire) {
	a := abXor.in[0]
	b := abXor.in[1]
	out := cXor.out["out"]
	carryOut := or.out["out"]
	a.out = slices.DeleteFunc(a.out, func(input GateInput) bool { return input.gate == abXor || input.gate == abAnd })
	b.out = slices.DeleteFunc(b.out, func(input GateInput) bool { return input.gate == abXor || input.gate == abAnd })
	delete(wires, or.in[0].name)
	delete(wires, or.in[1].name)
	delete(wires, abXor.out["out"].name)
	carryIn.out = slices.DeleteFunc(carryIn.out, func(input GateInput) bool { return input.gate == cXor || input.gate == cAnd })
	fullAdder := &FullAdder{InputBitWires: InputBitWires{in: make(map[int]*Wire), bits: make(map[int]bool)}, OutputBitWires: OutputBitWires{out: make(map[string]*Wire)}}
	a.ConnectToGateInput(fullAdder, 0)
	b.ConnectToGateInput(fullAdder, 1)
	carryIn.ConnectToGateInput(fullAdder, 2)
	out.ConnectToGateOutput(fullAdder, "out")
	carryOut.ConnectToGateOutput(fullAdder, "carry")
}

type Wire struct {
	name string
	in   GateOutput
	out  []GateInput
}

func (wire *Wire) SetBit(b bool) {
	for _, out := range wire.out {
		out.gate.SetInputBit(out.index, b)
	}
}

type GateInput struct {
	gate  ConnectableInput
	index int
}

type GateOutput struct {
	gate ConnectableOutput
	name string
}

func (w *Wire) ConnectToGateInput(gate ConnectableInput, index int) {
	w.out = append(w.out, GateInput{gate, index})
	gate.ConnectInput(index, w)
}

func (w *Wire) ConnectToGateOutput(gate ConnectableOutput, name string) {
	w.in = GateOutput{gate, name}
	gate.ConnectOutput(name, w)
}

type InputBitWires struct {
	in   map[int]*Wire
	bits map[int]bool
}

func (w *InputBitWires) ConnectInput(index int, wire *Wire) {
	w.in[index] = wire
}

type OutputBitWires struct {
	out map[string]*Wire
}

func (w *OutputBitWires) ConnectOutput(name string, wire *Wire) {
	w.out[name] = wire
}

type ConnectableInput interface {
	ConnectInput(index int, wire *Wire)
	SetInputBit(index int, b bool)
}

type ConnectableOutput interface {
	ConnectOutput(name string, wire *Wire)
}

type BitGate interface {
	ConnectableInput
	ConnectableOutput
}

type And struct {
	InputBitWires
	OutputBitWires
}

func (gate *And) SetInputBit(index int, b bool) {
	gate.bits[index] = b
	if len(gate.bits) == 2 {
		gate.out["out"].SetBit(gate.bits[0] && gate.bits[1])
	}
}

type Or struct {
	InputBitWires
	OutputBitWires
}

func (gate *Or) SetInputBit(index int, b bool) {
	gate.bits[index] = b
	if len(gate.bits) == 2 {
		gate.out["out"].SetBit(gate.bits[0] || gate.bits[1])
	}
}

type Xor struct {
	InputBitWires
	OutputBitWires
}

func (gate *Xor) SetInputBit(index int, b bool) {
	gate.bits[index] = b
	if len(gate.bits) == 2 {
		gate.out["out"].SetBit(gate.bits[0] != gate.bits[1])
	}
}

type HalfAdder struct {
	InputBitWires
	OutputBitWires
}

func (gate *HalfAdder) SetInputBit(index int, b bool) {
	gate.bits[index] = b
	if len(gate.bits) == 2 {
		gate.out["out"].SetBit(gate.bits[0] != gate.bits[1])
		gate.out["carry"].SetBit(gate.bits[0] && gate.bits[1])
	}
}

type FullAdder struct {
	InputBitWires
	OutputBitWires
}

func (gate *FullAdder) SetInputBit(index int, b bool) {
	gate.bits[index] = b
	if len(gate.bits) == 3 {
		xor := gate.bits[0] != gate.bits[1]
		gate.out["out"].SetBit(xor != gate.bits[2])
		gate.out["carry"].SetBit((gate.bits[0] && gate.bits[1]) || (gate.bits[2] && xor))
	}
}

type Mux struct {
	InputBitWires
}

func (g *Mux) SetInputBit(index int, b bool) {
	g.bits[index] = b
}

func (g *Mux) Get() int {
	var n int
	for i, b := range g.bits {
		if b {
			n |= 1 << i
		}
	}
	return n
}
