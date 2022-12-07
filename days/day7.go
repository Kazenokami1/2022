package days

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func Day7() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day7Sample.txt"
	} else {
		fileName = "inputfiles/Day7.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()

	var terminalOutput [][]string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		terminalOutput = append(terminalOutput, strings.Split(scanner.Text(), " "))
	}
	var directories []*Directory
	var currentDir *Directory
	for _, output := range terminalOutput {
		if output[0] == "$" {
			if output[1] == "ls" {
				continue
			} else if output[2] != ".." {
				if output[2] == "/" {
					currentDir = &Directory{Name: output[2]}
				} else {
					for _, dir := range currentDir.Directories {
						if dir.Name == output[2] {
							currentDir = dir
							break
						}
					}
				}
				directories = append(directories, currentDir)
			} else {
				currentDir.calcDirectorySize()
				currentDir = currentDir.Parent
			}
		} else if output[0] == "dir" {
			currentDir.Directories = append(currentDir.Directories, &Directory{Name: output[1], Parent: currentDir})
		} else {
			size, err := strconv.Atoi(output[0])
			Check(err)
			currentDir.Files = append(currentDir.Files, &File{Name: output[1], Size: size})
		}
	}
	for currentDir.Parent != nil {
		currentDir.calcDirectorySize()
		currentDir = currentDir.Parent
	}
	currentDir.calcDirectorySize()
	var sum int
	for _, dir := range directories {
		if dir.TotalSize <= 100000 {
			sum += dir.TotalSize
		}
	}
	fmt.Println("Day 7 Puzzle Solutions:")
	fmt.Printf("Part 1 - Sum of Directory Size: %d", sum)
	fmt.Println()
	maxSpace := 70000000
	neededSpace := 30000000
	unusedSpace := maxSpace - currentDir.TotalSize
	var smallestDirectory Directory
	for _, dir := range directories {
		if unusedSpace+dir.TotalSize >= neededSpace {
			if reflect.DeepEqual(smallestDirectory, Directory{}) || dir.TotalSize < smallestDirectory.TotalSize {
				smallestDirectory = *dir
			}
		}
	}
	fmt.Printf("Part 2 - Smallest Directory to Delete is Directory %s with size %d", smallestDirectory.Name, smallestDirectory.TotalSize)
	fmt.Println()
}
