package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("04/input2.txt")

	split := strings.Split(string(file), "\n")

	fmt.Println("part 1:", numOfContainedTasks(split, isFullyContainedIn))
	fmt.Println("part 2:", numOfContainedTasks(split, isPartiallyContainedIn))
}

func numOfContainedTasks(input []string, compare func(int, int, int, int) bool) int {
	sum := 0

	for _, pairs := range input {
		pair := strings.Split(pairs, ",")
		firstElfTasks := strings.Split(pair[0], "-")
		secondElfTasks := strings.Split(pair[1], "-")

		firstElfStart, _ := strconv.Atoi(firstElfTasks[0])
		firstElfEnd, _ := strconv.Atoi(firstElfTasks[1])
		secondElfStart, _ := strconv.Atoi(secondElfTasks[0])
		secondElfEnd, _ := strconv.Atoi(secondElfTasks[1])

		if compare(firstElfStart, firstElfEnd, secondElfStart, secondElfEnd) {
			sum += 1
		}
	}

	return sum
}

func isFullyContainedIn(aFirst, aLast, bFirst, bLast int) bool {
	if aFirst >= bFirst && aLast <= bLast {
		return true
	} else if bFirst >= aFirst && bLast <= aLast {
		return true
	} else {
		return false
	}
}

func isPartiallyContainedIn(aFirst, aLast, bFirst, bLast int) bool {
	if aFirst > bLast || bFirst > aLast {
		return false
	} else {
		return true
	}
}
