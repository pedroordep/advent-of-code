package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/pedroordep/advent-of-code/utils"
)

func main() {
	input := utils.GetInputFile(2025, 3)

	fmt.Println("Part 1", Part1(strings.Trim(string(input), "\n")))
	fmt.Println("Part 2", Part2(strings.Trim(string(input), "\n")))
}

func Part1(input string) int {
	banks := strings.Split(input, "\n")

	sum := 0
	for _, bank := range banks {
		first := 0
		second := 0
		for i := 0; i < len(bank); i++ {
			value, _ := strconv.Atoi(string(bank[i]))
			if i == len(bank)-1 {
				if value > second {
					second = value
				}
				continue
			}

			if value > first {
				first = value
				second = 0
			} else if value > second {
				second = value
			}
		}

		// fmt.Println("Found first", first, "and second", second)
		sum += first*10 + second
	}

	return sum
}

func Part2(input string) int64 {
	banks := strings.Split(input, "\n")

	sum := int64(0)
	for _, bank := range banks {
		result := int64(0)
		startingIndex := 0
		for i := 11; i >= 0; i-- {
			highest := int64(0)
			for j := startingIndex; j < len(bank)-i; j++ {
				value, _ := strconv.ParseInt(string(bank[j]), 10, 64)
				if value > highest {
					highest = value
					startingIndex = j + 1

					// fmt.Println("highest is", highest, "should start looking from", startingIndex, "on substring", bank[startingIndex:])
				}
			}

			result += highest * int64(math.Pow10(i))
			// fmt.Println("result for bank is", highest, "*", int64(math.Pow10(i)), "=", result)
		}

		// fmt.Println("Result is", result)
		sum += result
	}

	return sum
}
