package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type entry struct {
	number int
	marked bool
}

type board struct {
	entries [][]entry
}

func (b board) colIsWinning(col int) bool {
	for _, row := range b.entries {
		if !row[col].marked {
			return false
		}
	}

	return true
}

func (b board) rowIsWinning(row int) bool {
	for _, col := range b.entries[row] {
		if !col.marked {
			return false
		}
	}

	return true
}

func (b board) currentScore() int {
	score := 0
	for _, row := range b.entries {
		for _, col := range row {
			if !col.marked {
				score += col.number
			}
		}
	}

	return score
}

func bingo(boards []board, input []int) int {
	for _, number := range input {
		for _, board := range boards {
			for i, row := range board.entries {
				for j, col := range row {
					if col.number == number {
						board.entries[i][j].marked = true
						if board.colIsWinning(j) || board.rowIsWinning(i) {
							return board.currentScore() * number
						}
					}
				}
			}
		}
	}

	return 0
}

func removeBoard(boards []board, n int) []board {
	for i := range boards {
		if i == n {
			return append(boards[:i], boards[i+1:]...)
		}
	}

	return boards
}

func looserBingo(boards []board, input []int) int {
	for n, number := range input {
		for b, board := range boards {
			for i, row := range board.entries {
				for j, col := range row {
					if col.number == number {
						board.entries[i][j].marked = true
						if board.colIsWinning(j) || board.rowIsWinning(i) {
							if len(boards) == 1 {
								return board.currentScore() * number
							}

							return looserBingo(removeBoard(boards, b), input[n:])
						}
					}
				}
			}
		}
	}

	return 0
}

func main() {
	read, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer read.Close()

	scanner := bufio.NewScanner(read)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	input := make([]int, 0)
	for _, number := range strings.Split(scanner.Text(), ",") {
		n, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal(err)
		}

		input = append(input, n)
	}
	scanner.Scan()

	boards := make([]board, 0)
	current := board{entries: [][]entry{}}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if len(line) < 2 {
			boards = append(boards, current)
			current = board{entries: [][]entry{}}
			continue
		}

		col := make([]entry, 0)
		for _, number := range line {
			number = strings.TrimSpace(number)
			if number == "" {
				continue
			}

			number = strings.TrimSpace(number)
			n, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}

			col = append(col, entry{number: n})
		}

		current.entries = append(current.entries, col)
	}

	fmt.Println("part 1: ", bingo(boards, input))
	fmt.Println("part 2: ", looserBingo(boards, input))
}
