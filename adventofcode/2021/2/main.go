package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func navigate(commands []string) (int, int) {
	x, y := 0, 0
	for _, command := range commands {
		parts := strings.Split(command, " ")
		distance, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}

		switch parts[0] {
		case "forward":
			x += distance
		case "up":
			y -= distance
		case "down":
			y += distance
		default:
			log.Fatal("Unknown command: ", parts[0])
		}
	}

	return x, y
}

func navigateAndAim(commands []string) (int, int) {
	x, y, aim := 0, 0, 0
	for _, command := range commands {
		parts := strings.Split(command, " ")
		distance, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}

		switch parts[0] {
		case "forward":
			x += distance
			y += aim * distance
		case "up":
			aim -= distance
		case "down":
			aim += distance
		default:
			log.Fatal("Unknown command: ", parts[0])
		}
	}

	return x, y
}

func main() {
	read, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer read.Close()

	scanner := bufio.NewScanner(read)
	scanner.Split(bufio.ScanLines)

	input := make([]string, 0)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	x, y := navigate(input)
	fmt.Printf("Part 1: X = %d * Y = %d = %d\n", x, y, x*y)

	x, y = navigateAndAim(input)
	fmt.Printf("Part 2: X = %d * Y = %d = %d\n", x, y, x*y)
}
