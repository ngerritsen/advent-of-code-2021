package main

import (
	_ "embed"
	"encoding/json"
	"math"
	"reflect"
	"strings"
)

//go:embed input.txt
var input string

type node struct {
	left, right *node
	value       int
}

type state struct {
	left           *node
	carry          int
	done, exploded bool
}

func main() {
	lines := strings.Split(input, "\n")

	println(getMagnitude(sum(parseNodes(lines))))
	println(getMaxMagnitude(lines))
}

func getMaxMagnitude(lns []string) int {
	mm, l := math.MinInt, len(lns)

	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			mm = max(getMagnitude(sum(parseNodes(getPair(lns, i, j)))), mm)
			mm = max(getMagnitude(sum(parseNodes(getPair(lns, j, i)))), mm)
		}
	}

	return mm
}

func getMagnitude(n *node) int {
	if n.isLiteral() {
		return n.value
	}

	return getMagnitude(n.left)*3 + getMagnitude(n.right)*2
}

func sum(nodes []*node) *node {
	var sum *node

	for _, d := range nodes {
		if sum == nil {
			sum = d
		} else {
			sum = &node{sum, d, -1}
		}

		reduce(sum)
	}

	return sum
}

func reduce(n *node) {
	if explode(n, &state{carry: -1}, 0) || split(n) {
		reduce(n)
	}
}

func explode(n *node, s *state, d int) bool {
	if n == nil {
		return false
	}

	if n.isLiteral() && !s.exploded {
		s.left = n
	}

	if n.isLiteral() && s.exploded && !s.done {
		n.value += s.carry
		s.done = true
		return true
	}

	if !n.isLiteral() && d > 3 && !s.exploded {
		if s.left != nil {
			s.left.value += n.left.value
		}

		s.carry = n.right.value
		s.exploded = true
		n.explode()

		return true
	}

	l, r := explode(n.left, s, d+1), explode(n.right, s, d+1)

	return l || r
}

func split(n *node) bool {
	if n == nil {
		return false
	}

	if n.isLiteral() && n.value > 9 {
		n.left = &node{value: int(math.Floor(float64(n.value) / 2))}
		n.right = &node{value: int(math.Ceil(float64(n.value) / 2))}
		n.value = -1
		return true
	}

	return split(n.left) || split(n.right)
}

func parseNodes(lines []string) []*node {
	nodes := make([]*node, len(lines))

	for i, l := range lines {
		var d []interface{}
		err := json.Unmarshal([]byte(l), &d)

		if err != nil {
			panic(err)
		}

		nodes[i] = parseNode(d)
	}

	return nodes
}

func parseNode(d any) *node {
	if reflect.ValueOf(d).Kind() == reflect.Float64 {
		return &node{value: int(d.(float64))}
	}

	da := d.([]any)

	return &node{parseNode(da[0]), parseNode(da[1]), -1}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func getPair(lns []string, i, j int) []string {
	return []string{lns[i], lns[j]}
}

func (n *node) isLiteral() bool {
	return n.value > -1
}

func (n *node) explode() {
	n.left = nil
	n.right = nil
	n.value = 0
}
