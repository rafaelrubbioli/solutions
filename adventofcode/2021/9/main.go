package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func lowPoints(table [][]int) []int {
	low := make([]int, 0)

	for row, line := range table {
		for col, item := range line {
			switch row {
			case 0:
				if item >= table[row+1][col] {
					continue
				}

			case len(table) - 1:
				if item >= table[row-1][col] {
					continue
				}

			default:
				if item >= table[row-1][col] || item >= table[row+1][col] {
					continue
				}
			}

			switch col {
			case 0:
				if item >= table[row][col+1] {
					continue
				}

			case len(line) - 1:
				if item >= table[row][col-1] {
					continue
				}

			default:
				if item >= table[row][col-1] || item >= table[row][col+1] {
					continue
				}
			}

			low = append(low, item)
		}
	}

	return low
}

func calculateRisk(lowPoints []int) int {
	sum := 0
	for _, point := range lowPoints {
		sum += point + 1
	}

	return sum
}

func basins(table [][]int) int {
	sizes := make([]int, 3)
	for row, line := range table {
		for col, item := range line {
			switch row {
			case 0:
				if item >= table[row+1][col] {
					continue
				}

			case len(table) - 1:
				if item >= table[row-1][col] {
					continue
				}

			default:
				if item >= table[row-1][col] || item >= table[row+1][col] {
					continue
				}
			}

			switch col {
			case 0:
				if item >= table[row][col+1] {
					continue
				}

			case len(line) - 1:
				if item >= table[row][col-1] {
					continue
				}

			default:
				if item >= table[row][col-1] || item >= table[row][col+1] {
					continue
				}
			}

			currentSize := basinSize(make(map[string]struct{}), table, row, col)
			sizes = replaceLowest(sizes, currentSize)
		}
	}

	result := 1
	for _, basin := range sizes {
		result *= basin
	}

	return result
}

func replaceLowest(numbers []int, new int) []int {
	lowest := numbers[0]
	lowestIndex := 0
	for i, n := range numbers {
		if n < lowest {
			lowest = n
			lowestIndex = i
		}
	}

	if new > lowest {
		numbers[lowestIndex] = new
	}

	return numbers
}

func basinSize(seen map[string]struct{}, table [][]int, i, j int) int {
	_, ok := seen[fmt.Sprintf("%d,%d", i, j)]
	if ok || i == -1 || j == -1 || i == len(table) || j == len(table[i]) || table[i][j] == 9 {
		return 0
	}

	seen[fmt.Sprintf("%d,%d", i, j)] = struct{}{}
	size := 1 + basinSize(seen, table, i-1, j) + basinSize(seen, table, i, j-1) + basinSize(seen, table, i+1, j) + basinSize(seen, table, i, j+1)
	return size
}

func main() {
	read, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer read.Close()

	scanner := bufio.NewScanner(read)
	scanner.Split(bufio.ScanLines)

	var input [][]int
	for scanner.Scan() {
		var line []int
		for _, c := range scanner.Text() {
			num, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal(err)
			}

			line = append(line, num)
		}

		input = append(input, line)
	}

	fmt.Println("Part 1:", calculateRisk(lowPoints(input)))
	fmt.Println("Part 2:", basins(input))
}
