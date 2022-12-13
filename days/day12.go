package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"strings"
)

func Day12() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day12Sample.txt"
	} else {
		fileName = "inputfiles/Day12.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	start, end, possibleStarts := parseInput(scanner)
	findStepsTaken(start, end, false)
	fmt.Println("Day 12 Puzzle Solutions:")
	fmt.Printf("Part 1 - Shortest Path to End: %f\n", end.StepsTaken)
	for _, possibleStart := range possibleStarts {
		findStepsTaken(possibleStart, end, true)
	}
	fmt.Printf("Part 2 - Shortest Path to End from Start Point: %f\n", end.StepsTaken)
}

func parseInput(scanner *bufio.Scanner) (*GridSquare, *GridSquare, []*GridSquare) {
	var heightMap []string
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for scanner.Scan() {
		heightMap = append(heightMap, scanner.Text())
	}
	var gridSquares []*GridSquare
	var startSquare *GridSquare
	var endSquare *GridSquare
	var possibleStarts []*GridSquare
	for row, rowHeights := range heightMap {
		for col, height := range rowHeights {
			var start bool
			var end bool
			var possibleStart bool
			if height == 'S' {
				start = true
				height = 'a'
			} else if height == 'E' {
				end = true
				height = 'z'
			}
			intHeight := strings.Index(alphabet, string(height))
			if intHeight == 0 {
				possibleStart = true
			}
			position := []int{len(heightMap) - 1 - row, col}
			newSquare := &GridSquare{Height: intHeight, Position: position, Start: start, End: end, StepsTaken: math.Inf(1), PossibleStart: possibleStart}
			if start {
				startSquare = newSquare
				startSquare.StepsTaken = 0
			} else if end {
				endSquare = newSquare
			} else if possibleStart {
				possibleStarts = append(possibleStarts, newSquare)
			}
			gridSquares = append(gridSquares, newSquare)
		}
	}
	for _, square := range gridSquares {
		for _, square2 := range gridSquares {
			if math.Abs(float64(square.Position[0]-square2.Position[0])) < 2 && math.Abs(float64(square.Position[1]-square2.Position[1])) < 2 && !reflect.DeepEqual(square.Position, square2.Position) && (square.Position[1]-square2.Position[1] == 0 || square.Position[0]-square2.Position[0] == 0) && square2.Height-square.Height <= 1 {
				square.AccessibleNeighbors = append(square.AccessibleNeighbors, square2)
			}
		}
	}
	return startSquare, endSquare, possibleStarts
}

func findStepsTaken(current, end *GridSquare, part2 bool) {
	loopAgain := false
	current.Visited = true
	if part2 && current.PossibleStart {
		current.StepsTaken = 0
	}
	for _, neighbor := range current.AccessibleNeighbors {
		if current.StepsTaken+1 < end.StepsTaken && current.StepsTaken+1 < neighbor.StepsTaken {
			neighbor.StepsTaken = current.StepsTaken + 1
			neighbor.Visited = false
			loopAgain = true
		}
		if neighbor.End {
			return
		} else if loopAgain && !neighbor.Visited {
			findStepsTaken(neighbor, end, part2)
		}
	}
}
