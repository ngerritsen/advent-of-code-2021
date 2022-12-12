package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type fish map[int]int

func main() {
	println(countFish(runSim(parseFish(input), 80)))
	println(countFish(runSim(parseFish(input), 256)))
}

func parseFish(input string) fish {
	fish := make(fish)

	for _, x := range strings.Split(input, ",") {
		fish[toInt(x)]++
	}

	return fish
}

func countFish(fish fish) int {
	tot := 0

	for _, n := range fish {
		tot += n
	}

	return tot
}

func runSim(fish fish, days int) fish {
	for i := 0; i < days; i++ {
		fish = runDay(fish)
	}

	return fish
}

func runDay(fish fish) fish {
	ready := fish[0]

	for i := 1; i <= 8; i++ {
		fish[i-1] = fish[i]
	}

	fish[6] += ready
	fish[8] = ready

	return fish
}

func toInt(str string) int {
	n, err := strconv.Atoi(str)

	if err != nil {
		log.Fatalln(err)
	}

	return n
}
