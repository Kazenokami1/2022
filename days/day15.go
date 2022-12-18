package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day15() {
	var fileName string
	var part1Row float64
	var part2Min float64
	var part2Max float64
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day15Sample.txt"
		part1Row = 10
		part2Min = 0
		part2Max = 20
	} else {
		fileName = "inputfiles/Day15.txt"
		part1Row = 2000000
		part2Min = 0
		part2Max = 4000000
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var sensors []Sensor
	gridSquares := make(map[[2]float64]string)
	for scanner.Scan() {
		coordinates := strings.FieldsFunc(scanner.Text(), SplitOnNumber)
		var coordinateFloats []float64
		for _, value := range coordinates {
			number, err := strconv.ParseFloat(value, 64)
			Check(err)
			coordinateFloats = append(coordinateFloats, number)
		}
		point1 := [2]float64{coordinateFloats[0], coordinateFloats[1]}
		point2 := [2]float64{coordinateFloats[2], coordinateFloats[3]}
		newSensor := Sensor{Position: point1, ClosestBeacon: point2, DistanceToClosestBeacon: findManhattanDistance(point1, point2), XValuesPerYCovered: make(map[float64][2]float64)}
		sensors = append(sensors, newSensor)
		gridSquares[point1] = "Sensor"
		gridSquares[point2] = "Beacon"
	}
	for _, sensor := range sensors {
		minX := sensor.Position[0] - sensor.DistanceToClosestBeacon
		maxX := sensor.Position[0] + sensor.DistanceToClosestBeacon
		minY := sensor.Position[1] - sensor.DistanceToClosestBeacon
		maxY := sensor.Position[1] + sensor.DistanceToClosestBeacon
		x := minX
		y := sensor.Position[1]
		for y <= maxY {
			val, ok := sensor.XValuesPerYCovered[y]
			if !ok {
				sensor.XValuesPerYCovered[y] = [2]float64{x, maxX - (x - minX)}
			} else {
				if val[0] > x {
					sensor.XValuesPerYCovered[y] = [2]float64{x, val[1]}
				} else if val[1] < maxX-(x-minX) {
					sensor.XValuesPerYCovered[y] = [2]float64{val[0], maxX - (x - minX)}
				}
			}
			y++
			x++
		}
		x = minX
		y = sensor.Position[1]
		for y >= minY {
			val, ok := sensor.XValuesPerYCovered[y]
			if !ok {
				sensor.XValuesPerYCovered[y] = [2]float64{x, maxX - (x - minX)}
			} else {
				if val[0] > x {
					sensor.XValuesPerYCovered[y] = [2]float64{x, val[1]}
				} else if val[1] < maxX-(x-minX) {
					sensor.XValuesPerYCovered[y] = [2]float64{val[0], maxX - (x - minX)}
				}
			}
			y--
			x++
		}
	}
	var emptySquares int
	xValues := [2]float64{math.Inf(1), math.Inf(-1)}
	for _, s := range sensors {
		val, ok := s.XValuesPerYCovered[part1Row]
		if ok {
			if val[0] < xValues[0] {
				xValues[0] = val[0]
			}
			if val[1] > xValues[1] {
				xValues[1] = val[1]
			}
		}
	}
	//Note: This only works IF the row doesn't contain an unknown spot
	for x := xValues[0]; x <= xValues[1]; x++ {
		_, ok := gridSquares[[2]float64{x, part1Row}]
		if !ok {
			emptySquares++
		}
	}
	fmt.Println("Day 15 Puzzle Solutions:")
	fmt.Printf("Part 1: Number of Squares in Row %f: %d\n", part1Row, emptySquares)
	y := part2Min
	for y <= part2Max {
		fullyCovered := false
		x := part2Min
		lastX := x
		for !fullyCovered {
			for _, s := range sensors {
				val, ok := s.XValuesPerYCovered[y]
				if ok {
					if val[0] <= x && val[1] >= x {
						x = val[1] + 1
						break
					}
				}
			}
			if lastX == x {
				fmt.Printf("Position of Distress Beacon: %f, %f\n", x, y)
				fmt.Printf("Part 2 Solution: %f\n", x*4000000+y)
				y = part2Max + 1
				fullyCovered = true
			} else {
				lastX = x
			}
			if x > part2Max {
				fullyCovered = true
			}
		}
		y++
	}
}
