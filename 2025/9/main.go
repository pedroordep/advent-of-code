package main

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/pedroordep/advent-of-code/utils"
)

type Point struct {
	x, y int
}

type Rectangle struct {
	p1, p2 Point
	area   int
}

func main() {
	input := utils.GetInputFile(2025, 9)

	fmt.Println("Part 1", Part1(strings.Trim(string(input), "\n")))
	fmt.Println("Part 2", Part2(strings.Trim(string(input), "\n")))
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")

	points := []Point{}
	rectangles := []Rectangle{}

	for _, line := range lines {
		coordinates := strings.Split(line, ",")
		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])
		currentPoint := Point{x, y}

		for _, point := range points {
			area := calculateArea(point, currentPoint)

			rectangles = append(rectangles, Rectangle{
				p1:   point,
				p2:   currentPoint,
				area: area,
			})
		}
		points = append(points, currentPoint)
	}

	slices.SortFunc(rectangles, func(a, b Rectangle) int {
		return cmp.Compare(b.area, a.area)
	})

	// fmt.Println(rectangles)

	return rectangles[0].area
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")

	points := []Point{}
	rectangles := []Rectangle{}

	for _, line := range lines {
		coordinates := strings.Split(line, ",")
		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])
		currentPoint := Point{x, y}

		for _, point := range points {
			area := calculateArea(point, currentPoint)

			rectangles = append(rectangles, Rectangle{
				p1:   point,
				p2:   currentPoint,
				area: area,
			})
		}
		points = append(points, currentPoint)
	}

	slices.SortFunc(rectangles, func(a, b Rectangle) int {
		return cmp.Compare(b.area, a.area)
	})

	// fmt.Println(len(rectangles))

	for _, rectangle := range rectangles {
		// fmt.Println("- Checking rectangle number", i, rectangle)
		if isRectangleFullyInsidePolygon(rectangle.p1, rectangle.p2, points) {
			// fmt.Println("  - Rectangle", rectangle, "is fully inside polygon")
			return rectangle.area
		}
	}

	return rectangles[0].area
}

func isRectangleFullyInsidePolygon(p1, p2 Point, polygon []Point) bool {
	minX := min(p1.x, p2.x)
	maxX := max(p1.x, p2.x)
	minY := min(p1.y, p2.y)
	maxY := max(p1.y, p2.y)
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if x == minX || x == maxX || y == minY || y == maxY {
				if !isPointInsideAxisAlignedPolygon(Point{x, y}, polygon) {
					return false
				}
			}
		}
	}
	return true
}

func isPointInsideAxisAlignedPolygon(p Point, polygon []Point) bool {
	// fmt.Println("Checking if point", p, "is inside polygon")

	n := len(polygon)

	// Ray-casting: count crossings of vertical edges to the right of p
	crossings := 0
	numEdgedsOnPoint := 0
	for i := 0; i < n; i++ {
		a := polygon[i]
		b := polygon[(i+1)%n]

		// Only vertical edges
		if a.x == b.x && p.x < a.x {
			if p.y == a.y || p.y == b.y {
				// fmt.Println("Point", p, "is on vertex of edge from", a, "to", b, "but next is inside, so ignoring it")
				numEdgedsOnPoint++
				continue
			}
			if (p.y > min(a.y, b.y)) && (p.y < max(a.y, b.y)) {
				// fmt.Println("Edge from", a, "to", b, "crosses ray from", p)
				crossings++
			}
		}
	}

	if numEdgedsOnPoint == 1 {
		return true
	}

	return crossings%2 == 1
}

func calculateArea(p1, p2 Point) int {
	return (max(p1.x, p2.x) - min(p1.x, p2.x) + 1) * (max(p1.y, p2.y) - min(p1.y, p2.y) + 1)
}
