package days

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

func Day1() {
	start := time.Now()
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day1Sample.txt"
	} else {
		fileName = "inputfiles/Day1.txt"
	}
	f, err := os.Open(fileName)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var elves []elf
	var meals []int
	for scanner.Scan() {
		if scanner.Text() != "" {
			meal, err := strconv.Atoi(scanner.Text())
			check(err)
			meals = append(meals, meal)
		} else {
			var elf elf
			elf.meals = meals
			elf.calcTotalCalories()
			elves = append(elves, elf)
			meals = []int{}
		}
	}
	//Should figure out a better way to get the last elf added to my array
	var elf elf
	elf.meals = meals
	elf.calcTotalCalories()
	elves = append(elves, elf)
	var highestCalories int
	//Part 1 Answer
	for _, elf := range elves {
		if elf.totalCalories > highestCalories {
			highestCalories = elf.totalCalories
		}
	}
	fmt.Printf("Highest Elf has %d calories", highestCalories)
	fmt.Println()
	var totalCaloriesPerElf []int
	for _, elf := range elves {
		totalCaloriesPerElf = append(totalCaloriesPerElf, elf.totalCalories)
	}
	sort.Ints(totalCaloriesPerElf)
	var top3Calories int
	for i := len(totalCaloriesPerElf); i > len(totalCaloriesPerElf)-3; i-- {
		top3Calories += totalCaloriesPerElf[i-1]
	}
	fmt.Printf("Total Calories carried by top 3 elves is %d", top3Calories)
	fmt.Println()
	duration := time.Since(start)
	fmt.Print("Time Since Start: ")
	fmt.Println(duration)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
