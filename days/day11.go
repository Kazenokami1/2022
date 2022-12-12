package days

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day11() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day11Sample.txt"
	} else {
		fileName = "inputfiles/Day11.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var monkeyList []*Monkey
	for scanner.Scan() {
		monkeyNumberString := strings.Split(scanner.Text(), " ")
		monkeyNumber, err := strconv.Atoi(monkeyNumberString[1][0 : len(monkeyNumberString[1])-1])
		Check(err)
		scanner.Scan()
		items := strings.SplitAfterN(strings.TrimLeft(scanner.Text(), " "), " ", 3)[2]
		itemSlice := strings.Split(items, ", ")
		var monkeyItems []int
		for _, item := range itemSlice {
			itemWorryLevel, err := strconv.Atoi(item)
			Check(err)
			monkeyItems = append(monkeyItems, itemWorryLevel)
		}
		scanner.Scan()
		operation := strings.SplitAfterN(strings.TrimLeft(scanner.Text(), " "), " ", 5)[4]
		scanner.Scan()
		test := strings.SplitAfterN(strings.TrimLeft(scanner.Text(), " "), " ", 2)[1]
		scanner.Scan()
		testTrue := strings.TrimLeft(scanner.Text(), " ")
		scanner.Scan()
		testFalse := strings.TrimLeft(scanner.Text(), " ")
		monkeyTest := []string{test, testTrue, testFalse}
		monkeyList = append(monkeyList, &Monkey{Number: monkeyNumber, Items: monkeyItems, Operation: operation, Test: monkeyTest})
		scanner.Scan()
	}
	lcm := 1
	for _, monkey := range monkeyList {
		monkeyTrue, err := strconv.Atoi(monkey.Test[1][len(monkey.Test[1])-1 : len(monkey.Test[1])])
		Check(err)
		monkeyFalse, err := strconv.Atoi(monkey.Test[2][len(monkey.Test[2])-1 : len(monkey.Test[2])])
		Check(err)
		for _, tossMonkey := range monkeyList {
			if tossMonkey.Number == monkeyTrue {
				monkey.ThrowOnTrue = tossMonkey
			} else if tossMonkey.Number == monkeyFalse {
				monkey.ThrowOnFalse = tossMonkey
			}
		}
		multString := strings.Split(monkey.Test[0], "divisible by ")
		mult, err := strconv.Atoi(multString[1])
		Check(err)
		lcm *= mult
	}
	for round := 1; round <= 20; round++ {
		for _, monkey := range monkeyList {
			monkey.playRound(round, 1, lcm)
		}
	}
	var monkeyInspectors []int
	for _, monkey := range monkeyList {
		monkeyInspectors = append(monkeyInspectors, monkey.ItemsInspected)
	}
	sort.Ints(monkeyInspectors)
	fmt.Println("Day 11 Puzzle Solutions:")
	fmt.Printf("Part 1 Solution: %d\n", monkeyInspectors[len(monkeyInspectors)-1]*monkeyInspectors[len(monkeyInspectors)-2])
	day11part2(fileName)
}

func day11part2(fileName string) {
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var monkeyList []*Monkey
	for scanner.Scan() {
		monkeyNumberString := strings.Split(scanner.Text(), " ")
		monkeyNumber, err := strconv.Atoi(monkeyNumberString[1][0 : len(monkeyNumberString[1])-1])
		Check(err)
		scanner.Scan()
		items := strings.SplitAfterN(strings.TrimLeft(scanner.Text(), " "), " ", 3)[2]
		itemSlice := strings.Split(items, ", ")
		var monkeyItems []int
		for _, item := range itemSlice {
			itemWorryLevel, err := strconv.Atoi(item)
			Check(err)
			monkeyItems = append(monkeyItems, itemWorryLevel)
		}
		scanner.Scan()
		operation := strings.SplitAfterN(strings.TrimLeft(scanner.Text(), " "), " ", 5)[4]
		scanner.Scan()
		test := strings.SplitAfterN(strings.TrimLeft(scanner.Text(), " "), " ", 2)[1]
		scanner.Scan()
		testTrue := strings.TrimLeft(scanner.Text(), " ")
		scanner.Scan()
		testFalse := strings.TrimLeft(scanner.Text(), " ")
		monkeyTest := []string{test, testTrue, testFalse}
		monkeyList = append(monkeyList, &Monkey{Number: monkeyNumber, Items: monkeyItems, Operation: operation, Test: monkeyTest})
		scanner.Scan()
	}
	lcm := 1
	for _, monkey := range monkeyList {
		monkeyTrue, err := strconv.Atoi(monkey.Test[1][len(monkey.Test[1])-1 : len(monkey.Test[1])])
		Check(err)
		monkeyFalse, err := strconv.Atoi(monkey.Test[2][len(monkey.Test[2])-1 : len(monkey.Test[2])])
		Check(err)
		for _, tossMonkey := range monkeyList {
			if tossMonkey.Number == monkeyTrue {
				monkey.ThrowOnTrue = tossMonkey
			} else if tossMonkey.Number == monkeyFalse {
				monkey.ThrowOnFalse = tossMonkey
			}
		}
		multString := strings.Split(monkey.Test[0], "divisible by ")
		mult, err := strconv.Atoi(multString[1])
		Check(err)
		lcm *= mult
	}
	for round := 1; round <= 10000; round++ {
		for _, monkey := range monkeyList {
			monkey.playRound(round, 2, lcm)
		}
	}
	var monkeyInspectors []int
	for _, monkey := range monkeyList {
		monkeyInspectors = append(monkeyInspectors, monkey.ItemsInspected)
	}
	sort.Ints(monkeyInspectors)
	fmt.Printf("Part 2 Solution: %d\n", monkeyInspectors[len(monkeyInspectors)-1]*monkeyInspectors[len(monkeyInspectors)-2])
}
