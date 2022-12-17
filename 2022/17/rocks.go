package main

import (
	"fmt"
	"image"
	"os"
)

type Rock []image.Point
type Cave [10000][7]int

func (r Rock) Move(p image.Point) {
	for i := 0; i < len(r); i++ {
		r[i] = r[i].Add(p)
	}
}

func (r Rock) hasClonflict(c Cave) bool {
	for i := 0; i < len(r); i++ {
		if r[i].X < 0 || r[i].X > len(c[0])-1 {
			return true
		}
		if c[r[i].Y][r[i].X] > 0 {
			return true
		}
	}
	return false
}

func (c *Cave) Update(r Rock) {
	for i := 0; i < len(r); i++ {
		c[r[i].Y][r[i].X] = 1
	}
}

func (c *Cave) GetHighestY(r Rock) int {
	highestY := r[0].Y
	for i := highestY; i < len(c); i++ {
		foundHigher := false
		for x := 0; x < len(c[i]); x++ {
			if c[i][x] > 0 {
				highestY = i
				foundHigher = true
				break
			}
		}
		if !foundHigher {
			break
		}
	}
	return highestY
}

func (c Cave) String() string {
	highestY := c.GetHighestY([]image.Point{{0, 0}})
	str := ""
	for i := highestY; i >= 0; i-- {
		str += "|"
		for x := 0; x < len(c[i]); x++ {
			switch c[i][x] {
			case 0:
				str += "."
			case 1:
				str += "#"
			case 2:
				str += "-"
			}
		}
		str += "|\n"
	}
	return str
}

func main() {
	file, _ := os.ReadFile("./input2.txt")
	moves := string(file)
	rocks := []Rock{
		[]image.Point{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
		[]image.Point{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}},
		[]image.Point{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}},
		[]image.Point{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		[]image.Point{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
	}

	fmt.Println("part 1:", part1(rocks, moves, 2022))
}

func part1(rocks []Rock, moves string, n int) int {
	cave := Cave{}
	for i := 0; i < 7; i++ {
		cave[0][i] = 2
	}
	highestY := 0

	j := 0
	for i := 0; i < n; i++ {
		rock := make(Rock, len(rocks[i%len(rocks)]))
		copy(rock, rocks[i%len(rocks)])

		rock.Move(image.Point{2, highestY + 4})
		for {
			// air jet
			// fmt.Println("moving", string(moves[j%len(moves)]), rock)
			switch moves[j%len(moves)] {
			case '<':
				rock.Move(image.Point{-1, 0})
				if rock.hasClonflict(cave) {
					rock.Move(image.Point{1, 0})
				}
			case '>':
				rock.Move(image.Point{1, 0})
				if rock.hasClonflict(cave) {
					rock.Move(image.Point{-1, 0})
				}
			}

			j++

			// fall
			// fmt.Println("dropping", rock)
			rock.Move(image.Point{0, -1})
			if rock.hasClonflict(cave) {
				rock.Move(image.Point{0, 1})
				cave.Update(rock)
				highestY = cave.GetHighestY(rock)
				// fmt.Println("rested  ", rock)
				break
			}
		}
		// fmt.Println(cave)
	}
	return highestY
}
