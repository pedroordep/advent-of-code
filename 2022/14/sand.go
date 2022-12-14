package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func main() {
	file, _ := os.ReadFile("./input2.txt")
	lines := strings.Split(string(file), "\n")

	sandMap := parsePointSet(lines)
	fmt.Println("part 1:", part1(sandMap))

	sandMap = parsePointSet(lines)
	fmt.Println("part 2:", part2(sandMap))
}

func part1(sandMap map[image.Point]struct{}) int {
	highestY := gethighestY(sandMap)

	start := image.Point{500, 0}
	cur := start
	stalePoints := 0
	moves := []image.Point{{0, 1}, {-1, 1}, {1, 1}}

	for cur.Y < highestY {
		stale := true
		for _, move := range moves {
			_, exists := sandMap[cur.Add(move)]
			if !exists {
				stale = false
				cur = cur.Add(move)
				break
			}
		}
		if stale {
			sandMap[cur] = struct{}{}
			cur = start
			stalePoints++
		}
	}
	// fmt.Println(sandMap)
	return stalePoints
}

func part2(sandMap map[image.Point]struct{}) int {
	highestY := gethighestY(sandMap)
	ground := highestY + 2

	start := image.Point{500, 0}
	cur := start
	stalePoints := 0
	moves := []image.Point{{0, 1}, {-1, 1}, {1, 1}}

	for {
		stale := true
		for _, move := range moves {
			_, exists := sandMap[cur.Add(move)]
			if !exists && cur.Add(move).Y < ground {
				stale = false
				cur = cur.Add(move)
				break
			}
		}
		if stale {
			sandMap[cur] = struct{}{}
			stalePoints++
			cur = start
		}
		_, exists := sandMap[start]
		if exists {
			break
		}
	}
	// fmt.Println(sandMap, len(sandMap))
	return stalePoints
}

func parsePointSet(input []string) map[image.Point]struct{} {
	pointSet := map[image.Point]struct{}{}
	for _, line := range input {
		points := strings.Split(line, " -> ")

		for i := 1; i < len(points); i++ {
			p1 := image.Point{}
			p2 := image.Point{}

			fmt.Sscanf(points[i-1], "%d,%d", &p1.X, &p1.Y)
			fmt.Sscanf(points[i], "%d,%d", &p2.X, &p2.Y)

			if p1.X > p2.X {
				for x := p2.X; x <= p1.X; x++ {
					pointSet[image.Point{x, p1.Y}] = struct{}{}
				}
			} else if p1.X < p2.X {
				for x := p1.X; x <= p2.X; x++ {
					pointSet[image.Point{x, p1.Y}] = struct{}{}
				}
			} else if p1.Y > p2.Y {
				for y := p2.Y; y <= p1.Y; y++ {
					pointSet[image.Point{p1.X, y}] = struct{}{}
				}
			} else if p1.Y < p2.Y {
				for y := p1.Y; y <= p2.Y; y++ {
					pointSet[image.Point{p2.X, y}] = struct{}{}
				}
			}
		}
	}

	return pointSet
}

func gethighestY(sandMap map[image.Point]struct{}) int {
	highestY := 0
	for key := range sandMap {
		if key.Y > highestY {
			highestY = key.Y
		}
	}
	return highestY
}
