package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day14() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day14Sample.txt"
	} else {
		fileName = "inputfiles/Day14.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var rockMap [][]string
	for scanner.Scan() {
		rockMap = append(rockMap, strings.Split(scanner.Text(), " -> "))
	}
	rockFillMap := make(map[string]string)
	minX := math.Inf(1)
	maxX := math.Inf(-1)
	maxY := math.Inf(-1)
	for _, rockPath := range rockMap {
		for i := 0; i < len(rockPath)-1; i++ {
			rockPosition1 := strings.Split(rockPath[i], ",")
			rockPosition2 := strings.Split(rockPath[i+1], ",")
			rock1x, err := strconv.ParseFloat(rockPosition1[0], 64)
			Check(err)
			rock1y, err := strconv.ParseFloat(rockPosition1[1], 64)
			Check(err)
			rock2x, err := strconv.ParseFloat(rockPosition2[0], 64)
			Check(err)
			rock2y, err := strconv.ParseFloat(rockPosition2[1], 64)
			Check(err)
			if rock1x != rock2x {
				for j := math.Min(rock1x, rock2x); j <= math.Max(rock1x, rock2x); j++ {
					if j < minX {
						minX = j
					}
					if j > maxX {
						maxX = j
					}
					if rock1y > maxY {
						maxY = rock1y
					}
					rockFillMap[fmt.Sprint(int(j))+","+fmt.Sprint(int(rock1y))] = "rock"
				}
			} else {
				for j := math.Min(rock1y, rock2y); j < math.Max(rock1y, rock2y); j++ {
					if rock1x < minX {
						minX = rock1x
					}
					if rock1x > maxX {
						maxX = rock1x
					}
					if j > maxY {
						maxY = j
					}
					rockFillMap[fmt.Sprint(int(rock1x))+","+fmt.Sprint(int(j))] = "rock"
				}
			}
		}
	}
	unitsOfSand := pourSand(rockFillMap, minX, maxX, maxY, false)
	fmt.Println("Day 14 Puzzle Solution:")
	fmt.Printf("Part 1: Units of Sand that came to Rest: %d\n", unitsOfSand)
	unitsOfSand = pourSand(rockFillMap, minX, maxX, maxY, true)
	fmt.Printf("Part 2: Units of Sand that came to Rest: %d\n", unitsOfSand)
}

func pourSand(passedRockMap map[string]string, minX, maxX, maxY float64, part2 bool) int {
	var unitsOfSand int
	sandStartingLocation := [2]float64{500, 0}
	sandX := sandStartingLocation[0]
	sandY := sandStartingLocation[1]
	rockFillMap := make(map[string]string)
	for k, v := range passedRockMap {
		rockFillMap[k] = v
	}
	if part2 {
		minX = math.Inf(-1)
		maxX = math.Inf(1)
		maxY += 2
	}
	var stop bool
	for sandX <= maxX && sandX >= minX && sandY <= maxY && !stop {
		sandFalling := true
		sandLocation := sandStartingLocation
		for sandFalling {
			checkDown := fmt.Sprint(sandLocation[0]) + "," + fmt.Sprint(sandLocation[1]+1)
			checkDownLeft := fmt.Sprint(sandLocation[0]-1) + "," + fmt.Sprint(sandLocation[1]+1)
			checkDownRight := fmt.Sprint(sandLocation[0]+1) + "," + fmt.Sprint(sandLocation[1]+1)
			if sandLocation[1] == maxY-1 && part2 {
				sandFalling = false
				rockFillMap[fmt.Sprint(sandLocation[0])+","+fmt.Sprint(sandLocation[1])] = "sand"
				unitsOfSand++
			} else if sandLocation[0] < minX || sandLocation[0] > maxX || sandLocation[1] > maxY && !part2 {
				sandFalling = false
				sandX = sandLocation[0]
				sandY = sandLocation[1]
			} else if rockFillMap[checkDown] != "rock" && rockFillMap[checkDown] != "sand" {
				sandLocation[1]++
			} else if rockFillMap[checkDownLeft] != "rock" && rockFillMap[checkDownLeft] != "sand" {
				sandLocation[1]++
				sandLocation[0]--
			} else if rockFillMap[checkDownRight] != "rock" && rockFillMap[checkDownRight] != "sand" {
				sandLocation[1]++
				sandLocation[0]++
			} else {
				sandFalling = false
				sandX = sandLocation[0]
				sandY = sandLocation[1]
				rockFillMap[fmt.Sprint(sandLocation[0])+","+fmt.Sprint(sandLocation[1])] = "sand"
				unitsOfSand++
			}
		}
		if rockFillMap["500,0"] == "sand" {
			stop = true
		}
	}
	return unitsOfSand
}
