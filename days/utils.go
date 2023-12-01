package days

import (
	"math"
	"unicode"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func performArithmetic(operand string, value1 float64, value2 float64) float64 {
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
		return float64(int(value1) % int(value2))
	}
	return 0
}

func findManhattanDistance(point1, point2 [2]float64) float64 {
	return math.Abs(point1[0]-point2[0]) + math.Abs(point1[1]-point2[1])
}

func SplitOnNumber(r rune) bool {
	return !unicode.IsNumber(r) && r != '-'
}

func SplitOnNonNumber(r rune) bool {
	return unicode.IsNumber(r)
}

func SplitOnOperatorOrSpace(r rune) bool {
	if r == '+' || r == '-' || r == '*' || r == '/' || r == ' ' {
		return true
	}
	return false
}
