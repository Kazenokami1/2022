package days

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var priorityMap = make(map[string]int)

func init() {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for priority, letter := range alphabet {
		priorityMap[string(letter)] = priority + 1
	}
}

func Day3() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day3Sample.txt"
	} else {
		fileName = "inputfiles/Day3.txt"
	}
	f, err := os.Open(fileName)
	check(err)
	defer f.Close()

	var rucksacks []rucksack
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var rucksack rucksack
		rucksack.Items = scanner.Text()
		compartmentsize := len(rucksack.Items) / 2
		rucksack.Compartment1 = rucksack.Items[0:compartmentsize]
		rucksack.Compartment2 = rucksack.Items[compartmentsize:]
		rucksacks = append(rucksacks, rucksack)
	}
	var totalPriority int
	for _, rucksack := range rucksacks {
		for _, item := range rucksack.Compartment1 {
			if strings.Contains(rucksack.Compartment2, string(item)) {
				totalPriority += priorityMap[string(item)]
				break
			}
		}
	}
	fmt.Println("Day2 Puzzle Solutions:")
	fmt.Printf("Total Priority: %d", totalPriority)
	fmt.Println()
	totalPriority = 0
	for i := 0; i < len(rucksacks); i++ {
		for _, badge := range rucksacks[i].Items {
			if strings.Contains(rucksacks[i+1].Items, string(badge)) && strings.Contains(rucksacks[i+2].Items, string(badge)) {
				totalPriority += priorityMap[string(badge)]
				i += 2
				break
			}
		}

	}
	fmt.Printf("Total Group Priority: %d", totalPriority)
	fmt.Println()
}
