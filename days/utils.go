package days

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
