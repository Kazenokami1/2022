package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Day20() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day20Sample.txt"
	} else {
		fileName = "inputfiles/Day20.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	var unmixedFile []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		intToAdd, err := strconv.Atoi(scanner.Text())
		Check(err)
		unmixedFile = append(unmixedFile, intToAdd)
	}
	mixedFile, zeroIndex := decryptFile(unmixedFile, 1, 1)
	groveCoords, coordSum := findGroveCoordsAndSum(mixedFile, zeroIndex)
	fmt.Println("Day 20 Puzzle Solutions:")
	fmt.Printf("Part 1 - Grove Coordinates: %v, Sum: %d\n", groveCoords, coordSum)
	mixedFile, zeroIndex = decryptFile(unmixedFile, 811589153, 10)
	groveCoords, coordSum = findGroveCoordsAndSum(mixedFile, zeroIndex)
	fmt.Printf("Part 2 - Grove Coordinates: %v, Sum: %d\n", groveCoords, coordSum)
}

func decryptFile(unmixedFile []int, multiplier, decryptionCycles int) (map[int][2]int, int) {
	mixedFile := make(map[int][2]int)
	for i, val := range unmixedFile {
		mixedFile[i] = [2]int{i, val * multiplier}
	}
	var zeroIndex int
	for x := 0; x < decryptionCycles; x++ {
		for i := range unmixedFile {
			previousIndex := mixedFile[i][0]
			newIndex := previousIndex + mixedFile[i][1]
			if newIndex >= len(mixedFile) {
				newIndex %= len(mixedFile) - 1
			} else if newIndex <= 0 {
				newIndex %= len(mixedFile) - 1
				newIndex = len(mixedFile) + newIndex - 1
			}
			mixedFile[i] = [2]int{newIndex, mixedFile[i][1]}
			for k, v := range mixedFile {
				if v[0] > previousIndex && v[0] <= newIndex && i != k && previousIndex < newIndex {
					if v[0]-1 < 0 {
						mixedFile[k] = [2]int{len(mixedFile) + (v[0] - 1), v[1]}
					} else {
						mixedFile[k] = [2]int{v[0] - 1, v[1]}
					}
				} else if v[0] >= newIndex && v[0] < previousIndex && i != k && newIndex < previousIndex {
					if v[0]+1 > len(mixedFile) {
						mixedFile[k] = [2]int{v[0] + 1 - len(mixedFile), v[1]}
					} else {
						mixedFile[k] = [2]int{v[0] + 1, v[1]}
					}
				}
				if mixedFile[k][1] == 0 {
					zeroIndex = mixedFile[k][0]
				}
			}
		}
	}
	return mixedFile, zeroIndex
}

func findGroveCoordsAndSum(mixedFile map[int][2]int, zeroIndex int) ([]int, int) {
	var groveCoords []int
	for x := 0; x < 3; x++ {
		zeroIndex += 1000
		if zeroIndex >= len(mixedFile) {
			zeroIndex %= len(mixedFile)
		}
		for _, v := range mixedFile {
			if v[0] == zeroIndex {
				groveCoords = append(groveCoords, v[1])
				break
			}
		}
	}
	var coordSum int
	for _, val := range groveCoords {
		coordSum += val
	}
	return groveCoords, coordSum
}
