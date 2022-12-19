package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day18() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day18Sample.txt"
	} else {
		fileName = "inputfiles/Day18.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	var lavaDropletString []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lavaDropletString = append(lavaDropletString, scanner.Text())
	}
	minX := int(^uint(0) >> 1)
	var maxX int
	minY := int(^uint(0) >> 1)
	var maxY int
	minZ := int(^uint(0) >> 1)
	var maxZ int
	var lavaDropletCubes [][3]int
	for _, side := range lavaDropletString {
		sideString := strings.Split(side, ",")
		var sideCoords [3]int
		for i, val := range sideString {
			pos, err := strconv.Atoi(val)
			Check(err)
			sideCoords[i] = pos
			if i == 0 && pos < minX {
				minX = pos
			} else if i == 0 && pos > maxX {
				maxX = pos
			} else if i == 1 && pos < minY {
				minY = pos
			} else if i == 1 && pos > maxY {
				maxY = pos
			} else if i == 2 && pos < minZ {
				minZ = pos
			} else if i == 2 && pos > maxZ {
				maxZ = pos
			}
		}
		lavaDropletCubes = append(lavaDropletCubes, sideCoords)
	}
	var dropletCubes []*DropletCube
	for _, cube := range lavaDropletCubes {
		dropletCubes = append(dropletCubes, &DropletCube{Position: cube})
	}
	airCubes := make(map[[3]int]*DropletCube)
	for x := minX - 1; x <= maxX+1; x++ {
		for y := minY - 1; y <= maxY+1; y++ {
			for z := minZ - 1; z <= maxZ+1; z++ {
				airCubes[[3]int{x, y, z}] = &DropletCube{Position: [3]int{x, y, z}}
			}
		}
	}
	for _, cubeOne := range dropletCubes {
		for _, cubeTwo := range dropletCubes {
			var posDiff int
			for i := 0; i < 3; i++ {
				posDiff += int(math.Abs(float64(cubeOne.Position[i]) - float64(cubeTwo.Position[i])))
			}
			if posDiff == 1 {
				cubeOne.ConnectedCubes = append(cubeOne.ConnectedCubes, cubeTwo)
			}
		}
		delete(airCubes, cubeOne.Position)
	}
	var surfaceArea int
	for _, cube := range dropletCubes {
		surfaceArea = surfaceArea + 6 - len(cube.ConnectedCubes)
		if len(cube.ConnectedCubes) < 6 {
			determineAirCubes(cube, airCubes)
		}
	}
	fmt.Println("Day 18 Puzzle Solutions")
	fmt.Printf("Total Surface Area of Lava Droplet: %d\n", surfaceArea)
	surfaceArea = 0
	changed := true
	for changed {
		changed = false
		for _, airCube := range airCubes {
			var connectedCubes int
			for _, airCubeTwo := range airCubes {
				var posDiff int
				for i := 0; i < 3; i++ {
					posDiff += int(math.Abs(float64(airCube.Position[i]) - float64(airCubeTwo.Position[i])))
				}
				if posDiff == 1 {
					connectedCubes++
				}
			}
			for _, cube := range dropletCubes {
				var posDiff int
				for i := 0; i < 3; i++ {
					posDiff += int(math.Abs(float64(airCube.Position[i]) - float64(cube.Position[i])))
				}
				if posDiff == 1 {
					connectedCubes++
				}
			}
			if connectedCubes < 6 {
				delete(airCubes, airCube.Position)
				for _, cube := range dropletCubes {
					for i, cCube := range cube.ConnectedCubes {
						if cCube == airCube {
							cube.ConnectedCubes = append(cube.ConnectedCubes[:i], cube.ConnectedCubes[i+1:]...)
							changed = true
						}
					}
				}
			}
		}
	}
	for _, cube := range dropletCubes {
		surfaceArea = surfaceArea + 6 - len(cube.ConnectedCubes)
	}
	fmt.Printf("Total Exterior Surface Area: %d\n", surfaceArea)
}

func determineAirCubes(cube *DropletCube, airCubes map[[3]int]*DropletCube) {
	checkSidesMap := make(map[[3]int]interface{})
	checkSidesMap[[3]int{-1, 0, 0}] = ""
	checkSidesMap[[3]int{1, 0, 0}] = ""
	checkSidesMap[[3]int{0, -1, 0}] = ""
	checkSidesMap[[3]int{0, 1, 0}] = ""
	checkSidesMap[[3]int{0, 0, -1}] = ""
	checkSidesMap[[3]int{0, 0, 1}] = ""
	for _, connectedCube := range cube.ConnectedCubes {
		var diff [3]int
		for i := 0; i < 3; i++ {
			diff[i] = cube.Position[i] - connectedCube.Position[i]
		}
		delete(checkSidesMap, diff)
	}
	for key := range checkSidesMap {
		var airCube [3]int
		for i := 0; i < 3; i++ {
			airCube[i] = cube.Position[i] - key[i]
		}
		_, ok := airCubes[airCube]
		if !ok {
			airCubes[airCube] = &DropletCube{Position: airCube}
		}
		cube.ConnectedCubes = append(cube.ConnectedCubes, airCubes[airCube])
	}
}
