package main

import (
	"aoc2023/util"
	"fmt"
	"strconv"

	"golang.org/x/exp/maps"

	parsec "github.com/prataprc/goparsec"
)

func main() {
	lines := util.ReadLines()
	workflows := make(map[string]Rule)
	var parts []map[string]int

	for _, line := range lines {
		res, _ := parser()(parsec.NewScanner([]byte(line)))
		if res != nil {
			switch res := res.([]parsec.ParsecNode)[0].(type) {
			case Workflow:
				workflows[res.id] = res.rule
			case Part:
				parts = append(parts, res)
			}
		}
	}
	in := workflows["in"]

	var sumA int
	for _, part := range parts {
		if in.Evaluate(workflows, part) {
			for _, value := range part {
				sumA += value
			}
		}
	}
	fmt.Println(sumA)

	possibleParts := map[string][2]int{"x": {1, 4001}, "m": {1, 4001}, "a": {1, 4001}, "s": {1, 4001}}
	fmt.Println(in.CountMatching(workflows, possibleParts))
}

func parser() parsec.Parser {
	var rule parsec.Parser

	open := parsec.Atom("{", "OPEN")
	close := parsec.Atom("}", "CLOSE")
	comma := parsec.Maybe(nil, parsec.Atom(",", "COMMA"))
	colon := parsec.Atom(":", "COLON")
	equals := parsec.Atom("=", "EQUALS")

	id := parsec.Token(`[a-z]+`, "ID")
	operator := parsec.Token(`[<>]`, "OPERATOR")
	number := parsec.Token(`[0-9]+`, "NUMBER")

	accept := wrapParser(parsec.AtomExact("A", "ACCEPT"), func(_ parsec.ParsecNode) Accepted {
		return Accepted{}
	})

	reject := wrapParser(parsec.AtomExact("R", "REJECT"), func(_ parsec.ParsecNode) Rejected {
		return Rejected{}
	})

	gotoWorkflow := wrapParser(id, func(n parsec.ParsecNode) GotoWorkflow {
		return GotoWorkflow{n.(*parsec.Terminal).Value}
	})

	conditional := wrapParser(parsec.And(nil, id, operator, number, colon, &rule, comma, &rule), func(n parsec.ParsecNode) Conditional {
		ns := n.([]parsec.ParsecNode)
		value, _ := strconv.Atoi(ns[2].(*parsec.Terminal).Value)
		return Conditional{
			category: ns[0].(*parsec.Terminal).Value,
			operator: ns[1].(*parsec.Terminal).Value,
			value:    value,
			ifTrue:   ns[4].(Rule),
			ifFalse:  ns[6].(Rule),
		}
	})

	rule = wrapParser(parsec.OrdChoice(nil, accept, reject, conditional, gotoWorkflow), func(n parsec.ParsecNode) parsec.ParsecNode {
		return n.([]parsec.ParsecNode)[0]
	})

	workflow := wrapParser(parsec.And(nil, id, open, rule, close), func(n parsec.ParsecNode) Workflow {
		ns := n.([]parsec.ParsecNode)
		return Workflow{ns[0].(*parsec.Terminal).Value, ns[2].(Rule)}
	})

	rating := wrapParser(parsec.And(nil, id, equals, number, comma), func(n parsec.ParsecNode) Rating {
		ns := n.([]parsec.ParsecNode)
		value, _ := strconv.Atoi(ns[2].(*parsec.Terminal).Value)
		return Rating{ns[0].(*parsec.Terminal).Value, value}
	})

	ratings := wrapParser(parsec.Kleene(nil, rating, nil), func(n parsec.ParsecNode) map[string]int {
		m := make(map[string]int)
		for _, rating := range n.([]parsec.ParsecNode) {
			m[rating.(Rating).category] = rating.(Rating).value
		}
		return m
	})

	part := wrapParser(parsec.And(nil, open, ratings, close), func(n parsec.ParsecNode) Part {
		return Part(n.([]parsec.ParsecNode)[1].(map[string]int))
	})

	return parsec.OrdChoice(nil, workflow, part)
}

func wrapParser[T any](parser parsec.Parser, mapper func(node parsec.ParsecNode) T) parsec.Parser {
	return func(s parsec.Scanner) (parsec.ParsecNode, parsec.Scanner) {
		n, s := parser(s)
		if n == nil {
			return nil, s
		}
		return mapper(n), s
	}
}

type Rule interface {
	Evaluate(workflows map[string]Rule, part Part) bool
	CountMatching(workflows map[string]Rule, parts map[string][2]int) int
}

type Accepted struct{}

func (Accepted) Evaluate(_ map[string]Rule, _ Part) bool {
	return true
}

func (Accepted) CountMatching(_ map[string]Rule, parts map[string][2]int) int {
	count := 1
	for _, r := range parts {
		count *= r[1] - r[0]
	}
	return count
}

type Rejected struct{}

func (Rejected) Evaluate(_ map[string]Rule, _ Part) bool {
	return false
}

func (Rejected) CountMatching(_ map[string]Rule, _ map[string][2]int) int {
	return 0
}

type GotoWorkflow struct {
	id string
}

func (g GotoWorkflow) Evaluate(workflows map[string]Rule, part Part) bool {
	return workflows[g.id].Evaluate(workflows, part)
}

func (g GotoWorkflow) CountMatching(workflows map[string]Rule, parts map[string][2]int) int {
	return workflows[g.id].CountMatching(workflows, parts)
}

type Conditional struct {
	category string
	operator string
	value    int
	ifTrue   Rule
	ifFalse  Rule
}

func (c Conditional) Evaluate(workflows map[string]Rule, part Part) bool {
	var guard bool
	switch c.operator {
	case "<":
		guard = part[c.category] < c.value
	case ">":
		guard = part[c.category] > c.value
	}

	if guard {
		return c.ifTrue.Evaluate(workflows, part)
	} else {
		return c.ifFalse.Evaluate(workflows, part)
	}
}

func (c Conditional) CountMatching(workflows map[string]Rule, parts map[string][2]int) int {
	trueParts := maps.Clone(parts)
	falseParts := maps.Clone(parts)
	r := parts[c.category]

	switch c.operator {
	case "<":
		if c.value < r[0] {
			trueParts[c.category] = [2]int{0, 0}
		} else if r[1] <= c.value {
			falseParts[c.category] = [2]int{0, 0}
		} else {
			trueParts[c.category] = [2]int{trueParts[c.category][0], c.value}
			falseParts[c.category] = [2]int{c.value, falseParts[c.category][1]}
		}
	case ">":
		if c.value > r[1] {
			trueParts[c.category] = [2]int{0, 0}
		} else if r[0] >= c.value {
			falseParts[c.category] = [2]int{0, 0}
		} else {
			trueParts[c.category] = [2]int{c.value + 1, trueParts[c.category][1]}
			falseParts[c.category] = [2]int{falseParts[c.category][0], c.value + 1}
		}
	}

	return c.ifTrue.CountMatching(workflows, trueParts) + c.ifFalse.CountMatching(workflows, falseParts)
}

type Workflow struct {
	id   string
	rule Rule
}

type Rating struct {
	category string
	value    int
}

type Part map[string]int
