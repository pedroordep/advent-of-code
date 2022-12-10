package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.ReadFile("./input2.txt")
	instructions := strings.Split(string(file), "\n")

	score := []int{1}
	screen := [240]string{}
	for _, instruction := range instructions {
		if instruction == "noop" {
			score = append(score, score[len(score)-1])
		} else {
			val := 0
			fmt.Sscanf(instruction, "addx %d", &val)

			score = append(score, score[len(score)-1])
			score = append(score, score[len(score)-1]+val)
		}
	}
	for i := 0; i < len(score)-1; i++ {
		if score[i]-1 <= i%40 && score[i]+1 >= i%40 {
			screen[i] = "#"
		} else {
			screen[i] = " "
		}
	}
	fmt.Println("part 1", score[19]*20+score[59]*60+score[99]*100+score[139]*140+score[179]*180+score[219]*220)
	fmt.Println("part 2")
	printCRT(screen)
}

func printCRT(input [240]string) {
	for i, v := range input {
		if i%40 == 0 {
			fmt.Print("\n")
		}
		fmt.Print(v)
	}
	fmt.Print("\n")
}
