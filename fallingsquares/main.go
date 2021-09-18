package main

import "fmt"

type square struct {
	start  int
	end    int
	height int
}

func fallingSquares(positions [][]int) []int {
	result := make([]int, 0, len(positions))
	newPositions := make([]square, 0, len(positions))

	highest := 0
	for _, current := range positions {
		start := current[0]
		end := current[0] + current[1]
		height := current[1]

		finalHeight := current[1]
		for _, square := range newPositions {
			if (start >= square.start && start < square.end) || (end > square.start && end <= square.end) || start <= square.start && end >= square.end {
				if newHeight := height + square.height; newHeight > finalHeight {
					finalHeight = newHeight
				}
			}
		}

		if finalHeight > highest {
			highest = finalHeight
		}

		newPositions = append(newPositions, square{start: start, end: end, height: finalHeight})
		result = append(result, highest)
	}

	return result
}

func main() {
	output := fallingSquares([][]int{{1, 2}, {2, 3}, {6, 1}})
	// expected output: [2,5,5]
	fmt.Println(output)

	output = fallingSquares([][]int{{100, 100}, {200, 100}})
	// expected output: [100, 100]
	fmt.Println(output)

	output = fallingSquares([][]int{{6, 1}, {9, 2}, {2, 4}})
	// expected output: [1, 2, 4]
	fmt.Println(output)

	output = fallingSquares([][]int{{9, 7}, {1, 9}, {3, 1}})
	// expected output: [7, 16, 17]
	fmt.Println(output)

	output = fallingSquares([][]int{{4, 6}, {4, 2}, {4, 3}})
	// expected output: [6, 8, 11]
	fmt.Println(output)

	output = fallingSquares([][]int{{9, 1}, {6, 5}, {6, 7}})
	// expected output: [6, 8, 11]
	fmt.Println(output)
}
