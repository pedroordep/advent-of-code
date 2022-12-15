package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

type SensorData map[image.Point]image.Point
type SensorDistance map[image.Point]int
type PointSet map[image.Point]struct{}

func main() {
	file, _ := os.ReadFile("./input2.txt")
	lines := strings.Split(string(file), "\n")
	sensors, distances := parseSensors(lines)

	fmt.Println("part 1:", len(sensorInfoAtY(sensors, 2000000)))
	beaconPos := beaconPos(distances, 0, 0, 4000000, 4000000)
	fmt.Println("part 2:", beaconPos.X*4000000+beaconPos.Y)
}

func beaconPos(distances SensorDistance, xMin, yMin, xMax, yMax int) image.Point {
	for sensor, distance := range distances {
		for x1, y1 := sensor.X-distance, sensor.Y; x1 <= sensor.X; {
			if x1 >= xMin && x1 <= xMax && y1 >= yMin && y1 <= yMax && !sensorsReachPoint(distances, image.Point{x1 - 1, y1}) {
				return image.Point{x1 - 1, y1}
			}
			x1++
			y1++
		}
		for x1, y1 := sensor.X, sensor.Y+distance; y1 >= sensor.Y; {
			if x1 >= xMin && x1 <= xMax && y1 >= yMin && y1 <= yMax && !sensorsReachPoint(distances, image.Point{x1, y1 + 1}) {
				return image.Point{x1, y1 + 1}
			}
			x1++
			y1--
		}
		for x1, y1 := sensor.X+distance, sensor.Y; x1 >= sensor.X; {
			if x1 >= xMin && x1 <= xMax && y1 >= yMin && y1 <= yMax && !sensorsReachPoint(distances, image.Point{x1 + 1, y1}) {
				return image.Point{x1 + 1, y1}
			}
			x1--
			y1--
		}
		for x1, y1 := sensor.X, sensor.Y-distance; y1 <= sensor.Y; {
			if x1 >= xMin && x1 <= xMax && y1 >= yMin && y1 <= yMax && !sensorsReachPoint(distances, image.Point{x1, y1 - 1}) {
				return image.Point{x1, y1 - 1}
			}
			x1--
			y1++
		}
	}

	// for x := xMin; x <= xMax; x++ {
	// 	for y := yMin; y <= yMax; y++ {
	// 		found := false
	// 		for sensor, distance := range distances {
	// 			if distance >= Abs(x-sensor.X)+Abs(y-sensor.Y) {
	// 				found = true
	// 				break
	// 			}
	// 		}
	// 		if !found {
	// 			return image.Point{x, y}
	// 		}
	// 	}
	// }

	return image.Point{-1, -1}
}

func sensorsReachPoint(sensors SensorDistance, point image.Point) bool {
	for sensor, distance := range sensors {
		vector := sensor.Sub(point)
		if distance >= Abs(vector.X)+Abs(vector.Y) {
			return true
		}
	}
	return false
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

func parseSensors(lines []string) (SensorData, SensorDistance) {
	sensors := SensorData{}
	sensorsDist := SensorDistance{}
	for _, line := range lines {
		sx, sy, bx, by := 0, 0, 0, 0
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		sensors[image.Point{sx, sy}] = image.Point{bx, by}
		v := image.Point{sx, sy}.Sub(image.Point{bx, by})
		sensorsDist[image.Point{sx, sy}] = Abs(v.X) + Abs(v.Y)
	}
	return sensors, sensorsDist
}

func Abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
