package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/pedroordep/advent-of-code/utils"
)

func main() {
	input := utils.GetInputFile(2025, 2)

	fmt.Println("Part 1", Part1(strings.Trim(string(input), "\n")))
	fmt.Println("Part 2", Part2(strings.Trim(string(input), "\n")))
}

func Part1(input string) int {
	idRanges := strings.Split(input, ",")

	sum := 0
	for _, idRange := range idRanges {
		split := strings.Split(idRange, "-")

		start, _ := strconv.Atoi(split[0])
		end, _ := strconv.Atoi(split[1])
		for i := start; i <= end; i++ {
			value := strconv.Itoa(i)

			if value[:len(value)/2] == value[len(value)/2:] {
				// fmt.Println("Found", value)
				sum += i
			}
		}
	}

	return sum
}

var divisorsMap map[string][]int

func Part2(input string) int {
	idRanges := strings.Split(input, ",")

	sum := 0
	for _, idRange := range idRanges {
		split := strings.Split(idRange, "-")

		start, _ := strconv.Atoi(split[0])
		end, _ := strconv.Atoi(split[1])
		for current := start; current <= end; current++ {
			value := strconv.Itoa(current)

			divisors, exists := divisorsMap[value]
			if !exists {
				divisors = calcDivisors(len(value))
			}
			// fmt.Println("divisors for", value, "are", divisors)

			for _, divisor := range divisors {

				// fmt.Println("Checking for", value[:divisor], "ocurrences in", value)
				count := strings.Count(value, value[:divisor])

				if count == len(value)/divisor {
					sum += current

					// fmt.Println("Found", value)
					break
				}
			}
		}
	}

	return sum
}

func calcDivisors(n int) []int {
	var divs []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			divs = append(divs, i)
			if i != n/i {
				divs = append(divs, n/i)
			}
		}
	}
	sort.Ints(divs)
	return divs[:len(divs)-1]
}
