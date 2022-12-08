package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func main() {
	// file, err := os.Open("01/input1.txt")
	file, err := os.Open("./input2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Println(calculateTopCalories(file))
	fmt.Println(calculateTop3Calories(file))
}

func calculateTopCalories(input io.Reader) int {
	curCalories := 0
	topCalories := 0

	newReader := bufio.NewReader(input)
	for {
		lineWithDelim, errReader := newReader.ReadString('\n')
		line := lineWithDelim[:len(lineWithDelim)-1] // remove \n

		// fmt.Printf("Reading '%v'\n", line)

		if line == "" {
			if curCalories > topCalories {
				topCalories = curCalories
			}
			curCalories = 0
			continue
		}

		value, errAtoi := strconv.Atoi(line)
		if errAtoi != nil {
			panic(errAtoi)
		}

		if errReader == io.EOF {
			curCalories += value
			if curCalories > topCalories {
				topCalories = curCalories
			}
			return topCalories
		} else if errReader != nil {
			panic(errReader)
		}

		curCalories += value
	}
}

func calculateTop3Calories(input io.Reader) int {
	curCalories := 0
	calories := []int{}

	newReader := bufio.NewReader(input)
	for {
		lineWithDelim, errReader := newReader.ReadString('\n')
		line := lineWithDelim[:len(lineWithDelim)-1] // remove \n

		// fmt.Printf("Reading '%v'\n", line)

		if line == "" {
			calories = append(calories, curCalories)
			curCalories = 0
			continue
		}

		value, errAtoi := strconv.Atoi(line)
		if errAtoi != nil {
			panic(errAtoi)
		}

		if errReader == io.EOF {
			curCalories += value
			calories = append(calories, curCalories)

			sort.Slice(calories, func(a, b int) bool { return calories[a] < calories[b] })
			fmt.Println(calories)
			return sumLastValues(calories, 3)
		} else if errReader != nil {
			panic(errReader)
		}

		curCalories += value
	}
}

func sumLastValues(arr []int, num int) int {
	sum := 0
	for i := 0; i < num; i++ {
		sum += arr[len(arr)-i-1]
	}
	return sum
}
