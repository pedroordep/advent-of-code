package main

import (
	"fmt"
	"strings"

	"github.com/pedroordep/advent-of-code/utils"
)

type Coordinate struct{ x, y int }

func main() {
	input := utils.GetInputFile(2025, 4)

	fmt.Println("Part 1", Part1(strings.Trim(string(input), "\n")))
	fmt.Println("Part 2", Part2(strings.Trim(string(input), "\n")))
}

func Part1(input string) int {
	grid := strings.Split(input, "\n")

	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '@' && getAdjacentRolls(i, j, grid) < 4 {
				// fmt.Println("Found good location at", i, j)
				sum++
			}
		}
	}
	return sum
}

func Part2(input string) int {
	grid := strings.Split(input, "\n")

	sum := 0
	for true {
		total, positions := removeRolls(grid)
		if total == 0 {
			break
		}
		sum += total

		// fmt.Println("Total is", total, "Removing positions:", positions)
		for _, pos := range positions {
			grid[pos.x] = grid[pos.x][:pos.y] + "_" + grid[pos.x][pos.y+1:]
		}
		// PrintGrid(grid)
	}

	return sum
}

func removeRolls(grid []string) (total int, positions []Coordinate) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '@' && getAdjacentRolls(i, j, grid) < 4 {
				// fmt.Println("Found good location at", i, j)
				positions = append(positions, Coordinate{i, j})
				total++
			}
		}
	}
	return total, positions
}

func getRolls(grid []string) int {
	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '@' {
				sum++
			}
		}
	}
	return sum
}

func getAdjacentRolls(i, j int, grid []string) int {
	sum := 0
	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
				if x == i && y == j {
					continue
				}
				// fmt.Println("Checking", x, y, "value", string(grid[x][y]))
				if grid[x][y] == '@' {
					sum++
				}
			}
		}
	}
	return sum
}

func PrintGrid(grid []string) {
	for _, line := range grid {
		fmt.Println(line)
	}
}
