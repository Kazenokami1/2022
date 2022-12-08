package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Day8() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day8Sample.txt"
	} else {
		fileName = "inputfiles/Day8.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var forest [][]*Tree
	for scanner.Scan() {
		var heights []*Tree
		for _, char := range scanner.Text() {
			height, err := strconv.Atoi(string(char))
			Check(err)
			heights = append(heights, &Tree{Height: height, VisibleEast: true, VisibleWest: true, VisibleNorth: true, VisibleSouth: true})
		}
		forest = append(forest, heights)
	}
	var bestScenicScore int
	var visibleTrees int
	for i, row := range forest {
		for j, tree := range row {
			for x := j - 1; x >= 0; x-- {
				if row[x].Height >= tree.Height {
					tree.VisibleWest = false
					tree.ViewWest = j - x
					break
				}
				tree.ViewWest = j - x
			}
			for x := j + 1; x < len(row); x++ {
				if row[x].Height >= tree.Height {
					tree.VisibleEast = false
					tree.ViewEast = x - j
					break
				}
				tree.ViewEast = x - j
			}
			for y := i - 1; y >= 0; y-- {
				if forest[y][j].Height >= tree.Height {
					tree.VisibleNorth = false
					tree.ViewNorth = i - y
					break
				}
				tree.ViewNorth = i - y
			}
			for y := i + 1; y < len(forest); y++ {
				if forest[y][j].Height >= tree.Height {
					tree.VisibleSouth = false
					tree.ViewSouth = y - i
					break
				}
				tree.ViewSouth = y - i
			}
			if tree.VisibleEast || tree.VisibleNorth || tree.VisibleSouth || tree.VisibleWest {
				tree.Visible = true
				visibleTrees++
			}
			tree.calcScenicScore()
			if tree.ScenicScore > bestScenicScore {
				bestScenicScore = tree.ScenicScore
			}
		}
	}
	fmt.Println("Day 8 Puzzle Solution:")
	fmt.Printf("Part 1 - Number of Visible Trees from any Direction: %d\n", visibleTrees)
	fmt.Printf("Part 2 - Highest Scenic Score: %d\n", bestScenicScore)
}
