package days

import "strings"

type Elf struct {
	Meals               []int
	TotalCalories       int
	CleaningSectorStart int
	CleaningSectorEnd   int
}

func (e *Elf) calcTotalCalories() {
	for _, meal := range e.Meals {
		e.TotalCalories += meal
	}
}

type Rucksack struct {
	Compartment1 string
	Compartment2 string
	Items        string
}

type CrateInstructions struct {
	NumberOfCrates   int
	OriginStack      int
	DestinationStack int
}

type HandHeldDevice struct {
	Messages []string
}

func (hhd *HandHeldDevice) findMessageStart(datastream string, uniqueChars int) int {
	var packets string
	for i, packet := range datastream {
		if i == 0 {
			packets += string(packet)
		} else {
			if strings.Contains(packets, string(packet)) {
				for j, char := range packets {
					if packet == char {
						packets += string(packet)
						packets = packets[j+1:]
					}
				}
			} else {
				packets += string(packet)
				if len(packets) == uniqueChars {
					return i + 1
				}
			}
		}
	}
	return 0
}
