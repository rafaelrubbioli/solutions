package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func numberOfFishBruteForce(fishes []int, days int) int {
	for i := 0; i < days; i++ {
		newFish := make([]int, 0)
		for f, fish := range fishes {
			fishes[f]--
			if fish == 0 {
				fishes[f] = 6
				newFish = append(newFish, 8)
			}

		}

		fishes = append(fishes, newFish...)
	}

	return len(fishes)
}

func numberOfFish(fishes []int, days int) int {
	dayWindow := make([]int, 9)
	for _, fish := range fishes {
		dayWindow[fish] += 1
	}

	for day := 0; day < days; day++ {
		currentCycle := day % 9
		newFish := dayWindow[currentCycle]
		dayWindow[(currentCycle+7)%9] += newFish
	}

	total := 0
	for _, fish := range dayWindow {
		total += fish
	}

	return total
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

	fmt.Println("part 1: ", numberOfFishBruteForce(input, 80))
	fmt.Println("part 2: ", numberOfFish(input, 256))
}
