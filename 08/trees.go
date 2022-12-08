package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	North = iota
	East
	South
	West
)

func main() {
	file, _ := os.ReadFile("08/input2.txt")

	split := strings.Split(string(file), "\n")

	fmt.Println("part 1:", visibleTrees(split))
	fmt.Println("part 2:", highestScenicScore(split))
}

func highestScenicScore(input []string) int {
	highestScore := 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			score := scoreFrom(input, North, i, j) * scoreFrom(input, East, i, j) * scoreFrom(input, South, i, j) * scoreFrom(input, West, i, j)
			if score > highestScore {
				highestScore = score
			}
		}
	}

	return highestScore
}

func visibleTrees(input []string) int {
	sum := 0

	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[i])-1; j++ {
			if visibleFrom(input, North, i, j) || visibleFrom(input, East, i, j) || visibleFrom(input, South, i, j) || visibleFrom(input, West, i, j) {
				// fmt.Println(i, j, string(input[i][j]))
				sum += 1
			}
		}
	}

	sum += len(input)*2 - 4
	sum += len(input[0]) * 2

	return sum
}

func visibleFrom(input []string, direction int, x, y int) bool {
	switch direction {
	case North:
		for i := x - 1; i >= 0; i-- {
			if input[i][y] >= input[x][y] {
				return false
			}
		}
	case South:
		for i := x + 1; i < len(input); i++ {
			if input[i][y] >= input[x][y] {
				return false
			}
		}
	case East:
		for i := y + 1; i < len(input[x]); i++ {
			if input[x][i] >= input[x][y] {
				return false
			}
		}
	case West:
		for i := y - 1; i >= 0; i-- {
			if input[x][i] >= input[x][y] {
				return false
			}
		}
	}
	return true
}

func scoreFrom(input []string, direction int, x, y int) int {
	score := 0

	switch direction {
	case North:
		for i := x - 1; i >= 0; i-- {
			score += 1
			if input[i][y] >= input[x][y] {
				return score
			}
		}
	case South:
		for i := x + 1; i < len(input); i++ {
			score += 1
			if input[i][y] >= input[x][y] {
				return score
			}
		}
	case East:
		for i := y + 1; i < len(input[x]); i++ {
			score += 1
			if input[x][i] >= input[x][y] {
				return score
			}
		}
	case West:
		for i := y - 1; i >= 0; i-- {
			score += 1
			if input[x][i] >= input[x][y] {
				return score
			}
		}
	}
	return score
}
