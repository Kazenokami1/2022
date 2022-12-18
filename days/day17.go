package days

import (
	"bufio"
	"fmt"
	"os"
)

var blockShapes = make(map[int][][2]int)
var isRockOrFloor = make(map[[2]int]bool)

func init() {
	blockShapes[0] = [][2]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {0, 0}}
	blockShapes[1] = [][2]int{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}}
	blockShapes[2] = [][2]int{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}
	blockShapes[3] = [][2]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 0}}
	blockShapes[4] = [][2]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}, {0, 0}}
	for i := 0; i < 7; i++ {
		isRockOrFloor[[2]int{i, 0}] = true
	}
}

func Day17() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day17Sample.txt"
	} else {
		fileName = "inputfiles/Day17.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var jets string
	for scanner.Scan() {
		jets = scanner.Text()
	}
	highestRock := runTetris(jets, 2022)
	fmt.Println("Day 17 Puzzle Solutions:")
	fmt.Printf("Highest Rock after Block 2022: %d\n", highestRock)
	isRockOrFloor = make(map[[2]int]bool)
	for i := 0; i < 7; i++ {
		isRockOrFloor[[2]int{i, 0}] = true
	}
	highestRock = runTetris(jets, 1000000000000)
	fmt.Printf("Highest Rock after Block 1000000000000: %d\n", highestRock)
}

func runTetris(jets string, rockHeight int) int {
	jetPosition := 0
	var highestRock int
	jetMap := make(map[int][]int)
	blockHeightAtBlock := make(map[int]int)
	for block := 0; block < rockHeight; block++ {
		if block%5 == 0 {
			if val, ok := jetMap[jetPosition]; ok {
				loops := block - val[0]
				heightDifference := highestRock - val[1]
				remainingLoops := rockHeight - block
				multiple := remainingLoops / loops
				modulus := remainingLoops % loops
				highestRock = highestRock + multiple*heightDifference
				highestRock = highestRock + blockHeightAtBlock[val[0]+modulus] - val[1]
				break
			} else {
				jetMap[jetPosition] = []int{block, highestRock}
			}
		}
		blockHeightAtBlock[block] = highestRock
		rockFalling := true
		var currentBlock [5][2]int
		copy(currentBlock[:], blockShapes[block%5])
		for i, position := range currentBlock {
			currentBlock[i] = [2]int{position[0] + 2, position[1] + highestRock + 4}
		}
		for rockFalling {
			moveLeft := canMoveLeft(currentBlock)
			moveRight := canMoveRight(currentBlock)
			for i, position := range currentBlock {
				if jets[jetPosition] == '<' && moveLeft {
					currentBlock[i] = [2]int{position[0] - 1, position[1]}
				} else if jets[jetPosition] == '>' && moveRight {
					currentBlock[i] = [2]int{position[0] + 1, position[1]}
				}
			}
			jetPosition++
			if jetPosition == len(jets) {
				jetPosition = 0
			}
			moveDown := canMoveDown(currentBlock)
			for i, position := range currentBlock {
				if moveDown {
					currentBlock[i] = [2]int{position[0], position[1] - 1}
				} else {
					rockFalling = false
					isRockOrFloor[position] = true
				}
			}
		}
		for _, position := range currentBlock {
			if position[1] > highestRock {
				highestRock = position[1]
			}
		}
	}
	return highestRock
}

func canMoveLeft(block [5][2]int) bool {
	for _, position := range block {
		if position[0] == 0 || isRockOrFloor[[2]int{position[0] - 1, position[1]}] {
			return false
		}
	}
	return true
}

func canMoveRight(block [5][2]int) bool {
	for _, position := range block {
		if position[0] == 6 || isRockOrFloor[[2]int{position[0] + 1, position[1]}] {
			return false
		}
	}
	return true
}

func canMoveDown(block [5][2]int) bool {
	for _, position := range block {
		if isRockOrFloor[[2]int{position[0], position[1] - 1}] {
			return false
		}
	}
	return true
}
