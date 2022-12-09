package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

func main() {
	file, _ := os.ReadFile("./input2.txt")
	instructions := strings.Split(string(file), "\n")

	rope := []Point{{0, 0}, {0, 0}}
	fmt.Println("part 1:", crossBridge(instructions, rope))
	rope = []Point{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	fmt.Println("part 2:", crossBridge(instructions, rope))
}

func crossBridge(instructions []string, rope []Point) int {
	visited := map[Point]bool{{0, 0}: true}
	for _, instruction := range instructions {
		dir, amount := "", 0
		fmt.Sscanf(instruction, "%s %d", &dir, &amount)

		for i := 0; i < amount; i++ {
			moveInDirection(&rope[0], dir)
			// fmt.Println(headP, tailP)
			for index := 1; index < len(rope); index++ {
				checkAndMoveTailInDirection(&rope[index-1], &rope[index])
			}
			visited[rope[len(rope)-1]] = true
			// fmt.Println(dir, rope)
		}
	}
	return len(visited)
}

func checkAndMoveTailInDirection(headP *Point, tailP *Point) {
	if headP.x-tailP.x == 0 && headP.y-tailP.y == 2 {
		moveInDirection(tailP, "U")
	} else if headP.x-tailP.x == 1 && headP.y-tailP.y == 2 {
		moveInDirection(tailP, "UR")
	} else if headP.x-tailP.x == 2 && headP.y-tailP.y == 2 {
		moveInDirection(tailP, "UR")
	} else if headP.x-tailP.x == 2 && headP.y-tailP.y == 1 {
		moveInDirection(tailP, "UR")
	} else if headP.x-tailP.x == 2 && headP.y-tailP.y == 0 {
		moveInDirection(tailP, "R")
	} else if headP.x-tailP.x == 2 && headP.y-tailP.y == -1 {
		moveInDirection(tailP, "DR")
	} else if headP.x-tailP.x == 2 && headP.y-tailP.y == -2 {
		moveInDirection(tailP, "DR")
	} else if headP.x-tailP.x == 1 && headP.y-tailP.y == -2 {
		moveInDirection(tailP, "DR")
	} else if headP.x-tailP.x == 0 && headP.y-tailP.y == -2 {
		moveInDirection(tailP, "D")
	} else if headP.x-tailP.x == -1 && headP.y-tailP.y == -2 {
		moveInDirection(tailP, "DL")
	} else if headP.x-tailP.x == -2 && headP.y-tailP.y == -2 {
		moveInDirection(tailP, "DL")
	} else if headP.x-tailP.x == -2 && headP.y-tailP.y == -1 {
		moveInDirection(tailP, "DL")
	} else if headP.x-tailP.x == -2 && headP.y-tailP.y == 0 {
		moveInDirection(tailP, "L")
	} else if headP.x-tailP.x == -2 && headP.y-tailP.y == 1 {
		moveInDirection(tailP, "UL")
	} else if headP.x-tailP.x == -2 && headP.y-tailP.y == 2 {
		moveInDirection(tailP, "UL")
	} else if headP.x-tailP.x == -1 && headP.y-tailP.y == 2 {
		moveInDirection(tailP, "UL")
	}
}

func moveInDirection(p *Point, d string) {
	switch d {
	case "U":
		p.y++
	case "D":
		p.y--
	case "R":
		p.x++
	case "L":
		p.x--
	case "UR":
		p.x++
		p.y++
	case "DR":
		p.x++
		p.y--
	case "DL":
		p.x--
		p.y--
	case "UL":
		p.x--
		p.y++
	}
}
