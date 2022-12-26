package main

import (
	_ "embed"
	"math"
	"strconv"
	"strings"
)

type packet struct {
	version  int
	typeid   int
	value    int
	children []*packet
}

//go:embed input.txt
var input string

func main() {
	p, _ := parseBits(hexBin(input))

	println(getTotalVersion(p))
	println(getValue(p))
}

func parseBits(bs string) (*packet, int) {
	p := new(packet)
	p.version = binToInt(bs[:3])
	p.typeid = binToInt(bs[3:6])
	tl := 6

	if p.typeid == 4 {
		v, l := parseLiteral(bs[tl:])
		p.value = v
		tl += l
	} else {
		c, l := parseOperator(bs[tl:])
		p.children = c
		tl += l
	}

	return p, tl
}

func parseOperator(bs string) ([]*packet, int) {
	if bs[0] == "0"[0] {
		return parseLengthOperator(bs)
	}

	return parseCountOperator(bs)
}

func parseCountOperator(bs string) ([]*packet, int) {
	var c []*packet
	tl := 12

	pc := binToInt(bs[1:12])

	for i := 0; i < pc; i++ {
		p, l := parseBits(bs[tl:])
		tl += l
		c = append(c, p)
	}

	return c, tl
}

func parseLengthOperator(bs string) ([]*packet, int) {
	var c []*packet
	var o int

	pl := binToInt(bs[1:16])

	for pl > o {
		p, l := parseBits(bs[16+o:])
		o += l
		c = append(c, p)
	}

	return c, pl + 16
}

func parseLiteral(bs string) (int, int) {
	i, bn := 0, ""

	for true {
		s := i * 5
		bn += bs[s+1 : s+5]
		if bs[s] == "0"[0] {
			break
		}
		i++
	}

	return binToInt(bn), (i + 1) * 5
}

func binToInt(s string) int {
	i, _ := strconv.ParseInt(s, 2, 0)
	return int(i)
}

func hexBin(hs string) string {
	bs := ""

	for _, h := range hs {
		i, _ := strconv.ParseInt(string(h), 16, 0)
		b := strconv.FormatInt(i, 2)
		bs += strings.Repeat("0", 4-len(b)) + b
	}

	return bs
}

func getTotalVersion(p *packet) int {
	v := p.version

	for _, c := range p.children {
		v += getTotalVersion(c)
	}

	return v
}

func getValue(p *packet) int {
	var vs []int
	var v int

	for _, c := range p.children {
		vs = append(vs, getValue(c))
	}

	switch p.typeid {
	case 0:
		for _, cv := range vs {
			v += cv
		}
	case 1:
		for i, cv := range vs {
			if i == 0 {
				v = cv
			} else {
				v *= cv
			}
		}
	case 2:
		v = math.MaxInt
		for _, cv := range vs {
			if cv < v {
				v = cv
			}
		}
	case 3:
		for _, cv := range vs {
			if cv > v {
				v = cv
			}
		}
	case 4:
		v = p.value
	case 5:
		if vs[0] > vs[1] {
			v = 1
		}
	case 6:
		if vs[0] < vs[1] {
			v = 1
		}
	case 7:
		if vs[0] == vs[1] {
			v = 1
		}
	}

	return v
}
