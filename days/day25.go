package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func Day25() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day25Sample.txt"
	} else {
		fileName = "inputfiles/Day25.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	var fuel []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fuel = append(fuel, scanner.Text())
	}
	var decFuel []float64
	for _, code := range fuel {
		decFuel = append(decFuel, snafuToDec(code))
	}
	var decSum float64
	for _, val := range decFuel {
		decSum += val
	}
	snafuSum := decToSnafu(decSum)
	fmt.Println("Day 25 Puzzle Solutions")
	fmt.Printf("Part 1 Solution: %s\n", snafuSum)
}

func snafuToDec(code string) float64 {
	var dec float64
	for pos, char := range code {
		multiplier := math.Pow(5, float64(len(code)-pos-1))
		switch char {
		case '=':
			dec -= 2 * multiplier
		case '-':
			dec -= multiplier
		case '2':
			dec += 2 * multiplier
		case '1':
			dec += multiplier
		case '0':
		}
	}
	return dec
}

func decToSnafu(decSum float64) string {
	var expFound bool
	var exp float64
	var snafuSum string
	for !expFound {
		if math.Pow(5, exp) < decSum {
			exp++
		} else {
			expFound = true
			exp--
		}
	}
	for exp >= 0 {
		if decSum/math.Pow(5, exp) >= 1.5 {
			snafuSum += "2"
			decSum -= 2 * math.Pow(5, exp)
		} else if decSum/math.Pow(5, exp) >= 0.5 {
			snafuSum += "1"
			decSum -= math.Pow(5, exp)
		} else if decSum/math.Pow(5, exp) <= -1.5 {
			snafuSum += "="
			decSum += 2 * math.Pow(5, exp)
		} else if decSum/math.Pow(5, exp) <= -0.5 {
			snafuSum += "-"
			decSum += math.Pow(5, exp)
		} else {
			snafuSum += "0"
		}
		exp--
	}
	return snafuSum
}
