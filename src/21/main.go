package main

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

const (
	linearEndScore = 1000
	diracEndScore  = 21
)

var diracFrequencies = map[int]int{
	3: 1,
	4: 3,
	5: 6,
	6: 7,
	7: 6,
	8: 3,
	9: 1,
}

type die struct {
	max, val int
}

type player struct {
	pos, score int
}

type game struct {
	round  int
	p1, p2 player
}

func main() {
	g := parseGame(input)

	println(getLinearResult(g))
	println(getMaxDiracWins(g))
}

func getMaxDiracWins(g game) int {
	p1, p2 := findMaxDiracWins(g)

	if p1 > p2 {
		return p1
	}

	return p2
}

func findMaxDiracWins(g game) (int, int) {
	res := [2]int{0, 0}

	for n, f := range diracFrequencies {
		p := g.get()
		p = p.move(n)

		if p.score >= diracEndScore {
			res[g.round%2] += f
			continue
		}

		w1, w2 := findMaxDiracWins(g.set(p))

		res[0] += w1 * f
		res[1] += w2 * f
	}

	return res[0], res[1]
}

func getLinearResult(g game) int {
	d := die{100, 0}

	for true {
		p := g.get()

		p = p.move(d.roll(3))
		g = g.set(p)

		if p.score >= linearEndScore {
			return g.get().score * (g.round * 3)
		}
	}

	return -1
}

func (d *die) roll(n int) int {
	s := 0

	for i := 0; i < n; i++ {
		d.val = (d.val + 1) % d.max
		s += d.val
	}

	return s
}

func (p player) move(n int) player {
	p.pos = (p.pos + n) % 10
	p.score += p.pos + 1

	return p
}

func (g game) get() player {
	if g.round%2 == 0 {
		return g.p1
	}

	return g.p2
}

func (g game) set(p player) game {
	if g.round%2 == 0 {
		return game{g.round + 1, p, g.p2}
	}

	return game{g.round + 1, g.p1, p}
}

func parseGame(input string) game {
	lns := strings.Split(input, "\n")

	return game{
		round: 0,
		p1:    player{pos: parseLine(lns[0]) - 1},
		p2:    player{pos: parseLine(lns[1]) - 1},
	}
}

func parseLine(l string) int {
	n, _ := strconv.Atoi(strings.Split(l, ": ")[1])
	return n
}
