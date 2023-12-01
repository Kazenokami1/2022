package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var facingMap = make(map[int][2]int)
var maxColLen int
var rows int
var facing = 0

func init() {
	facingMap[0] = [2]int{1, 0}
	facingMap[1] = [2]int{0, 1}
	facingMap[2] = [2]int{-1, 0}
	facingMap[3] = [2]int{0, -1}
	rows = 1
}

func Day22() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day22Sample.txt"
	} else {
		fileName = "inputfiles/Day22.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	pathMap := make(map[[2]int]string)
	var path string
	var currentPosition [2]int
	for scanner.Scan() {
		if scanner.Text() != "" {
			for x, char := range scanner.Text() {
				if char == ' ' || char == '.' || char == '#' {
					pathMap[[2]int{x + 1, rows}] = string(char)
					if char == '.' && currentPosition == [2]int{0, 0} {
						currentPosition = [2]int{x + 1, rows}
					}
				} else {
					path += string(char)
				}
			}
			if len(scanner.Text()) > maxColLen {
				maxColLen = len(scanner.Text())
			}
		}
		rows++
	}
	for row := 1; row <= rows; row++ {
		for x := 1; x <= maxColLen; x++ {
			if _, ok := pathMap[[2]int{x, row}]; !ok {
				pathMap[[2]int{x, row}] = " "
			}
		}
	}
	pathSteps := strings.FieldsFunc(path, SplitOnNumber)
	var pathStepsInt []int
	for _, steps := range pathSteps {
		intSteps, err := strconv.Atoi(steps)
		Check(err)
		pathStepsInt = append(pathStepsInt, intSteps)
	}
	pathTurns := strings.FieldsFunc(path, SplitOnNonNumber)
	for i, steps := range pathStepsInt {
		if i != len(pathSteps)-1 {
			currentPosition = moveAlongPath(currentPosition, facing, steps, pathMap)
			if pathTurns[i] == "R" {
				facing = (facing + 1) % 4
			} else {
				facing = (facing + 3) % 4
			}
		} else {
			currentPosition = moveAlongPath(currentPosition, facing, steps, pathMap)
		}
	}
	fmt.Println("Day 22 Puzzle Solutions")
	fmt.Printf("Part 1 Solution: %d\n", currentPosition[0]*4+currentPosition[1]*1000+facing)
}

func moveAlongPath(currentPosition [2]int, facing, steps int, pathMap map[[2]int]string) [2]int {
	direction := facingMap[facing]
	nextPosition := [2]int{currentPosition[0] + direction[0], currentPosition[1] + direction[1]}
	for i := 0; i < steps; i++ {
		if _, ok := pathMap[nextPosition]; ok {
			if pathMap[nextPosition] == "." {
				currentPosition = nextPosition
				nextPosition = [2]int{nextPosition[0] + direction[0], nextPosition[1] + direction[1]}
			} else if pathMap[nextPosition] == "#" {
				return currentPosition
			} else {
				for {
					nextPosition = [2]int{nextPosition[0] + direction[0], nextPosition[1] + direction[1]}
					if nextPosition[0] > maxColLen {
						nextPosition[0] = 1
					} else if nextPosition[0] < 1 {
						nextPosition[0] = maxColLen
					}
					if nextPosition[1] > rows {
						nextPosition[1] = 1
					} else if nextPosition[1] < 1 {
						nextPosition[1] = rows
					}
					if pathMap[nextPosition] != " " {
						break
					}
				}
				i--
			}
		} else {
			if nextPosition[0] > maxColLen {
				nextPosition[0] = 1
			} else if nextPosition[0] < 1 {
				nextPosition[0] = maxColLen
			}
			if nextPosition[1] > rows {
				nextPosition[1] = 1
			} else if nextPosition[1] < 1 {
				nextPosition[1] = rows
			}
			i--
		}
	}
	return currentPosition
}
