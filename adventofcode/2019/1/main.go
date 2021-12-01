package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func calculateFuel(mass int) int {
	return (mass / 3) - 2
}

func main() {
	read, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(read)
	scanner.Split(bufio.ScanWords)

	var sum int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		sum += calculateFuel(num)
	}

	fmt.Println("Result: ", sum)
}
