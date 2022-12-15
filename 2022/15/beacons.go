package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

type SensorData map[image.Point]image.Point
type PointSet map[image.Point]struct{}

func main() {
	file, _ := os.ReadFile("./input2.txt")
	lines := strings.Split(string(file), "\n")
	sensors := parseSensors(lines)

	// sensorsAt10 := sensorInfoAtY(sensors, 10)
	// fmt.Println(sensorsAt10, len(sensorsAt10))
	fmt.Println("part 1:", len(sensorInfoAtY(sensors, 2000000)))
}

func sensorInfoAtY(sensors SensorData, atY int) PointSet {
	pointsAtY := PointSet{}
	for sensor, beacon := range sensors {
		vector := beacon.Sub(sensor)
		dist := Abs(vector.X) + Abs(vector.Y)
		for x := sensor.X - dist; x <= sensor.X+dist; x++ {
			tmpVector := image.Point{x, atY}.Sub(sensor)
			tmpDist := Abs(tmpVector.X) + Abs(tmpVector.Y)
			if tmpDist <= dist {
				pointsAtY[image.Point{x, atY}] = struct{}{}
			}
		}
	}

	// remove beacons
	for _, beacon := range sensors {
		_, exists := pointsAtY[beacon]
		if exists {
			delete(pointsAtY, beacon)
		}
	}
	return pointsAtY
}

func parseSensors(lines []string) SensorData {
	sensors := SensorData{}
	for _, line := range lines {
		sx, sy, bx, by := 0, 0, 0, 0
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		sensors[image.Point{sx, sy}] = image.Point{bx, by}
	}
	return sensors
}

func Abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
