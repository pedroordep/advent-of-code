package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pedroordep/advent-of-code/utils"
)

func main() {
	input := utils.GetInputFile(2025, 6)

	fmt.Println("Part 1", Part1(strings.Trim(string(input), "\n")))
	fmt.Println("Part 2", Part2(strings.Trim(string(input), "\n")))
}

func Part1(input string) int64 {
	parts := strings.Split(input, "\n")

	matrix := [][]string{}
	for _, line := range parts {
		fields := strings.Fields(line)
		row := []string{}
		for _, field := range fields {
			row = append(row, strings.ReplaceAll(field, " ", ""))
		}
		matrix = append(matrix, row)
	}

	transform := [][]string{}
	for i := 0; i < len(matrix[0]); i++ {
		column := []string{}
		for j := 0; j < len(matrix); j++ {
			column = append(column, matrix[j][i])
		}
		transform = append(transform, column)
	}

	return calculateMatrix(transform)
}

func Part2(input string) int64 {
	matrixXLength := strings.Index(input, "\n") + 1
	matrixYLength := strings.Count(input, "\n") + 1

	// fmt.Println("Matrix dimensions:", matrixXLength, "x", matrixYLength)

	matrix := [][]string{}
	matrixRow := make([]string, 0)
	// Remove \n from matrixXLength
	for x := matrixXLength - 2; x >= 0; x-- {
		numberString := ""
		for y := 0; y < matrixYLength-1; y++ {
			// fmt.Println("Building numberString", string(input[x+(matrixXLength*y)]))
			numberString += string(input[x+(matrixXLength*y)])
		}
		// fmt.Println("Adding number", strings.TrimSpace(numberString), "to row", matrixRow)
		matrixRow = append(matrixRow, strings.TrimSpace(numberString))
		if input[x+(matrixXLength*(matrixYLength-1))] == '*' || input[x+(matrixXLength*(matrixYLength-1))] == '+' {
			// fmt.Println("Adding operator", string(input[x+(matrixXLength*(matrixYLength-1))]), "to row")
			matrixRow = append(matrixRow, string(input[x+(matrixXLength*(matrixYLength-1))]))
			// fmt.Println("Adding slice", matrixRow, "with length", len(matrixRow), "to matrix", matrix)
			matrix = append(matrix, matrixRow)
			matrixRow = make([]string, 0)
			x--
		}
	}

	return calculateMatrix(matrix)
}

func calculateMatrix(matrix [][]string) (sum int64) {
	for i := 0; i < len(matrix); i++ {
		columnSum := int64(0)
		switch matrix[i][len(matrix[i])-1] {
		case "+":
			for j := 0; j < len(matrix[i])-1; j++ {
				val, _ := strconv.ParseInt(matrix[i][j], 10, 64)
				columnSum += val
			}
		case "*":
			columnSum = 1
			for j := 0; j < len(matrix[i])-1; j++ {
				val, _ := strconv.ParseInt(matrix[i][j], 10, 64)
				columnSum *= val
			}
		default:
			panic(fmt.Sprintf("unknown operator: %s", matrix[i][len(matrix[i])-1]))
		}
		// fmt.Println("Column", matrix[i], "sum:", columnSum)
		sum += columnSum
	}

	return sum
}
