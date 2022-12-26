package main

import (
	"container/heap"
	_ "embed"
	"math"
	"strings"
)

type nodeQueue []*node
type node struct {
	dist  int
	coord coord
}

type coord [2]int
type grid [][]int

//go:embed input.txt
var input string
var vecs = []coord{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func main() {
	g := parseInput(input)

	println(getShortestDist(g))
	println(getShortestDist(extendGrid(g, 5)))
}

func getShortestDist(g grid) int {
	dists := makeGrid(g, math.MaxInt)
	start := coord{0, 0}
	end := coord{g.width() - 1, g.height() - 1}
	dists.set(start, 0)

	q := nodeQueue{&node{
		dist:  0,
		coord: start,
	}}

	for q.Len() > 0 {
		item := heap.Pop(&q)
		c := *item.(*node)
		d := dists.get(c.coord)

		for _, vec := range vecs {
			n := c.coord.add(vec)

			if !g.has(n) {
				continue
			}

			nd := g.get(n) + d

			if dists.get(n) > nd {
				dists.set(n, nd)

				heap.Push(&q, &node{
					dist:  dists.get(n),
					coord: n,
				})
			}
		}
	}

	return dists.get(end)
}

func makeGrid(g grid, v int) grid {
	d := make(grid, g.height())

	for y, _ := range d {
		d[y] = make([]int, g.width())
		for x, _ := range d[y] {
			d[y][x] = v
		}
	}

	return d
}

func extendGrid(g grid, m int) grid {
	w, h := g.width(), g.height()
	eg := make(grid, h*m)

	for y := 0; y < h*m; y++ {
		eg[y] = make([]int, w*m)

		for x := 0; x < w*m; x++ {
			d := (x / w) + (y / h)
			v := g[y%(h)][x%(w)]
			v = ((v + d - 1) % 9) + 1
			eg[y][x] = v
		}
	}

	return eg
}

func parseInput(input string) grid {
	lns := strings.Split(input, "\n")
	g := make(grid, len(lns))

	for y, ln := range lns {
		g[y] = make([]int, len(ln))
		for x, c := range ln {
			g[y][x] = int(c - 48)
		}
	}

	return g
}

func (g grid) has(c coord) bool {
	return c[0] >= 0 && c[0] < g.width() && c[1] >= 0 && c[1] < g.height()
}

func (g grid) width() int { return len(g[0]) }

func (g grid) height() int { return len(g) }

func (g grid) get(c coord) int { return g[c[1]][c[0]] }

func (g grid) set(c coord, v int) { g[c[1]][c[0]] = v }

func (c coord) add(o coord) coord { return coord{c[0] + o[0], c[1] + o[1]} }

func (q nodeQueue) Len() int { return len(q) }

func (q nodeQueue) Less(i, j int) bool { return q[i].dist < q[j].dist }

func (q nodeQueue) Swap(i, j int) { q[i], q[j] = q[j], q[i] }

func (q *nodeQueue) Pop() any {
	old := *q
	item := old[len(old)-1]
	*q = old[0 : len(old)-1]
	return item
}

func (q *nodeQueue) Push(x any) {
	*q = append(*q, x.(*node))
}
