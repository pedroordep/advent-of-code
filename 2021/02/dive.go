package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.ReadFile("./input2.txt")
	lines := strings.Split(string(file), "\n")

	fmt.Println("part1", part1(lines))
	fmt.Println("part2", part2(lines))
}

func part1(input []string) int {
	posH, posD := 0, 0
	command, val := "", 0

	for _, line := range input {
		fmt.Sscanf(line, "%s %d", &command, &val)
		switch command {
		case "forward":
			posH += val
		case "down":
			posD -= val
		case "up":
			posD += val
		}
	}

	return posH * -posD
}

func part2(input []string) int {
	posH, posD, aim := 0, 0, 0
	command, val := "", 0

	for _, line := range input {
		fmt.Sscanf(line, "%s %d", &command, &val)
		switch command {
		case "forward":
			posH += val
			posD += aim * val
		case "down":
			aim += val
		case "up":
			aim -= val
		}
	}

	return posH * posD
}
