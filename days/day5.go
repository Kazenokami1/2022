package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	scanner := bufio.NewScanner(f)
	stacksInput := make(map[int][]string)
	var instructions []CrateInstructions
	for scanner.Scan() {
		if strings.Split(scanner.Text(), " ")[0] == "move" {
			stringInts := strings.Split(scanner.Text(), " ")
			numberOfCrates, err := strconv.Atoi(stringInts[1])
			Check(err)
			originStack, err := strconv.Atoi(stringInts[3])
			Check(err)
			destinationStack, err := strconv.Atoi(stringInts[5])
			Check(err)
			instructions = append(instructions, CrateInstructions{NumberOfCrates: numberOfCrates, OriginStack: originStack, DestinationStack: destinationStack})
		} else if len(scanner.Text()) > 1 && strings.Trim(scanner.Text(), " ")[0] == '[' {
			var j int
			for i := 1; i < len(scanner.Text()); i += 4 {
				j++
				if scanner.Text()[i] != ' ' {
					stacksInput[j] = append([]string{string(scanner.Text()[i])}, stacksInput[j]...)
				}
			}
		}
	}
	originalStacks := make(map[int][]string)
	for i := 1; i <= len(stacksInput); i++ {
		originalStacks[i] = stacksInput[i]
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
	for i := 1; i <= len(originalStacks); i++ {
		stacksInput[i] = originalStacks[i]
	}
	stacksInput = originalStacks
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
