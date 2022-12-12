package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type board [][]int

type game struct {
	draws  []int
	boards []board
}

func main() {
	game := parseGame(input)

	scores := play(game)

	println(scores[0])
	println(scores[len(scores)-1])
}

func play(game game) []int {
	var scores []int
	won := make([]bool, len(game.boards))

	for _, n := range game.draws {
		for i, board := range game.boards {
			if won[i] {
				continue
			}

			board.mark(n)

			if board.hasWon() {
				scores = append(scores, board.getScore()*n)
				won[i] = true
			}
		}
	}

	return scores
}

func (board *board) mark(n int) {
	for _, row := range *board {
		for i, col := range row {
			if col == n {
				row[i] = 0
			}
		}
	}
}

func (board *board) getScore() int {
	s := 0

	for _, row := range *board {
		for _, col := range row {
			s += col
		}
	}

	return s
}

func (board *board) hasWon() bool {
	for _, row := range *board {
		s := 0

		for _, col := range row {
			s += col
		}

		if s == 0 {
			return true
		}
	}

	for i := range (*board)[0] {
		s := 0

		for _, row := range *board {
			s += row[i]
		}

		if s == 0 {
			return true
		}
	}

	return false
}

func parseGame(input string) game {
	var game game
	parts := strings.Split(input, "\n\n")

	for _, d := range strings.Split(parts[0], ",") {
		game.draws = append(game.draws, toInt(d))
	}

	for _, b := range parts[1:] {
		game.boards = append(game.boards, parseBoard(b))
	}

	return game
}

func parseBoard(raw string) board {
	rows := strings.Split(raw, "\n")
	board := make(board, len(rows))

	for i, row := range rows {
		cols := strings.Fields(row)

		for _, c := range cols {
			board[i] = append(board[i], toInt(c))
		}
	}

	return board
}

func toInt(str string) int {
	n, err := strconv.Atoi(str)

	if err != nil {
		log.Fatal(err)
	}

	return n
}
