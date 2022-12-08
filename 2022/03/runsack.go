package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("./input2.txt")
	if err != nil {
		panic(err)
	}

	split := strings.Split(string(file), "\n")

	fmt.Println("first:", firstPriority(split))
	fmt.Println("second:", secondPriority(split))
}

func firstPriority(input []string) int {
	sum := 0

	for _, line := range input {
		itemInFirstCompartment := map[byte]bool{}
		for _, item := range []byte(line[:len(line)/2]) {
			itemInFirstCompartment[item] = true
		}
		for _, item := range []byte(line[len(line)/2:]) {
			if itemInFirstCompartment[item] {
				sum += GetPriorityByByte(item)
				break
			}
		}
	}

	return sum
}

func secondPriority(input []string) int {
	sum := 0

	for i := 0; i < len(input); i += 3 {
		itemQuantInDiffRunsacks := map[byte]int{}
	each:
		for j := 0; j < 3; j++ {
			itemChecked := map[byte]bool{}
			for _, item := range []byte(input[i+j]) {
				if itemChecked[item] {
					continue
				}
				itemChecked[item] = true
				itemQuantInDiffRunsacks[item] += 1
				if itemQuantInDiffRunsacks[item] == 3 {
					sum += GetPriorityByByte(item)
					break each
				}
			}
		}
	}

	return sum
}

func GetPriorityByByte(input byte) int {
	val := int(input) - 96
	if val > 0 {
		return val
	}
	return val + 29*2
}
