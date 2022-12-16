package days

import "math"

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func performArithmetic(operand string, value1 int, value2 int) int {
	switch operand {
	case "*":
		return value1 * value2
	case "+":
		return value1 + value2
	case "-":
		return value1 - value2
	case "/":
		return value1 / value2
	case "%":
		return value1 % value2
	}
	return 0
}

func findManhattanDistance(point1, point2 [2]float64) float64 {
	return math.Abs(point1[0]-point2[0]) + math.Abs(point1[1]-point2[1])
}
