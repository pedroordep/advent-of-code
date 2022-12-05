package main

import (
	"fmt"
	"os"
	"strings"
)

const forbiddenChars = "123456789"

func main() {
	file, _ := os.ReadFile("05/input2.txt")

	split := strings.Split(string(file), "\n")

	stack := parseInputStack(split)
	stack2 := parseInputStack(split)

	execMovesOnStack(split, stack, true)
	execMovesOnStack(split, stack2, false)

	fmt.Println("first:", printTopBoxes(stack))
	fmt.Println("second:", printTopBoxes(stack2))
}

func parseInputStack(input []string) [][]string {
	stack := [][]string{}

	for i := 0; i < len(input); i++ {
		for j := 1; j < len(input[i]); j += 4 {
			if i == 0 {
				stack = append(stack, []string{})
			}
			if strings.IndexByte(forbiddenChars, input[i][j]) != -1 {
				return stack
			} else if input[i][j] != ' ' {
				stack[(j-1)/4] = append(stack[(j-1)/4], string(input[i][j]))
			}
		}
	}

	return stack
}

func printTopBoxes(stacks [][]string) string {
	result := ""

	for _, stack := range stacks {
		result += stack[0]
	}

	return result
}

func execMovesOnStack(input []string, stack [][]string, singleBoxOp bool) {
	for _, line := range input {
		if len(line) < 4 || line[:4] != "move" {
			continue
		}
		stacksToMove := 0
		fromStack := 0
		toStack := 0

		fmt.Sscanf(line, "move %d from %d to %d", &stacksToMove, &fromStack, &toStack)
		// fmt.Println(stacksToMove, fromStack-1, toStack-1)

		boxesToMove := make([]string, stacksToMove)
		copy(boxesToMove, stack[fromStack-1][:stacksToMove])
		// fmt.Printf("boxesToMove: %v, len %d, cap: %d\n", boxesToMove, len(boxesToMove), cap(boxesToMove))

		if singleBoxOp {
			reverseStack(boxesToMove)
		}

		stack[toStack-1] = append(boxesToMove, stack[toStack-1]...)
		stack[fromStack-1] = stack[fromStack-1][stacksToMove:]

		// fmt.Println(stack)
		// fmt.Println("-------")
	}
}

func reverseStack(input []string) {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
}
