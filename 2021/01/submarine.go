package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("./input2.txt")
	lines := strings.Split(string(file), "\n")

	measurements := parseInput(lines)

	fmt.Println("part1:", part1(measurements))
	fmt.Println("part2:", part2(measurements))
}

func parseInput(input []string) []int {
	result := make([]int, len(input))
	for i, line := range input {
		val, _ := strconv.Atoi(line)
		result[i] = val
	}
	return result
}

func part1(input []int) int {
	sum := 0
	highest := input[0]

	for _, val := range input[1:] {
		if val > highest {
			sum += 1
		}
		highest = val
	}

	return sum
}

func part2(input []int) int {
	sum := 0
	highest := input[0] + input[1] + input[2]

	for i := 1; i < len(input)-2; i++ {
		val := input[i] + input[i+1] + input[i+2]
		if val > highest {
			sum += 1
		}
		highest = val
	}

	return sum
}
