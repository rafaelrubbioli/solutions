package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type Stack []string

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(elem string) {
	*s = append(*s, elem)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		i := len(*s) - 1
		elem := (*s)[i]
		*s = (*s)[:i]
		return elem, true
	}
}

func (s *Stack) Peek() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		return (*s)[len(*s)-1], true
	}
}

var closer = map[rune]string{
	'{': "}",
	'[': "]",
	'(': ")",
	'<': ">",
}

func calculateErrScore(lines []string) int {
	var points = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	errScore := 0
	for _, line := range lines {
		opened := NewStack()
		for _, c := range line {
			if c == '(' || c == '[' || c == '{' || c == '<' {
				opened.Push(closer[c])
			} else {
				expected, ok := opened.Peek()
				if !ok {
					continue
				}

				if expected == string(c) {
					opened.Pop()
					continue
				}

				errScore += points[c]
				break
			}
		}
	}

	return errScore
}

func completeLines(lines []string) int {
	var points = map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	completeScores := make([]int, 0)
	for _, line := range lines {
		opened := NewStack()
		corrupted := false
		for _, c := range line {
			if c == '(' || c == '[' || c == '{' || c == '<' {
				opened.Push(closer[c])
			} else {
				expected, ok := opened.Peek()
				if !ok {
					continue
				}

				if expected == string(c) {
					opened.Pop()
					continue
				}

				corrupted = true
				break
			}
		}

		if corrupted {
			continue
		}

		lineScore := 0
		for !opened.IsEmpty() {
			expected, _ := opened.Pop()
			lineScore = (lineScore * 5) + points[expected]
		}

		if lineScore > 0 {
			completeScores = append(completeScores, lineScore)
		}
	}

	sort.Ints(completeScores)
	return completeScores[len(completeScores)/2]
}

func main() {
	read, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer read.Close()

	scanner := bufio.NewScanner(read)
	scanner.Split(bufio.ScanLines)

	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Println("Part 1: ", calculateErrScore(input))
	fmt.Println("Part 2: ", completeLines(input))
}
