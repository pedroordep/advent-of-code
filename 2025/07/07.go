package main

import (
	"fmt"
	"strings"

	"github.com/pedroordep/advent-of-code/utils"
)

var timelinesFrom = map[string]int{}

func main() {
	input := utils.GetInputFile(2025, 7)

	fmt.Println("Part 1", Part1(strings.Trim(string(input), "\n")))
	fmt.Println("Part 2", Part2(strings.Trim(string(input), "\n")))
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")

	count := 0
	for nLine := range lines {
		for i := 0; i < len(lines[nLine]); i++ {
			if lines[nLine][i] == 'S' {
				lines[nLine+1] = replaceAtIndex(lines[nLine+1], '|', i)
			} else if lines[nLine][i] == '^' && lines[nLine-1][i] == '|' {
				lines[nLine] = replaceAtIndex(lines[nLine], '|', i-1)
				lines[nLine] = replaceAtIndex(lines[nLine], '|', i+1)
				count++
			} else if nLine > 0 && lines[nLine][i] == '.' && lines[nLine-1][i] == '|' {
				lines[nLine] = replaceAtIndex(lines[nLine], '|', i)
			}
		}
	}

	// fmt.Println(strings.Join(lines, "\n"))

	return count
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")

	return countTimelines(lines, 1, strings.Index(lines[0], "S")) + 1
}

func countTimelines(lines []string, y, x int) int {
	if timelinesFrom[fmt.Sprintf("%d-%d", x, y)] != 0 {
		return timelinesFrom[fmt.Sprintf("%d-%d", x, y)]
	}
	if y >= len(lines)-1 {
		return 0
	}

	if lines[y+1][x] == '^' {
		// fmt.Println("Found branch right at", y+1, x)
		return 1 + countTimelines(lines, y+1, x-1) + countTimelines(lines, y+1, x+1)
	} else {
		result := countTimelines(lines, y+1, x)
		timelinesFrom[fmt.Sprintf("%d-%d", x, y)] = result
		return result
	}
}

func replaceAtIndex(str string, replacement rune, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
}
