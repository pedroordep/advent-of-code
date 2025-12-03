package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pedroordep/advent-of-code/utils"
)

var debug = false

func main() {
	input := utils.GetInputFile(2025, 1)

	fmt.Println("Part 1", Part1(strings.Trim(string(input), "\n")))
	fmt.Println("Part 2", Part2(strings.Trim(string(input), "\n")))
}

func Part1(input string) int {
	dial := 50
	zeroOccurences := 0

	for _, line := range strings.Split(input, "\n") {
		rotation := string(line[0])
		rotationAmount, _ := strconv.Atoi(line[1:])

		if rotation == "L" {
			dial = (dial - rotationAmount + 100) % 100
		} else {
			dial = (dial + rotationAmount) % 100
		}

		if dial == 0 {
			zeroOccurences += 1
		}
	}

	return zeroOccurences
}

func Part2(input string) int {
	dial := 50
	zeroOccurences := 0

	for _, line := range strings.Split(input, "\n") {
		rotation := string(line[0])
		rotationAmount, _ := strconv.Atoi(line[1:])

		// fmt.Println("---Rotating", line)

		zeroOccurences += rotationAmount / 100
		// fmt.Println("Adding rotations", rotationAmount/100)

		rotationAmount = rotationAmount % 100

		if rotation == "L" {
			if dial != 0 && dial-rotationAmount <= 0 {
				// fmt.Println("dial-rotationAmount <= 0, adding 1 rotation")
				zeroOccurences += 1
			}
			dial = ((dial - rotationAmount) + 100) % 100
		} else {
			if dial != 0 && dial+rotationAmount >= 100 {
				// fmt.Println("dial+rotationAmount >= 100, adding 1 rotation")
				zeroOccurences += 1
			}
			dial = (dial + rotationAmount) % 100
		}

		// fmt.Println("Dial is", dial)
		// fmt.Println("Zero occurences", zeroOccurences)
	}

	return zeroOccurences
}
