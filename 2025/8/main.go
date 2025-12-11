package main

import (
	"cmp"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/pedroordep/advent-of-code/utils"
)

type JunctionBox struct {
	x, y, z int
}

type JunctionBoxPair struct {
	box1, box2 JunctionBox
	distance   float64
}

// Go has no native set type :(
type Circuits []map[JunctionBox]struct{}

func main() {
	input := utils.GetInputFile(2025, 8)

	fmt.Println("Part 1", Part1(strings.Trim(string(input), "\n"), 1000))
	fmt.Println("Part 2", Part2(strings.Trim(string(input), "\n"), -1))
}

func Part1(input string, numConnections int) int {
	circuits, _ := connectJunctionBoxes(input, numConnections)

	return len(circuits[0]) * len(circuits[1]) * len(circuits[2])
}

func Part2(input string, numConnections int) int {
	_, lastConnection := connectJunctionBoxes(input, numConnections)

	return lastConnection.box1.x * lastConnection.box2.x
}

// If numconnections is -1, connect all junction boxes
func connectJunctionBoxes(input string, numConnections int) (circuits Circuits, lastConnection JunctionBoxPair) {
	lines := strings.Split(input, "\n")

	junctionBoxes := []JunctionBox{}
	junctionBoxPairs := []JunctionBoxPair{}

	for _, line := range lines {
		coordinates := strings.Split(line, ",")
		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])
		z, _ := strconv.Atoi(coordinates[2])

		for _, junctionBox := range junctionBoxes {
			distance := calculateDistance(x, y, z, junctionBox.x, junctionBox.y, junctionBox.z)

			if x+y+z < junctionBox.x+junctionBox.y+junctionBox.z {
				junctionBoxPairs = append(junctionBoxPairs, JunctionBoxPair{
					box1:     JunctionBox{x, y, z},
					box2:     junctionBox,
					distance: distance,
				})
			} else {
				junctionBoxPairs = append(junctionBoxPairs, JunctionBoxPair{
					box1:     junctionBox,
					box2:     JunctionBox{x, y, z},
					distance: distance,
				})
			}
		}
		junctionBoxes = append(junctionBoxes, JunctionBox{x, y, z})
	}

	slices.SortFunc(junctionBoxPairs, func(a, b JunctionBoxPair) int {
		return cmp.Compare(a.distance, b.distance)
	})

	// for _, pair := range junctionBoxPairs {
	// 	fmt.Println("Pair:", pair)
	// }

	for i, pair := range junctionBoxPairs {
		if i >= numConnections && numConnections != -1 {
			break
		}

		// fmt.Println("---", i+1, "Connecting pair", pair.box1, "<->", pair.box2)

		box1InCircuit := -1
		box2InCircuit := -1
		for j := range circuits {
			if _, ok := circuits[j][pair.box1]; ok {
				// fmt.Println(pair.box1, "exists in circuit", j, circuits[j])
				box1InCircuit = j
			}
			if _, ok := circuits[j][pair.box2]; ok {
				// fmt.Println(pair.box2, "exists in circuit", j, circuits[j])
				box2InCircuit = j
			}
		}

		if box1InCircuit == box2InCircuit && box1InCircuit != -1 {
			// fmt.Println("Both boxes already in the same circuit, skipping")
		} else if box1InCircuit != -1 && box2InCircuit != -1 {
			// fmt.Println("Merging circuits", circuits[box1InCircuit], "and", circuits[box2InCircuit])
			for box := range circuits[box2InCircuit] {
				circuits[box1InCircuit][box] = struct{}{}
			}
			circuits = append(circuits[:box2InCircuit], circuits[box2InCircuit+1:]...)
			lastConnection = pair
		} else if box1InCircuit != -1 {
			// fmt.Println("Adding", pair.box2, "to circuit", circuits[box1InCircuit])
			circuits[box1InCircuit][pair.box2] = struct{}{}
			lastConnection = pair
		} else if box2InCircuit != -1 {
			// fmt.Println("Adding", pair.box1, "to circuit", circuits[box2InCircuit])
			circuits[box2InCircuit][pair.box1] = struct{}{}
			lastConnection = pair
		} else {
			// fmt.Println("Creating new circuit with", pair.box1, "and", pair.box2)
			circuits = append(circuits, map[JunctionBox]struct{}{
				pair.box1: {},
				pair.box2: {},
			})
		}
	}

	slices.SortFunc(circuits, func(a, b map[JunctionBox]struct{}) int {
		return cmp.Compare(len(b), len(a))
	})

	// for _, connection := range circuits {
	// 	fmt.Println("Have connection", connection)
	// }

	return circuits, lastConnection
}

func calculateDistance(x1, y1, z1, x2, y2, z2 int) float64 {
	return math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2) + (z1-z2)*(z1-z2)))
}
