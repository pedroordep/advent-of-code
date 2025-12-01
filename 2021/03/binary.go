package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	file, _ := os.ReadFile("./input2.txt")
	lines := strings.Split(string(file), "\n")

	fmt.Println("part1:", part1(lines))
	fmt.Println("part2:", part2(lines))
}

func part1(input []string) int {
	gamma, epsilon := 0, 0
	numOfOnes := make([]int, len(input[0]))

	for _, line := range input {
		for i := 0; i < len(line); i++ {
			if line[i] == '1' {
				numOfOnes[i] += 1
			}
		}
	}

	for i, val := range numOfOnes {
		binaryValue := int(math.Pow(2, float64(len(numOfOnes)-i-1)))
		if val >= len(input)/2 {
			gamma += binaryValue
		} else {
			epsilon += binaryValue
		}
	}

	return gamma * epsilon
}

func part2(input []string) int {
	oxigen := slices.Clone(input)
	co2 := slices.Clone(input)

	for i := 0; i < len(input[0]); i++ {
		// fmt.Println("oxigen", oxigen)
		// fmt.Println("co2", oxigen)
		if len(oxigen) > 1 {
			removeElsAtIndex(&oxigen, i, true)
		}
		if len(co2) > 1 {
			removeElsAtIndex(&co2, i, false)
		}
	}

	// fmt.Println(oxigen, co2)

	return getBinaryValue(oxigen[0]) * getBinaryValue(co2[0])
}

func removeElsAtIndex(input *[]string, index int, maintainHighest bool) {
	onesSlice := []string{}
	zerosSlice := []string{}
	for i := 0; i < len(*input); i++ {
		if (*input)[i][index] == '1' {
			onesSlice = append(onesSlice, (*input)[i])
		} else {
			zerosSlice = append(zerosSlice, (*input)[i])
		}
	}
	if maintainHighest {
		if len(onesSlice) >= len(zerosSlice) {
			*input = onesSlice
		} else {
			*input = zerosSlice
		}
	} else {
		if len(onesSlice) >= len(zerosSlice) {
			*input = zerosSlice
		} else {
			*input = onesSlice
		}
	}
}

func getBinaryValue(input string) int {
	sum := 0
	for i := range input {
		if input[i] == '1' {
			sum += int(math.Pow(2, float64(len(input)-i-1)))
		}
	}
	return sum
}
