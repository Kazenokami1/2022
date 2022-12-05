package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func init() {
}

func Day4() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day4Sample.txt"
	} else {
		fileName = "inputfiles/Day4.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()

	var elfPairs [][]elf
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var elf1 elf
		var elf2 elf
		cleaningAssignments := strings.Split(scanner.Text(), ",")
		elf1CleaningAssignments := strings.Split(cleaningAssignments[0], "-")
		elf1.cleaningSectorStart, err = strconv.Atoi(string(elf1CleaningAssignments[0]))
		Check(err)
		elf1.cleaningSectorEnd, err = strconv.Atoi(string(elf1CleaningAssignments[1]))
		Check(err)
		elf2CleaningAssignments := strings.Split(cleaningAssignments[1], "-")
		elf2.cleaningSectorStart, err = strconv.Atoi(string(elf2CleaningAssignments[0]))
		Check(err)
		elf2.cleaningSectorEnd, err = strconv.Atoi(string(elf2CleaningAssignments[1]))
		Check(err)
		elfPairs = append(elfPairs, []elf{elf1, elf2})
	}
	var fullyOverlappedPairs int
	var partialOverlappedPairs int
	for _, elfPair := range elfPairs {
		if elfPair[0].cleaningSectorStart >= elfPair[1].cleaningSectorStart && elfPair[0].cleaningSectorEnd <= elfPair[1].cleaningSectorEnd {
			fullyOverlappedPairs++
		} else if elfPair[1].cleaningSectorStart >= elfPair[0].cleaningSectorStart && elfPair[1].cleaningSectorEnd <= elfPair[0].cleaningSectorEnd {
			fullyOverlappedPairs++
		} else if elfPair[0].cleaningSectorStart >= elfPair[1].cleaningSectorStart && elfPair[0].cleaningSectorStart <= elfPair[1].cleaningSectorEnd {
			partialOverlappedPairs++
		} else if elfPair[1].cleaningSectorStart >= elfPair[0].cleaningSectorStart && elfPair[1].cleaningSectorStart <= elfPair[0].cleaningSectorEnd {
			partialOverlappedPairs++
		}
	}
	fmt.Println("Day 4 Puzzle Solutions:")
	fmt.Printf("Numbers of Pairs that should be reconsidered: %d", fullyOverlappedPairs)
	fmt.Println()
	fmt.Printf("Number of Pairs that overlap at all: %d", fullyOverlappedPairs+partialOverlappedPairs)
	fmt.Println()
}
