package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var bridgeSteps = make(map[string]interface{})

func Day9() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day9Sample.txt"
	} else {
		fileName = "inputfiles/Day9.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var moveList []string
	for scanner.Scan() {
		moveList = append(moveList, scanner.Text())
	}
	bridgeSteps[fmt.Sprint(0)+","+fmt.Sprint(0)] = nil
	headPos := []int{0, 0}
	tailPos := []int{0, 0}
	for _, movement := range moveList {
		move := strings.Split(movement, " ")
		direction := move[0]
		distance, err := strconv.Atoi(move[1])
		Check(err)
		for i := 1; i <= distance; i++ {
			headPos = moveToNextBridgeStep(direction, headPos)
			distanceApart := math.Sqrt(math.Pow(float64(headPos[0]-tailPos[0]), 2) + math.Pow(float64(headPos[1]-tailPos[1]), 2))
			if distanceApart >= float64(2) {
				tailPos = moveTail(tailPos, headPos, 1, 1)
			}
		}
	}
	fmt.Println("Day 9 Puzzle Solution:")
	fmt.Printf("Part 1: The tail visited %d steps\n", len(bridgeSteps))
	day9part2(moveList)
}

func day9part2(moveList []string) {
	bridgeSteps = make(map[string]interface{})
	bridgeSteps[fmt.Sprint(0)+","+fmt.Sprint(0)] = nil
	var knots [][]int
	for i := 0; i < 10; i++ {
		knots = append(knots, []int{0, 0})
	}
	for _, movement := range moveList {
		move := strings.Split(movement, " ")
		direction := move[0]
		distance, err := strconv.Atoi(move[1])
		Check(err)
		for i := 1; i <= distance; i++ {
			knots[0] = moveToNextBridgeStep(direction, knots[0])
			for knot := 0; knot < 9; knot++ {
				distanceApart := math.Sqrt(math.Pow(float64(knots[knot][0]-knots[knot+1][0]), 2) + math.Pow(float64(knots[knot][1]-knots[knot+1][1]), 2))
				if distanceApart >= float64(2) {
					knots[knot+1] = moveTail(knots[knot+1], knots[knot], 9, knot+1)
				}
			}
		}
	}
	fmt.Printf("Part 2: The tail visited %d steps\n", len(bridgeSteps))
}

func moveToNextBridgeStep(direction string, headPos []int) []int {
	switch direction {
	case "U":
		return []int{headPos[0], headPos[1] + 1}
	case "D":
		return []int{headPos[0], headPos[1] - 1}
	case "L":
		return []int{headPos[0] - 1, headPos[1]}
	case "R":
		return []int{headPos[0] + 1, headPos[1]}
	}
	return headPos
}

func moveTail(tail []int, head []int, numberOfKnots int, knot int) []int {
	x := head[0]
	y := head[1]
	if head[1]-tail[1] == 2 {
		y--
	} else if head[1]-tail[1] == -2 {
		y++
	}
	if head[0]-tail[0] == 2 {
		x--
	} else if head[0]-tail[0] == -2 {
		x++
	}
	if numberOfKnots == knot {
		bridgeSteps[fmt.Sprint(x)+","+fmt.Sprint(y)] = nil
	}
	return []int{x, y}
}
