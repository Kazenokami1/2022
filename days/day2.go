package days

import (
	"bufio"
	"fmt"
	"os"
)

func Day2() {
	part1results := makeResults("Part1")
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
		totalScore += play(part1results[game], game[0])
	}
	fmt.Printf("Total Score Part 1: %d", totalScore)
	fmt.Println()
	totalScore = 0
	part2results := makeResults("Part2")
	for _, game := range games {
		totalScore += play(part2results[string(game[2])], game[0])
	}
	fmt.Printf("Total Score Part 2: %d", totalScore)
	fmt.Println()
}

func play(result string, opp byte) int {
	switch result {
	case "loss":
		if opp == 'A' {
			return 3
		} else if opp == 'B' {
			return 1
		} else if opp == 'C' {
			return 2
		}
	case "draw":
		if opp == 'A' {
			return 4
		} else if opp == 'B' {
			return 5
		} else if opp == 'C' {
			return 6
		}
	case "win":
		if opp == 'A' {
			return 8
		} else if opp == 'B' {
			return 9
		} else if opp == 'C' {
			return 7
		}
	}
	return 0
}

func makeResults(part string) map[string]string {
	results := make(map[string]string)
	if part == "Part1" {
		results["A X"] = "draw"
		results["A Y"] = "win"
		results["A Z"] = "loss"
		results["B X"] = "loss"
		results["B Y"] = "draw"
		results["B Z"] = "win"
		results["C X"] = "win"
		results["C Y"] = "loss"
		results["C Z"] = "draw"
	} else if part == "Part2" {
		results["X"] = "loss"
		results["Y"] = "draw"
		results["Z"] = "win"
	}
	return results
}
