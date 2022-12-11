package days

import (
	"strings"
)

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
	Messages      []string
	Cycles        int
	XRegister     int
	ScreenDisplay [6]string
}

func (hhd *HandHeldDevice) findMessageStart(datastream string, uniqueChars int) int {
	var packets string
	for i, packet := range datastream {
		if i == 0 {
			packets += string(packet)
		} else {
			duplicateIndex := strings.Index(packets, string(packet))
			if duplicateIndex != -1 {
				packets += string(packet)
				packets = packets[duplicateIndex+1:]
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

func (hhd *HandHeldDevice) performCPUOperation(command string, xRegister int) {
	if command == "noop" {
		hhd.Cycles++
	} else if command == "addx" {
		hhd.Cycles += 2
		hhd.XRegister += xRegister
	}
}

func (hhd *HandHeldDevice) addScreenPixel(runTime int) {
	cycles := hhd.Cycles + runTime
	screenRow := cycles / 40
	if hhd.XRegister-1 == cycles%40 || hhd.XRegister == cycles%40 || hhd.XRegister+1 == cycles%40 {
		hhd.ScreenDisplay[screenRow] += "#"
	} else {
		hhd.ScreenDisplay[screenRow] += "."
	}
}

type Directory struct {
	Name        string
	Parent      *Directory
	Directories []*Directory
	Files       []*File
	TotalSize   int
}

func (d *Directory) calcDirectorySize() {
	var size int
	for _, dir := range d.Directories {
		size += dir.TotalSize
	}
	for _, file := range d.Files {
		size += file.Size
	}
	d.TotalSize = size
}

type File struct {
	Name string
	Size int
}

type Tree struct {
	Height       int
	Visible      bool
	VisibleEast  bool
	VisibleWest  bool
	VisibleNorth bool
	VisibleSouth bool
	ViewEast     int
	ViewWest     int
	ViewNorth    int
	ViewSouth    int
	ScenicScore  int
}

func (t *Tree) calcScenicScore() {
	t.ScenicScore = t.ViewEast * t.ViewNorth * t.ViewSouth * t.ViewWest
}
