package days

import (
	"bufio"
	"fmt"
	"os"
)

func Day6() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day6Sample.txt"
	} else {
		fileName = "inputfiles/Day6.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()

	var datastreams []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		datastreams = append(datastreams, scanner.Text())
	}
	var hhd *HandHeldDevice
	fmt.Println("Day 6 Puzzle Solution:")
	for _, datastream := range datastreams {
		messageStart := hhd.findMessageStart(datastream, 4)
		fmt.Printf("Part 1 Packet Start: %d", messageStart)
		fmt.Println()
	}
	for _, datastream := range datastreams {
		messageStart := hhd.findMessageStart(datastream, 14)
		fmt.Printf("Part 2 Packet Start: %d", messageStart)
		fmt.Println()
	}
}
