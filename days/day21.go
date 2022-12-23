package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day21() {
	var fileName string
	if os.Getenv("MODE") == "TEST" {
		fileName = "inputfiles/Day21Sample.txt"
	} else {
		fileName = "inputfiles/Day21.txt"
	}
	f, err := os.Open(fileName)
	Check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	monkeyJobs := make(map[string]interface{})
	monkeyJobs2 := make(map[string]interface{})
	for scanner.Scan() {
		monkey := strings.Split(scanner.Text(), ": ")
		val, err := strconv.Atoi(monkey[1])
		if err == nil {
			monkeyJobs[monkey[0]] = float64(val)
			monkeyJobs2[monkey[0]] = float64(val)
		} else {
			monkeyJobs[monkey[0]] = monkey[1]
			monkeyJobs2[monkey[0]] = monkey[1]
		}
	}
	var allNumeric bool
	for !allNumeric {
		allNumeric = true
		for monkey, job := range monkeyJobs {
			if _, ok := job.(float64); ok {
				continue
			} else {
				allNumeric = false
				monkeyJob := strings.Split(fmt.Sprintf("%v", job), " ")
				monkey1, ok := monkeyJobs[monkeyJob[0]].(float64)
				monkey2, ok2 := monkeyJobs[monkeyJob[2]].(float64)
				if ok && ok2 {
					monkeyJobs[monkey] = performArithmetic(monkeyJob[1], monkey1, monkey2)
				}
			}
		}
	}
	fmt.Println("Day 21 Puzzle Solutions")
	fmt.Printf("Part 1: Monkey root will shout: %f\n", monkeyJobs["root"])
	monkeyJobs2["root"] = strings.FieldsFunc(monkeyJobs2["root"].(string), SplitOnOperatorOrSpace)
	isHumanSide := findHuman(monkeyJobs2["root"].([]string)[0], monkeyJobs2)
	var human string
	var monkey string
	if isHumanSide {
		human = monkeyJobs2["root"].([]string)[0]
		monkey = monkeyJobs2["root"].([]string)[1]
	} else {
		human = monkeyJobs2["root"].([]string)[1]
		monkey = monkeyJobs2["root"].([]string)[0]
	}
	allMonkey := findShout(monkeyJobs2[monkey].(string), monkeyJobs2, 0)
	humanShout := float64(1)
	initialEquality := findShout(monkeyJobs2[human].(string), monkeyJobs2, humanShout)
	allMonkeyGreater := allMonkey > initialEquality
	changeBy := int(math.Abs(allMonkey - initialEquality))
	for {
		equality := findShout(monkeyJobs2[human].(string), monkeyJobs2, humanShout)
		if allMonkey == equality {
			break
		} else {
			if allMonkeyGreater && allMonkey > equality {
				humanShout += float64(changeBy)
			} else if allMonkeyGreater {
				humanShout -= float64(changeBy)
				changeBy /= 2
			} else if !allMonkeyGreater && allMonkey < equality {
				humanShout += float64(changeBy)
			} else {
				humanShout -= float64(changeBy)
				changeBy /= 2
			}
		}
	}
	fmt.Printf("Part 2 Puzzle Solution: Human must shout %f\n", humanShout)
}

func findHuman(monkey string, monkeyJobs map[string]interface{}) bool {
	var isHuman bool
	switch monkeyJobs[monkey].(type) {
	case int:
		return false
	case string:
		monkeyPass := strings.FieldsFunc(monkeyJobs[monkey].(string), SplitOnOperatorOrSpace)
		for _, monkeyShout := range monkeyPass {
			if monkeyShout == "humn" {
				return true
			}
			isHuman = findHuman(monkeyShout, monkeyJobs)
			if isHuman {
				return true
			}
		}
	}
	return isHuman
}

func findShout(monkey string, monkeyJobs map[string]interface{}, humanShout float64) float64 {
	var val1 float64
	var val2 float64
	monkeyMath := strings.Split(monkey, " ")
	switch monkeyJobs[monkeyMath[0]].(type) {
	case float64:
		if monkeyMath[0] == "humn" {
			val1 = humanShout
		} else {
			val1 = monkeyJobs[monkeyMath[0]].(float64)
		}
	case string:
		val1 = findShout(monkeyJobs[monkeyMath[0]].(string), monkeyJobs, humanShout)
	}
	switch monkeyJobs[monkeyMath[2]].(type) {
	case float64:
		if monkeyMath[2] == "humn" {
			val2 = humanShout
		} else {
			val2 = monkeyJobs[monkeyMath[2]].(float64)
		}
	case string:
		val2 = findShout(monkeyJobs[monkeyMath[2]].(string), monkeyJobs, humanShout)
	}
	if monkeyMath[0] != "humn" || monkeyMath[2] != "humn" {
		monkeyJobs[monkey] = performArithmetic(monkeyMath[1], val1, val2)
	}
	return performArithmetic(monkeyMath[1], val1, val2)
}
