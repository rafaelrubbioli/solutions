package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func countEasyDigits(input [][]string) int {
	count := 0
	for _, line := range input {
		for _, digit := range line {
			length := len(digit)
			if length == 2 || length == 3 || length == 4 || length == 7 {
				count++
			}
		}
	}

	return count
}

func decode(input, output []string) int {
	dict := make(map[string]int)
	var one, four string
	for _, digit := range append(input, output...) {
		if _, ok := dict[digit]; ok {
			continue
		}

		switch len(digit) {
		case 2:
			dict[digit] = 1
			one = digit
		case 3:
			dict[digit] = 7
		case 4:
			dict[digit] = 4
			four = digit
		case 7:
			dict[digit] = 8
		}
	}

	num := 0
	for i, digit := range output {
		number, ok := dict[digit]
		if !ok {
			length := len(digit)
			switch length {
			case 5:
				// 2, 3, 5
				if lenAfterRemove(digit, one) == 3 {
					dict[digit] = 3
					number = 3
				} else if lenAfterRemove(digit, four) == 2 {
					dict[digit] = 5
					number = 5
				} else {
					dict[digit] = 2
					number = 2
				}
			case 6:
				// 0, 6, 9
				if lenAfterRemove(digit, one) == 5 {
					dict[digit] = 6
					number = 6
				} else if lenAfterRemove(digit, four) == 2 {
					dict[digit] = 9
					number = 9
				} else {
					dict[digit] = 0
					number = 0
				}
			}
		}

		num += number * int(math.Pow(10, float64(len(output)-i-1)))
	}

	return num
}

func lenAfterRemove(digit, toRemove string) int {
	for _, c := range toRemove {
		digit = strings.ReplaceAll(digit, string(c), "")
	}

	return len(digit)
}

func main() {
	read, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer read.Close()

	scanner := bufio.NewScanner(read)
	scanner.Split(bufio.ScanLines)

	var input [][]string
	var results [][]string
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		input = append(input, strings.Split(parts[0], " "))
		results = append(results, strings.Split(parts[1], " "))
	}

	fmt.Println("Part 1: ", countEasyDigits(results))

	sum := 0
	for i, line := range input {
		sum += decode(append(line, results[i]...), results[i])
	}
	fmt.Println("Part 2: ", sum)
}
