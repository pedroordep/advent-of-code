package main

import (
	"fmt"
	"os"
	"strings"
)

type Cube struct {
	x, y, z int
}

type Drop map[Cube]struct{}

func main() {
	file, _ := os.ReadFile("./input1.txt")
	split := strings.Split(string(file), "\n")

	drop := Drop{}
	for _, line := range split {
		cube := Cube{}
		fmt.Sscanf(line, "%d,%d,%d", &cube.x, &cube.y, &cube.z)
		drop[cube] = struct{}{}
	}

	moves := []Cube{{-1, 0, 0}, {1, 0, 0}, {0, -1, 0}, {0, 1, 0}, {0, 0, -1}, {0, 0, 1}}

	sum := 0
	for cube := range drop {
		for _, move := range moves {
			_, exists := drop[cube.Add(move)]
			if !exists {
				sum += 1
			}
		}
	}

	fmt.Println(sum)
}

func (c Cube) Add(cube Cube) Cube {
	c.x += cube.x
	c.y += cube.y
	c.z += cube.z
	return c
}
