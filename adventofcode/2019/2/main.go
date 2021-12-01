package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	read, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(read), ",")
	commands := make([]int, 0)
	for _, i := range input {
		num, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}

		commands = append(commands, num)
	}

	commands[1] = 12
	commands[2] = 2

	for i := 0; i < len(commands); i += 4 {
		command := commands[i]
		if command == 99 {
			fmt.Println(commands[0])
			return
		}

		first := commands[commands[i+1]]
		second := commands[commands[i+2]]

		if command == 1 {
			commands[commands[i+3]] = first + second
		} else if command == 2 {
			commands[commands[i+3]] = first * second
		}
	}
}
