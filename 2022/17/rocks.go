package main

import (
	"fmt"
	"image"
	"os"
	"strconv"
)

type Rock []image.Point
type Cave [7]int
type CaveHistory map[Cave]struct{}

// var history CaveHistory = CaveHistory{}

func (r Rock) Move(p image.Point) {
	for i := 0; i < len(r); i++ {
		r[i] = r[i].Add(p)
	}
}

func (r Rock) hasClonflict(c Cave) bool {
	for i := 0; i < len(r); i++ {
		if r[i].X < 0 || r[i].X > len(c)-1 {
			return true
		}
		if c[r[i].X] >= r[i].Y {
			return true
		}
	}
	return false
}

func (c *Cave) Update(r Rock) {
	for i := 0; i < len(r); i++ {
		c[r[i].X] = r[i].Y
	}
}

func (c *Cave) GetHighestY(r Rock) int {
	highestY := 0
	for i := 0; i < len(c); i++ {
		if c[i] > highestY {
			highestY = c[i]
		}
	}
	return highestY
}

func (c Cave) String() string {
	str := ""
	for i := 0; i < len(c); i++ {
		str += strconv.Itoa(c[i]) + " "
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
