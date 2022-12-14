package main

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type entry struct {
	patterns, outputs []string
}

func main() {
	lines := strings.Split(input, "\n")
	entries := parseEntries(lines)

	println(getUniqueOutputCount(entries))
	println(getTotalOutput(entries))
}

func getUniqueOutputCount(entries []entry) int {
	total := 0

	for _, entry := range entries {
		for _, output := range entry.outputs {
			length := len(output)

			if (length >= 2 && length <= 4) || length == 7 {
				total++
			}
		}
	}

	return total
}

func getTotalOutput(entries []entry) int {
	tot := 0

	for _, e := range entries {
		tot += getOutput(e)
	}

	return tot
}

func getOutput(entry entry) int {
	res := analyzePatterns(entry.patterns)
	var str string

	for _, o := range entry.outputs {
		str += strconv.Itoa(res[sortPattern(o)])
	}

	val, _ := strconv.Atoi(str)

	return val
}

func analyzePatterns(patterns []string) map[string]int {
	sort.Slice(patterns, func(i, j int) bool {
		return len(patterns[i]) < len(patterns[j])
	})

	matches := map[int]string{
		1: patterns[0],
		7: patterns[1],
		4: patterns[2],
		8: patterns[9],
	}

	for _, p := range patterns[3:6] {
		if fitsIn(matches[1], p) {
			matches[3] = p
		}
	}

	var inFive rune

	for _, c := range matches[4] {
		if !strings.ContainsRune(matches[3], c) {
			inFive = c
			break
		}
	}

	for _, p := range patterns[3:6] {
		if strings.ContainsRune(p, inFive) {
			matches[5] = p
		} else if p != matches[3] {
			matches[2] = p
		}
	}

	for _, p := range patterns[6:9] {
		if !fitsIn(matches[1], p) {
			matches[6] = p
		} else if fitsIn(matches[4], p) {
			matches[9] = p
		} else {
			matches[0] = p
		}
	}

	res := make(map[string]int)

	for i, p := range matches {
		res[sortPattern(p)] = i
	}

	return res
}

func sortPattern(pattern string) string {
	arr := []rune(pattern)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	return string(arr)
}

func fitsIn(a, b string) bool {
	for _, c := range a {
		if !strings.ContainsRune(b, c) {
			return false
		}
	}

	return true
}

func parseEntries(lines []string) []entry {
	var entries = make([]entry, len(lines))

	for i, line := range lines {
		var entry entry
		parts := strings.Split(line, " | ")

		entry.patterns = strings.Fields(parts[0])
		entry.outputs = strings.Fields(parts[1])

		entries[i] = entry
	}

	return entries
}
