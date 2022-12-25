package main

import (
	_ "embed"
	"math"
	"strings"
)

type rules map[string]string
type pairs map[string]int
type chars map[rune]int

//go:embed input.txt
var input string

func main() {
	parts := strings.Split(input, "\n\n")
	template := parts[0]
	rules := parseRules(parts[1])

	println(scorePolymer(buildPolymer(template, rules, 10)))
	println(scorePolymer(buildPolymer(template, rules, 40)))
}

func scorePolymer(cf chars) int {
	min := math.MaxInt
	max := math.MinInt

	for _, s := range cf {
		if s < min {
			min = s
		}
		if s > max {
			max = s
		}
	}

	return max - min
}

func buildPolymer(tmp string, rules rules, n int) chars {
	pf := buildPairs(tmp)

	for i := 0; i < n; i++ {
		f := make(pairs)
		for p, n := range pf {
			c, exists := rules[p]

			if !exists {
				continue
			}

			f[string(p[0])+c] += n
			f[c+string(p[1])] += n
		}

		pf = f
	}

	return toChars(pf, tmp)
}

func toChars(pf pairs, tmp string) chars {
	cf := make(chars)

	for p, n := range pf {
		for _, c := range p {
			cf[c] += n
		}
	}

	for c, n := range cf {
		cf[c] = n / 2

		if rune(tmp[0]) == c {
			cf[c]++
		}

		if rune(tmp[len(tmp)-1]) == c {
			cf[c]++
		}
	}

	return cf
}

func buildPairs(polymer string) pairs {
	pf := make(pairs)

	for j := 1; j < len(polymer); j++ {
		p := string(polymer[j-1]) + string(polymer[j])
		pf[p]++
	}

	return pf
}

func parseRules(input string) rules {
	lines := strings.Split(input, "\n")
	rules := make(rules)

	for _, l := range lines {
		parts := strings.Split(l, " -> ")
		rules[parts[0]] = parts[1]
	}

	return rules
}
