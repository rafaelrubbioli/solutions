package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func polymer(base string, rules map[string]string, step, target int) string {
	if step == target {
		return base
	}

	newWord := ""
	for i, c := range base {
		if i == len(base)-1 {
			newWord += string(c)
			break
		}

		newWord += string(c)
		if replacement, ok := rules[string(c)+string(base[i+1])]; ok {
			newWord += replacement
		}
	}

	return polymer(newWord, rules, step+1, target)
}

func mostAndLeastCommon(s string) int {
	var occurences = make(map[rune]int)
	for _, c := range s {
		occurences[c]++
	}

	most, least := math.MinInt, math.MaxInt
	for _, v := range occurences {
		if v > most {
			most = v
		}

		if v < least {
			least = v
		}
	}

	return most - least
}

func calculateMostAndLeastCommon(base string, rules map[string]string, steps int) int {
	var occurences = make(map[rune]int)
	for _, c := range base {
		occurences[c]++
	}

	pairs := make(map[string]int)
	for i, c := range base {
		if i == len(base)-1 {
			break
		}

		pairs[string(c)+string(base[i+1])]++
	}

	for i := 0; i < steps; i++ {
		newPairs := make(map[string]int)
		for pair, count := range pairs {
			replacement, ok := rules[pair]
			if ok {
				occurences[rune(replacement[0])] += count
				newPairs[string(pair[0])+replacement] += count
				newPairs[replacement+string(pair[1])] += count
			}
		}

		pairs = newPairs
	}

	most, least := math.MinInt, math.MaxInt
	for _, ocurrence := range occurences {
		if ocurrence > most {
			most = ocurrence
		}

		if ocurrence < least {
			least = ocurrence
		}
	}

	return most - least
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
	base := scanner.Text()

	rules := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")
		rules[parts[0]] = parts[1]
	}

	fmt.Println("Part 1: ", mostAndLeastCommon(polymer(base, rules, 0, 10)))

	fmt.Println("Part 2: ", calculateMostAndLeastCommon(base, rules, 40))
}
