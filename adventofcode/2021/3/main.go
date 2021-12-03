package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func toInt(str string) int {
	digits := len(str)
	result := 0
	for i, n := range str {
		number, err := strconv.Atoi(string(n))
		if err != nil {
			log.Fatal(err)
		}

		result += int(math.Exp2(float64(digits-i-1)) * float64(number))
	}

	return result
}

func powerConsumption(input []string) (int, int) {
	rows := len(input[0])
	gama, epsilon := "", ""

	for i := 0; i < rows; i++ {
		ones := 0
		for _, line := range input {
			if line[i] == '1' {
				ones++
			}
		}

		if ones > len(input)/2 {
			gama += "1"
			epsilon += "0"
		} else {
			gama += "0"
			epsilon += "1"
		}
	}

	return toInt(gama), toInt(epsilon)
}

func oxygenRating(entries []string, pos int) string {
	if len(entries) == 1 {
		return entries[0]
	}

	oneCount := 0
	zeroCount := 0
	ones := make([]string, 0, len(entries))
	zeroes := make([]string, 0, len(entries))

	for _, line := range entries {
		if line[pos] == '1' {
			oneCount++
			ones = append(ones, line)
			continue
		}

		zeroCount++
		zeroes = append(zeroes, line)
	}

	if oneCount >= zeroCount {
		return oxygenRating(ones, pos+1)
	} else {
		return oxygenRating(zeroes, pos+1)
	}
}

func carbonDioxideRating(entries []string, pos int) string {
	if len(entries) == 1 {
		return entries[0]
	}

	zeroCount := 0
	oneCount := 0
	ones := make([]string, 0, len(entries))
	zeroes := make([]string, 0, len(entries))

	for _, line := range entries {
		if line[pos] == '1' {
			oneCount++
			ones = append(ones, line)
			continue
		}

		zeroCount++
		zeroes = append(zeroes, line)
	}

	if zeroCount > oneCount {
		return carbonDioxideRating(ones, pos+1)
	} else {
		return carbonDioxideRating(zeroes, pos+1)
	}
}

func lifeSupportRating(input []string) (int, int) {
	oxygen := oxygenRating(input, 0)
	carbonDioxide := carbonDioxideRating(input, 0)

	return toInt(oxygen), toInt(carbonDioxide)
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

	gama, epsilon := powerConsumption(input)
	fmt.Println(gama * epsilon)

	oxygen, carbonDioxide := lifeSupportRating(input)
	fmt.Println(oxygen * carbonDioxide)
}
