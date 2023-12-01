package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var dirMap = make(map[int]string)

func init() {
	dirMap[0] = "E"
	dirMap[1] = "W"
	dirMap[2] = "S"
	dirMap[3] = "N"
}

func Day23() {
	var fileName string
	minX := math.MaxInt
	maxX := math.MinInt
	minY := math.MaxInt
	maxY := math.MinInt
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day23Sample.txt"
	} else {
		fileName = "inputfiles/Day23.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var y int
	elfMap := make(map[[2]int]interface{})
	for scanner.Scan() {
		for i, char := range scanner.Text() {
			if char == '#' {
				elfMap[[2]int{i, y}] = ""
				if i < minX {
					minX = i
				}
				if i > maxX {
					maxX = i
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
			}
		}
		y++
	}
	var direction int
	var doneMoving bool
	var round int
	for !doneMoving {
		round++
		if round == 11 {
			fmt.Println("Day 23 Puzzle Solutions:")
			fmt.Printf("Part 1: Number of Empty Ground Tiles: %d\n", ((maxX-minX+1)*(maxY-minY+1) - len(elfMap)))
		}
		elvesToMove := make(map[[2]int]string)
		for k := range elfMap {
			for k2 := range elfMap {
				_, ok := elvesToMove[k]
				if ok {
					break
				}
				if k != k2 && math.Abs(float64(k[0]-k2[0])) <= 1 && math.Abs(float64(k[1]-k2[1])) <= 1 {
					elvesToMove[k] = ""
				}
			}
		}
		if len(elvesToMove) == 0 {
			doneMoving = true
			fmt.Printf("Part 2 Solution took %d Rounds\n", round)
		}
		for x := 0; x < 4; x++ {
			for k := range elvesToMove {
				switch dirMap[direction%4] {
				case "N":
					_, nw := elfMap[[2]int{k[0] - 1, k[1] - 1}]
					_, n := elfMap[[2]int{k[0], k[1] - 1}]
					_, ne := elfMap[[2]int{k[0] + 1, k[1] - 1}]
					if !(nw || n || ne) {
						elvesToMove[k] = "N"
					}
				case "S":
					_, sw := elfMap[[2]int{k[0] - 1, k[1] + 1}]
					_, s := elfMap[[2]int{k[0], k[1] + 1}]
					_, se := elfMap[[2]int{k[0] + 1, k[1] + 1}]
					if !(sw || s || se) {
						elvesToMove[k] = "S"
					}
				case "E":
					_, ne := elfMap[[2]int{k[0] + 1, k[1] - 1}]
					_, e := elfMap[[2]int{k[0] + 1, k[1]}]
					_, se := elfMap[[2]int{k[0] + 1, k[1] + 1}]
					if !(ne || e || se) {
						elvesToMove[k] = "E"
					}
				case "W":
					_, nw := elfMap[[2]int{k[0] - 1, k[1] - 1}]
					_, w := elfMap[[2]int{k[0] - 1, k[1]}]
					_, sw := elfMap[[2]int{k[0] - 1, k[1] + 1}]
					if !(nw || w || sw) {
						elvesToMove[k] = "W"
					}
				}

			}
			direction++
		}
		direction--
		for k, v := range elvesToMove {
			switch v {
			case "N":
				_, ok := elfMap[[2]int{k[0], k[1] - 1}]
				if ok {
					elfMap[[2]int{k[0], k[1] - 2}] = ""
					delete(elfMap, [2]int{k[0], k[1] - 1})
				} else {
					elfMap[[2]int{k[0], k[1] - 1}] = ""
					delete(elfMap, k)
					if k[1]-1 < minY {
						minY = k[1] - 1
					}
				}
			case "S":
				_, ok := elfMap[[2]int{k[0], k[1] + 1}]
				if ok {
					elfMap[[2]int{k[0], k[1] + 2}] = ""
					delete(elfMap, [2]int{k[0], k[1] + 1})
				} else {
					elfMap[[2]int{k[0], k[1] + 1}] = ""
					delete(elfMap, k)
					if k[1]+1 > maxY {
						maxY = k[1] + 1
					}
				}
			case "E":
				_, ok := elfMap[[2]int{k[0] + 1, k[1]}]
				if ok {
					elfMap[[2]int{k[0] + 2, k[1]}] = ""
					delete(elfMap, [2]int{k[0] + 1, k[1]})
				} else {
					elfMap[[2]int{k[0] + 1, k[1]}] = ""
					delete(elfMap, k)
					if k[0]+1 > maxX {
						maxX = k[0] + 1
					}
				}
			case "W":
				_, ok := elfMap[[2]int{k[0] - 1, k[1]}]
				if ok {
					elfMap[[2]int{k[0] - 2, k[1]}] = ""
					delete(elfMap, [2]int{k[0] - 1, k[1]})
				} else {
					elfMap[[2]int{k[0] - 1, k[1]}] = ""
					delete(elfMap, k)
					if k[0]-1 < minX {
						minX = k[0] - 1
					}
				}
			}
		}
		// for y := minY; y <= maxY; y++ {
		// 	for x := minX; x <= maxX; x++ {
		// 		_, ok := elfMap[[2]int{x, y}]
		// 		if ok {
		// 			fmt.Print("#")
		// 		} else {
		// 			fmt.Print(".")
		// 		}
		// 	}
		// 	fmt.Println()
		// }
		// fmt.Println()
	}
}
