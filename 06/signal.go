package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.ReadFile("06/input2.txt")

	fmt.Println("part 1:", checkNonRepeatedCharIndex(file, 4))
	fmt.Println("part 2:", checkNonRepeatedCharIndex(file, 14))
}

func checkNonRepeatedCharIndex(input []byte, size int) int {
	indexOfByte := map[byte]int{}
	numUniqueChars := 0

	for index := 0; index < len(input); index++ {
		oldIndex, exists := indexOfByte[input[index]]
		// fmt.Printf("[%d] checking val %v, map %v\n", index, input[index], indexOfByte)
		if exists {
			// fmt.Printf("[%d] val %v already exists in %v, reseting map\n", index, input[index], indexOfByte)
			numUniqueChars = 0
			indexOfByte = map[byte]int{}
			index = oldIndex
			continue
		}
		indexOfByte[input[index]] = index
		numUniqueChars += 1
		if numUniqueChars >= size {
			return index + 1
		}
	}
	return 0
}
