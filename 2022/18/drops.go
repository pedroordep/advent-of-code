package main

import (
	"fmt"
	"os"
	"strings"
)

type Cube struct {
	x, y, z int
}

type CubeSet map[Cube]struct{}
type Drop CubeSet

var moves []Cube = []Cube{{-1, 0, 0}, {1, 0, 0}, {0, -1, 0}, {0, 1, 0}, {0, 0, -1}, {0, 0, 1}}
var drop Drop
var fill CubeSet

func main() {
	file, _ := os.ReadFile("./input2.txt")
	split := strings.Split(string(file), "\n")

	drop = Drop{}
	for _, line := range split {
		cube := Cube{}
		fmt.Sscanf(line, "%d,%d,%d", &cube.x, &cube.y, &cube.z)
		drop[cube] = struct{}{}
	}

	fill = CubeSet{}
	minCube, maxCube := drop.MinMaxCubes()
	floodFill(minCube, minCube, maxCube)

	fmt.Println("part 1:", getExposedSum(drop, false))
	fmt.Println("part 2:", getExposedSum(drop, true))
}

func getExposedSum(drop Drop, checkFill bool) int {
	sum := 0
	for cube := range drop {
		for _, move := range moves {
			_, exists := drop[cube.Add(move)]
			if checkFill {
				_, existsInFill := fill[cube.Add(move)]
				if !exists && existsInFill {
					sum += 1
				}
			} else {
				if !exists {
					sum += 1
				}
			}
		}
	}

	return sum
}

func floodFill(cur, minCube, maxCube Cube) {
	if cur.x < minCube.x || cur.y < minCube.y || cur.z < minCube.z || cur.x > maxCube.x || cur.y > maxCube.y || cur.z > maxCube.z {
		return
	}
	_, exists := drop[cur]
	if exists {
		return
	}
	_, exists = fill[cur]
	if exists {
		return
	}
	fill[cur] = struct{}{}
	for _, move := range moves {
		floodFill(cur.Add(move), minCube, maxCube)
	}
}

func (c Cube) Add(cube Cube) Cube {
	c.x += cube.x
	c.y += cube.y
	c.z += cube.z
	return c
}

func (d Drop) MinMaxCubes() (Cube, Cube) {
	var min, max Cube
	for cube := range d {
		if (Cube{}) == min {
			min = cube
		}
		if (Cube{}) == max {
			max = cube
		}

		if cube.x > max.x {
			max.x = cube.x
		}
		if cube.y > max.y {
			max.y = cube.y
		}
		if cube.z > max.z {
			max.z = cube.z
		}

		if cube.x < min.x {
			min.x = cube.x
		}
		if cube.y < min.y {
			min.y = cube.y
		}
		if cube.z < min.z {
			min.z = cube.z
		}
	}

	return min.Add(Cube{-1, -1, -1}), max.Add(Cube{1, 1, 1})
}
