package main

import (
	"aoc2023/util"
	"fmt"
	"regexp"
	"strings"
)

var rLine = regexp.MustCompile(`([%&])?(\w+) -> ([a-z, ]+)`)

func main() {
	lines := util.ReadLines()

	network := parseNetwork(lines)
	var low, high int
	for i := 0; i < 1000; i++ {
		addLow, addHigh := network.PushButton()
		low += addLow
		high += addHigh
	}
	fmt.Println(low * high)

	productB := 1
	for _, target := range analyzeNetwork(network) {
		productB *= target
	}
	fmt.Println(productB)
}

func parseNetwork(lines []string) Network {
	moduleIds := make(map[string]int)
	nodes := make([]*Node, 0, len(lines))
	var connections [][2]string

	for _, line := range lines {
		m := rLine.FindStringSubmatch(line)
		moduleIds[m[2]] = len(nodes)
		id := uint64(1) << len(nodes)

		switch m[1] {
		case "%":
			nodes = append(nodes, &Node{module: new(FlipFlop), id: id})
		case "&":
			nodes = append(nodes, &Node{module: new(Conjunction), id: id})
		default:
			nodes = append(nodes, &Node{module: new(Broadcast), id: id})
		}

		for _, to := range strings.Split(m[3], ", ") {
			connections = append(connections, [2]string{m[2], to})
		}
	}

	for _, c := range connections {
		from := nodes[moduleIds[c[0]]]
		toId, hasTo := moduleIds[c[1]]
		if hasTo {
			to := nodes[toId]
			from.to = append(from.to, to)
			to.module.ConnectIn(from.id)
		} else {
			from.to = append(from.to, nil)
			continue
		}
	}

	return Network{nodes[moduleIds["broadcaster"]]}
}

type Module interface {
	ConnectIn(uint64)
	Interact(uint64, bool) (bool, bool)
}

type FlipFlop struct {
	on bool
}

func (*FlipFlop) ConnectIn(_ uint64) {}

func (f *FlipFlop) Interact(sender uint64, high bool) (bool, bool) {
	if high {
		return false, false
	}
	f.on = !f.on
	return true, f.on
}

type Conjunction struct {
	target uint64
	memory uint64
}

func (c *Conjunction) ConnectIn(sender uint64) {
	c.target |= sender
}

func (c *Conjunction) Interact(sender uint64, high bool) (bool, bool) {
	if high {
		c.memory |= sender
	} else {
		c.memory &= ^sender
	}
	return true, c.memory != c.target
}

type Broadcast struct{}

func (*Broadcast) ConnectIn(_ uint64) {}

func (*Broadcast) Interact(_ uint64, high bool) (bool, bool) {
	return true, high
}

type QueuedSignal struct {
	sender uint64
	to     *Node
	high   bool
}

type Node struct {
	module Module
	id     uint64
	to     []*Node
}

func (n *Node) Interact(sender uint64, high bool) []QueuedSignal {
	hasOutput, out := n.module.Interact(sender, high)
	if !hasOutput {
		return nil
	}

	signals := make([]QueuedSignal, len(n.to))
	for i, to := range n.to {
		signals[i] = QueuedSignal{n.id, to, out}
	}
	return signals
}

type Network struct {
	broadcaster *Node
}

func (n Network) PushButton() (int, int) {
	var low, high int
	queue := []QueuedSignal{{0, n.broadcaster, false}}
	for len(queue) > 0 {
		signal := queue[0]
		queue = queue[1:]

		if signal.high {
			high++
		} else {
			low++
		}

		if signal.to != nil {
			queue = append(queue, signal.to.Interact(signal.sender, signal.high)...)
		}
	}

	return low, high
}

func analyzeNetwork(network Network) []int {
	targets := make([]int, len(network.broadcaster.to))
	for i, node := range network.broadcaster.to {
		var target int
		bits := make(map[uint64]int)
		for node != nil {
			value := 1 << len(bits)
			bits[node.id] = value
			var nextNode *Node
			for _, to := range node.to {
				switch to.module.(type) {
				case *FlipFlop:
					nextNode = to
				case *Conjunction:
					target |= value
				}
			}
			node = nextNode
		}
		targets[i] = target
	}
	return targets
}
