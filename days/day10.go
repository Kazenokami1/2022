package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day10() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day10Sample.txt"
	} else {
		fileName = "inputfiles/Day10.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var commandList []string
	for scanner.Scan() {
		commandList = append(commandList, scanner.Text())
	}
	var sumXTimesCycles int
	hhd := HandHeldDevice{XRegister: 1}
	for _, command := range commandList {
		commandSplit := strings.Split(command, " ")
		if commandSplit[0] == "noop" {
			hhd.addScreenPixel(0)
			if hhd.Cycles%40 == 19 {
				cycles := hhd.Cycles + 1
				sumXTimesCycles += hhd.XRegister * cycles
			}
			hhd.performCPUOperation(commandSplit[0], 0)
		} else {
			hhd.addScreenPixel(0)
			hhd.addScreenPixel(1)
			cycles := hhd.Cycles % 40
			if cycles == 18 || cycles == 19 {
				cycles = hhd.Cycles - cycles + 20
				sumXTimesCycles += hhd.XRegister * cycles
			}
			addx, err := strconv.Atoi(commandSplit[1])
			Check(err)
			hhd.performCPUOperation(commandSplit[0], addx)
		}
	}
	fmt.Println("Day 10 Puzzle Solutions:")
	fmt.Printf("Part 1: %d\n", sumXTimesCycles)
	for _, row := range hhd.ScreenDisplay {
		fmt.Println(row)
	}
}
