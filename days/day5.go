package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var stacksInput = make(map[int][]string)

func initializeStacks() {
	/*stacksInput[1] = []string{"Z", "N"}
	stacksInput[2] = []string{"M", "C", "D"}
	stacksInput[3] = []string{"P"}*/
	stacksInput[1] = []string{"S", "L", "W"}
	stacksInput[2] = []string{"J", "T", "N", "Q"}
	stacksInput[3] = []string{"S", "C", "H", "F", "J"}
	stacksInput[4] = []string{"T", "R", "M", "W", "N", "G", "B"}
	stacksInput[5] = []string{"T", "R", "L", "S", "D", "H", "Q", "B"}
	stacksInput[6] = []string{"M", "J", "B", "V", "F", "H", "R", "L"}
	stacksInput[7] = []string{"D", "W", "R", "N", "J", "M"}
	stacksInput[8] = []string{"B", "Z", "T", "F", "H", "N", "D", "J"}
	stacksInput[9] = []string{"H", "L", "Q", "N", "B", "F", "T"}
}

func Day5() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day5Sample.txt"
	} else {
		fileName = "inputfiles/Day5.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()

	initializeStacks()

	scanner := bufio.NewScanner(f)
	var instructions []CrateInstructions
	for scanner.Scan() {
		stringInts := strings.Split(scanner.Text(), " ")
		numberOfCrates, err := strconv.Atoi(stringInts[1])
		Check(err)
		originStack, err := strconv.Atoi(stringInts[3])
		Check(err)
		destinationStack, err := strconv.Atoi(stringInts[5])
		Check(err)
		instructions = append(instructions, CrateInstructions{NumberOfCrates: numberOfCrates, OriginStack: originStack, DestinationStack: destinationStack})
	}

	for _, instruction := range instructions {
		for i := 0; i < instruction.NumberOfCrates; i++ {
			destinationStack := stacksInput[instruction.DestinationStack]
			originStack := stacksInput[instruction.OriginStack]
			stacksInput[instruction.DestinationStack] = append(destinationStack, originStack[len(originStack)-1])
			stacksInput[instruction.OriginStack] = originStack[0 : len(originStack)-1]
		}
	}
	fmt.Println("Day 5 Puzzle Solutions:")
	fmt.Print("Crates on top of Stacks (Part 1): ")
	for i := 1; i <= len(stacksInput); i++ {
		fmt.Print(stacksInput[i][len(stacksInput[i])-1])
	}
	fmt.Println()
	initializeStacks()
	for _, instruction := range instructions {
		var originStack []string
		for i := 0; i < instruction.NumberOfCrates; i++ {
			destinationStack := stacksInput[instruction.DestinationStack]
			originStack = stacksInput[instruction.OriginStack]
			stacksInput[instruction.DestinationStack] = append(destinationStack, originStack[len(originStack)-instruction.NumberOfCrates+i])
		}
		stacksInput[instruction.OriginStack] = originStack[0 : len(originStack)-instruction.NumberOfCrates]
	}
	fmt.Print("Crates on top of Stacks (Part 2): ")
	for i := 1; i <= len(stacksInput); i++ {
		fmt.Print(stacksInput[i][len(stacksInput[i])-1])
	}
	fmt.Println()
}
