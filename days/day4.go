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

	var elfPairs [][]Elf
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var elf1 Elf
		var elf2 Elf
		cleaningAssignments := strings.Split(scanner.Text(), ",")
		elf1CleaningAssignments := strings.Split(cleaningAssignments[0], "-")
		elf1.CleaningSectorStart, err = strconv.Atoi(string(elf1CleaningAssignments[0]))
		Check(err)
		elf1.CleaningSectorEnd, err = strconv.Atoi(string(elf1CleaningAssignments[1]))
		Check(err)
		elf2CleaningAssignments := strings.Split(cleaningAssignments[1], "-")
		elf2.CleaningSectorStart, err = strconv.Atoi(string(elf2CleaningAssignments[0]))
		Check(err)
		elf2.CleaningSectorEnd, err = strconv.Atoi(string(elf2CleaningAssignments[1]))
		Check(err)
		elfPairs = append(elfPairs, []Elf{elf1, elf2})
	}
	var fullyOverlappedPairs int
	var partialOverlappedPairs int
	for _, elfPair := range elfPairs {
		if elfPair[0].CleaningSectorStart >= elfPair[1].CleaningSectorStart && elfPair[0].CleaningSectorEnd <= elfPair[1].CleaningSectorEnd {
			fullyOverlappedPairs++
		} else if elfPair[1].CleaningSectorStart >= elfPair[0].CleaningSectorStart && elfPair[1].CleaningSectorEnd <= elfPair[0].CleaningSectorEnd {
			fullyOverlappedPairs++
		} else if elfPair[0].CleaningSectorStart >= elfPair[1].CleaningSectorStart && elfPair[0].CleaningSectorStart <= elfPair[1].CleaningSectorEnd {
			partialOverlappedPairs++
		} else if elfPair[1].CleaningSectorStart >= elfPair[0].CleaningSectorStart && elfPair[1].CleaningSectorStart <= elfPair[0].CleaningSectorEnd {
			partialOverlappedPairs++
		}
	}
	fmt.Println("Day 4 Puzzle Solutions:")
	fmt.Printf("Numbers of Pairs that should be reconsidered: %d", fullyOverlappedPairs)
	fmt.Println()
	fmt.Printf("Number of Pairs that overlap at all: %d", fullyOverlappedPairs+partialOverlappedPairs)
	fmt.Println()
}
