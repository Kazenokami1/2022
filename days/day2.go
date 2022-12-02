package days

import (
	"bufio"
	"fmt"
	"os"
)

var scores = make(map[string]int)
var part1results = make(map[string]string)
var part2results = make(map[byte]string)

func init() {
	scores["lossA"] = 3
	scores["drawA"] = 4
	scores["winA"] = 8
	scores["lossB"] = 1
	scores["drawB"] = 5
	scores["winB"] = 9
	scores["lossC"] = 2
	scores["drawC"] = 6
	scores["winC"] = 7
	part1results["A X"] = "draw"
	part1results["A Y"] = "win"
	part1results["A Z"] = "loss"
	part1results["B X"] = "loss"
	part1results["B Y"] = "draw"
	part1results["B Z"] = "win"
	part1results["C X"] = "win"
	part1results["C Y"] = "loss"
	part1results["C Z"] = "draw"
	part2results['X'] = "loss"
	part2results['Y'] = "draw"
	part2results['Z'] = "win"
}

func Day2() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day2Sample.txt"
	} else {
		fileName = "inputfiles/Day2.txt"
	}
	f, err := os.Open(fileName)
	check(err)
	defer f.Close()

	var games []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		games = append(games, scanner.Text())
	}
	var totalScore int
	for _, game := range games {
		totalScore += scores[part1results[game]+string(game[0])]
	}
	fmt.Println("Day2 Puzzle Solutions:")
	fmt.Printf("Total Score Part 1: %d", totalScore)
	fmt.Println()

	totalScore = 0
	for _, game := range games {
		totalScore += scores[part2results[game[2]]+string(game[0])]
	}
	fmt.Printf("Total Score Part 2: %d", totalScore)
	fmt.Println()
}
