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
	// fmt.Println("part 1:", part1(instructions))
	fmt.Println("part 1:", part2(instructions, rope))
	rope = []Point{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}
	fmt.Println("part 2:", part2(instructions, rope))
}

// func part1(instructions []string) int {
// 	visited := map[Point]bool{{0, 0}: true}
// 	tailP := Point{0, 0}
// 	headP := Point{0, 0}
// 	for _, instruction := range instructions {
// 		dir, amount := "", 0
// 		fmt.Sscanf(instruction, "%s %d", &dir, &amount)

// 		for i := 0; i < amount; i++ {
// 			// fmt.Println(headP, tailP)
// 			switch dir {
// 			case "U":
// 				moveInDirection(&headP, "U")
// 				if headP.y-tailP.y > 1 && headP.x-tailP.x == 0 {
// 					moveInDirection(&tailP, "U")
// 					visited[tailP] = true
// 				} else if headP.y-tailP.y > 1 && headP.x-tailP.x > 0 {
// 					moveInDirection(&tailP, "UR")
// 					visited[tailP] = true
// 				} else if headP.y-tailP.y > 1 && headP.x-tailP.x < 0 {
// 					moveInDirection(&tailP, "UL")
// 					visited[tailP] = true
// 				}
// 			case "D":
// 				moveInDirection(&headP, "D")
// 				if headP.y-tailP.y < -1 && headP.x-tailP.x == 0 {
// 					moveInDirection(&tailP, "D")
// 					visited[tailP] = true
// 				} else if headP.y-tailP.y < -1 && headP.x-tailP.x > 0 {
// 					moveInDirection(&tailP, "DR")
// 					visited[tailP] = true
// 				} else if headP.y-tailP.y < -1 && headP.x-tailP.x < 0 {
// 					moveInDirection(&tailP, "DL")
// 					visited[tailP] = true
// 				}
// 			case "R":
// 				moveInDirection(&headP, "R")
// 				if headP.x-tailP.x > 1 && headP.y-tailP.y == 0 {
// 					moveInDirection(&tailP, "R")
// 					visited[tailP] = true
// 				} else if headP.x-tailP.x > 1 && headP.y-tailP.y > 0 {
// 					moveInDirection(&tailP, "UR")
// 					visited[tailP] = true
// 				} else if headP.x-tailP.x > 1 && headP.y-tailP.y < 0 {
// 					moveInDirection(&tailP, "DR")
// 					visited[tailP] = true
// 				}
// 			case "L":
// 				moveInDirection(&headP, "L")
// 				if headP.x-tailP.x < -1 && headP.y-tailP.y == 0 {
// 					moveInDirection(&tailP, "L")
// 					visited[tailP] = true
// 				} else if headP.x-tailP.x < -1 && headP.y-tailP.y > 0 {
// 					moveInDirection(&tailP, "UL")
// 					visited[tailP] = true
// 				} else if headP.x-tailP.x < -1 && headP.y-tailP.y < 0 {
// 					moveInDirection(&tailP, "DL")
// 					visited[tailP] = true
// 				}
// 			}
// 		}
// 	}
// 	return len(visited)
// }

func part2(instructions []string, rope []Point) int {
	visited := map[Point]bool{{0, 0}: true}
	for _, instruction := range instructions {
		dir, amount := "", 0
		fmt.Sscanf(instruction, "%s %d", &dir, &amount)

		for i := 0; i < amount; i++ {
			moveInDirection(&rope[0], dir)
			// fmt.Println(headP, tailP)
			for index := 1; index < len(rope); index++ {
				checkAnMoveTailInDirection(&rope[index-1], &rope[index], dir)
			}
			visited[rope[len(rope)-1]] = true
		}
	}
	return len(visited)
}

func checkAnMoveTailInDirection(headP *Point, tailP *Point, dir string) {
	switch dir {
	case "U":
		// moveInDirection(headP, "U")
		if headP.y-tailP.y > 1 && headP.x-tailP.x == 0 {
			moveInDirection(tailP, "U")
		} else if headP.y-tailP.y > 1 && headP.x-tailP.x > 0 {
			moveInDirection(tailP, "UR")
		} else if headP.y-tailP.y > 1 && headP.x-tailP.x < 0 {
			moveInDirection(tailP, "UL")
		}
	case "D":
		// moveInDirection(headP, "D")
		if headP.y-tailP.y < -1 && headP.x-tailP.x == 0 {
			moveInDirection(tailP, "D")
		} else if headP.y-tailP.y < -1 && headP.x-tailP.x > 0 {
			moveInDirection(tailP, "DR")
		} else if headP.y-tailP.y < -1 && headP.x-tailP.x < 0 {
			moveInDirection(tailP, "DL")
		}
	case "R":
		// moveInDirection(headP, "R")
		if headP.x-tailP.x > 1 && headP.y-tailP.y == 0 {
			moveInDirection(tailP, "R")
		} else if headP.x-tailP.x > 1 && headP.y-tailP.y > 0 {
			moveInDirection(tailP, "UR")
		} else if headP.x-tailP.x > 1 && headP.y-tailP.y < 0 {
			moveInDirection(tailP, "DR")
		}
	case "L":
		// moveInDirection(headP, "L")
		if headP.x-tailP.x < -1 && headP.y-tailP.y == 0 {
			moveInDirection(tailP, "L")
		} else if headP.x-tailP.x < -1 && headP.y-tailP.y > 0 {
			moveInDirection(tailP, "UL")
		} else if headP.x-tailP.x < -1 && headP.y-tailP.y < 0 {
			moveInDirection(tailP, "DL")
		}
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
